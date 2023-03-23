package azopenai

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

var (
	DefaultStreamReaderPrefix = []byte("data: ")
	DefaultStreamReaderStop   = []byte("[DONE]")
	DefaultStreamReaderDelim  = byte('\n')
)

// StreamReader is a stream reader for Azure OpenAI API when stream enabled.
type StreamReader[T any] struct {
	Reader io.ReadCloser

	Prefix []byte
	Stop   []byte
	Delim  byte

	UnMarshaler func([]byte, any) error
}

// RecvChan returns a channel that receives the stream data.
func (s *StreamReader[T]) RecvChan(errCallback func(error)) (<-chan T, error) {
	if err := s.defaults(); err != nil {
		return nil, err
	}
	if errCallback == nil {
		errCallback = func(error) {}
	}

	ch := make(chan T)
	reader := bufio.NewReader(s.Reader)

	var delta T
	var err error
	go func() {
		defer close(ch)
		defer func() {
			if err != nil && err != io.EOF {
				errCallback(err)
			}
		}()
		defer s.Reader.Close()

		for err == nil {
			delta, err = s.read(reader)
			if err != nil {
				return
			}
			ch <- delta
		}
	}()
	return ch, nil
}

func (s *StreamReader[T]) defaults() error {
	if s.Reader == nil {
		return fmt.Errorf("StreamReader.Reader is nil")
	}
	if s.Prefix == nil {
		s.Prefix = DefaultStreamReaderPrefix
	}
	if s.Stop == nil {
		s.Stop = DefaultStreamReaderStop
	}
	if s.Delim == 0 {
		s.Delim = DefaultStreamReaderDelim
	}
	if s.UnMarshaler == nil {
		s.UnMarshaler = json.Unmarshal
	}
	return nil
}

func (s *StreamReader[T]) read(reader *bufio.Reader) (T, error) {
	var rv T
	var line []byte
	var err error
	for len(line) == 0 {
		line, err = reader.ReadBytes(s.Delim)
		if err != nil {
			return rv, err
		}
		line = bytes.TrimSpace(line)
	}
	line = bytes.TrimPrefix(line, s.Prefix)
	if bytes.Equal(line, s.Stop) {
		return rv, io.EOF
	}
	if err := s.UnMarshaler(line, &rv); err != nil {
		return rv, err
	}
	return rv, nil
}

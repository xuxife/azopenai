//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenai

// ChatCompletionsClientCreateResponse contains the response from method ChatCompletionsClient.Create.
type ChatCompletionsClientCreateResponse struct {
	ChatCompletions
}

// ChatCompletionsClientCreateStreamResponse contains the response from method ChatCompletionsClient.CreateStream.
type ChatCompletionsClientCreateStreamResponse struct {
	*StreamReader[ChatCompletions]
}

// CompletionsClientCreateResponse contains the response from method CompletionsClient.Create.
type CompletionsClientCreateResponse struct {
	Completions
	// ApimRequestID contains the information returned from the apim-request-id header response.
	ApimRequestID *string
}

// CompletionsClientCreateStreamResponse contains the response from method CompletionsClient.CreateStream.
type CompletionsClientCreateStreamResponse struct {
	*StreamReader[Completions]
}

// EmbeddingsClientCreateResponse contains the response from method EmbeddingsClient.Create.
type EmbeddingsClientCreateResponse struct {
	Embeddings
}

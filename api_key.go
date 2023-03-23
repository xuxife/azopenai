package azopenai

import (
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	azpolicy "github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

const (
	apiKeyHeaderName = "Api-Key"
)

// APIKeyClient is an alternative authentication method to use with the Azure OpenAI API.
type APIKeyClient struct {
	key string
	pl  runtime.Pipeline
}

// NewAPIKeyClient creates a new instance of APIKeyClient with the specified values.
func NewAPIKeyClient(apiKey string, options *azpolicy.ClientOptions) *APIKeyClient {
	return &APIKeyClient{
		key: apiKey,
		pl:  NewAPIKeyPipeline(moduleName, moduleVersion, apiKey, runtime.PipelineOptions{}, options),
	}
}

// NewAPIKeyPipeline creates a new pipeline for the API key authentication.
func NewAPIKeyPipeline(module, version string, apiKey string, plOpts runtime.PipelineOptions, options *azpolicy.ClientOptions) runtime.Pipeline {
	authPolicy := APIKeyPolicy(apiKey)
	plOpts.PerRetry = append(plOpts.PerRetry, authPolicy)
	return runtime.NewPipeline(module, version, plOpts, options)
}

// APIKeyPolicy is a policy that adds an API key to the request.
type APIKeyPolicy string

// Do implement azpolicy.Policy interface.
func (p APIKeyPolicy) Do(req *azpolicy.Request) (*http.Response, error) {
	req.Raw().Header.Set(apiKeyHeaderName, string(p))
	return req.Next()
}

func getPipeline(apiKey *APIKeyClient, internal *arm.Client) runtime.Pipeline {
	if apiKey != nil {
		return apiKey.pl
	}
	return internal.Pipeline()
}

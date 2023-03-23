//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azopenai

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// EmbeddingsClient contains the methods for the Embeddings group.
// Don't use this type directly, use NewEmbeddingsClient() instead.
type EmbeddingsClient struct {
	internal *arm.Client
}

// NewEmbeddingsClient creates a new instance of EmbeddingsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEmbeddingsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*EmbeddingsClient, error) {
	cl, err := arm.NewClient(moduleName+".EmbeddingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EmbeddingsClient{
	internal: cl,
	}
	return client, nil
}

// Create - Get a vector representation of a given input that can be easily consumed by machine learning models and algorithms.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-15-preview
//   - deploymentID - The deployment id of the model which was deployed.
//   - options - EmbeddingsClientCreateOptions contains the optional parameters for the EmbeddingsClient.Create method.
func (client *EmbeddingsClient) Create(ctx context.Context, deploymentID string, body Paths13PiqocDeploymentsDeploymentIDEmbeddingsPostRequestbodyContentApplicationJSONSchema, options *EmbeddingsClientCreateOptions) (EmbeddingsClientCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, deploymentID, body, options)
	if err != nil {
		return EmbeddingsClientCreateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EmbeddingsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return EmbeddingsClientCreateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *EmbeddingsClient) createCreateRequest(ctx context.Context, deploymentID string, body Paths13PiqocDeploymentsDeploymentIDEmbeddingsPostRequestbodyContentApplicationJSONSchema, options *EmbeddingsClientCreateOptions) (*policy.Request, error) {
	host := "https://{endpoint}/openai"
	host = strings.ReplaceAll(host, "{endpoint}", endpoint)
	urlPath := "/deployments/{deployment-id}/embeddings"
	if deploymentID == "" {
		return nil, errors.New("parameter deploymentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deployment-id}", url.PathEscape(deploymentID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// createHandleResponse handles the Create response.
func (client *EmbeddingsClient) createHandleResponse(resp *http.Response) (EmbeddingsClientCreateResponse, error) {
	result := EmbeddingsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Paths15Cw454DeploymentsDeploymentIDEmbeddingsPostResponses200ContentApplicationJSONSchema); err != nil {
		return EmbeddingsClientCreateResponse{}, err
	}
	return result, nil
}


//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azopenai

// ChatCompletionsClientCreateResponse contains the response from method ChatCompletionsClient.Create.
type ChatCompletionsClientCreateResponse struct {
	Paths1H0F83DeploymentsDeploymentIDChatCompletionsPostResponses200ContentApplicationJSONSchema
}

// CompletionsClientCreateResponse contains the response from method CompletionsClient.Create.
type CompletionsClientCreateResponse struct {
	PathsMaorw9DeploymentsDeploymentIDCompletionsPostResponses200ContentApplicationJSONSchema
	// ApimRequestID contains the information returned from the apim-request-id header response.
	ApimRequestID *string
}

// EmbeddingsClientCreateResponse contains the response from method EmbeddingsClient.Create.
type EmbeddingsClientCreateResponse struct {
	Paths15Cw454DeploymentsDeploymentIDEmbeddingsPostResponses200ContentApplicationJSONSchema
}

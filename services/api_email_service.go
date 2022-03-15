/*
 * OpenAPI proxy service
 *
 * This is the OpenAPI definition of the proxy service.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package services

import (
	"chaingate/proxy-service/proxyApi"
	"context"
	"errors"
	"net/http"
)

// EmailApiService is a service that implements the logic for the EmailApiServicer
// This service should implement the business logic for every endpoint for the EmailApi API.
// Include any external packages or services that will be required by this service.
type EmailApiService struct {
}

// NewEmailApiService creates a default api service
func NewEmailApiService() proxyApi.EmailApiServicer {
	return &EmailApiService{}
}

// SendEmail - send email
func (s *EmailApiService) SendEmail(ctx context.Context, email proxyApi.Email) (proxyApi.ImplResponse, error) {
	// TODO - update SendEmail with the required logic for this service method.
	// Add api_email_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	//return Response(200, nil),nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return proxyApi.Response(http.StatusNotImplemented, nil), errors.New("SendEmail method not implemented")
}

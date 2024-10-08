/*
Apache JAMES Web Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// CassandraExtraAPIService CassandraExtraAPI service
type CassandraExtraAPIService service

type ApiPerformActionOnCassandraMappingsRequest struct {
	ctx        context.Context
	ApiService *CassandraExtraAPIService
	action     *string
}

// The action to perform on mappings_sources table
func (r ApiPerformActionOnCassandraMappingsRequest) Action(action string) ApiPerformActionOnCassandraMappingsRequest {
	r.action = &action
	return r
}

func (r ApiPerformActionOnCassandraMappingsRequest) Execute() (*PerformActionOnCassandraMappings201Response, *http.Response, error) {
	return r.ApiService.PerformActionOnCassandraMappingsExecute(r)
}

/*
PerformActionOnCassandraMappings Perform an action on mappings_sources table

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiPerformActionOnCassandraMappingsRequest
*/
func (a *CassandraExtraAPIService) PerformActionOnCassandraMappings(ctx context.Context) ApiPerformActionOnCassandraMappingsRequest {
	return ApiPerformActionOnCassandraMappingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PerformActionOnCassandraMappings201Response
func (a *CassandraExtraAPIService) PerformActionOnCassandraMappingsExecute(r ApiPerformActionOnCassandraMappingsRequest) (*PerformActionOnCassandraMappings201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PerformActionOnCassandraMappings201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CassandraExtraAPIService.PerformActionOnCassandraMappings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/cassandra/mappings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.action == nil {
		return localVarReturnValue, nil, reportError("action is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

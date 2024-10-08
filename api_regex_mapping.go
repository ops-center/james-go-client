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
	"strings"
)

// RegexMappingAPIService RegexMappingAPI service
type RegexMappingAPIService service

type ApiAddRegexMappingRequest struct {
	ctx           context.Context
	ApiService    *RegexMappingAPIService
	mappingSource string
	regex         string
}

func (r ApiAddRegexMappingRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddRegexMappingExecute(r)
}

/*
AddRegexMapping Add a regex mapping

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param mappingSource Mapping source (recipient address)
	@param regex Java Regular Expression (regex)
	@return ApiAddRegexMappingRequest
*/
func (a *RegexMappingAPIService) AddRegexMapping(ctx context.Context, mappingSource string, regex string) ApiAddRegexMappingRequest {
	return ApiAddRegexMappingRequest{
		ApiService:    a,
		ctx:           ctx,
		mappingSource: mappingSource,
		regex:         regex,
	}
}

// Execute executes the request
func (a *RegexMappingAPIService) AddRegexMappingExecute(r ApiAddRegexMappingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegexMappingAPIService.AddRegexMapping")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mappings/regex/{mappingSource}/targets/{regex}"
	localVarPath = strings.Replace(localVarPath, "{"+"mappingSource"+"}", url.PathEscape(parameterValueToString(r.mappingSource, "mappingSource")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"regex"+"}", url.PathEscape(parameterValueToString(r.regex, "regex")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRemoveRegexMappingRequest struct {
	ctx           context.Context
	ApiService    *RegexMappingAPIService
	mappingSource string
	regex         string
}

func (r ApiRemoveRegexMappingRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveRegexMappingExecute(r)
}

/*
RemoveRegexMapping Remove a regex mapping

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param mappingSource Mapping source (recipient address)
	@param regex Java Regular Expression (regex)
	@return ApiRemoveRegexMappingRequest
*/
func (a *RegexMappingAPIService) RemoveRegexMapping(ctx context.Context, mappingSource string, regex string) ApiRemoveRegexMappingRequest {
	return ApiRemoveRegexMappingRequest{
		ApiService:    a,
		ctx:           ctx,
		mappingSource: mappingSource,
		regex:         regex,
	}
}

// Execute executes the request
func (a *RegexMappingAPIService) RemoveRegexMappingExecute(r ApiRemoveRegexMappingRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RegexMappingAPIService.RemoveRegexMapping")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mappings/regex/{mappingSource}/targets/{regex}"
	localVarPath = strings.Replace(localVarPath, "{"+"mappingSource"+"}", url.PathEscape(parameterValueToString(r.mappingSource, "mappingSource")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"regex"+"}", url.PathEscape(parameterValueToString(r.regex, "regex")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

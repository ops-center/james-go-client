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

// MessagesAPIService MessagesAPI service
type MessagesAPIService service

type ApiReindexEmailRequest struct {
	ctx        context.Context
	ApiService *MessagesAPIService
	messageId  string
	task       *string
}

// Task type for reindexing
func (r ApiReindexEmailRequest) Task(task string) ApiReindexEmailRequest {
	r.task = &task
	return r
}

func (r ApiReindexEmailRequest) Execute() (*PerformActionsOnMailboxes201Response, *http.Response, error) {
	return r.ApiService.ReindexEmailExecute(r)
}

/*
ReindexEmail Reindex a single mail by messageId

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param messageId Implementation-dependent valid messageId
	@return ApiReindexEmailRequest
*/
func (a *MessagesAPIService) ReindexEmail(ctx context.Context, messageId string) ApiReindexEmailRequest {
	return ApiReindexEmailRequest{
		ApiService: a,
		ctx:        ctx,
		messageId:  messageId,
	}
}

// Execute executes the request
//
//	@return PerformActionsOnMailboxes201Response
func (a *MessagesAPIService) ReindexEmailExecute(r ApiReindexEmailRequest) (*PerformActionsOnMailboxes201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PerformActionsOnMailboxes201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesAPIService.ReindexEmail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages/{messageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"messageId"+"}", url.PathEscape(parameterValueToString(r.messageId, "messageId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.task == nil {
		return localVarReturnValue, nil, reportError("task is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "task", r.task, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ScheduleTask400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

type ApiScheduleTaskRequest struct {
	ctx               context.Context
	ApiService        *MessagesAPIService
	task              *string
	messagesPerSecond *int32
	mode              *string
}

// Task type for fixing message inconsistencies
func (r ApiScheduleTaskRequest) Task(task string) ApiScheduleTaskRequest {
	r.task = &task
	return r
}

// Rate of messages to be processed per second
func (r ApiScheduleTaskRequest) MessagesPerSecond(messagesPerSecond int32) ApiScheduleTaskRequest {
	r.messagesPerSecond = &messagesPerSecond
	return r
}

// Reindexing mode used
func (r ApiScheduleTaskRequest) Mode(mode string) ApiScheduleTaskRequest {
	r.mode = &mode
	return r
}

func (r ApiScheduleTaskRequest) Execute() (*PerformActionsOnMailboxes201Response, *http.Response, error) {
	return r.ApiService.ScheduleTaskExecute(r)
}

/*
ScheduleTask Schedule a task for fixing message inconsistencies

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiScheduleTaskRequest
*/
func (a *MessagesAPIService) ScheduleTask(ctx context.Context) ApiScheduleTaskRequest {
	return ApiScheduleTaskRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return PerformActionsOnMailboxes201Response
func (a *MessagesAPIService) ScheduleTaskExecute(r ApiScheduleTaskRequest) (*PerformActionsOnMailboxes201Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PerformActionsOnMailboxes201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MessagesAPIService.ScheduleTask")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/messages"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.task == nil {
		return localVarReturnValue, nil, reportError("task is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "task", r.task, "")
	if r.messagesPerSecond != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "messagesPerSecond", r.messagesPerSecond, "")
	}
	if r.mode != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mode", r.mode, "")
	}
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ScheduleTask400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
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

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
	"time"
)

// TaskAPIService TaskAPI service
type TaskAPIService service

type ApiAwaitTaskCompletionRequest struct {
	ctx        context.Context
	ApiService *TaskAPIService
	taskId     string
	timeout    *string
}

// The timeout duration in the format Nunit (e.g., 30s, 5m, 7d, 1y)
func (r ApiAwaitTaskCompletionRequest) Timeout(timeout string) ApiAwaitTaskCompletionRequest {
	r.timeout = &timeout
	return r
}

func (r ApiAwaitTaskCompletionRequest) Execute() (*ExecutionReport, *http.Response, error) {
	return r.ApiService.AwaitTaskCompletionExecute(r)
}

/*
AwaitTaskCompletion Await the completion of a task

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taskId The ID of the task
	@return ApiAwaitTaskCompletionRequest
*/
func (a *TaskAPIService) AwaitTaskCompletion(ctx context.Context, taskId string) ApiAwaitTaskCompletionRequest {
	return ApiAwaitTaskCompletionRequest{
		ApiService: a,
		ctx:        ctx,
		taskId:     taskId,
	}
}

// Execute executes the request
//
//	@return ExecutionReport
func (a *TaskAPIService) AwaitTaskCompletionExecute(r ApiAwaitTaskCompletionRequest) (*ExecutionReport, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExecutionReport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskAPIService.AwaitTaskCompletion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tasks/{taskId}/await"
	localVarPath = strings.Replace(localVarPath, "{"+"taskId"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.timeout != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timeout", r.timeout, "")
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

type ApiCancelTaskRequest struct {
	ctx        context.Context
	ApiService *TaskAPIService
	taskId     string
}

func (r ApiCancelTaskRequest) Execute() (*http.Response, error) {
	return r.ApiService.CancelTaskExecute(r)
}

/*
CancelTask Cancel a task

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taskId The ID of the task
	@return ApiCancelTaskRequest
*/
func (a *TaskAPIService) CancelTask(ctx context.Context, taskId string) ApiCancelTaskRequest {
	return ApiCancelTaskRequest{
		ApiService: a,
		ctx:        ctx,
		taskId:     taskId,
	}
}

// Execute executes the request
func (a *TaskAPIService) CancelTaskExecute(r ApiCancelTaskRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskAPIService.CancelTask")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tasks/{taskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taskId"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

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

type ApiGetTaskRequest struct {
	ctx        context.Context
	ApiService *TaskAPIService
	taskId     string
}

func (r ApiGetTaskRequest) Execute() (*ExecutionReport, *http.Response, error) {
	return r.ApiService.GetTaskExecute(r)
}

/*
GetTask Get a task's details

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param taskId The ID of the task
	@return ApiGetTaskRequest
*/
func (a *TaskAPIService) GetTask(ctx context.Context, taskId string) ApiGetTaskRequest {
	return ApiGetTaskRequest{
		ApiService: a,
		ctx:        ctx,
		taskId:     taskId,
	}
}

// Execute executes the request
//
//	@return ExecutionReport
func (a *TaskAPIService) GetTaskExecute(r ApiGetTaskRequest) (*ExecutionReport, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ExecutionReport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskAPIService.GetTask")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tasks/{taskId}"
	localVarPath = strings.Replace(localVarPath, "{"+"taskId"+"}", url.PathEscape(parameterValueToString(r.taskId, "taskId")), -1)

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

type ApiListTasksRequest struct {
	ctx             context.Context
	ApiService      *TaskAPIService
	status          *string
	type_           *string
	submittedBefore *time.Time
	submittedAfter  *time.Time
	startedBefore   *time.Time
	startedAfter    *time.Time
	completedBefore *time.Time
	completedAfter  *time.Time
	failedBefore    *time.Time
	failedAfter     *time.Time
	offset          *int32
	limit           *int32
}

// The status of the tasks (waiting, inProgress, canceledRequested, completed, canceled, failed)
func (r ApiListTasksRequest) Status(status string) ApiListTasksRequest {
	r.status = &status
	return r
}

// The type of the tasks
func (r ApiListTasksRequest) Type_(type_ string) ApiListTasksRequest {
	r.type_ = &type_
	return r
}

// Return tasks submitted before this date
func (r ApiListTasksRequest) SubmittedBefore(submittedBefore time.Time) ApiListTasksRequest {
	r.submittedBefore = &submittedBefore
	return r
}

// Return tasks submitted after this date
func (r ApiListTasksRequest) SubmittedAfter(submittedAfter time.Time) ApiListTasksRequest {
	r.submittedAfter = &submittedAfter
	return r
}

// Return tasks started before this date
func (r ApiListTasksRequest) StartedBefore(startedBefore time.Time) ApiListTasksRequest {
	r.startedBefore = &startedBefore
	return r
}

// Return tasks started after this date
func (r ApiListTasksRequest) StartedAfter(startedAfter time.Time) ApiListTasksRequest {
	r.startedAfter = &startedAfter
	return r
}

// Return tasks completed before this date
func (r ApiListTasksRequest) CompletedBefore(completedBefore time.Time) ApiListTasksRequest {
	r.completedBefore = &completedBefore
	return r
}

// Return tasks completed after this date
func (r ApiListTasksRequest) CompletedAfter(completedAfter time.Time) ApiListTasksRequest {
	r.completedAfter = &completedAfter
	return r
}

// Return tasks failed before this date
func (r ApiListTasksRequest) FailedBefore(failedBefore time.Time) ApiListTasksRequest {
	r.failedBefore = &failedBefore
	return r
}

// Return tasks failed after this date
func (r ApiListTasksRequest) FailedAfter(failedAfter time.Time) ApiListTasksRequest {
	r.failedAfter = &failedAfter
	return r
}

// Number of tasks to skip in the response
func (r ApiListTasksRequest) Offset(offset int32) ApiListTasksRequest {
	r.offset = &offset
	return r
}

// Maximum number of tasks to return in one call
func (r ApiListTasksRequest) Limit(limit int32) ApiListTasksRequest {
	r.limit = &limit
	return r
}

func (r ApiListTasksRequest) Execute() ([]ExecutionReport, *http.Response, error) {
	return r.ApiService.ListTasksExecute(r)
}

/*
ListTasks List tasks

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListTasksRequest
*/
func (a *TaskAPIService) ListTasks(ctx context.Context) ApiListTasksRequest {
	return ApiListTasksRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []ExecutionReport
func (a *TaskAPIService) ListTasksExecute(r ApiListTasksRequest) ([]ExecutionReport, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []ExecutionReport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TaskAPIService.ListTasks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/tasks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.submittedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "submittedBefore", r.submittedBefore, "")
	}
	if r.submittedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "submittedAfter", r.submittedAfter, "")
	}
	if r.startedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startedBefore", r.startedBefore, "")
	}
	if r.startedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startedAfter", r.startedAfter, "")
	}
	if r.completedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "completedBefore", r.completedBefore, "")
	}
	if r.completedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "completedAfter", r.completedAfter, "")
	}
	if r.failedBefore != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "failedBefore", r.failedBefore, "")
	}
	if r.failedAfter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "failedAfter", r.failedAfter, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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

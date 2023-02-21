# {{classname}}

All URIs are relative to *https://gus.gokrazy.org/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Attempt**](UpdateApi.md#Attempt) | **Post** /attempt | note that an update will be attempted
[**Update**](UpdateApi.md#Update) | **Post** /update | check if an update is available

# **Attempt**
> AttemptResponse Attempt(ctx, optional)
note that an update will be attempted

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateApiAttemptOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateApiAttemptOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of AttemptRequest**](AttemptRequest.md)| JSON-encoded update request | 

### Return type

[**AttemptResponse**](attemptResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> UpdateResponse Update(ctx, optional)
check if an update is available

The gokrazy/selfupdate program calls this request periodically to see if an update is available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateApiUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdateApiUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of UpdateRequest**](UpdateRequest.md)| JSON-encoded update request | 

### Return type

[**UpdateResponse**](updateResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


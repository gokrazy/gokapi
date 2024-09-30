# \SuperviseApi

All URIs are relative to *http://gokrazy*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Index**](SuperviseApi.md#Index) | **Get** / | Index
[**Status**](SuperviseApi.md#Status) | **Get** /status | Status


# **Index**
> IndexResult Index(ctx, )
Index

TODO: description

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IndexResult**](IndexResult.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Status**
> Service Status(ctx, path)
Status

TODO: description

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **path** | **string**| path to the service | 

### Return type

[**Service**](Service.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


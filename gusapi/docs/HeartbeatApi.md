# {{classname}}

All URIs are relative to *https://gus.gokrazy.org/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Heartbeat**](HeartbeatApi.md#Heartbeat) | **Post** /heartbeat | device sends a heartbeat

# **Heartbeat**
> HeartbeatResponse Heartbeat(ctx, optional)
device sends a heartbeat

The heartbeat indicates that the device is active (for display in the GUS user interface) and contains: 1. the SBOM (Software Bill Of Materials) 1. human readable system information, like model or kernel version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***HeartbeatApiHeartbeatOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HeartbeatApiHeartbeatOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of HeartbeatRequest**](HeartbeatRequest.md)| JSON-encoded heartbeat request | 

### Return type

[**HeartbeatResponse**](heartbeatResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


# {{classname}}

All URIs are relative to *https://gus.gokrazy.org/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Ingest**](IngestApi.md#Ingest) | **Post** /ingest | ingest a gokrazy build into the GUS database
[**Push**](IngestApi.md#Push) | **Put** /push | push a gokrazy build into the local disk registry

# **Ingest**
> IngestResponse Ingest(ctx, optional)
ingest a gokrazy build into the GUS database

Ingesting a gokrazy build makes it available for update on the machines matching the specified `machine_id_pattern`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IngestApiIngestOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IngestApiIngestOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of IngestRequest**](IngestRequest.md)| JSON-encoded ingest request | 

### Return type

[**IngestResponse**](ingestResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Push**
> PushResponse Push(ctx, optional)
push a gokrazy build into the local disk registry

If the GUS server was started with a non-empty `--image_dir`, you can push gokrazy builds into the image directory. Example: ``` gok push --gaf disk.gaf --server http://localhost:8655 ```  This is the most convenient form of build registry (registry type `localdisk`), as no other service providers are involved — you directly upload builds to the GUS server, devices download the builds from the GUS server.  A /push request is typically followed by a /ingest request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***IngestApiPushOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IngestApiPushOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Object**](Object.md)| build as a GAF file (gokrazy archive format) | 

### Return type

[**PushResponse**](pushResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/octet-stream
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


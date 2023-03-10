# Go API client for swagger

OpenAPI for GUS (gokrazy update service)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.4.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *https://gus.gokrazy.org/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*HeartbeatApi* | [**Heartbeat**](docs/HeartbeatApi.md#heartbeat) | **Post** /heartbeat | device sends a heartbeat
*IngestApi* | [**Ingest**](docs/IngestApi.md#ingest) | **Post** /ingest | ingest a gokrazy build into the GUS database
*IngestApi* | [**Push**](docs/IngestApi.md#push) | **Put** /push | push a gokrazy build into the local disk registry
*UpdateApi* | [**Attempt**](docs/UpdateApi.md#attempt) | **Post** /attempt | note that an update will be attempted
*UpdateApi* | [**Update**](docs/UpdateApi.md#update) | **Post** /update | check if an update is available

## Documentation For Models

 - [AttemptRequest](docs/AttemptRequest.md)
 - [AttemptResponse](docs/AttemptResponse.md)
 - [HeartbeatRequest](docs/HeartbeatRequest.md)
 - [HeartbeatRequestHumanReadable](docs/HeartbeatRequestHumanReadable.md)
 - [HeartbeatRequestSbom](docs/HeartbeatRequestSbom.md)
 - [HeartbeatResponse](docs/HeartbeatResponse.md)
 - [IngestRequest](docs/IngestRequest.md)
 - [IngestResponse](docs/IngestResponse.md)
 - [PathHash](docs/PathHash.md)
 - [PushResponse](docs/PushResponse.md)
 - [UpdateRequest](docs/UpdateRequest.md)
 - [UpdateResponse](docs/UpdateResponse.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author

michael+gokrazy@stapelberg.ch

# IngestRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineIdPattern** | **string** | The regexp pattern (go regexp syntax) to match the machineIDs to which apply the provided version. The most base case is a plain machine_id. | [optional] [default to null]
**SbomHash** | **string** |  | [optional] [default to null]
**RegistryType** | **string** | the type of registry on which the build is stored. see download_link | [optional] [default to null]
**DownloadLink** | **string** | relative (localdisk registry) or absolute download link with which gokrazy devices can download the build | [optional] [default to null]
**IngestionPolicy** | **string** | the policy with which the version (sbom_hash) is set for the matching machine_id_pattern. At the moment these are the two possible values: - \&quot;\&quot; (empty string): this is the default value for this property. And it means the provided version (sbom_hash) is temporarily set for the matching machine_id_pattern.   when a new image is ingested for this matching pattern, the new image will become the new desired version for this matching pattern. - \&quot;permanent\&quot;: the provided version (sbom_hash) will stay the desired one for the matching machine_id_pattern until it gets explicitly unset/changed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


/*
 * GUS (gokrazy update service)
 *
 * OpenAPI for GUS (gokrazy update service)
 *
 * API version: 1.4.0
 * Contact: michael+gokrazy@stapelberg.ch
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package gusapi

type UpdateResponse struct {
	SbomHash string `json:"sbom_hash,omitempty"`
	// the type of registry on which the build is stored. see download_link
	RegistryType string `json:"registry_type,omitempty"`
	// relative (localdisk registry) or absolute download link with which gokrazy devices can download the build
	DownloadLink string `json:"download_link,omitempty"`
}

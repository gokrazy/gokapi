/*
 * gokrazy on-device API
 *
 * OpenAPI for the gokrazy on-device API
 *
 * API version: 1.4.0
 * Contact: michael+gokrazy@stapelberg.ch
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ondeviceapi

// index of what is supervised
type IndexResult struct {
	// The Linux kernel version
	Kernel string `json:"Kernel,omitempty"`
	// Model name of the machine on which gokrazy is running
	Model string `json:"Model,omitempty"`
	// Build timestamp of the running gokrazy instance
	BuildTimestamp string `json:"BuildTimestamp,omitempty"`
	// running services
	Services []Service `json:"Services,omitempty"`
}

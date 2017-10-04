/* 
 * Ringcentral API
 *
 * RingCentral Connect Platform API
 *
 * OpenAPI spec version: v1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package ringcentral

type DeviceInfoExtensionInfo struct {

	// Internal identifier of an extension
	Id string `json:"id,omitempty"`

	// Canonical URI of an extension
	Uri string `json:"uri,omitempty"`

	// Number of department extension
	ExtensionNumber string `json:"extensionNumber,omitempty"`

	// For Partner Applications Internal identifier of an extension created by partner. The RingCentral supports the mapping of accounts and stores the corresponding account ID/extension ID for each partner ID of a client application. In request URIs partner IDs are accepted instead of regular RingCentral native IDs as path parameters using pid = XXX clause. Though in response URIs contain the corresponding account IDs and extension IDs. In all request and response bodies these values are reflected via partnerId attributes of account and extension
	PartnerId string `json:"partnerId,omitempty"`
}
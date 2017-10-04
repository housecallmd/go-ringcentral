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

type ShippingMethod struct {

	// Method identifier. The default value is \"1\" (Ground)
	Id string `json:"id,omitempty"`

	// Method name, corresponding to the identifier
	Name string `json:"name,omitempty"`
}
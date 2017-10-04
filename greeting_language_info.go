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

type GreetingLanguageInfo struct {

	// Internal identifier of a greeting language
	Id string `json:"id,omitempty"`

	// Localization code of a greeting language
	LocaleCode string `json:"localeCode,omitempty"`

	// Official name of a greeting language
	Name string `json:"name,omitempty"`
}
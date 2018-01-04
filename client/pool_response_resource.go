/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ringcentral

import (
	"os"
)

type PoolResponseResource struct {

	RendererId string `json:"rendererId,omitempty"`

	MessageId string `json:"messageId,omitempty"`

	ExtensionId string `json:"extensionId,omitempty"`

	ExtensionNumber string `json:"extensionNumber,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Street string `json:"street,omitempty"`

	City string `json:"city,omitempty"`

	State string `json:"state,omitempty"`

	Zip string `json:"zip,omitempty"`

	Country string `json:"country,omitempty"`

	CountryId string `json:"countryId,omitempty"`

	CompanyName string `json:"companyName,omitempty"`

	FaxNumber string `json:"faxNumber,omitempty"`

	ContactPhone string `json:"contactPhone,omitempty"`

	Email string `json:"email,omitempty"`

	LanguageCode string `json:"languageCode,omitempty"`

	LanguageId string `json:"languageId,omitempty"`

	FooterType float32 `json:"footerType,omitempty"`

	CoverIndex float32 `json:"coverIndex,omitempty"`

	CoverPageText string `json:"coverPageText,omitempty"`

	SourceFiles []*os.File `json:"sourceFiles,omitempty"`

	MessageServers []string `json:"messageServers,omitempty"`

	TouchInterval int64 `json:"touchInterval,omitempty"`
}

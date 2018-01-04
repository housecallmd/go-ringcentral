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
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type UserContactsApiService service


/* UserContactsApiService Get Favorite Contacts
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 @return */
func (a *UserContactsApiService) ContactFavorite(ctx context.Context, accountId string, extensionId string) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/favorite"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* UserContactsApiService Create Contact
 &lt;p style&#x3D;&#39;font-style:italic;&#39;&gt;&lt;/p&gt;&lt;p&gt;&lt;/p&gt;&lt;h4&gt;Required Permissions&lt;/h4&gt;&lt;table class&#x3D;&#39;fullwidth&#39;&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Permission&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;/thead&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td class&#x3D;&#39;code&#39;&gt;Contacts&lt;/td&gt;&lt;td&gt;Creating, viewing, editing and deleting user personal contacts&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td class&#x3D;&#39;code&#39;&gt;ReadContacts&lt;/td&gt;&lt;td&gt;Viewing user personal contacts&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;h4&gt;Usage Plan Group&lt;/h4&gt;&lt;p&gt;Heavy&lt;/p&gt;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (PersonalContactResource) 
 @return PersonalContactResource*/
func (a *UserContactsApiService) CreateContact(ctx context.Context, accountId string, extensionId string, localVarOptionals map[string]interface{}) (PersonalContactResource,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PersonalContactResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(PersonalContactResource); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* UserContactsApiService Update Favorite Contacts
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (FavoriteCollection) 
 @return */
func (a *UserContactsApiService) CreateContacts(ctx context.Context, accountId string, extensionId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/favorite"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(FavoriteCollection); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* UserContactsApiService Delete Contact(s) by ID
 &lt;p style&#x3D;&#39;font-style:italic;&#39;&gt;&lt;/p&gt;&lt;p&gt;&lt;/p&gt;&lt;h4&gt;Required Permissions&lt;/h4&gt;&lt;table class&#x3D;&#39;fullwidth&#39;&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Permission&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;/thead&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td class&#x3D;&#39;code&#39;&gt;Contacts&lt;/td&gt;&lt;td&gt;Creating, viewing, editing and deleting user personal contacts&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td class&#x3D;&#39;code&#39;&gt;ReadContacts&lt;/td&gt;&lt;td&gt;Viewing user personal contacts&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;h4&gt;Usage Plan Group&lt;/h4&gt;&lt;p&gt;Heavy&lt;/p&gt;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 @param contactId Internal identifier of a contact record in the RingCentral database
 @return */
func (a *UserContactsApiService) DeleteContact(ctx context.Context, accountId string, extensionId string, contactId int32) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contactId"+"}", fmt.Sprintf("%v", contactId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* UserContactsApiService Get Contacts
 &lt;p style&#x3D;&#39;font-style:italic;&#39;&gt;&lt;/p&gt;&lt;p&gt;&lt;/p&gt;&lt;h4&gt;Required Permissions&lt;/h4&gt;&lt;table class&#x3D;&#39;fullwidth&#39;&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Permission&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;/thead&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td class&#x3D;&#39;code&#39;&gt;ReadContacts&lt;/td&gt;&lt;td&gt;Viewing user personal contacts&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;h4&gt;Usage Plan Group&lt;/h4&gt;&lt;p&gt;Heavy&lt;/p&gt;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "startsWith" (string) If specified, only contacts whose First name or Last name start with the mentioned substring are returned. Case-insensitive
     @param "sortBy" ([]string) Sorts results by the specified property. The default is &#39;First Name&#39;
     @param "page" (int32) Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39;
     @param "perPage" (int32) Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default
     @param "phoneNumber" ([]string) 
 @return ContactList*/
func (a *UserContactsApiService) ListContacts(ctx context.Context, accountId string, extensionId string, localVarOptionals map[string]interface{}) (ContactList,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ContactList
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["startsWith"], "string", "startsWith"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["page"], "int32", "page"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["perPage"], "int32", "perPage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["startsWith"].(string); localVarOk {
		localVarQueryParams.Add("startsWith", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["sortBy"].([]string); localVarOk {
		localVarQueryParams.Add("sortBy", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["page"].(int32); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["perPage"].(int32); localVarOk {
		localVarQueryParams.Add("perPage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["phoneNumber"].([]string); localVarOk {
		localVarQueryParams.Add("phoneNumber", parameterToString(localVarTempParam, "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* UserContactsApiService Get Contact(s) by ID
 &lt;p style&#x3D;&#39;font-style:italic;&#39;&gt;&lt;/p&gt;&lt;p&gt;&lt;/p&gt;&lt;h4&gt;Required Permissions&lt;/h4&gt;&lt;table class&#x3D;&#39;fullwidth&#39;&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Permission&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;/thead&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td class&#x3D;&#39;code&#39;&gt;ReadContacts&lt;/td&gt;&lt;td&gt;Viewing user personal contacts&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;h4&gt;Usage Plan Group&lt;/h4&gt;&lt;p&gt;Heavy&lt;/p&gt;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 @param contactId Internal identifier of a contact record in the RingCentral database
 @return PersonalContactResource*/
func (a *UserContactsApiService) LoadContact(ctx context.Context, accountId string, extensionId string, contactId int32) (PersonalContactResource,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PersonalContactResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contactId"+"}", fmt.Sprintf("%v", contactId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* UserContactsApiService Address Book Synchronization
 &lt;p style&#x3D;&#39;font-style:italic;&#39;&gt;&lt;/p&gt;&lt;p&gt;&lt;/p&gt;&lt;h4&gt;Required Permissions&lt;/h4&gt;&lt;table class&#x3D;&#39;fullwidth&#39;&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Permission&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;/thead&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td class&#x3D;&#39;code&#39;&gt;ReadContacts&lt;/td&gt;&lt;td&gt;Viewing user personal contacts&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;h4&gt;Usage Plan Group&lt;/h4&gt;&lt;p&gt;Heavy&lt;/p&gt;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "syncType" ([]string) Type of synchronization. The default value is &#39;FSync&#39;
     @param "syncToken" (string) Value of syncToken property of the last sync request response
     @param "perPage" (int32) Number of records per page to be returned. The max number of records is 250, which is also the default. For FSync ??? if the number of records exceeds the parameter value (either specified or default), all of the pages can be retrieved in several requests. For ISync ??? if the number of records exceeds the page size, the number of incoming changes to this number is limited
     @param "pageId" (int32) Internal identifier of a page. It can be obtained from the &#39;nextPageId&#39; parameter passed in response body
 @return */
func (a *UserContactsApiService) SyncAddressBook(ctx context.Context, accountId string, extensionId string, localVarOptionals map[string]interface{}) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book-sync"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["syncToken"], "string", "syncToken"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["perPage"], "int32", "perPage"); err != nil {
		return nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageId"], "int32", "pageId"); err != nil {
		return nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["syncType"].([]string); localVarOk {
		localVarQueryParams.Add("syncType", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["syncToken"].(string); localVarOk {
		localVarQueryParams.Add("syncToken", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["perPage"].(int32); localVarOk {
		localVarQueryParams.Add("perPage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageId"].(int32); localVarOk {
		localVarQueryParams.Add("pageId", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* UserContactsApiService Update Contact(s) by ID
 &lt;p style&#x3D;&#39;font-style:italic;&#39;&gt;&lt;/p&gt;&lt;p&gt;&lt;/p&gt;&lt;h4&gt;Required Permissions&lt;/h4&gt;&lt;table class&#x3D;&#39;fullwidth&#39;&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Permission&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;/thead&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td class&#x3D;&#39;code&#39;&gt;Contacts&lt;/td&gt;&lt;td&gt;Creating, viewing, editing and deleting user personal contacts&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td class&#x3D;&#39;code&#39;&gt;ReadContacts&lt;/td&gt;&lt;td&gt;Viewing user personal contacts&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;h4&gt;Usage Plan Group&lt;/h4&gt;&lt;p&gt;Heavy&lt;/p&gt;
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 @param contactId Internal identifier of a contact record in the RingCentral database
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "body" (PersonalContactResource) 
 @return PersonalContactResource*/
func (a *UserContactsApiService) UpdateContact(ctx context.Context, accountId string, extensionId string, contactId int32, localVarOptionals map[string]interface{}) (PersonalContactResource,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PersonalContactResource
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contactId"+"}", fmt.Sprintf("%v", contactId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["body"].(PersonalContactResource); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}


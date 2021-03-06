# CompanyAnsweringRuleInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal identifier of an answering rule | [optional] [default to null]
**Uri** | **string** | Canonical URI of an answering rule | [optional] [default to null]
**Enabled** | **bool** | Specifies if the rule is active or inactive. The default value is &#39;True&#39; | [optional] [default to null]
**Type_** | **string** | Type of an answering rule, the default value is &#39;Custom&#39; &#x3D; [&#39;BusinessHours&#39;, &#39;AfterHours&#39;, &#39;Custom&#39;] | [optional] [default to null]
**Name** | **string** | Name of an answering rule specified by user. Max number of symbols is 30. The default value is &#39;My Rule N&#39; where &#39;N&#39; is the first free number | [optional] [default to null]
**Callers** | [**[]CompanyAnsweringRuleCallersInfoRequest**](CompanyAnsweringRuleCallersInfoRequest.md) | Answering rule will be applied when calls are received from the specified caller(s) | [optional] [default to null]
**CalledNumbers** | [**[]CompanyAnsweringRuleCalledNumberInfoRequest**](CompanyAnsweringRuleCalledNumberInfoRequest.md) | Answering rule will be applied when calling the specified number(s) | [optional] [default to null]
**Schedule** | [***CompanyAnsweringRuleScheduleInfo**](CompanyAnsweringRuleScheduleInfo.md) | Schedule when an answering rule should be applied , | [optional] [default to null]
**CallHandlingAction** | **string** | Specifies how incoming calls are forwarded. The default value is &#39;Operator&#39; &#39;Operator&#39; - play company greeting and forward to operator extension &#39;Disconnect&#39; - play company greeting and disconnect &#39;Bypass&#39; - bypass greeting to go to selected extension &#x3D; [&#39;Operator&#39;, &#39;Disconnect&#39;, &#39;Bypass&#39;] | [optional] [default to null]
**Extension** | [***CompanyAnsweringRuleCallersInfoRequest**](CompanyAnsweringRuleCallersInfoRequest.md) | Extension to which the call is forwarded in &#39;Bypass&#39; mode | [optional] [default to null]
**Greetings** | [**[]GreetingInfo**](GreetingInfo.md) | Greetings applied for an answering rule; only predefined greetings can be applied, see Dictionary Greeting List | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



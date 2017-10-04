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

import (
	"time"
)

type PresenceLineNotification struct {

	// Universally unique identifier of a notification
	Uuid string `json:"uuid,omitempty"`

	// Event filter URI
	Event string `json:"event,omitempty"`

	// Internal identifier of a subscription
	SubscriptionId string `json:"subscriptionId,omitempty"`

	// The datetime of sending a notification in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z
	Timestamp time.Time `json:"timestamp,omitempty"`

	// Notification payload body
	Body PresenceLineEvent `json:"body,omitempty"`
}
/*
 * messagemediamessages_lib
 *
 *
 */

package models_pkg


import uuid "github.com/satori/go.uuid"

/*
 * Structure for the custom type CheckDeliveryReportsResponse
 */
type CheckDeliveryReportsResponse struct {
    DeliveryReports  []interface{}   `json:"delivery_reports" form:"delivery_reports"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ConfirmRepliesAsReceivedRequest8
 */
type ConfirmRepliesAsReceivedRequest8 struct {
    ReplyIds        []uuid.UUID     `json:"reply_ids" form:"reply_ids"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ConfirmRepliesAsReceivedRequest
 */
type ConfirmRepliesAsReceivedRequest struct {
    ReplyIds        []string        `json:"reply_ids" form:"reply_ids"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type CheckRepliesResponse
 */
type CheckRepliesResponse struct {
    Replies         []interface{}   `json:"replies" form:"replies"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type CancelScheduledMessageRequest
 */
type CancelScheduledMessageRequest struct {
    Status          string          `json:"status" form:"status"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type SendMessagesResponse
 */
type SendMessagesResponse struct {
    Messages        []interface{}   `json:"messages" form:"messages"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type SendMessagesRequest
 */
type SendMessagesRequest struct {
    Messages        []interface{}   `json:"messages" form:"messages"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ConfirmDeliveryReportsAsReceivedRequest
 */
type ConfirmDeliveryReportsAsReceivedRequest struct {
    DeliveryReportIds   []string        `json:"delivery_report_ids" form:"delivery_report_ids"` //TODO: Write general description for this field
}

/*
 * Structure for the custom type ConfirmDeliveryReportsAsReceivedRequest11
 */
type ConfirmDeliveryReportsAsReceivedRequest11 struct {
    DeliveryReportIds   []uuid.UUID     `json:"delivery_report_ids" form:"delivery_report_ids"` //TODO: Write general description for this field
}

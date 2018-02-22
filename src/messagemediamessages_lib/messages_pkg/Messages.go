/*
 * messagemediamessages_lib
 *
 * This file was automatically generated for MessageMedia by APIMATIC v2.0 ( https://apimatic.io )
 */

package messages_pkg

import "messagemediamessages_lib/models_pkg"

/*
 * Interface for the MESSAGES_IMPL
 */
type MESSAGES interface {
    UpdateCancelScheduledMessage (string, *models_pkg.CancelScheduledMessageRequest) (interface{}, error)

    GetMessageStatus (string) (interface{}, error)

    CreateSendMessages (*models_pkg.SendMessagesRequest) (*models_pkg.SendMessagesResponse, error)
}

/*
 * Factory for the MESSAGES interaface returning MESSAGES_IMPL
 */
func NewMESSAGES() MESSAGES {
    return &MESSAGES_IMPL{}
}
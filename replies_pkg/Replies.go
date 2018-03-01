/*
 * messagemedia_messages_sdk
 *
 *
 */

package replies_pkg

import "github.com/messagemedia/messages-go-sdk/models_pkg"

/*
 * Interface for the REPLIES_IMPL
 */
type REPLIES interface {
    CreateConfirmRepliesAsReceived (*models_pkg.ConfirmRepliesAsReceivedRequest) (interface{}, error)

    GetCheckReplies () (*models_pkg.CheckRepliesResponse, error)
}

/*
 * Factory for the REPLIES interaface returning REPLIES_IMPL
 */
func NewREPLIES() REPLIES {
    return &REPLIES_IMPL{}
}

/*
 * messagemediamessages_lib
 *
 *
 */

package replies_pkg

import "messagemediamessages_lib/models_pkg"

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
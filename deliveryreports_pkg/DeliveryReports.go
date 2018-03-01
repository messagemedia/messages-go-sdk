/*
 * messagemedia_messages_sdk
 *
 *
 */

package deliveryreports_pkg

import "github.com/messagemedia/messages-go-sdk/models_pkg"

/*
 * Interface for the DELIVERYREPORTS_IMPL
 */
type DELIVERYREPORTS interface {
    GetCheckDeliveryReports () (*models_pkg.CheckDeliveryReportsResponse, error)

    CreateConfirmDeliveryReportsAsReceived (*models_pkg.ConfirmDeliveryReportsAsReceivedRequest) (interface{}, error)
}

/*
 * Factory for the DELIVERYREPORTS interaface returning DELIVERYREPORTS_IMPL
 */
func NewDELIVERYREPORTS() DELIVERYREPORTS {
    return &DELIVERYREPORTS_IMPL{}
}

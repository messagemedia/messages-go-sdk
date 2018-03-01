/*
 * messagemediamessages_lib
 *
 * This file was automatically generated for MessageMedia by APIMATIC v2.0 ( https://apimatic.io )
 */

package deliveryreports_pkg

import "messagemediamessages_lib/models_pkg"

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
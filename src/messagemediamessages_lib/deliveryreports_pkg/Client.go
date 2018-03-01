/*
 * messagemediamessages_lib
 *
 *
 */
package deliveryreports_pkg


import(
	"encoding/json"
	"messagemediamessages_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"messagemediamessages_lib"
	"messagemediamessages_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type DELIVERYREPORTS_IMPL struct { }

/**
 * Check for any delivery reports that have been received.
 * Delivery reports are a notification of the change in status of a message as it is being processed.
 * Each request to the check delivery reports endpoint will return any delivery reports received that
 * have not yet been confirmed using the confirm delivery reports endpoint. A response from the check
 * delivery reports endpoint will have the following structure:
 * ```json
 * {
 *     "delivery_reports": [
 *         {
 *             "callback_url": "https://my.callback.url.com",
 *             "delivery_report_id": "01e1fa0a-6e27-4945-9cdb-18644b4de043",
 *             "source_number": "+61491570157",
 *             "date_received": "2017-05-20T06:30:37.642Z",
 *             "status": "enroute",
 *             "delay": 0,
 *             "submitted_date": "2017-05-20T06:30:37.639Z",
 *             "original_text": "My first message!",
 *             "message_id": "d781dcab-d9d8-4fb2-9e03-872f07ae94ba",
 *             "vendor_account_id": {
 *                 "vendor_id": "MessageMedia",
 *                 "account_id": "MyAccount"
 *             },
 *             "metadata": {
 *                 "key1": "value1",
 *                 "key2": "value2"
 *             }
 *         },
 *         {
 *             "callback_url": "https://my.callback.url.com",
 *             "delivery_report_id": "0edf9022-7ccc-43e6-acab-480e93e98c1b",
 *             "source_number": "+61491570158",
 *             "date_received": "2017-05-21T01:46:42.579Z",
 *             "status": "enroute",
 *             "delay": 0,
 *             "submitted_date": "2017-05-21T01:46:42.574Z",
 *             "original_text": "My second message!",
 *             "message_id": "fbb3b3f5-b702-4d8b-ab44-65b2ee39a281",
 *             "vendor_account_id": {
 *                 "vendor_id": "MessageMedia",
 *                 "account_id": "MyAccount"
 *             },
 *             "metadata": {
 *                 "key1": "value1",
 *                 "key2": "value2"
 *             }
 *         }
 *     ]
 * }
 * ```
 * Each delivery report will contain details about the message, including any metadata specified
 * and the new status of the message (as each delivery report indicates a change in status of a
 * message) and the timestamp at which the status changed. Every delivery report will have a
 * unique delivery report ID for use with the confirm delivery reports endpoint.
 * *Note: The source number and destination number properties in a delivery report are the inverse of
 * those specified in the message that the delivery report relates to. The source number of the
 * delivery report is the destination number of the original message.*
 * Subsequent requests to the check delivery reports endpoint will return the same delivery reports
 * and a maximum of 100 delivery reports will be returned in each request. Applications should use the
 * confirm delivery reports endpoint in the following pattern so that delivery reports that have been
 * processed are no longer returned in subsequent check delivery reports requests.
 * 1. Call check delivery reports endpoint
 * 2. Process each delivery report
 * 3. Confirm all processed delivery reports using the confirm delivery reports endpoint
 * *Note: It is recommended to use the Webhooks feature to receive reply messages rather than
 * polling the check delivery reports endpoint.*
 * @return	Returns the *models_pkg.CheckDeliveryReportsResponse response from the API call
 */
func (me *DELIVERYREPORTS_IMPL) GetCheckDeliveryReports () (*models_pkg.CheckDeliveryReportsResponse, error) {
        //the base uri for api requests
    _queryBuilder := messagemediamessages_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1/delivery_reports"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "messagemedia-messages-go-sdk-1.0.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, messagemediamessages_lib.Config.BasicAuthUserName, messagemediamessages_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.CheckDeliveryReportsResponse = &models_pkg.CheckDeliveryReportsResponse{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Mark a delivery report as confirmed so it is no longer return in check delivery reports requests.
 * The confirm delivery reports endpoint is intended to be used in conjunction with the check delivery
 * reports endpoint to allow for robust processing of delivery reports. Once one or more delivery
 * reports have been processed, they can then be confirmed using the confirm delivery reports endpoint so they
 * are no longer returned in subsequent check delivery reports requests.
 * The confirm delivery reports endpoint takes a list of delivery report IDs as follows:
 * ```json
 * {
 *     "delivery_report_ids": [
 *         "011dcead-6988-4ad6-a1c7-6b6c68ea628d",
 *         "3487b3fa-6586-4979-a233-2d1b095c7718",
 *         "ba28e94b-c83d-4759-98e7-ff9c7edb87a1"
 *     ]
 * }
 * ```
 * Up to 100 delivery reports can be confirmed in a single confirm delivery reports request.
 * @param    *models_pkg.ConfirmDeliveryReportsAsReceivedRequest        body     parameter: Required
 * @return	Returns the interface{} response from the API call
 */
func (me *DELIVERYREPORTS_IMPL) CreateConfirmDeliveryReportsAsReceived (
            body *models_pkg.ConfirmDeliveryReportsAsReceivedRequest) (interface{}, error) {
        //the base uri for api requests
    _queryBuilder := messagemediamessages_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/v1/delivery_reports/confirmed"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "messagemedia-messages-go-sdk-1.0.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
    }

    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, body, messagemediamessages_lib.Config.BasicAuthUserName, messagemediamessages_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code == 400) {
        err = apihelper_pkg.NewAPIError("", _response.Code, _response.RawBody)
    } else if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
            err = apihelper_pkg.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
        }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal interface{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

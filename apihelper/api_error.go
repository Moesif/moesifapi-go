/*
 * moesifapi
 *
 */

package apihelper

/*
 * APIError structure for custom error handling in API invocation
 */
type APIError struct {
	ResponseCode int    //The HTTP response code from the API request
	Reason       string //The reason for throwing exception
	Response     []byte
}

/*
 * Initialization constructor
 * @param   string  reason  The reason for throwing exception
 * @param   int     code    The HTTP response code from the API request
 * @return  new APIException object
 */
func NewAPIError(reason string, code int, res []byte) *APIError {
	ex := new(APIError)
	ex.ResponseCode = code
	ex.Reason = reason
	ex.Response = res
	return ex
}

/*
 * Implementing the Error method for the error interface
 */
func (e *APIError) Error() string {
	return e.Reason
}

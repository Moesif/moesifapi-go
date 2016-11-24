/*
 * moesifapi
 *
 */
package health

import (
	"encoding/json"
	"github.com/apimatic/unirest-go"
	"github.com/moesif/moesifapi-go"
	"github.com/moesif/moesifapi-go/apihelper"
	"github.com/moesif/moesifapi-go/models"
)

/*
 * Client structure as interface implementation
 */
type HEALTH_IMPL struct{}

/**
 * Health Probe
 * @return	Returns the *models.StatusModel response from the API call
 */
func (me *HEALTH_IMPL) GetHealthProbe() (*models.StatusModel, error) {
	//the base uri for api requests
	_queryBuilder := moesifapi.BASEURI

	//prepare query string for API call
	_queryBuilder = _queryBuilder + "/health/probe"

	//variable to hold errors
	var err error = nil
	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return nil, err
	}

	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"accept":                  "application/json",
		"X-Moesif-Application-Id": moesifapi.Config.ApplicationId,
	}

	//prepare API request
	_request := unirest.Get(_queryBuilder, headers)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request)
	if err != nil {
		//error in API invocation
		return nil, err
	}

	//error handling using HTTP status codes
	if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return nil, err
	}

	//returning the response
	var retVal *models.StatusModel = &models.StatusModel{}
	err = json.Unmarshal(_response.RawBody, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil
}

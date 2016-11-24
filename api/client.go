/*
 * moesifapi
 *
 */
package api

import (
	"github.com/apimatic/unirest-go"
	"github.com/moesif/moesifapi-go"
	"github.com/moesif/moesifapi-go/apihelper"
	"github.com/moesif/moesifapi-go/models"
)

/*
 * Client structure as interface implementation
 */
type API_IMPL struct{}

/**
 * Add Single API Event Call
 * @param    *models.EventModel        body     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *API_IMPL) CreateEvent(
	body *models.EventModel) error {
	//the base uri for api requests
	_queryBuilder := moesifapi.BASEURI

	//prepare query string for API call
	_queryBuilder = _queryBuilder + "/v1/events"

	//variable to hold errors
	var err error = nil
	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return err
	}

	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"content-type":            "application/json; charset=utf-8",
		"X-Moesif-Application-Id": moesifapi.Config.ApplicationId,
	}

	//prepare API request
	_request := unirest.Post(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request)
	if err != nil {
		//error in API invocation
		return err
	}

	//error handling using HTTP status codes
	if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return err
	}

	//returning the response
	return nil
}

/**
 * Add multiple API Events in a single batch (batch size must be less than 250kb)
 * @param    []*models.EventModel        body     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *API_IMPL) CreateEventsBatch(
	body []*models.EventModel) error {
	//the base uri for api requests
	_queryBuilder := moesifapi.BASEURI

	//prepare query string for API call
	_queryBuilder = _queryBuilder + "/v1/events/batch"

	//variable to hold errors
	var err error = nil
	//validate and preprocess url
	_queryBuilder, err = apihelper.CleanUrl(_queryBuilder)
	if err != nil {
		//error in url validation or cleaning
		return err
	}

	//prepare headers for the outgoing request
	headers := map[string]interface{}{
		"content-type":            "application/json; charset=utf-8",
		"X-Moesif-Application-Id": moesifapi.Config.ApplicationId,
	}

	//prepare API request
	_request := unirest.Post(_queryBuilder, headers, body)
	//and invoke the API call request to fetch the response
	_response, err := unirest.AsString(_request)
	if err != nil {
		//error in API invocation
		return err
	}

	//error handling using HTTP status codes
	if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
		err = apihelper.NewAPIError("HTTP Response Not OK", _response.Code, _response.RawBody)
	}
	if err != nil {
		//error detected in status code validation
		return err
	}

	//returning the response
	return nil
}

/*
 * moesifapi-go
 */
package moesifapi

import (
	"net/http"

	"github.com/moesif/moesifapi-go/models"
)

/*
 * Interface for the Client
 */
type API interface {

	/**
	 * Queue Single API Event Call to be created
	 * @param    *models.EventModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	QueueEvent(*models.EventModel) error

	/**
	 * Queue multiple API Events to be added
	 * @param    []*models.EventModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	QueueEvents([]*models.EventModel) error

	/**
	 * Queue Single User to be updated
	 * @param    *models.UserModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	QueueUser(*models.UserModel) error

	/**
	 * Queue multiple Users to be updated
	 * @param    []*models.UserModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	QueueUsers([]*models.UserModel) error

	/**
	 * Queue Single Company to be added
	 * @param    *models.CompanyModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	QueueCompany(*models.CompanyModel) error

	/**
	 * Queue multiple companies to be added
	 * @param    []*models.CompanyModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	QueueCompanies([]*models.CompanyModel) error

	/**
	 * Queue Single Subscription to be added
	 * @param    *models.SubscriptionModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	QueueSubscription(*models.SubscriptionModel) error

	/**
	 * Queue multiple Subscriptions to be added
	 * @param    []*models.SubscriptionModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	QueueSubscriptions([]*models.SubscriptionModel) error

	/**
	 * Add Single API Event Call
	 * @param    *models.EventModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	CreateEvent(*models.EventModel) (http.Header, error)

	/**
	 * Add multiple API Events in a single batch (batch size must be less than 250kb)
	 * @param    []*models.EventModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	CreateEventsBatch([]*models.EventModel) (http.Header, error)

	/**
	 * Update a Single User
	 * @param    *models.UserModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	UpdateUser(*models.UserModel) error

	/**
	 * Update multiple Users in a single batch (batch size must be less than 250kb)
	 * @param    []*models.UserModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	UpdateUsersBatch([]*models.UserModel) error

	/**
	 * Get Application configuration
	 * @param    nil        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	GetAppConfig() (*http.Response, error)

	/**
	 * Update a Single Company
	 * @param    *models.CompanyModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	UpdateCompany(*models.CompanyModel) error

	/**
	 * Update multiple companies in a single batch (batch size must be less than 250kb)
	 * @param    []*models.CompanyModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	UpdateCompaniesBatch([]*models.CompanyModel) error

	/**
	 * Update a Single Subscription
	 * @param    *models.SubscriptionModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	UpdateSubscription(*models.SubscriptionModel) error

	/**
	 * Update multiple Subscriptions in a single batch (batch size must be less than 250kb)
	 * @param    []*models.SubscriptionModel        body     parameter: Required
	 * @return	Returns the  response from the API call
	 */
	UpdateSubscriptionsBatch([]*models.SubscriptionModel) error

	// GetGovernanceRules gets the collector's /v1/rules endpoint which contains
	// regex governance rules and user/company governance rule templates which are
	// templated with individual user and company info from the /v1/config endpoint
	GetGovernanceRules() (GovernanceRulesResponse, error)

	// SetEventsHeaderCallback sets a Moesif API response header value and callback
	// function which is called with
	SetEventsHeaderCallback(string, func(string))

	Flush()

	Close()
}

/*
 * Factory for the API interaface returning Client
 */
func NewAPI(moesifApplicationId string, apiEndpoint *string, eventQueueSize int, batchSize int, timerWakeupSeconds int) API {
	Config.MoesifApplicationId = moesifApplicationId

	/** Maximum number of events to be store in the queue, defaults to 1 million */
	defaultEventQueueSize := 1000000
	if eventQueueSize == 0 {
		eventQueueSize = defaultEventQueueSize
	}

	/** Maximum number of events to be sent to Moesif in a single batch, defaults to 200 */
	defaultBatchSize := 200
	if batchSize == 0 {
		batchSize = defaultBatchSize
	}

	/** Schedule events batch job periodically, defaults to 2 seconds */
	defaultTimerWakeUpSeconds := 2
	if timerWakeupSeconds == 0 {
		timerWakeupSeconds = defaultTimerWakeUpSeconds
	}

	/** The base Uri for API calls */
	defaultAPIEndpoint := "https://api.moesif.net"
	if apiEndpoint != nil && *apiEndpoint != "" {
		defaultAPIEndpoint = *apiEndpoint
	}

	/** Config Options */
	Config.BaseURI = defaultAPIEndpoint
	Config.EventQueueSize = eventQueueSize
	Config.BatchSize = batchSize
	Config.TimerWakeupSeconds = timerWakeupSeconds

	return NewClient()
}

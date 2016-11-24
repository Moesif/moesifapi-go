/*
 * moesifapi
 *
 */

package api

import "github.com/moesif/moesifapi-go/models"

/*
 * Interface for the API_IMPL
 */
type API interface {
	CreateEvent(*models.EventModel) error

	CreateEventsBatch([]*models.EventModel) error
}

/*
 * Factory for the API interaface returning API_IMPL
 */
func NewAPI() API {
	return &API_IMPL{}
}

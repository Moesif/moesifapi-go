/*
 * moesifapi
 *
 */

package health

import "github.com/moesif/moesifapi-go/models"

/*
 * Interface for the HEALTH_IMPL
 */
type HEALTH interface {
	GetHealthProbe() (*models.StatusModel, error)
}

/*
 * Factory for the HEALTH interaface returning HEALTH_IMPL
 */
func NewHEALTH() HEALTH {
	return &HEALTH_IMPL{}
}

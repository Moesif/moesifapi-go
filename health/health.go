/*
 * moesifapi-go
 */
package health

import (
	"github.com/moesif/moesifapi-go/models"
)

/*
 * Interface for the Client
 */
type Health interface {
	GetHealthProbe() (*models.StatusModel, error)
}

/*
 * Factory for the Health interaface returning Client
 */
func NewHealth() Health {
	return &Client{}
}

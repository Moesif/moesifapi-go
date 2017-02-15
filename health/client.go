/*
 * moesifapi-go
 */
package health

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	moesifapi "github.com/moesif/moesifapi-go"
	"github.com/moesif/moesifapi-go/models"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

/*
 * Client structure as interface implementation
 */
type Client struct{}

/**
 * Health Probe
 * @return	Returns the *models.StatusModel response from the API call
 */
func (me *Client) GetHealthProbe() (*models.StatusModel, error) {

	url := moesifapi.BaseURI + "/health/probe"

	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Moesif-Application-Id", moesifapi.Config.MoesifApplicationId)
	req.Header.Set("User-Agent", "moesifapi-go/"+moesifapi.Version)

	resp, err := ctxhttp.Do(ctx, http.DefaultClient, req)

	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//returning the response
	var retVal *models.StatusModel = &models.StatusModel{}
	err = json.Unmarshal(body, &retVal)

	if err != nil {
		//error in parsing
		return nil, err
	}
	return retVal, nil
}

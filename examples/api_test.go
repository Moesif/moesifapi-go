package moesifapi_test

import "github.com/moesif/moesifapi-go"
import "github.com/moesif/moesifapi-go/api"
import "github.com/moesif/moesifapi-go/models"
import "fmt"
import "time"

import "testing"

func TestCreateEvent(t *testing.T) {
	moesifapi.Config.ApplicationId = "eyJhcHAiOiIzNjU6NiIsInZlciI6IjIuMCIsIm9yZyI6IjM1OTo0IiwiaWF0IjoxNDczMzc5MjAwfQ.9WOx3D357PGMxrXzFm3pV3IzJSYNsO4oRudiMI8mQ3Q"
	apiClient := api.NewAPI()

	reqTime := time.Now().UTC()
	apiVersion := "1.0"
	ipAddress := "61.48.220.123"

	req := models.EventRequestModel{
		Time:       &reqTime,
		Uri:        "https://api.acmeinc.com/widgets",
		Verb:       "GET",
		ApiVersion: &apiVersion,
		IpAddress:  &ipAddress,
		Headers: map[string]interface{}{
			"ReqHeader1": "ReqHeaderValue1",
		},
		Body: nil,
	}

	rspTime := time.Now().UTC().Add(time.Duration(1) * time.Second)

	rsp := models.EventResponseModel{
		Time:      &rspTime,
		Status:    500,
		IpAddress: nil,
		Headers: map[string]interface{}{
			"RspHeader1": "RspHeaderValue1",
		},
		Body: map[string]interface{}{
			"Key1": "Value1",
			"Key2": 12,
			"Key3": map[string]interface{}{
				"Key3_1": "SomeValue",
			},
		},
	}

	sessiomnToken := "23jdf0owekfmcn4u3qypxg09w4d8ayrcdx8nu2ng]s98y18cx98q3yhwmnhcfx43f"
	userId := "my_user_id"

	event := models.EventModel{
		Request:      req,
		Response:     rsp,
		SessionToken: &sessiomnToken,
		Tags:         nil,
		UserId:       &userId,
	}

	fmt.Printf("Event.\n%#v", event)

	result := apiClient.CreateEvent(&event)

	fmt.Printf("Created Event.\n%#v", result)
}

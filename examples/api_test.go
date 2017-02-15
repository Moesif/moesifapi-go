/*
 * moesifapi-go
 */
package moesifapi_test

import (
	"fmt"
	"testing"
	"time"

	moesifapi "github.com/moesif/moesifapi-go"
	"github.com/moesif/moesifapi-go/models"
)

func TestCreateEvent(t *testing.T) {
	appId := "eyJhcHAiOiIxMzk6MCIsInZlciI6IjIuMCIsIm9yZyI6IjI2NTowIiwiaWF0IjoxNDg3MDMwNDAwfQ.IaoyV5EHbcBH23EaCZZc5fzzlV1yGmkU7TwykE0viK8"
	apiClient := moesifapi.NewAPI(appId)

	event := genEvent()

	fmt.Printf("Event.\n%#v", event)

	result := apiClient.CreateEvent(&event)

	if result != nil {
		t.Fail()
	}
}

func TestCreateEventBatch(t *testing.T) {
	appId := "eyJhcHAiOiIxMzk6MCIsInZlciI6IjIuMCIsIm9yZyI6IjI2NTowIiwiaWF0IjoxNDg3MDMwNDAwfQ.IaoyV5EHbcBH23EaCZZc5fzzlV1yGmkU7TwykE0viK8"
	apiClient := moesifapi.NewAPI(appId)

	events := make([]*models.EventModel, 20)
	for i := 0; i < 20; i++ {
		e := genEvent()
		events[i] = &e
	}

	result := apiClient.CreateEventsBatch(events)

	if result != nil {
		t.Fail()
	}
}

func TestQueueEvent(t *testing.T) {
	appId := "eyJhcHAiOiIxMzk6MCIsInZlciI6IjIuMCIsIm9yZyI6IjI2NTowIiwiaWF0IjoxNDg3MDMwNDAwfQ.IaoyV5EHbcBH23EaCZZc5fzzlV1yGmkU7TwykE0viK8"
	apiClient := moesifapi.NewAPI(appId)

	event := genEvent()

	fmt.Printf("Event.\n%#v", event)

	result := apiClient.QueueEvent(&event)
	apiClient.Close()

	if result != nil {
		t.Fail()
	}
}

func TestQueueEvents(t *testing.T) {
	appId := "eyJhcHAiOiIxMzk6MCIsInZlciI6IjIuMCIsIm9yZyI6IjI2NTowIiwiaWF0IjoxNDg3MDMwNDAwfQ.IaoyV5EHbcBH23EaCZZc5fzzlV1yGmkU7TwykE0viK8"
	apiClient := moesifapi.NewAPI(appId)

	events := make([]*models.EventModel, 5000)
	for i := 0; i < 5000; i++ {
		e := genEvent()
		events[i] = &e
	}

	result := apiClient.QueueEvents(events)

	apiClient.Close()

	if result != nil {
		t.Fail()
	}
}

func genEvent() models.EventModel {
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
			"RspHeader1":     "RspHeaderValue1",
			"Content-Type":   "application/json",
			"Content-Length": "1000",
		},
		Body: map[string]interface{}{
			"Key1": "Value1",
			"Key2": 12,
			"Key3": map[string]interface{}{
				"Key3_1": "SomeValue",
			},
		},
	}

	sessionToken := "23jdf0owekfmcn4u3qypxg09w4d8ayrcdx8nu2ng]s98y18cx98q3yhwmnhcfx43f"
	userId := "end_user_id"

	event := models.EventModel{
		Request:      req,
		Response:     rsp,
		SessionToken: &sessionToken,
		Tags:         nil,
		UserId:       &userId,
	}
	return event
}

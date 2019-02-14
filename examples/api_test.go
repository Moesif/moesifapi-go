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

var applicationId = "Your Application Id"

func TestCreateEvent(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId)

	event := genEvent()

	fmt.Printf("Event.\n%#v", event)

	result, err := apiClient.CreateEvent(&event)

	if err != nil {
		t.Fail()
	}

	fmt.Printf("Event.\n%#v", result)
}

func TestCreateEventBatch(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId)

	events := make([]*models.EventModel, 20)
	for i := 0; i < 20; i++ {
		e := genEvent()
		events[i] = &e
	}

	result, err := apiClient.CreateEventsBatch(events)

	if err != nil {
		t.Fail()
	}

	fmt.Printf("Events Batch.\n%#v", result)
}

func TestQueueEvent(t *testing.T) {
	appId := applicationId
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
	appId := applicationId
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

func TestCreateBase64Event(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId)

	event := genBase64Event()

	fmt.Printf("Event.\n%#v", event)

	result, err := apiClient.CreateEvent(&event)

	if err != nil {
		t.Fail()
	}

	fmt.Printf("Event.\n%#v", result)
}

func TestCreateBase64EventBatch(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId)

	events := make([]*models.EventModel, 20)
	for i := 0; i < 20; i++ {
		e := genBase64Event()
		events[i] = &e
	}

	result, err := apiClient.CreateEventsBatch(events)

	if err != nil {
		t.Fail()
	}

	fmt.Printf("Events Batch.\n%#v", result)
}

func TestUpdateUser(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId)

	user := genUser()

	fmt.Printf("User.\n%#v", user)

	result := apiClient.UpdateUser(&user)

	if result != nil {
		t.Fail()
	}
}

func TestUpdateUserBatch(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId)

	users := make([]*models.UserModel, 1)
	for i := 0; i < 1; i++ {
		u := genUser()
		users[i] = &u
	}

	fmt.Printf("%v", users)

	result := apiClient.UpdateUsersBatch(users)

	if result != nil {
		t.Fail()
	}
}

func TestQueueUser(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId)

	user := genUser()

	fmt.Printf("User.\n%#v", user)

	result := apiClient.QueueUser(&user)
	apiClient.Close()

	if result != nil {
		t.Fail()
	}
}

func TestQueueUsers(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId)

	users := make([]*models.UserModel, 2)
	for i := 0; i < 2; i++ {
		u := genUser()
		users[i] = &u
	}

	result := apiClient.QueueUsers(users)

	apiClient.Close()

	if result != nil {
		t.Fail()
	}
}

func TestGetAppConfig(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId)

	fmt.Printf(applicationId)

	result, err := apiClient.GetAppConfig()

	if err != nil {
		t.Fail()
	}

	fmt.Printf("AppConfig.\n%#v", result)
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
	metadata := map[string]interface{}{
		"Key1": "metadata",
		"Key2": 12,
		"Key3": map[string]interface{}{
			"Key3_1": "SomeValue",
		},
	}

	event := models.EventModel{
		Request:      req,
		Response:     rsp,
		SessionToken: &sessionToken,
		Tags:         nil,
		UserId:       &userId,
		Metadata: 	  &metadata,
	}
	return event
}

func genBase64Event() models.EventModel {
	reqTime := time.Now().UTC()
	apiVersion := "1.0"
	ipAddress := "61.48.220.123"
	transferEncoding := "base64"

	req := models.EventRequestModel{
		Time:       &reqTime,
		Uri:        "https://api.acmeinc.com/widgets",
		Verb:       "GET",
		ApiVersion: &apiVersion,
		IpAddress:  &ipAddress,
		Headers: map[string]interface{}{
			"ReqHeader1": "ReqHeaderValue1",
		},
		Body:             nil,
		TransferEncoding: &transferEncoding,
	}

	rspTime := time.Now().UTC().Add(time.Duration(1) * time.Second)

	var rspBody interface{} = "eyJzdGF0dXMiOnRydWUsInJlZ2lvbiI6Indlc3R1cyJ9"
	rsp := models.EventResponseModel{
		Time:      &rspTime,
		Status:    500,
		IpAddress: nil,
		Headers: map[string]interface{}{
			"RspHeader1":     "RspHeaderValue1",
			"Content-Type":   "application/json",
			"Content-Length": "1000",
		},
		Body:             rspBody,
		TransferEncoding: &transferEncoding,
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

func genUser() models.UserModel {
	
	modifiedTime := time.Now().UTC()

	metadata := map[string]interface{}{
		"email": "johndoe@acmeinc.com",
		"Key1": "metadata",
		"Key2": 42,
		"Key3": map[string]interface{}{
			"Key3_1": "SomeValue",
		},
	}
	
	user := models.UserModel{
		ModifiedTime: 	  &modifiedTime,
		SessionToken:     nil,
		IpAddress:		  nil,
		UserId:			  "end_user_id",	
		UserAgentString:  nil,
		Metadata:		  &metadata,
	}
	return user
}

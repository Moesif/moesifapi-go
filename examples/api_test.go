/*
 * moesifapi-go
 */
package moesifapi_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	moesifapi "github.com/moesif/moesifapi-go"
	"github.com/moesif/moesifapi-go/models"
)

var applicationId = "Your Moesif Application Id"
var apiEndpoint string
var batchSize int
var eventQueueSize int
var timerWakeupSeconds int

func TestCreateEvent(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

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
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

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
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

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
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

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
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

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
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

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
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

	user := genUser()

	fmt.Printf("User.\n%#v", user)

	result := apiClient.UpdateUser(&user)

	if result != nil {
		t.Fail()
	}
}

func TestUpdateUserBatch(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

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
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

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
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

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
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

	fmt.Printf(applicationId)

	result, err := apiClient.GetAppConfig()

	if err != nil {
		t.Fail()
	}

	fmt.Printf("AppConfig.\n%#v", result)
}

func TestUpdateCompany(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

	company := genCompany("12345")

	fmt.Printf("Company.\n%#v", &company)

	result := apiClient.UpdateCompany(&company)

	if result != nil {
		t.Fail()
	}
}

func TestUpdateCompaniesBatch(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

	companies := make([]*models.CompanyModel, 2)
	for i := 0; i < 2; i++ {
		c := genCompany(strconv.Itoa(i))
		companies[i] = &c
	}

	fmt.Printf("Companies.\n%#v", companies)

	result := apiClient.UpdateCompaniesBatch(companies)

	if result != nil {
		t.Fail()
	}
}

func TestQueueCompany(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

	company := genCompany("1")

	fmt.Printf("Company.\n%#v", company)

	result := apiClient.QueueCompany(&company)
	apiClient.Close()

	if result != nil {
		t.Fail()
	}
}

func TestQueueCompanies(t *testing.T) {
	appId := applicationId
	apiClient := moesifapi.NewAPI(appId, &apiEndpoint, eventQueueSize, batchSize, timerWakeupSeconds)

	companies := make([]*models.CompanyModel, 2)
	for i := 0; i < 2; i++ {
		c := genCompany(strconv.Itoa(i))
		companies[i] = &c
	}

	result := apiClient.QueueCompanies(companies)

	apiClient.Close()

	if result != nil {
		t.Fail()
	}
}

func genEvent() models.EventModel {
	reqTime := time.Now().UTC()
	apiVersion := "1.0"
	ipAddress := "61.48.220.123"

	reqLength := int64(0)
	req := models.EventRequestModel{
		Time:       &reqTime,
		Uri:        "https://api.acmeinc.com/widgets",
		Verb:       "GET",
		ApiVersion: &apiVersion,
		IpAddress:  &ipAddress,
		Headers: map[string]interface{}{
			"ReqHeader1": "ReqHeaderValue1",
		},
		Body:          nil,
		ContentLength: &reqLength,
	}

	rspTime := time.Now().UTC().Add(time.Duration(1) * time.Second)

	rspLength := int64(1000)
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
		ContentLength: &rspLength,
	}

	sessionToken := "23jdf0owekfmcn4u3qypxg09w4d8ayrcdx8nu2ng]s98y18cx98q3yhwmnhcfx43f"
	userId := "my_user_id"
	companyId := "my_company_id"
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
		CompanyId:    &companyId,
		Metadata:     &metadata,
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
	rspLength := int64(len(rspBody.(string)))
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
		ContentLength:    &rspLength,
	}

	sessionToken := "23jdf0owekfmcn4u3qypxg09w4d8ayrcdx8nu2ng]s98y18cx98q3yhwmnhcfx43f"
	userId := "my_user_id"
	companyId := "my_company_id"

	event := models.EventModel{
		Request:      req,
		Response:     rsp,
		SessionToken: &sessionToken,
		Tags:         nil,
		UserId:       &userId,
		CompanyId:    &companyId,
	}
	return event
}

func genUser() models.UserModel {

	modifiedTime := time.Now().UTC()

	metadata := map[string]interface{}{
		"email": "johndoe@acmeinc.com",
		"Key1":  "metadata",
		"Key2":  42,
		"Key3": map[string]interface{}{
			"Key3_1": "SomeValue",
		},
	}

	utmSource := "Newsletter"
	utmMedium := "Email"

	campaign := models.CampaignModel{
		UtmSource: &utmSource,
		UtmMedium: &utmMedium,
	}

	companyId := "67890"

	user := models.UserModel{
		ModifiedTime:    &modifiedTime,
		SessionToken:    nil,
		IpAddress:       nil,
		UserId:          "12345",
		CompanyId:       &companyId,
		UserAgentString: nil,
		Metadata:        &metadata,
		Campaign:        &campaign,
	}
	return user
}

func genCompany(companyId string) models.CompanyModel {

	modifiedTime := time.Now().UTC()

	metadata := map[string]interface{}{
		"email": "johndoe@acmeinc.com",
		"Key1":  "metadata",
		"Key2":  42,
		"Key3": map[string]interface{}{
			"Key3_1": "SomeValue",
		},
	}

	utmSource := "Adwords"
	utmMedium := "Twitter"

	campaign := models.CampaignModel{
		UtmSource: &utmSource,
		UtmMedium: &utmMedium,
	}

	company := models.CompanyModel{
		ModifiedTime:  &modifiedTime,
		SessionToken:  nil,
		IpAddress:     nil,
		CompanyId:     companyId,
		CompanyDomain: nil,
		Metadata:      &metadata,
		Campaign:      &campaign,
	}
	return company
}

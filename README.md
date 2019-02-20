# MoesifApi Lib for Golang

Send REST API Events to Moesif for error analysis

[Source Code on GitHub](https://github.com/moesif/moesifapi-go)

## Introduction

This lib has both synchronous and async methods:

- The synchronous methods call the Moesif API directly
- The async methods (Which start with _QueueXXX_) will queue the requests into batches
and flush buffers periodically.

Async methods are expected to be the common use case

## How to install
Run the following commands:

```bash
go get github.com/moesif/moesifapi-go;
```

## How to use

(See examples/api_test.go for usage examples)

### Create single event

```go
import "github.com/moesif/moesifapi-go"
import "github.com/moesif/moesifapi-go/models"
import "time"

apiClient := moesifapi.NewAPI("my_application_id")

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

// Queue the events
err := apiClient.QueueEvent(&event)

// Create the events synchronously
err := apiClient.CreateEvent(&event)

```

### Create batches of events with bulk API


```go
import "github.com/moesif/moesifapi-go"
import "github.com/moesif/moesifapi-go/models"
import "time"

apiClient := moesifapi.NewAPI("my_application_id")

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

events := make([]*models.EventModel, 20)
for i := 0; i < 20; i++ {
	events[i] = &event
}

// Queue the events
err := apiClient.QueueEvents(events)

// Create the events batch synchronously
err := apiClient.CreateEventsBatch(event)

```

### How To Update User

```go
import "github.com/moesif/moesifapi-go"
import "github.com/moesif/moesifapi-go/models"
import "time"

apiClient := moesifapi.NewAPI("my_application_id")

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

// Queue the user
err := apiClient.QueueUser(&user)

// Update the user synchronously
err := apiClient.UpdateUser(&user)

```

### Update batches of users with bulk API

```go
import "github.com/moesif/moesifapi-go"
import "github.com/moesif/moesifapi-go/models"
import "time"

apiClient := moesifapi.NewAPI("my_application_id")

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

users := make([]*models.UserModel, 5)
	for i := 0; i < 5; i++ {
		u := genUser()
		users[i] = &u
	}

// Queue the users
err := apiClient.QueueUsers(users)

// Update the users synchronously
err := apiClient.UpdateUsersBatch(users)

```


### Update company

```go
import "github.com/moesif/moesifapi-go"
import "github.com/moesif/moesifapi-go/models"
import "time"

apiClient := moesifapi.NewAPI("my_application_id")

modifiedTime := time.Now().UTC()

metadata := map[string]interface{}{
	"email": "johndoe@acmeinc.com",
	"Key1": "metadata",
	"Key2": 42,
	"Key3": map[string]interface{}{
		"Key3_1": "SomeValue",
	},
}

company := models.CompanyModel{
	ModifiedTime: 	  &modifiedTime,
	SessionToken:     nil,
	IpAddress:		  nil,
	CompanyId:		  "1",	
	CompanyDomain:    nil,
	Metadata:		  &metadata,
}

// Queue the company
err := apiClient.QueueCompany(&company)

// Update the company synchronously
err := apiClient.UpdateCompany(&company)
```

### Update batches of companies with bulk API

```go
import "github.com/moesif/moesifapi-go"
import "github.com/moesif/moesifapi-go/models"
import "time"
import "strconv"

apiClient := moesifapi.NewAPI("my_application_id")

modifiedTime := time.Now().UTC()

companies := make([]*models.CompanyModel, 2)
	for i := 0; i < 2; i++ {
		c:= genCompany(strconv.Itoa(i)) // Generate company model
		companies[i] = &c
	}

// Queue the companies
err := apiClient.QueueCompanies(companies)

// Update the companies synchronously
err := apiClient.UpdateCompaniesBatch(companies)
```

### Health Check

```bash
go get github.com/moesif/moesifapi-go/health;
```

## How To Test:
```bash
git clone https://github.com/moesif/moesifapi-go
cd moesifapi-go/examples
# Be sure to edit the examples/api_test.go to change the application id to your real one obtained from Moesif.
# var applicationId = "Your Application Id"
go test  -v
```


## Other integrations

To view more more documentation on integration options, please visit __[the Integration Options Documentation](https://www.moesif.com/docs/getting-started/integration-options/).__

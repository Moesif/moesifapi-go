# MoesifApi Lib for Golang


[Source Code on GitHub](https://github.com/moesif/moesifapi-go)

__Check out Moesif's [Developer Documentation](https://www.moesif.com/docs) and [Golang API Reference](https://www.moesif.com/docs/api?go) to learn more__

## How To Install:

The code depends on unirest-go http library. Run the following commands:

```bash
go get github.com/moesif/moesifapi-go;
go get github.com/apimatic/unirest-go
```

## How To Use:

(See examples/api_test.go for usage examples)

```go
import "github.com/moesif/moesifapi-go"
import "github.com/moesif/moesifapi-go/api"
import "github.com/moesif/moesifapi-go/models"
import "time"

moesifapi.Config.ApplicationId = "my_application_id"
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

sessionToken := "23jdf0owekfmcn4u3qypxg09w4d8ayrcdx8nu2ng]s98y18cx98q3yhwmnhcfx43f"
userId := "end_user_id"

event := models.EventModel{
	Request:      req,
	Response:     rsp,
	SessionToken: &sessionToken,
	Tags:         nil,
	UserId:       &userId,
}

result := apiClient.CreateEvent(&event)

```

## How To Test:

```bash
git clone https://github.com/moesif/moesifapi-go
cd moesifapi-go/examples
go test  -v
```

/*
 * moesifapi-go
 */
package models

import "time"

/*
 * Structure for the custom type EventRequestModel
 */
type EventRequestModel struct {
	Time             *time.Time   `json:"time" form:"time"`                                               //Time when request was made
	Uri              string       `json:"uri" form:"uri"`                                                 //full uri of request such as https://www.example.com/my_path?param=1
	Verb             string       `json:"verb" form:"verb"`                                               //verb of the API request such as GET or POST
	Headers          interface{}  `json:"headers" form:"headers"`                                         //Key/Value map of request headers
	ApiVersion       *string      `json:"api_version,omitempty" form:"api_version,omitempty"`             //Optionally tag the call with your API or App version
	IpAddress        *string      `json:"ip_address,omitempty" form:"ip_address,omitempty"`               //IP Address of the client if known.
	Body             *interface{} `json:"body,omitempty" form:"body,omitempty"`                           //Request body
	TransferEncoding *string      `json:"transfer_encoding,omitempty" form:"transfer_encoding,omitempty"` //Transfer Encoding of Body, such as 'base64'
	ContentLength    *int64       `json:"content_length,omitempty" form:"content_length,omitempty"`       //Content Length of the body
}

/*
 * Structure for the custom type EventModel
 */
type EventModel struct {
	Request      EventRequestModel  `json:"request" form:"request"`                                 //API request object
	Response     EventResponseModel `json:"response,omitempty" form:"response,omitempty"`           //API response Object
	SessionToken *string            `json:"session_token,omitempty" form:"session_token,omitempty"` //End user's auth/session token
	Tags         *string            `json:"tags,omitempty" form:"tags,omitempty"`                   //comma separated list of tags, see documentation
	UserId       *string            `json:"user_id,omitempty" form:"user_id,omitempty"`             //End user's user_id string from your app
	CompanyId    *string            `json:"company_id,omitempty" form:"company_id,omitempty"`       //company_id string
	Metadata     interface{}        `json:"metadata,omitempty" form:"metadata,omitempty"`           //User Metadata
	Direction    *string            `json:"direction,omitempty" form:"direction,omitempty"`         // Direction of an API call
	Weight       *int               `json:"weight,omitempty" form:"weight,omitempty"`               // Weight of an API call
	AiContext    *AiContextModel    `json:"ai_context,omitempty" form:"ai_context,omitempty"`       // AI context for the API call
	A2a          *A2aModel          `json:"a2a,omitempty" form:"a2a,omitempty"`                     // A2A analytics for the API call
}

// AiContextModel /*
type AiContextModel struct {
	Provider     *string                `json:"provider,omitempty" form:"provider,omitempty"`           //AI provider such as openai
	Model        *string                `json:"model,omitempty" form:"model,omitempty"`                 //Model name
	ModelType    *string                `json:"model_type,omitempty" form:"model_type,omitempty"`       //Model type such as chat
	Usage        *AiUsageModel          `json:"usage,omitempty" form:"usage,omitempty"`                 //Token usage details
	Attribution  map[string]interface{} `json:"attribution,omitempty" form:"attribution,omitempty"`     //Attribution details (dynamic key-value pairs)
	CostMetrics  *AiCostMetricsModel    `json:"cost_metrics,omitempty" form:"cost_metrics,omitempty"`   //Cost metrics
	FinishReason *string                `json:"finish_reason,omitempty" form:"finish_reason,omitempty"` //Finish reason
}

// AiUsageModel /*
type AiUsageModel struct {
	PromptTokens             *int `json:"prompt_tokens,omitempty" form:"prompt_tokens,omitempty"`                             //Prompt tokens
	CompletionTokens         *int `json:"completion_tokens,omitempty" form:"completion_tokens,omitempty"`                     //Completion tokens
	TotalTokens              *int `json:"total_tokens,omitempty" form:"total_tokens,omitempty"`                               //Total tokens
	CachedTokens             *int `json:"cached_tokens,omitempty" form:"cached_tokens,omitempty"`                             //Cached tokens
	ReasoningTokens          *int `json:"reasoning_tokens,omitempty" form:"reasoning_tokens,omitempty"`                       //Reasoning tokens
	CacheReadInputTokens     *int `json:"cache_read_input_tokens,omitempty" form:"cache_read_input_tokens,omitempty"`         //Cache read input tokens
	CacheCreationInputTokens *int `json:"cache_creation_input_tokens,omitempty" form:"cache_creation_input_tokens,omitempty"` //Cache creation input tokens
}

// AiCostMetricsModel /*
type AiCostMetricsModel struct {
	TotalInputCost    *float64 `json:"total_input_cost,omitempty" form:"total_input_cost,omitempty"`       //Total input cost
	CachedInputCost   *float64 `json:"cached_input_cost,omitempty" form:"cached_input_cost,omitempty"`     //Cached input cost
	CachedReadCost    *float64 `json:"cached_read_cost,omitempty" form:"cached_read_cost,omitempty"`       //Cached read cost
	CacheCreationCost *float64 `json:"cache_creation_cost,omitempty" form:"cache_creation_cost,omitempty"` //Cache creation cost
	ReasoningCost     *float64 `json:"reasoning_cost,omitempty" form:"reasoning_cost,omitempty"`           //Reasoning cost
	OutputCost        *float64 `json:"output_cost,omitempty" form:"output_cost,omitempty"`                 //Output cost
	TotalCost         *float64 `json:"total_cost,omitempty" form:"total_cost,omitempty"`                   //Total cost
	PriceVersion      *string  `json:"price_version,omitempty" form:"price_version,omitempty"`             //Price version
}

// A2aModel
type A2aModel struct {
	Operation       *string           `json:"operation,omitempty" form:"operation,omitempty"`               //A2A operation (e.g. SendMessage, GetTask)
	Transport       *string           `json:"transport,omitempty" form:"transport,omitempty"`               //Wire binding (JSONRPC, GRPC, HTTP+JSON)
	RequestType     *string           `json:"request_type,omitempty" form:"request_type,omitempty"`         //Traffic classifier (operation, agentCard, preflight)
	ProtocolVersion *string           `json:"protocol_version,omitempty" form:"protocol_version,omitempty"` //A2A protocol version (e.g. 1.0)
	Request         *A2aRequestModel  `json:"request,omitempty" form:"request,omitempty"`                   //A2A request-side analytics
	Response        *A2aResponseModel `json:"response,omitempty" form:"response,omitempty"`                 //A2A response-side analytics
	Terminal        *bool             `json:"terminal,omitempty" form:"terminal,omitempty"`                 //True if the observed task state is terminal
	Outcome         *string           `json:"outcome,omitempty" form:"outcome,omitempty"`                   //Derived outcome (SUCCESS, FAILURE, UNKNOWN)
	FailureOrigin   *string           `json:"failure_origin,omitempty" form:"failure_origin,omitempty"`     //Layer responsible for a failure (CLIENT, POLICY, GATEWAY, UPSTREAM, UNKNOWN)
}

// A2aRequestModel
type A2aRequestModel struct {
	MessageId         *string `json:"message_id,omitempty" form:"message_id,omitempty"`                   //Opaque client-generated message id
	TaskId            *string `json:"task_id,omitempty" form:"task_id,omitempty"`                         //Opaque task id
	ContextId         *string `json:"context_id,omitempty" form:"context_id,omitempty"`                   //Opaque context id
	InputPartCount    *int    `json:"input_part_count,omitempty" form:"input_part_count,omitempty"`       //Number of parts in the request message
	ReturnImmediately *bool   `json:"return_immediately,omitempty" form:"return_immediately,omitempty"`   //From SendMessageConfiguration.return_immediately
	HistoryLength     *int    `json:"history_length,omitempty" form:"history_length,omitempty"`           //From SendMessageConfiguration.history_length
}

// A2aResponseModel
type A2aResponseModel struct {
	IsError            *bool   `json:"is_error,omitempty" form:"is_error,omitempty"`                             //Whether the response carried an error
	ErrorCode          *int    `json:"error_code,omitempty" form:"error_code,omitempty"`                         //JSON-RPC -32xxx, gRPC status 0-16, or HTTP status >= 400
	IsStreaming        *bool   `json:"is_streaming,omitempty" form:"is_streaming,omitempty"`                     //Whether delivered via SSE / gRPC stream
	TimeToFirstEventMs *int64  `json:"time_to_first_event_ms,omitempty" form:"time_to_first_event_ms,omitempty"` //Request start to first stream event, ms (streaming only)
	StreamDurationMs   *int64  `json:"stream_duration_ms,omitempty" form:"stream_duration_ms,omitempty"`         //First frame to stream end, ms (streaming only)
	PayloadType        *string `json:"payload_type,omitempty" form:"payload_type,omitempty"`                     //A2A proto response type (task, message, status_update, etc.)
	TaskId             *string `json:"task_id,omitempty" form:"task_id,omitempty"`                               //Task id observed in the response
	ContextId          *string `json:"context_id,omitempty" form:"context_id,omitempty"`                         //Context id observed in the response
	TaskState          *string `json:"task_state,omitempty" form:"task_state,omitempty"`                         //Latest observed A2A TaskState enum value
}

/*
 * Structure for the custom type EventResponseModel
 */
type EventResponseModel struct {
	Time             *time.Time  `json:"time" form:"time"`                                               //Time when response received
	Status           int         `json:"status" form:"status"`                                           //HTTP Status code such as 200
	Headers          interface{} `json:"headers" form:"headers"`                                         //Key/Value map of response headers
	Body             interface{} `json:"body" form:"body"`                                               //Response body
	IpAddress        *string     `json:"ip_address,omitempty" form:"ip_address,omitempty"`               //IP Address from the response, such as the server IP Address
	TransferEncoding *string     `json:"transfer_encoding,omitempty" form:"transfer_encoding,omitempty"` //Transfer Encoding of Body, such as 'base64'
	ContentLength    *int64      `json:"content_length,omitempty" form:"content_length,omitempty"`       //Content Length of the body
}

/*
 * Structure for the custom type StatusModel
 */
type StatusModel struct {
	Status bool   `json:"status" form:"status"` //Status of Call
	Region string `json:"region" form:"region"` //Location
}

/*
 * Structure for the custom type CampaignModel
 */
type CampaignModel struct {
	UtmSource       *string `json:"utm_source,omitempty" form:"utm_source,omitempty"`             //The Utm source
	UtmMedium       *string `json:"utm_medium,omitempty" form:"utm_medium,omitempty"`             //The Utm Medium
	UtmCampaign     *string `json:"utm_campaign,omitempty" form:"utm_campaign,omitempty"`         //The Utm Campaign
	UtmTerm         *string `json:"utm_term,omitempty" form:"utm_term,omitempty"`                 //The Utm Term
	UtmContent      *string `json:"utm_content,omitempty" form:"utm_content,omitempty"`           //The Utm Content
	Referrer        *string `json:"referrer,omitempty" form:"referrer,omitempty"`                 //The Referrer
	ReferringDomain *string `json:"referring_domain,omitempty" form:"referring_domain,omitempty"` //The Referring Domain
	Gclid           *string `json:"gclid,omitempty" form:"gclid,omitempty"`                       //The Gclid
}

/*
 * Structure for the custom type UserModel
 */
type UserModel struct {
	ModifiedTime    *time.Time     `json:"modified_time" form:"modified_time"`                             //Time when request was made
	SessionToken    *string        `json:"session_token,omitempty" form:"session_token,omitempty"`         //End user's auth/session token
	IpAddress       *string        `json:"ip_address,omitempty" form:"ip_address,omitempty"`               //IP Address of the client if known.
	UserId          string         `json:"user_id" form:"user_id"`                                         //End user's user_id string from your app
	CompanyId       *string        `json:"company_id,omitempty" form:"company_id,omitempty"`               //CompanyId associated with the user if known
	UserAgentString *string        `json:"user_agent_string,omitempty" form:"user_agent_string,omitempty"` //End user's user agent string
	Metadata        interface{}    `json:"metadata,omitempty" form:"metadata,omitempty"`                   //User Metadata
	Campaign        *CampaignModel `json:"campaign,omitempty" form:"campaign,omitempty"`                   //The Campaign Object
}

/*
 * Structure for the custom type CompanyModel
 */
type CompanyModel struct {
	ModifiedTime  *time.Time     `json:"modified_time" form:"modified_time"`                       //Time when request was made
	SessionToken  *string        `json:"session_token,omitempty" form:"session_token,omitempty"`   //End user's auth/session token
	IpAddress     *string        `json:"ip_address,omitempty" form:"ip_address,omitempty"`         //IP Address of the client if known.
	CompanyId     string         `json:"company_id" form:"company_id"`                             //Company Id string from your app
	CompanyDomain *string        `json:"company_domain,omitempty" form:"company_domain,omitempty"` //Company Domain string
	Metadata      interface{}    `json:"metadata,omitempty" form:"metadata,omitempty"`             //User Metadata
	Campaign      *CampaignModel `json:"campaign,omitempty" form:"campaign,omitempty"`             //The Campaign Object
}

/*
 * Structure for the custom type SubscriptionModel
 */
type SubscriptionModel struct {
	SubscriptionId     string      `json:"subscription_id" form:"subscription_id"`           //Subscription Id
	CompanyId          string      `json:"company_id" form:"company_id"`                     //Company Id
	CurrentPeriodStart *time.Time  `json:"current_period_start" form:"current_period_start"` //Current Period Start
	CurrentPeriodEnd   *time.Time  `json:"current_period_end" form:"current_period_end"`     //Current Period End
	Status             *string     `json:"status,omitempty" form:"status,omitempty"`         //Status
	Metadata           interface{} `json:"metadata,omitempty" form:"metadata,omitempty"`     //Subscription Metadata
}

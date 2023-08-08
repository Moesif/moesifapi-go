package moesifapi

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type GovernanceRulesResponse struct {
	Rules []GovernanceRule
	ETag  string
}

// implement support for regex AND user/company cohort.  Logically both conditions must be true
// also make example of regex with OR/AND conditions between expressions
type GovernanceRule struct {
	ID                string `json:"_id"`
	Name              string `json:"name"`
	Type              string `json:"type"`
	Block             bool   `json:"block"`
	ApplyTo           string `json:"applied_to"`              // "matching" or "not_matching"
	ApplyUnidentified bool   `json:"applied_to_unidentified"` // true if rule applies to unidentified entities
	// RegexConfigOr is a boolean expression tree with a fixed depth of two,
	// a slice of slices of regex patterns.  The regex rule as a whole is a match,
	// evals to true, if one OR more of the top level slice items is true.
	// One of the top level slices is a slice of regex patterns to match, and
	// one slice is true of one pattern AND the others are all matches.
	RegexConfigOr     []RegexConditionsAnd `json:"regex_config"`
	ResponseOverrides ResponseOverrides    `json:"response"`
	Variables         []Variable           `json:"variables"`
	OrgID             string               `json:"org_id"`
	AppID             string               `json:"app_id"`
	CreatedAt         string               `json:"created_at"`
}

type RegexConditionsAnd struct {
	Conditions []RegexCondition `json:"conditions"`
}

type RegexCondition struct {
	Path  string `json:"path"`
	Value string `json:"value"`
}

type ResponseOverrides struct {
	Body    BodyTemplate      `json:"body"`
	Headers map[string]string `json:"headers"`
	Status  int               `json:"status"`
}

type BodyTemplate string

type Variable struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (b *BodyTemplate) UnmarshalJSON(data []byte) error {
	if b == nil {
		// This error can't actually happen except in tests or custom extension code
		// BodyTemplate is a value in ResponseOverrides and therefore is always allocated
		return errors.New("Template: UnmarshalJSON on nil pointer")
	}
	*b = BodyTemplate(data)
	return nil
}

// Template replaces substitution variables in the request body template
// vars is a map of variable name to variable substitution value
func Template(t string, vars map[string]string) string {
	// build a replacer with all variable {{name}}, value... pairs in vars
	v := make([]string, 0, len(vars)*2)
	for name, value := range vars {
		v = append(v, "{{"+name+"}}", value)
	}
	r := strings.NewReplacer(v...)
	return r.Replace(string(t))
}

func (c *Client) GetGovernanceRules() (r GovernanceRulesResponse, err error) {
	url := Config.BaseURI + "/v1/rules"

	ctx, cancel := context.WithTimeout(c.ctx, c.timeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("X-Moesif-Application-Id", Config.MoesifApplicationId)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Governance rules request error: %v", err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Governance rules response body read error: %v", err)
		return
	}

	if err = json.Unmarshal(body, &r.Rules); err != nil {
		log.Printf("Governance rules response body malformed: %v", err)
		return
	}

	if values, ok := resp.Header["X-Moesif-Rules-Tag"]; ok {
		r.ETag = values[0]
	}
	return
}

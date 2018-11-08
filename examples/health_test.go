/*
 * moesifapi-go
 */
package moesifapi_test

import (
	"fmt"
	"testing"

	"github.com/moesif/moesifapi-go/health"
)

func TestHealth(t *testing.T) {
	healthClient := health.NewHealth()

	result, err := healthClient.GetHealthProbe()

	if err != nil {
		t.Fail()
	}

	fmt.Printf("%#v\n", result)
}

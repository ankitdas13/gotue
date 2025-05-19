// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gotue_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/gotue-go"
	"github.com/stainless-sdks/gotue-go/internal/testutil"
	"github.com/stainless-sdks/gotue-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := gotue.NewClient(
		option.WithBaseURL(baseURL),
		option.WithUsername("My Username"),
		option.WithPassword("My Password"),
	)
	orders, err := client.Orders.List(context.TODO(), gotue.OrderListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", orders.Count)
}

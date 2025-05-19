// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gotue_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/gotue-go"
	"github.com/stainless-sdks/gotue-go/internal/testutil"
	"github.com/stainless-sdks/gotue-go/option"
)

func TestSettlementReconGetWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Settlements.Recon.Get(context.TODO(), gotue.SettlementReconGetParams{
		Month: 0,
		Year:  0,
		Count: gotue.Int(1),
		Day:   gotue.Int(0),
		Skip:  gotue.Int(0),
	})
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

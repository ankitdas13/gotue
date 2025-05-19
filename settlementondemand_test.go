// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gotue_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/ankitdas13/gotue"
	"github.com/ankitdas13/gotue/internal/testutil"
	"github.com/ankitdas13/gotue/option"
)

func TestSettlementOndemandNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Settlements.Ondemand.New(context.TODO(), gotue.SettlementOndemandNewParams{
		Amount:      200000,
		Description: gotue.String("Need this to make vendor payments."),
		Notes: gotue.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
		SettleFullBalance: gotue.Bool(false),
	})
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSettlementOndemandGet(t *testing.T) {
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
	_, err := client.Settlements.Ondemand.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

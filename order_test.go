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

func TestOrderNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.New(context.TODO(), gotue.OrderNewParams{
		Amount:                1000,
		Currency:              "INR",
		FirstPaymentMinAmount: gotue.Int(0),
		Notes: gotue.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
		OfferID:        gotue.String("offer_JHD834hjbxzhd38d"),
		PartialPayment: gotue.Bool(true),
		Receipt:        gotue.String("rcptid_11"),
		TestNew:        gotue.String("test_new"),
	})
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderGet(t *testing.T) {
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
	_, err := client.Orders.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.Update(
		context.TODO(),
		"id",
		gotue.OrderUpdateParams{
			Notes: gotue.NotesUnionParam{
				OfStringMap: map[string]string{
					"key1": "value3",
					"key2": "value2",
				},
			},
		},
	)
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderListWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.List(context.TODO(), gotue.OrderListParams{
		Authorized: true,
		Count:      gotue.Int(1),
		Expand:     []string{"payments"},
		From:       gotue.Int(0),
		Receipt:    gotue.String("receipt"),
		Skip:       gotue.Int(0),
		To:         gotue.Int(0),
	})
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestOrderListPaymentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Orders.ListPayments(
		context.TODO(),
		"id",
		gotue.OrderListPaymentsParams{
			Expand: gotue.OrderListPaymentsParamsExpandCard,
		},
	)
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

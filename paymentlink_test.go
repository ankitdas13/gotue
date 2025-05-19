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

func TestPaymentLinkNewWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.New(context.TODO(), gotue.PaymentLinkNewParams{
		Amount:         1000,
		Currency:       "INR",
		AcceptPartial:  gotue.Bool(true),
		CallbackMethod: gotue.String("get"),
		CallbackURL:    gotue.String("https://example-callback-url.com/"),
		Customer: gotue.PaymentLinkNewParamsCustomer{
			Contact: gotue.String("+919000090000"),
			Email:   gotue.String("gaurav.kumar@example.com"),
			Name:    gotue.String("Gaurav Kumar"),
		},
		Description:           gotue.String("Payment for policy no"),
		ExpireBy:              gotue.Int(1691097057),
		FirstMinPartialAmount: gotue.Int(100),
		Notes: gotue.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
		Notify: gotue.PaymentLinkNewParamsNotify{
			Email: gotue.Bool(true),
			SMS:   gotue.Bool(true),
		},
		ReferenceID:    gotue.String("TS1989"),
		ReminderEnable: gotue.Bool(true),
		UpiLink:        gotue.Bool(true),
	})
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkGet(t *testing.T) {
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
	_, err := client.PaymentLinks.Get(context.TODO(), "id")
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.Update(
		context.TODO(),
		"id",
		gotue.PaymentLinkUpdateParams{
			AcceptPartial: gotue.Bool(false),
			ExpireBy:      gotue.Int(1653347540),
			Notes: gotue.NotesUnionParam{
				OfStringMap: map[string]string{
					"key1": "value3",
					"key2": "value2",
				},
			},
			ReferenceID: gotue.String("TS35"),
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

func TestPaymentLinkListWithOptionalParams(t *testing.T) {
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
	_, err := client.PaymentLinks.List(context.TODO(), gotue.PaymentLinkListParams{
		PaymentID:   gotue.String("payment_id"),
		ReferenceID: gotue.String("reference_id"),
	})
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPaymentLinkNotify(t *testing.T) {
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
	_, err := client.PaymentLinks.Notify(
		context.TODO(),
		gotue.PaymentLinkNotifyParamsMediumSMS,
		gotue.PaymentLinkNotifyParams{
			ID: "id",
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

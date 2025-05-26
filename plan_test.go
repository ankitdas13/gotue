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
	"github.com/ankitdas13/gotue/packages/param"
)

func TestPlanNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Plans.New(context.TODO(), gotue.PlanNewParams{
		Interval: 1,
		Item: gotue.PlanNewParamsItem{
			ID:           gotue.String("item_00000000000001"),
			Active:       gotue.Bool(true),
			Amount:       gotue.Int(69900),
			CreatedAt:    gotue.Int(1580219935),
			Currency:     gotue.String("INR"),
			Description:  gotue.String("Description for the test plan - Weekly"),
			HsnCode:      param.Null[string](),
			Name:         gotue.String("Test plan - Weekly"),
			SacCode:      param.Null[string](),
			TaxGroupID:   param.Null[string](),
			TaxID:        param.Null[string](),
			TaxInclusive: gotue.Bool(false),
			TaxRate:      param.Null[float64](),
			Type:         gotue.String("plan"),
			Unit:         param.Null[string](),
			UnitAmount:   gotue.Int(69900),
			UpdatedAt:    gotue.Int(1580219935),
		},
		Period: gotue.PlanNewParamsPeriodWeekly,
		Notes: gotue.NotesUnionParam{
			OfStringMap: map[string]string{
				"key1": "value3",
				"key2": "value2",
			},
		},
	})
	if err != nil {
		var apierr *gotue.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

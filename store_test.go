// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waldeedle_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/waldeedle-go"
	"github.com/stainless-sdks/waldeedle-go/internal/testutil"
	"github.com/stainless-sdks/waldeedle-go/option"
	"github.com/stainless-sdks/waldeedle-go/shared"
)

func TestStoreNewOrderWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := waldeedle.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.NewOrder(context.TODO(), waldeedle.StoreNewOrderParams{
		Order: shared.OrderParam{
			ID:       waldeedle.F(int64(10)),
			PetID:    waldeedle.F(int64(198772)),
			Quantity: waldeedle.F(int64(7)),
			ShipDate: waldeedle.F(time.Now()),
			Status:   waldeedle.F(shared.OrderStatusApproved),
			Complete: waldeedle.F(true),
		},
	})
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStoreInventory(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := waldeedle.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Store.Inventory(context.TODO())
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

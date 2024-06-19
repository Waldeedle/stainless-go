// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waldeedle_test

import (
	"context"
	"os"
	"testing"

	"github.com/Waldeedle/stainless-go"
	"github.com/Waldeedle/stainless-go/internal/testutil"
	"github.com/Waldeedle/stainless-go/option"
	"github.com/Waldeedle/stainless-go/shared"
)

func TestUsage(t *testing.T) {
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
	order, err := client.Store.NewOrder(context.TODO(), waldeedle.StoreNewOrderParams{
		Order: shared.OrderParam{},
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", order.ID)
}

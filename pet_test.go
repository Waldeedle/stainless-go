// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waldeedle_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/undefined/stainless-go"
	"github.com/undefined/stainless-go/internal/testutil"
	"github.com/undefined/stainless-go/option"
)

func TestPetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.New(context.TODO(), waldeedle.PetNewParams{
		Pet: waldeedle.PetParam{
			ID:   waldeedle.F(int64(10)),
			Name: waldeedle.F("doggie"),
			Category: waldeedle.F(waldeedle.PetCategoryParam{
				ID:   waldeedle.F(int64(1)),
				Name: waldeedle.F("Dogs"),
			}),
			PhotoURLs: waldeedle.F([]string{"string", "string", "string"}),
			Tags: waldeedle.F([]waldeedle.PetTagParam{{
				ID:   waldeedle.F(int64(0)),
				Name: waldeedle.F("string"),
			}, {
				ID:   waldeedle.F(int64(0)),
				Name: waldeedle.F("string"),
			}, {
				ID:   waldeedle.F(int64(0)),
				Name: waldeedle.F("string"),
			}}),
			Status: waldeedle.F(waldeedle.PetStatusAvailable),
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

func TestPetGet(t *testing.T) {
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
	_, err := client.Pets.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.Update(context.TODO(), waldeedle.PetUpdateParams{
		Pet: waldeedle.PetParam{
			ID:   waldeedle.F(int64(10)),
			Name: waldeedle.F("doggie"),
			Category: waldeedle.F(waldeedle.PetCategoryParam{
				ID:   waldeedle.F(int64(1)),
				Name: waldeedle.F("Dogs"),
			}),
			PhotoURLs: waldeedle.F([]string{"string", "string", "string"}),
			Tags: waldeedle.F([]waldeedle.PetTagParam{{
				ID:   waldeedle.F(int64(0)),
				Name: waldeedle.F("string"),
			}, {
				ID:   waldeedle.F(int64(0)),
				Name: waldeedle.F("string"),
			}, {
				ID:   waldeedle.F(int64(0)),
				Name: waldeedle.F("string"),
			}}),
			Status: waldeedle.F(waldeedle.PetStatusAvailable),
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

func TestPetDelete(t *testing.T) {
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
	err := client.Pets.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByStatus(context.TODO(), waldeedle.PetFindByStatusParams{
		Status: waldeedle.F(waldeedle.PetFindByStatusParamsStatusAvailable),
	})
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetFindByTagsWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.FindByTags(context.TODO(), waldeedle.PetFindByTagsParams{
		Tags: waldeedle.F([]string{"string", "string", "string"}),
	})
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUpdateByIDWithOptionalParams(t *testing.T) {
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
	err := client.Pets.UpdateByID(
		context.TODO(),
		int64(0),
		waldeedle.PetUpdateByIDParams{
			Name:   waldeedle.F("string"),
			Status: waldeedle.F("string"),
		},
	)
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPetUploadImageWithOptionalParams(t *testing.T) {
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
	_, err := client.Pets.UploadImage(
		context.TODO(),
		int64(0),
		waldeedle.PetUploadImageParams{
			Image:              io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			AdditionalMetadata: waldeedle.F("string"),
		},
	)
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

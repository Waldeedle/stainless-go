// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waldeedle_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Waldeedle/stainless-go"
	"github.com/Waldeedle/stainless-go/internal/testutil"
	"github.com/Waldeedle/stainless-go/option"
)

func TestUserNewWithOptionalParams(t *testing.T) {
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
	err := client.User.New(context.TODO(), waldeedle.UserNewParams{
		User: waldeedle.UserParam{
			ID:         waldeedle.F(int64(10)),
			Username:   waldeedle.F("theUser"),
			FirstName:  waldeedle.F("John"),
			LastName:   waldeedle.F("James"),
			Email:      waldeedle.F("john@email.com"),
			Password:   waldeedle.F("12345"),
			Phone:      waldeedle.F("12345"),
			UserStatus: waldeedle.F(int64(1)),
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

func TestUserGet(t *testing.T) {
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
	_, err := client.User.Get(context.TODO(), "string")
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserUpdateWithOptionalParams(t *testing.T) {
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
	err := client.User.Update(
		context.TODO(),
		"string",
		waldeedle.UserUpdateParams{
			User: waldeedle.UserParam{
				ID:         waldeedle.F(int64(10)),
				Username:   waldeedle.F("theUser"),
				FirstName:  waldeedle.F("John"),
				LastName:   waldeedle.F("James"),
				Email:      waldeedle.F("john@email.com"),
				Password:   waldeedle.F("12345"),
				Phone:      waldeedle.F("12345"),
				UserStatus: waldeedle.F(int64(1)),
			},
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

func TestUserDelete(t *testing.T) {
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
	err := client.User.Delete(context.TODO(), "string")
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserNewWithList(t *testing.T) {
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
	_, err := client.User.NewWithList(context.TODO(), waldeedle.UserNewWithListParams{
		Items: []waldeedle.UserParam{{
			ID:         waldeedle.F(int64(10)),
			Username:   waldeedle.F("theUser"),
			FirstName:  waldeedle.F("John"),
			LastName:   waldeedle.F("James"),
			Email:      waldeedle.F("john@email.com"),
			Password:   waldeedle.F("12345"),
			Phone:      waldeedle.F("12345"),
			UserStatus: waldeedle.F(int64(1)),
		}, {
			ID:         waldeedle.F(int64(10)),
			Username:   waldeedle.F("theUser"),
			FirstName:  waldeedle.F("John"),
			LastName:   waldeedle.F("James"),
			Email:      waldeedle.F("john@email.com"),
			Password:   waldeedle.F("12345"),
			Phone:      waldeedle.F("12345"),
			UserStatus: waldeedle.F(int64(1)),
		}, {
			ID:         waldeedle.F(int64(10)),
			Username:   waldeedle.F("theUser"),
			FirstName:  waldeedle.F("John"),
			LastName:   waldeedle.F("James"),
			Email:      waldeedle.F("john@email.com"),
			Password:   waldeedle.F("12345"),
			Phone:      waldeedle.F("12345"),
			UserStatus: waldeedle.F(int64(1)),
		}},
	})
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLoginWithOptionalParams(t *testing.T) {
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
	_, err := client.User.Login(context.TODO(), waldeedle.UserLoginParams{
		Password: waldeedle.F("string"),
		Username: waldeedle.F("string"),
	})
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestUserLogout(t *testing.T) {
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
	err := client.User.Logout(context.TODO())
	if err != nil {
		var apierr *waldeedle.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

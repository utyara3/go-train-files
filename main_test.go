package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	if got := rec.Body.String(); got != "OK\n" {
		t.Fatalf("expected body %q, got %q", "OK\n", got)
	}
}

func TestGetUsersHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()

	getUsersHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	gotContentType := res.Header.Get("Content-Type")
	wantContentType := "application/json"

	if gotContentType != wantContentType {
		t.Fatalf("expected content type '%s', got '%s'", wantContentType, gotContentType)
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Reading Error: %v", err)
	}

	var users []User

	if err := json.Unmarshal(bodyBytes, &users); err != nil {
		t.Fatalf("Invalid JSON: %v. Raw body: %s", err, string(bodyBytes))
	}

	if len(users) == 0 {
		t.Error("expected to get users, but got an empty list")
	}
}

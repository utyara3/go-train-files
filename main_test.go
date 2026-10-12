package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type HealthStatus struct {
	Status string `json:"status"`
}

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	var health HealthStatus
	expected := HealthStatus{Status: "ok"}

	if err := json.NewDecoder(res.Body).Decode(&health); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	if health.Status != expected.Status {
		t.Fatalf("wrong json, expected %v, got %v", expected, health)
	}
}

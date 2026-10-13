package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}

	expectedContentType := "application/json"
	gotContentType := res.Header.Get("Content-Type")
	if gotContentType != expectedContentType {
		t.Fatalf("expected %s content type, got %s", expectedContentType, gotContentType)
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

func TestHealthHandlerUnavailable(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	healthHandler(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	log.Printf("Header: %v", res.Header)

	if res.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected %v status code, got %v", http.StatusServiceUnavailable, res.StatusCode)
	}

	if res.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("expected content type %s, got %s", "application/json", res.Header.Get("Content-Type"))
	}

	var health HealthStatus
	if err := json.NewDecoder(res.Body).Decode(&health); err != nil {
		t.Fatalf("invalid JSON, got %v", err)
	}

	if health.Status != "unavailable" {
		t.Fatalf("expected %s status, got %s", "unavailable", health.Status)
	}
}

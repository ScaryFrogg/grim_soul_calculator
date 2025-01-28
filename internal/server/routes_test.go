package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/ScaryFrogg/grim_soul_calculator/internal/common"
	"github.com/ScaryFrogg/grim_soul_calculator/internal/database"
)

func setupServer() *Server {
	dbConn := database.New()
	return &Server{
		db: dbConn,
	}
}

func TestMain(m *testing.M) {
	absPath, err := filepath.Abs("../../gs_calc_test.db")
	if err != nil {
		fmt.Printf("error resolving database path: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Using database path: %s\n", absPath)
	os.Setenv("BLUEPRINT_DB_URL", absPath)

	code := m.Run()
	os.Exit(code)
}

const (
	expectedName = "candle"
)

func TestTrades(t *testing.T) {
	s := setupServer()
	server := httptest.NewServer(http.HandlerFunc(s.TradesHandler))
	defer server.Close()

	// Make a test request
	resp, err := http.Get(server.URL + "/trades")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()

	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	fmt.Printf("resp.Body: %v\n", resp.Body)
}

func TestGetItemHandler(t *testing.T) {
	s := setupServer()
	server := httptest.NewServer(http.HandlerFunc(s.GetItemHandler))
	defer server.Close()

	// Make a test request
	resp, err := http.Get(server.URL + "/design/1")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()

	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", resp.Status)
	}

	fmt.Printf("resp.Body: %v\n", resp.Body)
	var result types.Blueprint
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("error decoding response body. Err: %v", err)
	}
	if result.Name != expectedName {
		t.Errorf("expected Name to be %v; got %v", expectedName, result.Name)
	}
}

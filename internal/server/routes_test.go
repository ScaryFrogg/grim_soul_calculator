package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/ScaryFrogg/grim_soul_calculator/internal/common"
	"github.com/ScaryFrogg/grim_soul_calculator/internal/database"
)

var server *httptest.Server

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

	s := setupServer()

	mux := s.RegisterRoutes()
	server = httptest.NewServer(mux) // Use mux, not HandlerFunc
	defer server.Close()

	code := m.Run()
	os.Exit(code)
}

const (
	expectedDesign = "weapon shelf lvl2"
)

func TestTrades(t *testing.T) {
	// Make a test request
	resp, err := http.Get(server.URL + "/trades")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v\n", err)
	}
	defer resp.Body.Close()

	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v\n", resp.Status)
	}

	var result []types.Trade
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		t.Fatalf("error decoding response body. Err: %v\n", err)
	}
	if len(result) == 0 {
		t.Log("expected to find trades")
	}

	t.Logf("found %d trades\n", len(result))
}

func TestGetItemHandler(t *testing.T) {
	// Make a test request
	resp, err := http.Get(server.URL + "/design/1")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v\n", err)
	}
	defer resp.Body.Close()

	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v\n", resp.Status)
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Printf("Raw Response: %s\n", string(bodyBytes))

	var result types.Blueprint
	parseErr := json.Unmarshal(bodyBytes, &result) // Use json.Unmarshal instead of decoding from resp.Body
	if parseErr != nil {
		t.Fatalf("Error decoding response body. Err: %v\n", parseErr)
	}
	if result.Name != expectedDesign {
		t.Fatalf("expected Name to be %v; got %v\n", expectedDesign, result.Name)
	}
}

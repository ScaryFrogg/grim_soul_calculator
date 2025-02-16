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
	absPath, err := filepath.Abs("../../gs_calc.db")
	if err != nil {
		fmt.Printf("error resolving database path: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Using database path: %s\n", absPath)
	os.Setenv("BLUEPRINT_DB_URL", absPath)

	s := setupServer()

	mux := s.RegisterRoutes()
	server = httptest.NewServer(mux)
	defer server.Close()

	code := m.Run()
	os.Exit(code)
}

func TestDesigns(t *testing.T) {
	var result []types.BuildBaseInfo
	RequestAndParseResponse(t, "/designs", &result)

	if len(result) == 0 {
		t.Fatal("expected to find designs")
	}

}

func TestCook(t *testing.T) {
	var result []types.Recipe
	RequestAndParseResponse(t, "/cook", &result)

	if len(result) == 0 {
		t.Fatal("expected to find designs")
	}

}

func TestSets(t *testing.T) {
	var result []types.ArmorData
	RequestAndParseResponse(t, "/armor", &result)

	if len(result) == 0 {
		t.Fatal("expected to find sets")
	}

}

func TestSetPieces(t *testing.T) {
	var result []types.ArmorData
	RequestAndParseResponse(t, "/armor/set/12", &result)

	if len(result) != 5 {
		t.Fatal("expected to find 5 set item")
	}

}

func TestRequirementById(t *testing.T) {
	var result []types.Requirement
	RequestAndParseResponse(t, "/requirement/1", &result)

	if len(result) == 0 {
		t.Fatal("expected to find requirements")
	}

}

func TestItemById(t *testing.T) {
	expectedName := "candle"
	var result types.MaterialInfo
	RequestAndParseResponse(t, "/item/1", &result)

	if result.Name != expectedName {
		t.Fatalf("expected Name to be %v; got %v\n", expectedName, result.Name)
	}
}

func TestTrades(t *testing.T) {
	var result []types.Trade
	RequestAndParseResponse(t, "/trades", &result)
	if len(result) == 0 {
		t.Fatal("expected to find trades")
	}

	t.Logf("found %d trades\n", len(result))
}

func TestShields(t *testing.T) {
	var result []types.Shield
	RequestAndParseResponse(t, "/shields", &result)
	if len(result) == 0 {
		t.Fatal("expected to find shields")
	}

	t.Logf("found %d shields\n", len(result))
}

func TestGetItemHandler(t *testing.T) {
	expectedDesign := "weapon shelf lvl2"
	var result types.Blueprint
	RequestAndParseResponse(t, "/design/1", &result)

	if result.Name != expectedDesign {
		t.Fatalf("expected Name to be %v; got %v\n", expectedDesign, result.Name)
	}
}

func TestTradesForItem(t *testing.T) {
	expected := "alchemical gold"
	var response []types.Trade
	RequestAndParseResponse(t, "/trades/69", &response)
	if response == nil || len(response) == 0 {
		t.Fatal("No Trades found for item id 69")
	}
	if response[1].GetName != expected {
		t.Fatalf("expected Name to be %v; got %v\n", expected, response[1].GetName)
	}
}

func RequestAndParseResponse(t *testing.T, path string, result interface{}) {
	resp, err := http.Get(server.URL + path)
	if err != nil {
		t.Fatalf("error making request to server. Err: %v\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status OK; got %v\n", resp.Status)
	}

	bodyBytes, _ := io.ReadAll(resp.Body)
	fmt.Printf("Raw Response: %s\n", string(bodyBytes))

	parseErr := json.Unmarshal(bodyBytes, &result)
	if parseErr != nil {
		t.Fatalf("Error decoding response body. Err: %v\n", parseErr)
	}
}

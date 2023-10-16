package api_handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/querybuilder/model"
    "github.com/joho/godotenv"
	"os"
)


func TestPostHandlerValidInput(t *testing.T) {

	originalWorkingDir, _ := os.Getwd()
	  defer os.Chdir(originalWorkingDir)
  
	  testEnvDir := "/Users/b0a04gl/Documents/#BA6/query-builder-go/"
	  os.Chdir(testEnvDir)
  
	  if err := godotenv.Load(".env.test"); err != nil {
		  panic("Error loading .env.test file")
	  }
	requestBody := model.QueryRequest{
		Select:  []string{"id", "name", "team", "role"},
		From:    "players",
		Where:   []string{"batsman"},
		OrderBy: "total_wickets",
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest("POST", "/fetchPlayers", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	PostHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("doesnt match expected reesponse code")
	}

	var response model.QueryResponse
	err := json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Error("failed in decoding response")
	}
}

func TestPostHandlerInvalidInput(t *testing.T) {

	originalWorkingDir, _ := os.Getwd()
	  defer os.Chdir(originalWorkingDir)
  
	  testEnvDir := "/Users/b0a04gl/Documents/#BA6/query-builder-go/"
	  os.Chdir(testEnvDir)
  
	  if err := godotenv.Load(".env.test"); err != nil {
		  panic("Error loading .env.test file")
	  }

	requestBody := "Invalid JSON"
	body := []byte(requestBody)

	req := httptest.NewRequest("POST", "/fetchPlayers", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	PostHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Error("doesnt match expected reesponse code")
	}
    
}

func TestPostHandlerUnsupportedMethod(t *testing.T) {

	originalWorkingDir, _ := os.Getwd()
	  defer os.Chdir(originalWorkingDir)
  
	  testEnvDir := "/Users/b0a04gl/Documents/#BA6/query-builder-go/"
	  os.Chdir(testEnvDir)
  
	  if err := godotenv.Load(".env.test"); err != nil {
		  panic("Error loading .env.test file")
	  }

	req := httptest.NewRequest("GET", "/fetchPlayers", nil)
	rr := httptest.NewRecorder()

	PostHandler(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Error("doesnt match expected reesponse code")
	}
}

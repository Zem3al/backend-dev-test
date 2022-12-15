package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"dev/internal/app/handler"
	"dev/internal/persistence"
	"dev/internal/service"
)

func TestCreateHubInvalidRequest(t *testing.T) {

	//ARRANGE
	body := []map[string]interface{}{}
	data, _ := json.Marshal(body)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/hubs", bytes.NewBuffer(data))

	//ACTION
	handler.HubCreate(rr, req)

	//ASSERT
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestCreateHubInvalidPayload(t *testing.T) {

	//ARRANGE
	persistence.MockHubRepo()
	body := service.CreateHubCommand{
		Name:     "",
		Location: "",
	}
	data, _ := json.Marshal(body)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/hubs", bytes.NewBuffer(data))

	//ACTION
	handler.HubCreate(rr, req)

	//ASSERT
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestCreateHub(t *testing.T) {
	persistence.MockHubRepo()
	//ARRANGE
	body := service.CreateHubCommand{
		Name:     "Test",
		Location: "Test Location",
	}
	data, _ := json.Marshal(body)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/hubs", bytes.NewBuffer(data))

	//ACTION
	handler.HubCreate(rr, req)

	//ASSERT
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestSearch(t *testing.T) {
	persistence.MockHubRepo()
	//ARRANGE
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("Get", "/search?location=Test&&type=test", nil)

	//ACTION
	handler.Search(rr, req)

	//ASSERT
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

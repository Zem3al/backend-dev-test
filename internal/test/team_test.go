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

func TestCreateTeamInvalidRequest(t *testing.T) {

	//ARRANGE
	body := []map[string]interface{}{}
	data, _ := json.Marshal(body)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/team", bytes.NewBuffer(data))

	//ACTION
	handler.TeamCreate(rr, req)

	//ASSERT
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestCreateTeamInvalidPayload(t *testing.T) {
	//ARRANGE
	persistence.MockTeamRepo()
	body := service.CreateTeamCommand{
		Name:  "",
		Type:  "",
		HubID: "",
	}
	data, _ := json.Marshal(body)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/team", bytes.NewBuffer(data))

	//ACTION
	handler.TeamCreate(rr, req)

	//ASSERT
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestCreateTeam(t *testing.T) {
	persistence.MockTeamRepo()
	//ARRANGE
	body := service.CreateTeamCommand{
		Name:  "Test",
		Type:  "test",
		HubID: "test_hub_id",
	}
	data, _ := json.Marshal(body)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/teams", bytes.NewBuffer(data))

	//ACTION
	handler.TeamCreate(rr, req)

	//ASSERT
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

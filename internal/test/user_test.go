package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"dev/internal/app/handler"
	"dev/internal/persistence"
	"dev/internal/persistence/mock"
	"dev/internal/service"
)

func TestCreateUserInvalidRequest(t *testing.T) {

	//ARRANGE
	body := []map[string]interface{}{}
	data, _ := json.Marshal(body)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(data))

	//ACTION
	handler.UserCreate(rr, req)

	//ASSERT
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestCreateUserInvalidPayload(t *testing.T) {
	mock.NewMockUser()
	//ARRANGE
	body := service.CreateUserCommand{
		Name:   "",
		Age:    0,
		TeamID: "",
	}
	data, _ := json.Marshal(body)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(data))

	//ACTION
	handler.UserCreate(rr, req)

	//ASSERT
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

func TestCreateUser(t *testing.T) {
	persistence.MockUserRepo()
	//ARRANGE
	body := service.CreateUserCommand{
		Name:   "Test",
		Age:    2323,
		TeamID: "test_team_ID",
	}
	data, _ := json.Marshal(body)
	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(data))

	//ACTION
	handler.UserCreate(rr, req)

	//ASSERT
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code. Expected: %d. Got: %d.", http.StatusOK, status)
	}
}

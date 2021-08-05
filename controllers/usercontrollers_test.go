package controllers

import (
	"bytes"
	_ "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/get", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllUser)
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)

	}
	expected := `[{"id":1,"name":"test","balance":600000,"created_time":"12:00:00","modified_time":"12:00:00"},{"id":2,"name":"test2","balance":700000,"created_time":"12:00:00","modified_time":"12:00:00"},{"id":3,"name":"test3","balance":800000,"created_time":"12:00:00","modified_time":"12:00:00"}]`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected bot: got %v want %v", w.Body.String(), expected)
	}
}

func TestGetUserById(t *testing.T) {
	req, err := http.NewRequest("GET", "/get/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserByID)
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)

	}
	expected := `[{"id":1,"name":"test","balance":600000,"created_time":"12:00:00","modified_time":"12:00:00"}]`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected bot: got %v want %v", w.Body.String(), expected)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	req, err := http.NewRequest("GET", "/get/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "55555")
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserByID)
	handler.ServeHTTP(w, req)
	if status := w.Code; status == http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestUpdateUserById(t *testing.T) {
	var jsonStr = []byte(`{"id":1,"name":"test_change","balance":600000,"created_time":"12:00:00","modified_time":"12:00:00"}`)
	req, err := http.NewRequest("PUT", "/update/{id}", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserByID)
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)

	}
	expected := `[{"id":1,"name":"_change","balance":600000,"created_time":"12:00:00","modified_time":"12:00:00"}]`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected bot: got %v want %v", w.Body.String(), expected)
	}
}
func TestCreateUser(t *testing.T) {
	var jsonStr = []byte(`{"id":1,"name":"test","balance":600000,"created_time":"12:00:00","modified_time":"12:00:00"}`)
	req, err := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)

	}
	expected := `[{"id":1,"name":"test","balance":600000,"created_time":"12:00:00","modified_time":"12:00:00"}]`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected bot: got %v want %v", w.Body.String(), expected)
	}
}
func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/delete/{id}", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("id", "1")
	req.URL.RawQuery = q.Encode()
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteUser)
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)

	}
	expected := `[{"id":1,"name":"test_change","balance":600000,"created_time":"12:00:00","modified_time":"12:00:00"}]`
	if w.Body.String() != expected {
		t.Errorf("handler returned unexpected bot: got %v want %v", w.Body.String(), expected)
	}
}
func TestWithdraw(t *testing.T) {

	var jsonStr = []byte(`{"id":1,"name":"Lan","amount":5000,"targetID":1}`)

	req, err := http.NewRequest("PUT", "/withdraw", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UserWithdraw)
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":1,"name":"Lan","balance":130000,"created_time":"2006-01-02 15:04:05","modified_time":"2006-01-02 15:10:00"}`
	if strings.TrimSpace(w.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}
func TestDeposit(t *testing.T) {

	var jsonStr = []byte(`{"id":1,"name":"Lan","amount":5000,"targetID":1}`)

	req, err := http.NewRequest("PUT", "/deposit", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UserDeposit)
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":1,"name":"Lan","balance":115000,"created_time":"2006-01-02 15:04:05","modified_time":"2006-01-02 15:10:00"}`
	if strings.TrimSpace(w.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}

func TestTransfer(t *testing.T) {

	var jsonStr = []byte(`{"id":1,"name":"Lan","amount":5000,"targetID":2}`)

	req, err := http.NewRequest("PUT", "/transfer", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UserTransfer)
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":1,"name":"Lan","balance":120000,"created_time":"2006-01-02 15:04:05","modified_time":"2006-01-02 15:10:00"},{"id":2,"name":"Hoa","balance":155000,"created_time":"2006-01-02 15:04:05","modified_time":"2021-08-04 15:58:31"}]`
	if strings.TrimSpace(w.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}
func TestWithdraw(t *testing.T) {

	var jsonStr = []byte(`{"id":1,"name":"Lan","amount":5000,"targetID":1}`)

	req, err := http.NewRequest("PUT", "/withdraw", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UserWithdraw)
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":1,"name":"Lan","balance":130000,"created_time":"2006-01-02 15:04:05","modified_time":"2006-01-02 15:10:00"}`
	if strings.TrimSpace(w.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}
func TestDeposit(t *testing.T) {

	var jsonStr = []byte(`{"id":1,"name":"Lan","amount":5000,"targetID":1}`)

	req, err := http.NewRequest("PUT", "/deposit", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UserDeposit)
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":1,"name":"Lan","balance":115000,"created_time":"2006-01-02 15:04:05","modified_time":"2006-01-02 15:10:00"}`
	if strings.TrimSpace(w.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}

func TestTransfer(t *testing.T) {

	var jsonStr = []byte(`{"id":1,"name":"Lan","amount":5000,"targetID":2}`)

	req, err := http.NewRequest("PUT", "/transfer", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(UserTransfer)
	handler.ServeHTTP(w, req)
	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `[{"id":1,"name":"Lan","balance":120000,"created_time":"2006-01-02 15:04:05","modified_time":"2006-01-02 15:10:00"},{"id":2,"name":"Hoa","balance":155000,"created_time":"2006-01-02 15:04:05","modified_time":"2021-08-04 15:58:31"}]`
	if strings.TrimSpace(w.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			w.Body.String(), expected)
	}
}

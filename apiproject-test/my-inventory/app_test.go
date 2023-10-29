package main

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	err := a.Initialise(DbUser, DbPassword, "test")
	if err != nil {
		log.Fatal("Error initialising")
	}
	m.Run()

}

func TestGetProduct(t *testing.T) {

	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected %d, got %d", expectedStatusCode, actualStatusCode)
	}
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, request)
	return recorder

}

func TestCreateProduct(t *testing.T) {
	var product = []byte(`{"name":"table","quantity":1, "price":200}`)
	request, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))
	response := sendRequest(request)
	checkStatusCode(t, http.StatusCreated, response.Code)
}

func TestUpdateProduct(t *testing.T) {
	var product = []byte(`{"name":"table","quantity":2, "price":200}`)
	request, _ := http.NewRequest("PUT", "/product/8", bytes.NewBuffer(product))
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)
}

func TestDeleteProduct(t *testing.T) {
	request, _ := http.NewRequest("DELETE", "/product/8", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)

}

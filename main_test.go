package main

import (
	"github.com/nbari/violetear"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/* Test Helpers */
func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func expectDeepEqual(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func TestHealthcheck(t *testing.T) {
	router := violetear.New()
	router.LogRequests = false
	err := router.HandleFunc("/_healthcheck_", healthCheck)
	expect(t, err, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_healthcheck_", nil)
	router.ServeHTTP(w, req)
	expect(t, w.Code, 200)
}

package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetTodos(t *testing.T) {

	h := Handler{
		Todos: &[]Todo{
			{
				Name: "Test",
				Done: false,
			},
		},
	}
	handler := http.HandlerFunc(h.HandleTodos)

	r, err := http.NewRequest("GET","/",strings.NewReader(""))
	if err != nil {
		t.Errorf("error %v",err)
	}

	w := httptest.NewRecorder()
	handler.ServeHTTP(w,r)
	if w.Code != http.StatusOK{
		t.Errorf("ALARM!!! Expected %d, got %d",http.StatusOK,w.Code)
	}

	expected := `[{"name":"Test","done":false}]`

	if w.Body.String() != expected {
		t.Errorf(`expected %s, got %s`,expected,w.Body.String())
	}
	t.Logf("Expected %s, got %s",expected,w.Body.String())
}

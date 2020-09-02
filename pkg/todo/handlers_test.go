package todo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTodoIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/todos", nil)
	if err != nil {
		t.Fatalf("could not created request: %v", err)
	}

	rec := httptest.NewRecorder()

	TodoIndex(rec, req)

	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	var got []*Todo
	err = json.Unmarshal(b, &got)
	if err != nil {
		t.Fatalf("could not unmarshall response %v", err)
	}

	expected := len(todos)

	if len(got) != expected {
		t.Errorf("expected %v todos, got: %v todo", todos, got)
	}
}

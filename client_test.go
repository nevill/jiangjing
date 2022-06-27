package jiangjing

import (
	"encoding/json"
	"net/http"
	"testing"
)

const (
	address  = "http://localhost:3002"
	username = "elastic"
	password = "elastic823"
)

func TestGetHealth(t *testing.T) {
	client, err := NewClient(Config{
		Address:  address,
		Username: username,
		Password: password,
	})
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}

	resp, err := client.Health()
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expect to get status: %d, but got: %d\n", http.StatusOK, resp.StatusCode)
	}

	var r map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		t.Fatalf("Error parsing the response body: %s", err)
	}

	switch name := r["name"].(type) {
	case string:
		if len(name) == 0 {
			t.Fatal("Expect to have a field `name` in response.")
		}
	default:
		t.Fatal("Expect to have a field `name` in response.")
	}

	_, ok := r["version"]
	if !ok {
		t.Fatal("Expect to have a field `version` in response.")
	}
}

func TestEnginesList(t *testing.T) {
	client, err := NewClient(Config{
		Address:  address,
		Username: username,
		Password: password,
	})
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}

	resp, err := client.AppSearch.Engines.List()
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expect to get status: %d, but got: %d\n", http.StatusOK, resp.StatusCode)
	}

	var r map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		t.Fatalf("Error parsing the response body: %s", err)
	}

	_, ok := r["results"]
	if !ok {
		t.Fatal("Expect to have a field `results` in response.")
	}

	switch results := r["results"].(type) {
	case []interface{}:
		if len(results) < 1 {
			t.Fatal("Expect to have some engines returned in response.")
		}
	default:
		t.Fatal("Expect to have some engines returned in response.")
	}
}

func TestEnginesGet(t *testing.T) {
	client, err := NewClient(Config{
		Address:  address,
		Username: username,
		Password: password,
	})
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}

	name := "national-parks-demo"
	resp, err := client.AppSearch.Engines.Get(name)
	defer resp.Body.Close()
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expect to get status: %d, but got: %d\n", http.StatusOK, resp.StatusCode)
	}

	var r map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		t.Fatalf("Error parsing the response body: %s", err)
	}

	switch ename := r["name"].(type) {
	case string:
		if name != ename {
			t.Fatalf("Expect to have: %s, but got: %s in response.", name, ename)
		}
	default:
		t.Fatal("Expect to have a field `name` returned in response.")
	}
}

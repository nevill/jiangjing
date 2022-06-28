package jiangjing

import (
	"encoding/json"
	"net/http"
	"strings"
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

func TestEngines(t *testing.T) {
	client, err := NewClient(Config{
		Address:  address,
		Username: username,
		Password: password,
	})
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}

	name := "search-engine-testing"
	t.Run("create an engine", func(t *testing.T) {
		resp, err := client.AppSearch.Engines.Create(name)
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
	})

	t.Run("list engines", func(t *testing.T) {
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
	})

	t.Run("get an engine", func(t *testing.T) {
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
	})

	t.Run("delete an engine", func(t *testing.T) {
		resp, err := client.AppSearch.Engines.Delete(name)
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
	})
}

func TestSynonyms(t *testing.T) {
	client, err := NewClient(Config{
		Address:  address,
		Username: username,
		Password: password,
	})
	if err != nil {
		t.Fatalf("Unexpected error: %s\n", err)
	}

	engine := "search-synonyms-testings"

	{
		// setup testing engine
		resp, err := client.AppSearch.Engines.Create(engine)
		defer resp.Body.Close()
		if err != nil {
			t.Fatalf("Unexpected error: %s\n", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expect to get status: %d, but got: %d\n", http.StatusOK, resp.StatusCode)
		}
	}

	var synId string

	t.Run("create a synonyms set", func(t *testing.T) {
		synonyms := []string{
			"ring", "rings", "fingerring", "fingerrings",
		}
		resp, err := client.AppSearch.Synonyms.Create(engine, synonyms)
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

		switch id := r["id"].(type) {
		case string:
			if !strings.HasPrefix(id, "syn-") {
				t.Fatalf("Expect to have syn- as prefix in ID, but got: %s in response.", id)
			}
			synId = id
		default:
			t.Fatal("Expect to have a field `id` returned in response.")
		}

		switch set := r["synonyms"].(type) {
		case []interface{}:
			if len(set) != len(synonyms) {
				t.Fatalf("Expect to have: %s, but got: %s in response.", synonyms, set)
			}
		default:
			t.Fatal("Expect to have a field `synonyms` returned in response.")
		}
	})

	t.Run("update a synonyms set", func(t *testing.T) {
		synonyms := []string{
			"ring", "rings",
		}
		resp, err := client.AppSearch.Synonyms.Update(engine, synId, synonyms)
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

		switch id := r["id"].(type) {
		case string:
			if id != synId {
				t.Fatalf("Expect to have: %s, but got: %s in response.", synId, id)
			}
		default:
			t.Fatal("Expect to have a field `id` returned in response.")
		}

		switch set := r["synonyms"].(type) {
		case []interface{}:
			if len(set) != len(synonyms) {
				t.Fatalf("Expect to have: %s, but got: %s in response.", synonyms, set)
			}
		default:
			t.Fatal("Expect to have a field `synonyms` returned in response.")
		}
	})

	t.Run("Get a synonyms set", func(t *testing.T) {
		resp, err := client.AppSearch.Synonyms.Get(engine, synId)
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

		switch id := r["id"].(type) {
		case string:
			if synId != id {
				t.Fatalf("Expect to get id: %s, but got: %s.", synId, id)
			}
		default:
			t.Fatal("Expect to have a field `id` returned in response.")
		}

	})

	t.Run("list synonyms set", func(t *testing.T) {
		resp, err := client.AppSearch.Synonyms.List(engine)
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

		var synonymsSet interface{}
		switch results := r["results"].(type) {
		case []interface{}:
			if len(results) != 1 {
				t.Fatalf("Expect to return one synonyms set.")
			}
			synonymsSet = results[0]
		default:
			t.Fatal("Expect to have a field `results` returned in response.")
		}

		switch set := synonymsSet.(type) {
		case map[string]interface{}:
			if set["id"] != synId {
				t.Fatalf("Expect to have a synonyms id %s returned in response, but got: %s.", synId, set["id"])
			}
		default:
			t.Fatalf("Expect to have a synonyms map returned in response, but got: %s.", set)
		}
	})

	t.Run("delete a synonyms set", func(t *testing.T) {
		resp, err := client.AppSearch.Synonyms.Delete(engine, synId)
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

		switch deleted := r["deleted"].(type) {
		case bool:
			if !deleted {
				t.Fatalf("Cannot delete synonyms set.")
			}
		default:
			t.Fatal("Expect to have a field `deleted` returned in response.")
		}

	})

	{
		// remove testing engine
		resp, err := client.AppSearch.Engines.Delete(engine)
		defer resp.Body.Close()
		if err != nil {
			t.Fatalf("Unexpected error: %s\n", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expect to get status: %d, but got: %d\n", http.StatusOK, resp.StatusCode)
		}
	}
}

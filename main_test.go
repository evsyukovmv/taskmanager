package main

import (
	"bytes"
	"github.com/evsyukovmv/taskmanager/handlers"
	"github.com/evsyukovmv/taskmanager/postgres"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupServer() *httptest.Server {
	configureLogger()
	configurePostgres()
	return httptest.NewServer(handlers.NewRouter())
}

func clearDBTables() {
	postgres.DB().Exec("TRUNCATE comments RESTART IDENTITY CASCADE;")
	postgres.DB().Exec("TRUNCATE tasks RESTART IDENTITY CASCADE;")
	postgres.DB().Exec("TRUNCATE columns RESTART IDENTITY CASCADE;")
	postgres.DB().Exec("TRUNCATE projects RESTART IDENTITY CASCADE;")
}

func compareResponse(t *testing.T, resp *http.Response, expectedStatus int, expectedJson string) {
	if resp.StatusCode != expectedStatus {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			resp.StatusCode,
			expectedStatus,
		)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	if string(body) != expectedJson {
		t.Errorf(
			"unexpected response:\n \tgot %v\n\twant %v",
			string(body),
			expectedJson,
		)
	}
}

func postRequest(t *testing.T, server *httptest.Server, path string, data string) *http.Response {
	resp, err := http.Post(
		server.URL+path,
		"application/json",
		bytes.NewBuffer([]byte(data)),
	)

	if err != nil {
		t.Fatal(err)
	}
	return resp
}

func getRequest(t *testing.T, server *httptest.Server, path string) *http.Response {
	resp, err := http.Get(server.URL + path)

	if err != nil {
		t.Fatal(err)
	}
	return resp
}

func deleteRequest(t *testing.T, server *httptest.Server, path string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", server.URL+path, nil)

	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return resp
}

func TestProjects(t *testing.T) {
	server := setupServer()
	defer server.Close()
	defer clearDBTables()

	aProjectRequest := `{ "name": "AProject", "description": "AProjectDescription" }`
	bProjectRequest := `{ "name": "BProject", "description": "BProjectDescription" }`

	aProjectResponse := `{"id":2,"name":"AProject","description":"AProjectDescription"}`
	bProjectResponse := `{"id":1,"name":"BProject","description":"BProjectDescription"}`

	// Create B project
	resp := postRequest(t, server, "/projects", bProjectRequest)
	compareResponse(t, resp, http.StatusOK, bProjectResponse)

	// Get B project by id
	resp = getRequest(t, server, "/projects/1")
	compareResponse(t, resp, http.StatusOK, bProjectResponse)

	// Create A project and get all projects sorted alphabetically
	postRequest(t, server, "/projects", aProjectRequest)
	resp = getRequest(t, server, "/projects")
	compareResponse(t, resp, http.StatusOK, `[`+aProjectResponse+`,`+bProjectResponse+`]`)

	// Delete B project and get all projects without deleted B project
	resp = deleteRequest(t, server, "/projects/1")
	compareResponse(t, resp, http.StatusOK, bProjectResponse)
	resp = getRequest(t, server, "/projects")
	compareResponse(t, resp, http.StatusOK, `[`+aProjectResponse+`]`)

	aProjectColumnResponse := `{"id":2,"name":"Default","position":0,"project_id":2}`

	// Column must be automatically created for project
	resp = getRequest(t, server, "/projects/2/columns")
	compareResponse(t, resp, http.StatusOK, `[`+aProjectColumnResponse+`]`)
}

func TestColumns(t *testing.T) {
	server := setupServer()
	defer server.Close()
	defer clearDBTables()
}

func TestNotFound(t *testing.T) {
	configureLogger()
	configurePostgres()

	server := httptest.NewServer(handlers.NewRouter())
	defer server.Close()

	resp, err := http.Get(server.URL + "/not_found")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf(
			"unexpected status: got (%v) want (%v)",
			resp.StatusCode,
			http.StatusNotFound,
		)
	}
}

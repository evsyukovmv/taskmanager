package main

import (
	"bytes"
	"github.com/evsyukovmv/taskmanager/handlers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/evsyukovmv/taskmanager/services/projectsvc"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupServer() *httptest.Server {
	configureLogger()
	configurePostgres()
	configureServices()
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

func putRequest(t *testing.T, server *httptest.Server, path string, data string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest("PUT", server.URL+path, bytes.NewBuffer([]byte(data)))

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

	// Update B project
	bRenameRequest := `{ "name": "BProjectRenamed", "description": "BProjectDescription" }`
	bProjectResponse = `{"id":1,"name":"BProjectRenamed","description":"BProjectDescription"}`
	resp = putRequest(t, server, "/projects/1", bRenameRequest)
	compareResponse(t, resp, http.StatusOK, bProjectResponse)

	// Delete B project and get all projects without deleted B project
	resp = deleteRequest(t, server, "/projects/1")
	compareResponse(t, resp, http.StatusOK, bProjectResponse)
	resp = getRequest(t, server, "/projects")
	compareResponse(t, resp, http.StatusOK, `[`+aProjectResponse+`]`)

	aProjectColumnResponse := `{"id":2,"project_id":2,"name":"Default","position":0}`
	// Column must be automatically created for project
	resp = getRequest(t, server, "/projects/2/columns")
	compareResponse(t, resp, http.StatusOK, `[`+aProjectColumnResponse+`]`)
}

func TestColumns(t *testing.T) {
	server := setupServer()
	defer server.Close()
	defer clearDBTables()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	err := projectsvc.Create(project)
	if err != nil {
		t.Fatal(err)
	}

	columnRequest := `{ "name": "TestColumn" }`

	columnResponse := `{"id":2,"project_id":1,"name":"TestColumn","position":0}`
	defaultResponse := `{"id":1,"project_id":1,"name":"Default","position":1}`

	// Create column
	resp := postRequest(t, server, "/projects/1/columns", columnRequest)
	compareResponse(t, resp, http.StatusOK, columnResponse)

	// Get all columns sorted by priority
	resp = getRequest(t, server, "/projects/1/columns")
	compareResponse(t, resp, http.StatusOK, `[`+columnResponse+`,`+defaultResponse+`]`)

	// Get column by id
	resp = getRequest(t, server, "/projects/1/columns/2")
	compareResponse(t, resp, http.StatusOK, columnResponse)

	// Update column name
	updateRequest := `{ "name": "RenamedColumn" }`
	defaultResponse = `{"id":1,"project_id":1,"name":"RenamedColumn","position":1}`
	resp = putRequest(t, server, "/projects/1/columns/1", updateRequest)
	compareResponse(t, resp, http.StatusOK, defaultResponse)

	// Move column position
	moveRequest := `{ "position": 0 }`
	columnResponse = `{"id":1,"project_id":1,"name":"RenamedColumn","position":0}`
	resp = putRequest(t, server, "/projects/1/columns/1/move", moveRequest)
	compareResponse(t, resp, http.StatusOK, columnResponse)

	// Delete column and get all projects without deleted column
	resp = deleteRequest(t, server, "/projects/1/columns/1")
	compareResponse(t, resp, http.StatusOK, columnResponse)
	resp = getRequest(t, server, "/projects/1/columns")
	columnResponse = `{"id":2,"project_id":1,"name":"TestColumn","position":1}`
	compareResponse(t, resp, http.StatusOK, `[`+columnResponse+`]`)

	// Deleting last column is not allowed
	resp = deleteRequest(t, server, "/projects/1/columns/2")
	compareResponse(t, resp, http.StatusBadRequest, `{ error: "deleting last column is not allowed" }`)
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

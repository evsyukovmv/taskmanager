package main

import (
	"bytes"
	"fmt"
	"github.com/evsyukovmv/taskmanager/handlers"
	"github.com/evsyukovmv/taskmanager/models"
	"github.com/evsyukovmv/taskmanager/postgres"
	"github.com/evsyukovmv/taskmanager/services/columnsvc"
	"github.com/evsyukovmv/taskmanager/services/projectsvc"
	"github.com/evsyukovmv/taskmanager/services/tasksvc"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

type testRequestResponse struct {
	message            string
	requestMethod      string
	requestPath        string
	requestData        string
	responseCode       int
	responseData       string
	responseSkipFields []string
}

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

func verifyResponse(t *testing.T, server *httptest.Server, trr testRequestResponse) error {
	resp := makeRequest(t, server, trr.requestMethod, trr.requestPath, trr.requestData)
	if resp.StatusCode != trr.responseCode {
		return fmt.Errorf(
			"%s\n unexpected status:\n \tgot %d\n\twant %d",
			trr.message,
			resp.StatusCode,
			trr.responseCode,
		)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatal(err)
	}

	sBody := string(body)
	for _, field := range trr.responseSkipFields {
		reg := regexp.MustCompile(`"` + field + `":"[\d\w\s-:\.]+"\,`)
		sBody = reg.ReplaceAllString(sBody, "")
	}

	if sBody != trr.responseData {
		return fmt.Errorf(
			"%s\n unexpected response:\n \tgot %v\n\twant %v",
			trr.message,
			sBody,
			trr.responseData,
		)
	}

	return nil
}

func makeRequest(t *testing.T, server *httptest.Server, method string, path string, data string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(method, server.URL+path, bytes.NewBuffer([]byte(data)))

	if err != nil {
		t.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return resp
}

var projectsTestData = [...]testRequestResponse{
	{
		message:       "should create project",
		requestMethod: "POST",
		requestPath:   "/projects",
		requestData:   `{ "name": "BProject", "description": "BProjectDescription" }`,
		responseCode:  http.StatusOK,
		responseData:  `{"id":1,"name":"BProject","description":"BProjectDescription"}`,
	},
	{
		message:       "should return project by id",
		requestMethod: "GET",
		requestPath:   "/projects/1",
		responseCode:  http.StatusOK,
		responseData:  `{"id":1,"name":"BProject","description":"BProjectDescription"}`,
	},
	{
		message:       "should create another project",
		requestMethod: "POST",
		requestPath:   "/projects",
		requestData:   `{ "name": "AProject", "description": "AProjectDescription" }`,
		responseCode:  http.StatusOK,
		responseData:  `{"id":2,"name":"AProject","description":"AProjectDescription"}`,
	},
	{
		message:       "should return projects sorted by name",
		requestMethod: "GET",
		requestPath:   "/projects",
		responseCode:  http.StatusOK,
		responseData:  `[{"id":2,"name":"AProject","description":"AProjectDescription"},{"id":1,"name":"BProject","description":"BProjectDescription"}]`,
	},
	{
		message:       "should update project name",
		requestMethod: "PUT",
		requestPath:   "/projects/1",
		requestData:   `{ "name": "BProjectRenamed", "description": "BProjectDescription" }`,
		responseCode:  http.StatusOK,
		responseData:  `{"id":1,"name":"BProjectRenamed","description":"BProjectDescription"}`,
	},
	{
		message:       "should delete project",
		requestMethod: "DELETE",
		requestPath:   "/projects/1",
		responseCode:  http.StatusOK,
		responseData:  `{"id":1,"name":"BProjectRenamed","description":"BProjectDescription"}`,
	},
	{
		message:       "should return projects without deleted",
		requestMethod: "GET",
		requestPath:   "/projects",
		responseCode:  http.StatusOK,
		responseData:  `[{"id":2,"name":"AProject","description":"AProjectDescription"}]`,
	},
}

func TestProjects(t *testing.T) {
	server := setupServer()
	defer server.Close()
	defer clearDBTables()

	for _, testData := range projectsTestData {
		if err := verifyResponse(t, server, testData); err != nil {
			t.Error(err)
			break
		}
	}
}

var columnsTestData = [...]testRequestResponse{
	{
		message:       "should create column",
		requestMethod: "POST",
		requestPath:   "/projects/1/columns",
		requestData:   `{ "name": "TestColumn" }`,
		responseCode:  http.StatusOK,
		responseData:  `{"id":2,"project_id":1,"name":"TestColumn","position":0}`,
	},
	{
		message:       "should return columns with default column created automatically sorted by priority",
		requestMethod: "GET",
		requestPath:   "/projects/1/columns",
		responseCode:  http.StatusOK,
		responseData:  `[{"id":2,"project_id":1,"name":"TestColumn","position":0},{"id":1,"project_id":1,"name":"default","position":1}]`,
	},
	{
		message:       "should return column by id",
		requestMethod: "GET",
		requestPath:   "/projects/1/columns/2",
		responseCode:  http.StatusOK,
		responseData:  `{"id":2,"project_id":1,"name":"TestColumn","position":0}`,
	},
	{
		message:       "should update columns name",
		requestMethod: "PUT",
		requestPath:   "/projects/1/columns/2",
		requestData:   `{ "name": "RenamedColumn" }`,
		responseCode:  http.StatusOK,
		responseData:  `{"id":2,"project_id":1,"name":"RenamedColumn","position":0}`,
	},
	{
		message:       "should move column position",
		requestMethod: "PUT",
		requestPath:   "/projects/1/columns/2/move",
		requestData:   `{ "position": 1 }`,
		responseCode:  http.StatusOK,
		responseData:  `{"id":2,"project_id":1,"name":"RenamedColumn","position":1}`,
	},
	{
		message:       "should delete column by id",
		requestMethod: "DELETE",
		requestPath:   "/projects/1/columns/2",
		responseCode:  http.StatusOK,
		responseData:  `{"id":2,"project_id":1,"name":"RenamedColumn","position":1}`,
	},
	{
		message:       "should return columns without deleted",
		requestMethod: "GET",
		requestPath:   "/projects/1/columns",
		responseCode:  http.StatusOK,
		responseData:  `[{"id":1,"project_id":1,"name":"default","position":0}]`,
	},
	{
		message:       "should return error when delete first column",
		requestMethod: "DELETE",
		requestPath:   "/projects/1/columns/1",
		responseCode:  http.StatusBadRequest,
		responseData:  `{ error: "deleting the first column is not allowed" }`,
	},
}

func TestColumns(t *testing.T) {
	server := setupServer()
	defer server.Close()
	defer clearDBTables()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	err := projectsvc.Create(project)
	if err != nil {
		t.Fatal(err)
		return
	}

	for _, testData := range columnsTestData {
		if err := verifyResponse(t, server, testData); err != nil {
			t.Error(err)
			break
		}
	}
}

var tasksTestData = [...]testRequestResponse{
	{
		message:       "should create task",
		requestMethod: "POST",
		requestPath:   "/projects/1/columns/2/tasks",
		requestData:   `{ "name": "TestTask", "description": "TestTask description" }`,
		responseCode:  http.StatusOK,
		responseData:  `{"id":1,"column_id":2,"name":"TestTask","description":"TestTask description","position":0}`,
	},
	{
		message:       "should get task by id",
		requestMethod: "GET",
		requestPath:   "/projects/1/columns/2/tasks/1",
		responseCode:  http.StatusOK,
		responseData:  `{"id":1,"column_id":2,"name":"TestTask","description":"TestTask description","position":0}`,
	},
	{
		message:       "should create another task",
		requestMethod: "POST",
		requestPath:   "/projects/1/columns/2/tasks",
		requestData:   `{ "name": "TestTask2", "description": "TestTask2 description" }`,
		responseCode:  http.StatusOK,
		responseData:  `{"id":2,"column_id":2,"name":"TestTask2","description":"TestTask2 description","position":0}`,
	},
	{
		message:       "should get tasks list",
		requestMethod: "GET",
		requestPath:   "/projects/1/columns/2/tasks",
		responseCode:  http.StatusOK,
		responseData:  `[{"id":2,"column_id":2,"name":"TestTask2","description":"TestTask2 description","position":0},{"id":1,"column_id":2,"name":"TestTask","description":"TestTask description","position":1}]`,
	},
	{
		message:       "should update task",
		requestMethod: "PUT",
		requestPath:   "/projects/1/columns/2/tasks/1",
		requestData:   `{ "name": "TestTask1", "description": "TestTask1 description" }`,
		responseCode:  http.StatusOK,
		responseData:  `{"id":1,"column_id":2,"name":"TestTask1","description":"TestTask1 description","position":1}`,
	},
	{
		message:       "should move task",
		requestMethod: "PUT",
		requestPath:   "/projects/1/columns/2/tasks/1/move",
		requestData:   `{ "position": 0 }`,
		responseCode:  http.StatusOK,
		responseData:  `{"id":1,"column_id":2,"name":"TestTask1","description":"TestTask1 description","position":0}`,
	},
	{
		message:       "should remove task",
		requestMethod: "DELETE",
		requestPath:   "/projects/1/columns/2/tasks/2",
		responseCode:  http.StatusOK,
		responseData:  `{"id":2,"column_id":2,"name":"TestTask2","description":"TestTask2 description","position":1}`,
	},
	{
		message:       "should create task in another column",
		requestMethod: "POST",
		requestPath:   "/projects/1/columns/1/tasks",
		requestData:   `{ "name": "TestTaskColumn1" }`,
		responseCode:  http.StatusOK,
		responseData:  `{"id":3,"column_id":1,"name":"TestTaskColumn1","description":"","position":0}`,
	},
	{
		message:       "should delete right column",
		requestMethod: "DELETE",
		requestPath:   "/projects/1/columns/1",
		responseCode:  http.StatusOK,
		responseData:  `{"id":1,"project_id":1,"name":"default","position":1}`,
	},
	{
		message:       "should append tasks from right deleted column to left",
		requestMethod: "GET",
		requestPath:   "/projects/1/columns/2/tasks",
		responseCode:  http.StatusOK,
		responseData:  `[{"id":1,"column_id":2,"name":"TestTask1","description":"TestTask1 description","position":0},{"id":3,"column_id":2,"name":"TestTaskColumn1","description":"","position":1}]`,
	},
}

func TestTasks(t *testing.T) {
	server := setupServer()
	defer server.Close()
	defer clearDBTables()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "TestProject"}}
	if err := projectsvc.Create(project); err != nil {
		t.Fatal(err)
		return
	}

	column := &models.Column{ColumnBase: models.ColumnBase{Name: "TestColumn"}, ProjectId: project.Id}
	if err := columnsvc.Create(column); err != nil {
		t.Fatal(err)
		return
	}

	for _, testData := range tasksTestData {
		if err := verifyResponse(t, server, testData); err != nil {
			t.Error(err)
			break
		}
	}
}

var commentsTestData = [...]testRequestResponse{
	{
		message:            "should create comment",
		requestMethod:      "POST",
		requestPath:        "/projects/1/columns/1/tasks/1/comments",
		requestData:        `{ "text": "TestComment1" }`,
		responseCode:       http.StatusOK,
		responseData:       `{"id":1,"task_id":1,"text":"TestComment1"}`,
		responseSkipFields: []string{"created_at"},
	},
	{
		message:            "should get comment by id",
		requestMethod:      "GET",
		requestPath:        "/projects/1/columns/1/tasks/1/comments/1",
		responseCode:       http.StatusOK,
		responseData:       `{"id":1,"task_id":1,"text":"TestComment1"}`,
		responseSkipFields: []string{"created_at"},
	},
	{
		message:            "should create another comment",
		requestMethod:      "POST",
		requestPath:        "/projects/1/columns/1/tasks/1/comments",
		requestData:        `{ "text": "TestComment2" }`,
		responseCode:       http.StatusOK,
		responseData:       `{"id":2,"task_id":1,"text":"TestComment2"}`,
		responseSkipFields: []string{"created_at"},
	},
	{
		message:            "should get comments list sorted by creation date",
		requestMethod:      "GET",
		requestPath:        "/projects/1/columns/1/tasks/1/comments",
		responseCode:       http.StatusOK,
		responseData:       `[{"id":2,"task_id":1,"text":"TestComment2"},{"id":1,"task_id":1,"text":"TestComment1"}]`,
		responseSkipFields: []string{"created_at"},
	},
	{
		message:            "should update comment",
		requestMethod:      "PUT",
		requestPath:        "/projects/1/columns/1/tasks/1/comments/1",
		requestData:        `{ "text": "UpdatedComment1" }`,
		responseCode:       http.StatusOK,
		responseData:       `{"id":1,"task_id":1,"text":"UpdatedComment1"}`,
		responseSkipFields: []string{"created_at"},
	},
	{
		message:            "should remove comment",
		requestMethod:      "DELETE",
		requestPath:        "/projects/1/columns/1/tasks/1/comments/1",
		responseCode:       http.StatusOK,
		responseData:       `{"id":1,"task_id":1,"text":"UpdatedComment1"}`,
		responseSkipFields: []string{"created_at"},
	},
	{
		message:            "should get comments list without deleted",
		requestMethod:      "GET",
		requestPath:        "/projects/1/columns/1/tasks/1/comments",
		responseCode:       http.StatusOK,
		responseData:       `[{"id":2,"task_id":1,"text":"TestComment2"}]`,
		responseSkipFields: []string{"created_at"},
	},
}

func TestComments(t *testing.T) {
	server := setupServer()
	defer server.Close()
	defer clearDBTables()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "TestProject"}}
	if err := projectsvc.Create(project); err != nil {
		t.Fatal(err)
		return
	}

	task := &models.Task{TaskBase: models.TaskBase{Name: "TestTask"}, ColumnId: 1}
	if err := tasksvc.Create(task); err != nil {
		t.Fatal(err)
		return
	}

	for _, testData := range commentsTestData {
		if err := verifyResponse(t, server, testData); err != nil {
			t.Error(err)
			break
		}
	}
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

package services

import (
	//"fmt"
	"github.com/evsyukovmv/taskmanager/models"
	"testing"
)

func TestProjectCreate(t *testing.T) {
	setupTests()
	defer clearTests()

	// Should return error if model invalid
	err := ForProject().Create(&models.Project{})
	expectedError := "Key: 'Project.ProjectBase.Name' Error:Field validation for 'Name' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should create if project valid
	err = ForProject().Create(&models.Project{ProjectBase: models.ProjectBase{Name: "Test"}})
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
}

func TestProjectDelete(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	err := ForProject().Create(project)
	if err != nil {
		t.Error(err.Error())
	}

	// Should return error if project doesn't exist
	_, err = ForProject().Delete(project.Id + 1000)
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should delete if project exists
	_, err = ForProject().Delete(project.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
}

func TestProjectGetById(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	err := ForProject().Create(project)
	if err != nil {
		t.Error(err.Error())
	}

	// Should return error if project doesn't exist
	_, err = ForProject().GetById(project.Id + 1000)
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should return project if exists
	result, err := ForProject().GetById(project.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
	if result.Id != project.Id {
		t.Error("expected:", project.Id, "got:", result.Id)
	}
}

func TestProjectGetList(t *testing.T) {
	setupTests()
	defer clearTests()

	// Should return empty array if projects are empty
	result, err := ForProject().GetList()
	if len(*result) != 0 {
		t.Error("expected:", 0, "got:", len(*result))
	}

	project1 := &models.Project{ProjectBase: models.ProjectBase{Name: "Test1"}}
	err = ForProject().Create(project1)
	if err != nil {
		t.Error(err.Error())
	}

	project2 := &models.Project{ProjectBase: models.ProjectBase{Name: "Test2"}}
	err = ForProject().Create(project2)
	if err != nil {
		t.Error(err.Error())
	}

	// Should return all projects
	result, err = ForProject().GetList()
	if len(*result) != 2 {
		t.Error("expected:", 2, "got:", len(*result))
	}
}

func TestProjectUpdate(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	err := ForProject().Create(project)
	if err != nil {
		t.Error(err.Error())
	}

	// Should return error if project doesn't exist
	_, err = ForProject().Update(project.Id+1000, &models.ProjectBase{})
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should return error if field for update are invalid
	_, err = ForProject().Update(project.Id, &models.ProjectBase{})
	expectedError = "Key: 'Project.ProjectBase.Name' Error:Field validation for 'Name' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should update project
	expectedName := "Updated"
	result, err := ForProject().Update(project.Id, &models.ProjectBase{Name: expectedName})
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
	if result.Name != expectedName {
		t.Error("expected:", expectedName, "got:", result.Name)
	}
}

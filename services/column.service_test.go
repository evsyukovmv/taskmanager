package services

import (
	"github.com/evsyukovmv/taskmanager/models"
	"testing"
)

func TestColumnCreate(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	// Should return error if model invalid
	err := ForColumn().Create(&models.Column{ProjectId: project.Id})
	expectedError := "Key: 'Column.ColumnBase.Name' Error:Field validation for 'Name' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should create if valid
	err = ForColumn().Create(&models.Column{ProjectId: project.Id, ColumnBase: models.ColumnBase{Name: "Test"}})
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
}

func TestColumnDelete(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	column := &models.Column{ProjectId: project.Id, ColumnBase: models.ColumnBase{Name: "Test"}}
	if err := ForColumn().Create(column); err != nil {
		t.Error(err.Error())
	}

	// Should return error if column doesn't exist
	_, err := ForColumn().Delete(column.Id + 100)
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should delete column if column exists and not the last
	_, err = ForColumn().Delete(column.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}

	// Shouldn't allow delete last column
	expectedError = "deleting the last column is not allowed"
	_, err = ForColumn().Delete(1)
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}
}

func TestColumnGetById(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	column := &models.Column{ProjectId: project.Id, ColumnBase: models.ColumnBase{Name: "Test"}}
	if err := ForColumn().Create(column); err != nil {
		t.Error(err.Error())
	}

	// Should return error if column doesn't exist
	_, err := ForColumn().GetById(column.Id + 100)
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should return column
	result, err := ForColumn().GetById(column.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}

	if result.Id != column.Id {
		t.Error("expected:", column.Id, "got:", result.Id)
	}
}

func TestColumnGetListByProjectId(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	// Should return automatically created column in the list
	result, err := ForColumn().GetListByProjectId(project.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err)
	}
	if len(*result) != 1 {
		t.Error("expected:", 1, "got:", len(*result))
	}

	column := &models.Column{ProjectId: project.Id, ColumnBase: models.ColumnBase{Name: "Test"}}
	if err := ForColumn().Create(column); err != nil {
		t.Error(err.Error())
	}

	// Should return all columns for the project
	result, err = ForColumn().GetListByProjectId(project.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err)
	}
	if len(*result) != 2 {
		t.Error("expected:", 2, "got:", len(*result))
	}
}

func TestColumnMove(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	column := &models.Column{ProjectId: project.Id, ColumnBase: models.ColumnBase{Name: "Test"}}
	if err := ForColumn().Create(column); err != nil {
		t.Error(err.Error())
	}

	// Should return error if column doesn't exist
	_, err := ForColumn().Move(column.Id+100, &models.ColumnPosition{Position: column.Position + 1})
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should return error if column position is less or more then allowed
	_, err = ForColumn().Move(column.Id, &models.ColumnPosition{Position: column.Position + 1000})
	expectedError = "position must be more or eq 0 and less than 1"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should move column position
	result, err := ForColumn().Move(column.Id, &models.ColumnPosition{Position: column.Position + 1})
	if err != nil {
		t.Error("expected:", nil, "got:", err)
	}
	if result.Position != column.Position+1 {
		t.Error("expected:", column.Position+1, "got:", result.Position)
	}
}

func TestColumnUpdate(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	column := &models.Column{ProjectId: project.Id, ColumnBase: models.ColumnBase{Name: "Test"}}
	if err := ForColumn().Create(column); err != nil {
		t.Error(err.Error())
	}

	// Should return error if column doesn't exist
	_, err := ForColumn().Update(column.Id+100, &models.ColumnBase{})
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should return error if column data invalid
	_, err = ForColumn().Update(column.Id, &models.ColumnBase{})
	expectedError = "Key: 'Column.ColumnBase.Name' Error:Field validation for 'Name' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should update column
	expectedName := "Updated"
	result, err := ForColumn().Update(column.Id, &models.ColumnBase{Name: expectedName})
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
	if result.Name != expectedName {
		t.Error("expected:", expectedName, "got:", result.Name)
	}
}

func TestColumnIsSameProject(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	column := &models.Column{ProjectId: project.Id, ColumnBase: models.ColumnBase{Name: "Test"}}
	if err := ForColumn().Create(column); err != nil {
		t.Error(err.Error())
	}

	// Should return true if same projects
	result, err := ForColumn().IsSameProject(1, column.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
	if result != true {
		t.Error("expected:", true, "got:", result)
	}

	anotherProject := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(anotherProject); err != nil {
		t.Error(err.Error())
	}

	anotherProjectColumn := &models.Column{ProjectId: anotherProject.Id, ColumnBase: models.ColumnBase{Name: "Test"}}
	if err := ForColumn().Create(anotherProjectColumn); err != nil {
		t.Error(err.Error())
	}

	result, err = ForColumn().IsSameProject(1, column.Id, anotherProjectColumn.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
	if result != false {
		t.Error("expected:", false, "got:", result)
	}
}

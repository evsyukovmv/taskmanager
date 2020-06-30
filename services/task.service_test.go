package services

import (
	"github.com/evsyukovmv/taskmanager/models"
	"testing"
)

func TestTaskCreate(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{ Name: "Test" }}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	defaultColumn, err := ForColumn().GetById(1)
	if err != nil {
		t.Error(err.Error())
	}

	// Should return error if model invalid
	err = ForTask().Create(&models.Task{ TaskColumn: models.TaskColumn{ColumnId: defaultColumn.Id }})
	expectedError := "Key: 'Task.TaskBase.Name' Error:Field validation for 'Name' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should create if valid
	err = ForTask().Create(&models.Task{ TaskBase: models.TaskBase{ Name: "Test" }, TaskColumn: models.TaskColumn{ColumnId: defaultColumn.Id }})
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
}

func TestTaskDelete(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{ Name: "Test" }}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	task := &models.Task{ TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task); err != nil {
		t.Error(err.Error())
	}

	// Should return error if doesn't exist
	_, err := ForTask().Delete(task.Id + 100)
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should delete if exists
	_, err = ForTask().Delete(task.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
}

func TestTaskGetById(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{ Name: "Test" }}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	task := &models.Task{ TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task); err != nil {
		t.Error(err.Error())
	}

	// Should return error if doesn't exist
	_, err := ForTask().GetById(task.Id + 100)
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should return task
	result, err := ForTask().GetById(task.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}

	if result.Id != task.Id {
		t.Error("expected:", task.Id, "got:", result.Id)
	}
}

func TestTaskGetListByColumnId(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{ Name: "Test" }}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	task1 := &models.Task{ TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task1); err != nil {
		t.Error(err.Error())
	}

	task2 := &models.Task{ TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task2); err != nil {
		t.Error(err.Error())
	}

	// Should return all tasks for the column
	result, err := ForTask().GetListByColumnId(task1.ColumnId)
	if err != nil {
		t.Error("expected:", nil, "got:", err)
	}
	if len(*result) != 2 {
		t.Error("expected:", 2, "got:", len(*result))
	}
}

func TestTaskUpdate(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{ Name: "Test" }}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	task := &models.Task{ TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task); err != nil {
		t.Error(err.Error())
	}

	// Should return error if doesn't exist
	_, err := ForTask().Update(task.Id + 100, &models.TaskBase{})
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should return error if data invalid
	_, err = ForTask().Update(task.Id, &models.TaskBase{})
	expectedError = "Key: 'Task.TaskBase.Name' Error:Field validation for 'Name' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should update column
	expectedName := "Updated"
	result, err := ForTask().Update(task.Id, &models.TaskBase{Name: expectedName })
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
	if result.Name != expectedName {
		t.Error("expected:", expectedName, "got:", result.Name)
	}
}

func TestTaskMove(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{ Name: "Test" }}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	task1 := &models.Task{ TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task1); err != nil {
		t.Error(err.Error())
	}

	task2 := &models.Task{ TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task2); err != nil {
		t.Error(err.Error())
	}

	// Should return error if doesn't exist
	_, err := ForTask().Move(task2.Id + 100, &models.TaskPosition{Position: task2.Position + 1 })
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should return error if position is less or more then allowed
	_, err = ForTask().Move(task2.Id, &models.TaskPosition{Position: task2.Position + 1000 })
	expectedError = "position must be more or eq 0 and less than 1"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should move position
	result, err := ForTask().Move(task2.Id, &models.TaskPosition{Position: task2.Position + 1 })
	if err != nil {
		t.Error("expected:", nil, "got:", err)
	}
	if result.Position != task2.Position + 1 {
		t.Error("expected:", task2.Position + 1, "got:", result.Position)
	}
}

func TestTaskShift(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{ Name: "Test" }}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	anotherProject := &models.Project{ProjectBase: models.ProjectBase{ Name: "Test" }}
	if err := ForProject().Create(anotherProject); err != nil {
		t.Error(err.Error())
	}

	newColumn := &models.Column{ProjectId: project.Id, ColumnBase: models.ColumnBase{ Name: "Test" }}
	if err := ForColumn().Create(newColumn); err != nil {
		t.Error(err.Error())
	}

	anotherProjectColumn := &models.Column{ProjectId: anotherProject.Id, ColumnBase: models.ColumnBase{ Name: "Test" }}
	if err := ForColumn().Create(anotherProjectColumn); err != nil {
		t.Error(err.Error())
	}

	task := &models.Task{ TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task); err != nil {
		t.Error(err.Error())
	}


	// Should return error if doesn't exist
	_, err := ForTask().Shift(task.Id + 100, &models.TaskColumn{ColumnId: newColumn.Id })
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should return error if column invalid
	_, err = ForTask().Shift(task.Id, &models.TaskColumn{ColumnId: anotherProjectColumn.Id })
	expectedError = "columns must be in the same project"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should shift to another column
	result, err := ForTask().Shift(task.Id, &models.TaskColumn{ColumnId: newColumn.Id })
	if err != nil {
		t.Error("expected:", nil, "got:", err)
	}
	if result.ColumnId != newColumn.Id {
		t.Error("expected:", newColumn.Id, "got:", result.ColumnId)
	}
}

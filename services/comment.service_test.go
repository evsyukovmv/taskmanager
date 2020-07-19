package services

import (
	"github.com/evsyukovmv/taskmanager/models"
	"testing"
)

func TestCommentCreate(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	task := &models.Task{TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task); err != nil {
		t.Error(err.Error())
	}

	// Should return error if model invalid
	err := ForComment().Create(&models.Comment{})
	expectedError := "Key: 'Comment.CommentBase.Text' Error:Field validation for 'Text' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should create if valid
	err = ForComment().Create(&models.Comment{TaskId: task.Id, CommentBase: models.CommentBase{Text: "Test"}})
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
}

func TestCommentDelete(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	task := &models.Task{TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task); err != nil {
		t.Error(err.Error())
	}

	comment := &models.Comment{TaskId: task.Id, CommentBase: models.CommentBase{Text: "Test"}}
	if err := ForComment().Create(comment); err != nil {
		t.Error(err.Error())
	}

	// Should return error if doesn't exist
	_, err := ForComment().Delete(comment.Id + 100)
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should delete if exists
	_, err = ForComment().Delete(comment.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
}

func TestCommentGetById(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	task := &models.Task{TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task); err != nil {
		t.Error(err.Error())
	}

	comment := &models.Comment{TaskId: task.Id, CommentBase: models.CommentBase{Text: "Test"}}
	if err := ForComment().Create(comment); err != nil {
		t.Error(err.Error())
	}

	// Should return error if doesn't exist
	_, err := ForComment().GetById(comment.Id + 100)
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should return comment
	result, err := ForComment().GetById(comment.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}

	if result.Id != comment.Id {
		t.Error("expected:", task.Id, "got:", result.Id)
	}
}

func TestCommentGetListByTaskId(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	task := &models.Task{TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task); err != nil {
		t.Error(err.Error())
	}

	comment1 := &models.Comment{TaskId: task.Id, CommentBase: models.CommentBase{Text: "Test"}}
	if err := ForComment().Create(comment1); err != nil {
		t.Error(err.Error())
	}

	comment2 := &models.Comment{TaskId: task.Id, CommentBase: models.CommentBase{Text: "Test"}}
	if err := ForComment().Create(comment2); err != nil {
		t.Error(err.Error())
	}

	// Should return all comments for the task
	result, err := ForComment().GetListByTaskId(task.Id)
	if err != nil {
		t.Error("expected:", nil, "got:", err)
	}
	if len(*result) != 2 {
		t.Error("expected:", 2, "got:", len(*result))
	}
}

func TestCommentUpdate(t *testing.T) {
	setupTests()
	defer clearTests()

	project := &models.Project{ProjectBase: models.ProjectBase{Name: "Test"}}
	if err := ForProject().Create(project); err != nil {
		t.Error(err.Error())
	}

	task := &models.Task{TaskBase: models.TaskBase{Name: "Test"}, TaskColumn: models.TaskColumn{ColumnId: 1}}
	if err := ForTask().Create(task); err != nil {
		t.Error(err.Error())
	}

	comment := &models.Comment{TaskId: task.Id, CommentBase: models.CommentBase{Text: "Test"}}
	if err := ForComment().Create(comment); err != nil {
		t.Error(err.Error())
	}

	// Should return error if doesn't exist
	_, err := ForComment().Update(comment.Id+100, &models.CommentBase{})
	expectedError := "sql: no rows in result set"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should return error if data invalid
	_, err = ForComment().Update(comment.Id, &models.CommentBase{})
	expectedError = "Key: 'Comment.CommentBase.Text' Error:Field validation for 'Text' failed on the 'required' tag"
	if err.Error() != expectedError {
		t.Error("expected:", expectedError, "got:", err.Error())
	}

	// Should update column
	expectedText := "Updated"
	result, err := ForComment().Update(comment.Id, &models.CommentBase{Text: expectedText})
	if err != nil {
		t.Error("expected:", nil, "got:", err.Error())
	}
	if result.Text != expectedText {
		t.Error("expected:", expectedText, "got:", result.Text)
	}
}

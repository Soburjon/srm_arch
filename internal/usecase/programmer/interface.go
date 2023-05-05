package programmer

import (
	"srm_arch/internal/entity"
	"srm_arch/internal/service/programmer"
)

type ProgrammerService interface {
	CreateTask(req entity.Task) error
	EditTask(req programmer.EditTaskRequest) error
	DeleteTask(taskID string) error
	UpdateTaskStatus(status string, id string) error
	GetProjectTasks(projectID string) ([]entity.Task, error)
	GetTask(taskID string) (entity.Task, error)
	GetMyTasks(req programmer.GetMyTaskRequest) ([]entity.Task, error)
	CreateCommit(req entity.Comment) error
	EditCommit(req programmer.EditCommitRequest, userID string) error
	DeleteCommit(createdAt string, userID string) error
	GetCommits(taskID string) ([]entity.Comment, error)
	CreateAttendance(req entity.Attendince) error
	UsersInProject(projectID string) ([]entity.ProjectsPeople, error)
	UserRoleInProject(req programmer.UserRoleInProjectRequest) (string, error)
	MyProjects(userID string) ([]entity.Project, error)
}

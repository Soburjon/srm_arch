package programmer

import (
	"srm_arch/internal/entity"
)

type ProgrammerRepo interface {
	CreateTask(req entity.Task) error
	EditTask(req EditTaskRequest) error
	DeleteTask(taskID string) error
	UpdateTaskStatus(status string, id string) error
	GetProjectTasks(projectID string) ([]entity.Task, error)
	GetTask(taskID string) (entity.Task, error)
	GetMyTasks(req GetMyTaskRequest) ([]entity.Task, error)
	CreateCommit(req entity.Comment) error
	EditCommit(req EditCommitRequest, userID string) error
	DeleteCommit(createdAt string, userID string) error
	GetCommits(taskID string) ([]entity.Comment, error)
	CreateAttendance(req entity.Attendince) error
	UsersInProject(projectID string) ([]entity.ProjectsPeople, error)
	UserRoleInProject(req UserRoleInProjectRequest) (string, error)
	MyProjects(userID string) ([]entity.Project, error)
}

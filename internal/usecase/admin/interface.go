package admin

import (
	"srm_arch/internal/entity"
	"srm_arch/internal/service/admin"
)

type AdminService interface {
	CreateUser(req entity.Users) error
	EditUser(req admin.EditUserRequest) error
	DeleteUser(userID string) error
	CreateProject(req entity.Project) error
	EditProject(req admin.EditProjectRequest) error
	DeleteProject(projectID string) error
	UpdateProjectStatus(status string, id string) error
	GetUsersList(role string) ([]entity.Users, error)
	GetProjectsList() ([]entity.Project, error)
	GetMyProjects(userID string) ([]entity.Project, error)
	AddPeopleProject(req entity.ProjectsPeople) error
	CheckTeamLead(req admin.CheckTeamLeadRequest) (bool, error)
	GetUserRole(userID string) (string, error)
	GetUser(userID string) (entity.Users, error)
	GetProject(projectID string) (entity.Project, error)
	GetAttendance(programmerID string) ([]entity.Attendince, error)
}

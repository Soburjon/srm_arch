package admin

import (
	"srm_arch/internal/entity"
)

type AdminService struct {
	repo AdminRepo
}

func NewService(repo AdminRepo) *AdminService {
	return &AdminService{
		repo: repo,
	}
}

func (a *AdminService) CreateUser(req entity.Users) error {
	return a.repo.CreateUser(req)
}
func (a *AdminService) EditUser(req EditUserRequest) error {
	return a.repo.EditUser(req)
}
func (a *AdminService) DeleteUser(userID string) error {
	return a.repo.DeleteUser(userID)
}
func (a *AdminService) CreateProject(req entity.Project) error {
	return a.repo.CreateProject(req)
}
func (a *AdminService) EditProject(req EditProjectRequest) error {
	return a.repo.EditProject(req)
}
func (a *AdminService) DeleteProject(projectID string) error {
	return a.repo.DeleteProject(projectID)
}
func (a *AdminService) UpdateProjectStatus(status string, id string) error {
	return a.repo.UpdateProjectStatus(status, id)
}
func (a *AdminService) GetUsersList(role string) ([]entity.Users, error) {
	return a.repo.GetUsersList(role)
}
func (a *AdminService) GetProjectsList() ([]entity.Project, error) {
	return a.repo.GetProjectsList()
}
func (a *AdminService) GetMyProjects(userID string) ([]entity.Project, error) {
	return a.repo.GetMyProjects(userID)
}
func (a *AdminService) AddPeopleProject(req entity.ProjectsPeople) error {
	return a.repo.AddPeopleProject(req)
}
func (a *AdminService) CheckTeamLead(req CheckTeamLeadRequest) (bool, error) {
	return a.repo.CheckTeamLead(req)
}
func (a *AdminService) GetUserRole(userID string) (string, error) {
	return a.repo.GetUserRole(userID)
}
func (a *AdminService) GetUser(userID string) (entity.Users, error) {
	return a.repo.GetUser(userID)
}
func (a *AdminService) GetProject(projectID string) (entity.Project, error) {
	return a.repo.GetProject(projectID)
}
func (a *AdminService) GetAttendance(programmerID string) ([]entity.Attendince, error) {
	return a.repo.GetAttendance(programmerID)
}

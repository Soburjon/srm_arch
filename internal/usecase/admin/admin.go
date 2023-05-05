package admin

import (
	"github.com/google/uuid"
	"srm_arch/internal/entity"
	"srm_arch/internal/service/admin"
)

type AdminUseCase struct {
	service AdminService
}

func NewUseCase(s AdminService) *AdminUseCase {
	return &AdminUseCase{
		service: s,
	}
}

func (a *AdminUseCase) CreateUser(req admin.CreateUserRequest, role string) error {
	user := entity.Users{}
	user.ID = uuid.NewString()
	user.FullName = req.FullName
	user.Avatar = &req.Avatar
	user.PhoneNumber = req.PhoneNumber
	user.Birthday = req.Birthday
	user.Password = req.Password
	user.Position = req.Position
	user.Role = role
	return a.service.CreateUser(user)

}

func (a *AdminUseCase) EditUser(req admin.EditUserRequest) error {
	return a.service.EditUser(req)
}

func (a *AdminUseCase) DeleteUser(req admin.DeleteUserRequest) error {
	return a.service.DeleteUser(req.UserID)
}

func (a *AdminUseCase) CreateProject(req admin.CreateProjectRequest, teamLeadID string) error {
	project := entity.Project{}
	project.ID = uuid.NewString()
	project.Name = req.Name
	if req.StartDate != "" {
		project.StartDate = &req.StartDate
	}
	if req.EndDate != "" {
		project.EndDate = &req.EndDate
	}
	if req.Attachments != "" {
		project.Attachments = &req.Attachments
	}
	project.Status = "new"
	project.TeamleadID = teamLeadID
	return a.service.CreateProject(project)
}

func (a *AdminUseCase) EditProject(req admin.EditProjectRequest) error {
	return a.service.EditProject(req)
}

func (a *AdminUseCase) DeleteProject(req admin.DeleteProjectRequest) error {
	return a.service.DeleteProject(req.ID)
}

func (a *AdminUseCase) UpdateProjectStatus(req admin.UpdateProjectStatusRequest) error {
	return a.service.UpdateProjectStatus(req.Status, req.ProjectID)
}

func (a *AdminUseCase) GetUsersList(role string) ([]entity.Users, error) {
	return a.service.GetUsersList(role)
}

func (a *AdminUseCase) GetProjectsList() ([]entity.Project, error) {
	return a.service.GetProjectsList()
}

func (a *AdminUseCase) GetMyProjects(userID string) ([]entity.Project, error) {
	return a.service.GetMyProjects(userID)
}

func (a *AdminUseCase) AddPeopleProject(req admin.AddPeopleProjectRequest) error {
	projectPeople := entity.ProjectsPeople{}
	projectPeople.ProjectID = req.ProjectID
	projectPeople.UserID = req.ProgrammerID
	projectPeople.Position = req.Position
	return a.service.AddPeopleProject(projectPeople)
}

func (a *AdminUseCase) CheckTeamLead(req admin.CheckTeamLeadRequest) (bool, error) {
	return a.service.CheckTeamLead(req)
}

func (a *AdminUseCase) GetUserRole(userID string) (string, error) {
	return a.service.GetUserRole(userID)
}

func (a *AdminUseCase) GetUser(userID string) (entity.Users, error) {
	return a.service.GetUser(userID)
}

func (a *AdminUseCase) GetProject(projectID string) (entity.Project, error) {
	return a.service.GetProject(projectID)
}

func (a *AdminUseCase) GetAttendance(programmerID string) ([]entity.Attendince, error) {
	return a.service.GetAttendance(programmerID)
}

package programmer

import (
	"srm_arch/internal/entity"
)

type ProgrammerService struct {
	repo ProgrammerRepo
}

func NewService(repo ProgrammerRepo) *ProgrammerService {
	return &ProgrammerService{
		repo,
	}
}

func (p *ProgrammerService) CreateTask(req entity.Task) error {
	return p.repo.CreateTask(req)
}
func (p *ProgrammerService) EditTask(req EditTaskRequest) error {
	return p.repo.EditTask(req)
}
func (p *ProgrammerService) DeleteTask(taskID string) error {
	return p.repo.DeleteTask(taskID)
}
func (p *ProgrammerService) UpdateTaskStatus(status string, id string) error {
	return p.repo.UpdateTaskStatus(status, id)
}
func (p *ProgrammerService) GetProjectTasks(projectID string) ([]entity.Task, error) {
	return p.repo.GetProjectTasks(projectID)
}
func (p *ProgrammerService) GetTask(taskID string) (entity.Task, error) {
	return p.repo.GetTask(taskID)
}
func (p *ProgrammerService) GetMyTasks(req GetMyTaskRequest) ([]entity.Task, error) {
	return p.repo.GetMyTasks(req)
}
func (p *ProgrammerService) CreateCommit(req entity.Comment) error {
	return p.repo.CreateCommit(req)
}
func (p *ProgrammerService) EditCommit(req EditCommitRequest, userID string) error {
	return p.repo.EditCommit(req, userID)
}
func (p *ProgrammerService) DeleteCommit(createdAt string, userID string) error {
	return p.repo.DeleteCommit(createdAt, userID)
}
func (p *ProgrammerService) GetCommits(taskID string) ([]entity.Comment, error) {
	return p.repo.GetCommits(taskID)
}
func (p *ProgrammerService) CreateAttendance(req entity.Attendince) error {
	return p.repo.CreateAttendance(req)
}
func (p *ProgrammerService) UsersInProject(projectID string) ([]entity.ProjectsPeople, error) {
	return p.repo.UsersInProject(projectID)
}
func (p *ProgrammerService) UserRoleInProject(req UserRoleInProjectRequest) (string, error) {
	return p.repo.UserRoleInProject(req)
}
func (p *ProgrammerService) MyProjects(userID string) ([]entity.Project, error) {
	return p.repo.MyProjects(userID)
}

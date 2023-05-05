package programmer

import (
	"github.com/google/uuid"
	"srm_arch/internal/entity"
	"srm_arch/internal/service/programmer"
)

type ProgrammerUseCase struct {
	service ProgrammerService
}

func NewUseCase(s ProgrammerService) *ProgrammerUseCase {
	return &ProgrammerUseCase{
		service: s,
	}
}

func (p *ProgrammerUseCase) CreateTask(req programmer.CreateTaskRequest) error {
	task := entity.Task{}
	task.ID = uuid.NewString()
	task.ProjectID = req.ProjectID
	task.Title = req.Title
	task.Description = req.Description
	if req.StartAt != "" {
		task.StartAt = &req.StartAt
	}
	if req.FinishAt != "" {
		task.FinishAt = &req.FinishAt
	}
	task.ProgrammerID = req.ProgrammerID
	if req.Attachments != "" {
		task.Attachments = &req.Attachments
	}
	task.Status = "new"

	return p.service.CreateTask(task)
}

func (p *ProgrammerUseCase) EditTask(req programmer.EditTaskRequest) error {
	return p.service.EditTask(req)
}

func (p *ProgrammerUseCase) DeleteTask(taskID string) error {
	return p.service.DeleteTask(taskID)
}

func (p *ProgrammerUseCase) UpdateTaskStatus(req programmer.UpdateTaskStatusRequest) error {
	return p.service.UpdateTaskStatus(req.Status, req.TaskID)
}

func (p *ProgrammerUseCase) GetProjectTasks(projectID string) ([]entity.Task, error) {
	return p.service.GetProjectTasks(projectID)
}

func (p *ProgrammerUseCase) GetTask(taskID string) (entity.Task, error) {
	return p.service.GetTask(taskID)
}

func (p *ProgrammerUseCase) GetMyTasks(req programmer.GetMyTaskRequest) ([]entity.Task, error) {
	return p.service.GetMyTasks(req)
}

func (p *ProgrammerUseCase) CreateCommit(req programmer.CreateCommitRequest) error {
	comment := entity.Comment{}
	comment.TaskID = req.TaskID
	comment.ProgrammerID = req.UserID
	comment.Text = req.Text
	return p.service.CreateCommit(comment)
}

func (p *ProgrammerUseCase) EditCommit(req programmer.EditCommitRequest, userID string) error {
	return p.service.EditCommit(req, userID)
}

func (p *ProgrammerUseCase) DeleteCommit(createdAt string, userID string) error {
	return p.service.DeleteCommit(createdAt, userID)
}

func (p *ProgrammerUseCase) GetCommits(taskID string) ([]entity.Comment, error) {
	return p.service.GetCommits(taskID)
}

func (p *ProgrammerUseCase) CreateAttendance(req programmer.CreateAttendanceRequest) error {
	attendance := entity.Attendince{}
	attendance.UserID = req.UserID
	attendance.Type = req.Type
	return p.service.CreateAttendance(attendance)
}

func (p *ProgrammerUseCase) UsersInProject(projectID string) ([]entity.ProjectsPeople, error) {
	return p.service.UsersInProject(projectID)
}

func (p *ProgrammerUseCase) UserRoleInProject(req programmer.UserRoleInProjectRequest) (string, error) {
	return p.service.UserRoleInProject(req)
}

func (p *ProgrammerUseCase) MyProjects(userID string) ([]entity.Project, error) {
	return p.service.MyProjects(userID)
}

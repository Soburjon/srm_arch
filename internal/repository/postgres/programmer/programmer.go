package programmer

import (
	"context"
	"github.com/uptrace/bun"
	"srm_arch/internal/entity"
	"srm_arch/internal/service/programmer"
	"time"
)

type programmerRepo struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *programmerRepo {
	return &programmerRepo{
		DB,
	}
}

func (p *programmerRepo) CreateTask(req entity.Task) error {
	_, err := p.NewInsert().Model(&req).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (p *programmerRepo) EditTask(req programmer.EditTaskRequest) error {
	columns := make([]string, 0)
	var task entity.Task
	task.ID = req.ID
	if req.Title != "" {
		task.Title = req.Title
		columns = append(columns, "title")
	}
	if req.Description != "" {
		task.Description = req.Description
		columns = append(columns, "description")
	}
	if req.StartAt != "" {
		task.StartAt = &req.StartAt
		columns = append(columns, "start_at")
	}
	if req.FinishAt != "" {
		task.FinishAt = &req.FinishAt
		columns = append(columns, "finish_at")
	}
	if req.ProgrammerID != "" {
		task.ProgrammerID = req.ProgrammerID
		columns = append(columns, "programmer_id")
	}
	if req.Attachments != "" {
		task.Attachments = &req.Attachments
		columns = append(columns, "attachments")
	}
	update := time.Now()
	task.UpdatedAt = &update
	_, err := p.NewUpdate().Column(columns...).Model(&task).WherePK("id").Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (p *programmerRepo) DeleteTask(taskID string) error {
	var task entity.Task
	task.ID = taskID
	deleted := time.Now()
	task.DeletedAt = &deleted
	_, err := p.NewUpdate().Column("deleted_at").Model(&task).WherePK("id").Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (p *programmerRepo) UpdateTaskStatus(status string, id string) error {
	var task entity.Task
	task.ID = id
	task.Status = status
	update := time.Now()
	task.UpdatedAt = &update
	_, err := p.NewUpdate().Column("status").Model(&task).WherePK("id").Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (p *programmerRepo) GetProjectTasks(projectID string) ([]entity.Task, error) {
	tasks := make([]entity.Task, 0)
	err := p.NewSelect().Model(&tasks).Where("project_id = ? and deleted_at IS NULL", projectID).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func (p *programmerRepo) GetTask(taskID string) (entity.Task, error) {
	var task entity.Task
	err := p.NewSelect().Model(&task).Where("id = ? and deleted_at IS NULL", taskID).Scan(context.Background())
	if err != nil {
		return entity.Task{}, err
	}
	return task, nil
}
func (p *programmerRepo) GetMyTasks(req programmer.GetMyTaskRequest) ([]entity.Task, error) {
	tasks := make([]entity.Task, 0)
	err := p.NewSelect().Model(&tasks).Where("project_id = ? and programmer_id = ? and deleted_at IS NULL", req.ProjectID, req.UserID).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func (p *programmerRepo) CreateCommit(req entity.Comment) error {
	_, err := p.NewInsert().Model(&req).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (p *programmerRepo) EditCommit(req programmer.EditCommitRequest, userID string) error {
	columns := make([]string, 0)
	var commit entity.Comment
	if req.Text != "" {
		commit.Text = req.Text
		columns = append(columns, "text")
	}

	update := time.Now()
	commit.UpdatedAt = &update
	_, err := p.NewUpdate().Column(columns...).Model(&commit).Where("created_at = ? and programmer_id = ?", req.CreatedAt, userID).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (p *programmerRepo) DeleteCommit(createdAt string, userID string) error {
	var commit entity.Comment
	deleted := time.Now()
	commit.DeletedAt = &deleted
	_, err := p.NewUpdate().Column("deleted_at").Model(&commit).Where("created_at = ? and programmer_id = ?", createdAt, userID).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (p *programmerRepo) GetCommits(taskID string) ([]entity.Comment, error) {
	comments := make([]entity.Comment, 0)
	err := p.NewSelect().Model(&comments).Where("task_id = ? and deleted_at IS NULL", taskID).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return comments, nil
}
func (p *programmerRepo) CreateAttendance(req entity.Attendince) error {
	_, err := p.NewInsert().Model(&req).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (p *programmerRepo) UsersInProject(projectID string) ([]entity.ProjectsPeople, error) {
	projectPeople := make([]entity.ProjectsPeople, 0)
	err := p.NewSelect().Model(&projectPeople).Where("project_id = ?", projectID).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return projectPeople, nil
}
func (p *programmerRepo) UserRoleInProject(req programmer.UserRoleInProjectRequest) (string, error) {
	var projectPeople entity.ProjectsPeople
	err := p.NewSelect().Model(&projectPeople).Where("project_id = ? and user_id = ?", req.ProjectID, req.UserID).Scan(context.Background())
	if err != nil {
		return "", err
	}
	return projectPeople.Position, nil
}
func (p *programmerRepo) MyProjects(userID string) ([]entity.Project, error) {
	projects := make([]entity.Project, 0)

	err := p.NewSelect().Model(&projects).ColumnExpr("project.*").
		Join("JOIN projects_people AS ps ").JoinOn("project.id = ps.project_id").
		Where("ps.user_id = ?", userID).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return projects, nil
}

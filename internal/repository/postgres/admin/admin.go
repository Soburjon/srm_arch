package admin

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"srm_arch/internal/entity"
	"srm_arch/internal/service/admin"
	"time"
)

type adminRepo struct {
	*bun.DB
}

func NewRepository(DB *bun.DB) *adminRepo {
	return &adminRepo{
		DB,
	}
}

func (a *adminRepo) CreateUser(req entity.Users) error {
	_, err := a.NewInsert().Model(&req).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (a *adminRepo) EditUser(req admin.EditUserRequest) error {
	columns := make([]string, 0)
	var user entity.Users
	user.ID = req.UserID
	if req.FullName != "" {
		user.FullName = req.FullName
		columns = append(columns, "full_name")
	}
	if req.Avatar != "" {
		user.Avatar = &req.Avatar
		columns = append(columns, "avatar")
	}
	if req.PhoneNumber != "" {
		user.PhoneNumber = req.PhoneNumber
		columns = append(columns, "phone_number")
	}
	if req.Birthday != "" {
		user.Birthday = req.Birthday
		columns = append(columns, "birthday")
	}
	if req.Password != "" {
		user.Password = req.Password
		columns = append(columns, "password")
	}
	if req.Position != "" {
		user.Position = req.Position
		columns = append(columns, "position")
	}
	fmt.Println(req)
	update := time.Now()
	user.UpdatedAt = &update
	_, err := a.NewUpdate().Column(columns...).Model(&user).WherePK("id").Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (a *adminRepo) DeleteUser(userID string) error {
	var user entity.Users
	user.ID = userID
	deleted := time.Now()
	user.DeletedAt = &deleted
	_, err := a.NewUpdate().Column("deleted_at").Model(&user).WherePK("id").Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (a *adminRepo) CreateProject(req entity.Project) error {
	_, err := a.NewInsert().Model(&req).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (a *adminRepo) EditProject(req admin.EditProjectRequest) error {
	columns := make([]string, 0)
	var project entity.Project
	project.ID = req.ID
	if req.Name != "" {
		project.Name = req.Name
		columns = append(columns, "name")
	}
	if req.Status != "" {
		project.Status = req.Status
		columns = append(columns, "status")
	}
	if req.StartDate != "" {
		project.StartDate = &req.StartDate
		columns = append(columns, "start_date")
	}
	if req.EndDate != "" {
		project.EndDate = &req.EndDate
		columns = append(columns, "end_date")
	}
	if req.Attachments != "" {
		project.Attachments = &req.Attachments
		columns = append(columns, "attachments")
	}
	update := time.Now()
	project.UpdatedAt = &update
	_, err := a.NewUpdate().Column(columns...).Model(&project).WherePK("id").Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (a *adminRepo) DeleteProject(projectID string) error {
	var project entity.Project
	project.ID = projectID
	deleted := time.Now()
	project.DeletedAt = &deleted
	_, err := a.NewUpdate().Column("deleted_at").Model(&project).WherePK("id").Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (a *adminRepo) UpdateProjectStatus(status string, id string) error {
	var project entity.Project
	project.ID = id
	project.Status = status
	update := time.Now()
	project.UpdatedAt = &update
	fmt.Println(project.ID)
	_, err := a.NewUpdate().Column("status").Model(&project).WherePK("id").Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (a *adminRepo) GetUsersList(role string) ([]entity.Users, error) {
	users := make([]entity.Users, 0)
	err := a.NewSelect().Model(&users).Where("role = ? and deleted_at IS NULL", role).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (a *adminRepo) GetProjectsList() ([]entity.Project, error) {
	projects := make([]entity.Project, 0)
	err := a.NewSelect().Model(&projects).Where("deleted_at IS NULL").Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return projects, nil
}
func (a *adminRepo) GetMyProjects(userID string) ([]entity.Project, error) {
	projects := make([]entity.Project, 0)
	err := a.NewSelect().Model(&projects).Where("teamlead_id = ? and deleted_at IS NULL", userID).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return projects, nil
}
func (a *adminRepo) AddPeopleProject(req entity.ProjectsPeople) error {
	_, err := a.NewInsert().Model(&req).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}
func (a *adminRepo) CheckTeamLead(req admin.CheckTeamLeadRequest) (bool, error) {
	projects := make([]entity.Project, 0)
	count, err := a.NewSelect().Model(&projects).Where("teamlead_id = ? and id = ? and deleted_at IS NULL", req.UserID, req.ProjectID).ScanAndCount(context.Background())
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err
	}
	return true, nil
}
func (a *adminRepo) GetUserRole(userID string) (string, error) {
	var users entity.Users
	err := a.NewSelect().Model(&users).Where("id = ? and deleted_at IS NULL", userID).Scan(context.Background())
	if err != nil {
		return "", err
	}
	return users.Role, nil
}
func (a *adminRepo) GetUser(userID string) (entity.Users, error) {
	var users entity.Users
	err := a.NewSelect().Model(&users).Where("id = ? and deleted_at IS NULL", userID).Scan(context.Background())
	if err != nil {
		return entity.Users{}, err
	}
	return users, nil
}
func (a *adminRepo) GetProject(projectID string) (entity.Project, error) {
	var project entity.Project
	err := a.NewSelect().Model(&project).Where("id = ? and deleted_at IS NULL", projectID).Scan(context.Background())
	if err != nil {
		return entity.Project{}, err
	}
	return project, nil
}
func (a *adminRepo) GetAttendance(programmerID string) ([]entity.Attendince, error) {
	attendances := make([]entity.Attendince, 0)
	err := a.NewSelect().Model(&attendances).Where("user_id = ?", programmerID).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return attendances, nil
}

package sqlite

import (
	"errors"
	"fmt"

	"github.com/2ips/todo/internal/app/port"
	"github.com/2ips/todo/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) NextID() (domain.TaskID, error) {
	var id int64

	if err := r.db.
		Raw("SELECT COALESCE(MAX(id), 0) + 1 FROM tasks").
		Scan(&id).Error; err != nil {
		return 0, fmt.Errorf("%w: get next id error: %v", port.ErrStorage, err)
	}

	return domain.NewTaskID(id)
}

func (r *TaskRepository) List() ([]*domain.Task, error) {
	var models []taskModel

	if err := r.db.
		Order("deadline").
		Order("id").
		Find(&models).Error; err != nil {
		return nil, fmt.Errorf("%w: list error: %v", port.ErrStorage, err)
	}

	tasks := make([]*domain.Task, len(models))
	for i, model := range models {
		task, err := fromModel(&model)
		if err != nil {
			return nil, fmt.Errorf("%w: list error: %v", port.ErrStorage, err)
		}
		tasks[i] = task
	}

	return tasks, nil
}

func (r *TaskRepository) Get(id domain.TaskID) (*domain.Task, error) {
	var model taskModel

	err := r.db.
		Where("id = ?", int64(id)).
		First(&model).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("get error: %w", &port.ErrTaskNotFound{Value: id})
	}
	if err != nil {
		return nil, fmt.Errorf("%w: get error for task id %d: %v", port.ErrStorage, id, err)
	}

	return fromModel(&model)
}

func (r *TaskRepository) Save(task *domain.Task) error {
	model := toModel(task)

	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"title", "status", "deadline"}),
	}).Create(&model).Error
	if err != nil {
		return fmt.Errorf("%w: save error for task id %d: %v", port.ErrStorage, task.ID, err)
	}

	return nil
}

func (r *TaskRepository) Delete(id domain.TaskID) error {
	result := r.db.
		Where("id = ?", int64(id)).
		Delete(&taskModel{})

	if err := result.Error; err != nil {
		return fmt.Errorf("%w: delete error for task id %d: %v", port.ErrStorage, id, err)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("get error: %w", &port.ErrTaskNotFound{Value: id})
	}

	return nil
}

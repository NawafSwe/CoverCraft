package repositories

import (
	"coverCraft/entities"
	"gorm.io/gorm"
)

// JobRepository defines the methods that a job repository should implement.
type JobRepository interface {
	Create(job entities.Job) error
	ListAllJobs() ([]entities.Job, error)
}

// GormJobRepository is an implementation of JobRepository using GORM.
type GormJobRepository struct {
	db *gorm.DB
}

// NewGormJobRepository creates a new GormJobRepository instance with the provided GORM DB connection.
func NewGormJobRepository(db *gorm.DB) *GormJobRepository {
	return &GormJobRepository{db: db}
}

func (r *GormJobRepository) Create(job entities.Job) error {
	if res := r.db.Create(&job); res.Error != nil {
		return res.Error
	} else {
		return nil
	}
}

func (r *GormJobRepository) ListAllJobs() ([]entities.Job, error) {
	var jobs []entities.Job
	if err := r.db.Find(&jobs).Error; err != nil {
		return nil, err
	}
	return jobs, nil
}

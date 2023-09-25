package aggregates

import (
	"coverCraft/config"
	"coverCraft/entities"
	"github.com/google/uuid"
)

type Job struct {
	Id             uuid.UUID
	Company        string
	Location       string
	Description    string
	Qualifications string
	Name           string
}

func CreateJob(job *Job) error {
	jobModel := &entities.Job{
		ID:             uuid.New(),
		Name:           job.Name,
		Location:       job.Location,
		Company:        job.Company,
		Description:    job.Description,
		Qualifications: job.Qualifications,
	}
	gorm := config.DB()
	if err := gorm.Create(&jobModel).Error; err != nil {
		return err
	}
	return nil

}

func ListAllJobs() ([]Job, error) {
	gorm := config.DB()
	var jobs []*entities.Job
	var convertedJobs []Job
	if res := gorm.Find(&jobs); res.Error != nil {
		return convertedJobs, res.Error
	}
	for _, v := range jobs {
		convertedJob := Job{
			Id:             v.ID,
			Company:        v.Company,
			Location:       v.Location,
			Description:    v.Description,
			Qualifications: v.Qualifications,
			Name:           v.Name,
		}
		convertedJobs = append(convertedJobs, convertedJob)
	}
	return convertedJobs, nil
}

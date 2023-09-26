package aggregates

import (
	"coverCraft/domain/repositories"
	"coverCraft/entities"
	"github.com/google/uuid"
)

type Job struct {
	Id             uuid.UUID `json:"id"`
	Company        string    `json:"company"`
	Location       string    `json:"location"`
	Description    string    `json:"description"`
	Qualifications string    `json:"qualifications"`
	Name           string    `json:"name"`
}

func CreateJob(job *Job, jobRepo repositories.JobRepository) error {
	jobModel := entities.Job{
		ID:             uuid.New(),
		Name:           job.Name,
		Location:       job.Location,
		Company:        job.Company,
		Description:    job.Description,
		Qualifications: job.Qualifications,
	}
	if err := jobRepo.Create(jobModel); err != nil {
		return err
	}
	return nil

}

func ListAllJobs(jobRepo repositories.JobRepository) ([]Job, error) {
	jobs, err := jobRepo.ListAllJobs()
	if err != nil {
		return nil, err
	}
	var convertedJobs []Job
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

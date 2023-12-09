package domain

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

// this is a hook to set the fields required by default
// is when the struct is initialized
func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// Job is the struct that represents a video. Like an Entity in DDD.
type Job struct {
	ID               string    `json:"job_id" valid:"uuid" gorm:"type:uuid;primary_key;"`
	OutputBucketPath string    `json:"output_bucket-path" valid:"notnull"`
	Status           string    `json:"status" valid:"notnull"`
	Video            *Video    `json:"video" valid:"-"`
	VideoID          string    `valid:"-" gorm:"column:video_id;type:uuid;notnull"`
	Error            string    `valid:"-"`
	CreatedAt        time.Time `json:"created_at" valid:"-"`
	UpdatedAt        time.Time `json:"updated_at" valid:"-"`
}

// this is a hook to set the fields required by default
// it works like a class constructor
func NewJob(output string, status string, video *Video) (*Job, error) {
	job := Job{
		OutputBucketPath: output,
		Status:           status,
		Video:            video,
	}

	// hook to inject the initial values
	job.prepare()

	err := job.Validate()

	if err != nil {
		return nil, err
	}

	return &job, nil
}

// prepare is a hook to set the initial values of a Job instance
func (j *Job) prepare() {
	j.ID = uuid.New().String()
	j.CreatedAt = time.Now()
	j.UpdatedAt = time.Now()
}

// Validate is a method that validates the Job struct using govalidator.
// It returns an error if the validation fails.
func (j *Job) Validate() error {
	_, err := govalidator.ValidateStruct(j)

	if err != nil {
		return err
	}

	return nil
}

func (j *Job) GenerateID() {
	j.ID = uuid.New().String()
}

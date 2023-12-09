package repository_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/application/repository"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/domain"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/framework/database"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("tmp/video.mp4", "pending", video)
	require.Nil(t, err)

	repo := repository.JobRepositoryDb{Db: db}
	repo.Insert(job)

	j, err := repo.Find(job.ID)

	require.Nil(t, err)
	require.NotEmpty(t, j.ID)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("tmp/video.mp4", "pending", video)
	require.Nil(t, err)

	repo := repository.JobRepositoryDb{Db: db}
	repo.Insert(job)

	job.Status = "complete"

	repo.Update(job)

	j, err := repo.Find(job.ID)

	require.Nil(t, err)
	require.NotEmpty(t, j.ID)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.Status, "complete")
}

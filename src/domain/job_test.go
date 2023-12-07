package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/domain"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.New().String()
	video.FilePath = "video.mp4"
	video.CreatedAt = time.Now()

	job, err := domain.NewJob("output-bucket", "PENDING", video)

	require.NotNil(t, job)
	require.Nil(t, err)
}

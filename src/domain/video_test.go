package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/domain"
)

func TestVideoValidationWhenVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidationWhenIDIsNotUUID(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "abc"
	video.ResourceID = "abc"
	video.FilePath = "abc"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidationWhenIsValid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = uuid.New().String()
	video.ResourceID = "abc"
	video.FilePath = "abc"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}

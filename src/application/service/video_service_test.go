package service_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/application/repository"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/application/service"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/domain"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/framework/database"
)

func init() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	log.Println("Loaded .env file")
}

func prepare() (*domain.Video, repository.VideoRepositoryDb) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "two-wolfs-ice.mp4"
	video.CreatedAt = time.Now()

	repo := repository.VideoRepositoryDb{Db: db}

	return video, repo
}

func TestVideoServiceDownload(t *testing.T) {
	video, repo := prepare()

	vs := service.NewVideoService()
	vs.Video = video
	vs.VideoRepository = repo

	err := vs.Download(os.Getenv("INPUT_BUCKET_NAME"))
	require.Nil(t, err)
}

func TestVideoServiceFragment(t *testing.T) {
	video, repo := prepare()

	vs := service.NewVideoService()
	vs.Video = video
	vs.VideoRepository = repo

	err := vs.Download(os.Getenv("INPUT_BUCKET_NAME"))
	require.Nil(t, err)

	err = vs.Fragment()
	require.Nil(t, err)
}

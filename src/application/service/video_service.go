package service

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"cloud.google.com/go/storage"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/application/repository"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/domain"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repository.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (vs *VideoService) Fragment() error {
	err := os.Mkdir(os.Getenv("LOCAL_STORAGE_PATH")+"/"+vs.Video.ID, os.ModePerm)
	if err != nil {
		return err
	}

	source := os.Getenv("LOCAL_STORAGE_PATH") + "/" + vs.Video.ID + ".mp4"
	target := os.Getenv("LOCAL_STORAGE_PATH") + "/" + vs.Video.ID + ".frag"

	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	printOutput(output)

	return nil
}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("==> Output: %s\n", string(out))
	}
}

func (vs *VideoService) Download(bucketName string) error {
	// ctx is the context.Context object for the current app engine request.
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	bkt := client.Bucket(bucketName)
	obj := bkt.Object(vs.Video.FilePath)

	r, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}
	defer r.Close()

	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	LOCAL_STORAGE_PATH := os.Getenv("LOCAL_STORAGE_PATH")
	filePath := LOCAL_STORAGE_PATH + "/" + vs.Video.ID + ".mp4"

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = f.Write(body)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Printf("Video %v downloaded.\n", vs.Video.ID)

	return nil
}

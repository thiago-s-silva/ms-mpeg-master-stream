package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/thiago-s-silva/ms-mpeg-master-stream/src/domain"
)

/*
This is the repository layer. It is responsible for the communication with the database.
It implements the interface VideoRepository.
The benefit of using an interface is that we can easily change the database without changing the application layer.
*/
type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

/*
This is the implementation of the interface VideoRepository.
It has a field Db of type *gorm.DB, which is the connection to the database.
*/
type VideoRepositoryDb struct {
	Db *gorm.DB
}

/*
This is the constructor of the VideoRepositoryDb struct.
It receives a *gorm.DB as parameter and returns a pointer to a VideoRepositoryDb.
*/
func NewVideoRepository(db *gorm.DB) *VideoRepositoryDb {
	return &VideoRepositoryDb{Db: db}
}

func (repo VideoRepositoryDb) Insert(video *domain.Video) (*domain.Video, error) {
	if video.ID == "" {
		video.GenerateID()
	}

	err := repo.Db.Create(video).Error
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (repo VideoRepositoryDb) Find(id string) (*domain.Video, error) {
	// Create an empty Video object.
	var video domain.Video

	// Find the video with the given id.
	repo.Db.Preload("Jobs").Find(&video, "id = ?", id)

	// Check if the video was found.
	if video.ID == "" {
		return nil, fmt.Errorf("video does not exist")
	}

	return &video, nil
}

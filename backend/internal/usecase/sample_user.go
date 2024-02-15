package usecase

import "github.com/uptrace/bun"

type SampleUserUsecase interface {
	CreateSampleUser() error
}

type SampleUser struct {
	db *bun.DB
}

func NewSampleUser(db *bun.DB) *SampleUser {
	return &SampleUser{
		db: db,
	}
}

func (s *SampleUser) CreateSampleUser() error {
	return nil
}

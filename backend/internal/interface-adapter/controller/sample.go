package controller

import (
	"github.com/GenkiHirano/go-grpc-base/internal/usecase"
)

type Sample struct {
	sampleUserUsecase usecase.SampleUserUsecase
}

func NewSample(sampleUserUsecase usecase.SampleUserUsecase) *Sample {
	return &Sample{
		sampleUserUsecase: sampleUserUsecase,
	}
}

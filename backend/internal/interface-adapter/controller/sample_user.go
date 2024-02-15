package controller

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	samplev1 "github.com/GenkiHirano/go-grpc-base/internal/gen/sample/v1"
)

func (s *Sample) SampleCreateUser(ctx context.Context, req *connect.Request[samplev1.SampleCreateUserRequest]) (*connect.Response[samplev1.SampleCreateUserResponse], error) {
	fmt.Println("🌈")
	return nil, nil
}

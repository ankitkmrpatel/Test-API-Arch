package grpc

import (
	"context"
	"test-api-arch/internal/usecase"
)

type UserGrpcHandler struct {
	Usecase usecase.UserUsecase
}

func (h *UserGrpcHandler) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	user, err := h.Usecase.GetUserByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}

	return &GetUserResponse{
		Id:    int32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

package login

import (
	"context"
	"errors"

	"github.com/skripov-ds-ai/social_network/generated"
	"github.com/skripov-ds-ai/social_network/internal/models"
	"github.com/skripov-ds-ai/social_network/internal/services/user"
)

type Handler struct {
	service user.UserService
}

func NewLoginHandler(service user.UserService) *Handler {
	return &Handler{
		service: service,
	}
}

// LoginPost implements POST /login operation.
//
// Упрощенный процесс аутентификации путем передачи
// идентификатор пользователя и получения токена для
// дальнейшего прохождения авторизации.
//
// POST /login
func (h *Handler) LoginPost(ctx context.Context, req generated.OptLoginPostReq) (generated.LoginPostRes, error) {
	userCredentials, ok := req.Get()
	if !ok {
		return &generated.LoginPostBadRequest{}, nil
	}

	id, ok := userCredentials.ID.Get()
	if !ok {
		return &generated.LoginPostBadRequest{}, nil
	}

	password, ok := userCredentials.Password.Get()
	if !ok {
		return &generated.LoginPostBadRequest{}, nil
	}

	token, err := h.service.Login(ctx, string(id), password)
	if errors.Is(err, models.ErrNotFound) {
		return &generated.LoginPostNotFound{}, nil
	}
	if err != nil {
		return &generated.LoginPostInternalServerError{}, err
	}

	return &generated.LoginPostOK{
		Token: generated.NewOptString(token),
	}, nil
}

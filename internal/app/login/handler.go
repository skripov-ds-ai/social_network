package login

import (
	"context"
	"github.com/skripov-ds-ai/social_network/generated"
)

type Handler struct {
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

	return &generated.LoginPostOK{}, nil
}

package user

import (
	"context"
	"github.com/skripov-ds-ai/social_network/generated"
)

type Handler struct {
}

// UserGetIDGet implements GET /user/get/{id} operation.
//
// Получение анкеты пользователя.
//
// GET /user/get/{id}
func (h *Handler) UserGetIDGet(ctx context.Context, params generated.UserGetIDGetParams) (generated.UserGetIDGetRes, error) {
	return &generated.User{}, nil
}

// UserRegisterPost implements POST /user/register operation.
//
// Регистрация нового пользователя.
//
// POST /user/register
func (h *Handler) UserRegisterPost(ctx context.Context, req generated.OptUserRegisterPostReq) (generated.UserRegisterPostRes, error) {
	return &generated.UserRegisterPostOK{}, nil
}

// UserSearchGet implements GET /user/search operation.
//
// Поиск анкет.
//
// GET /user/search
func (h *Handler) UserSearchGet(ctx context.Context, params generated.UserSearchGetParams) (generated.UserSearchGetRes, error) {
	// TODO
	return &generated.UserSearchGetOKApplicationJSON{}, nil
}

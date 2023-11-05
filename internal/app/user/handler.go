package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/skripov-ds-ai/social_network/generated"
	"github.com/skripov-ds-ai/social_network/internal/models"
	"github.com/skripov-ds-ai/social_network/internal/services/user"
)

type Handler struct {
	service user.UserService
}

func NewUserHandler(service user.UserService) *Handler {
	return &Handler{
		service: service,
	}
}

// UserGetIDGet implements GET /user/get/{id} operation.
//
// Получение анкеты пользователя.
//
// GET /user/get/{id}
func (h *Handler) UserGetIDGet(ctx context.Context, params generated.UserGetIDGetParams) (generated.UserGetIDGetRes, error) {
	// TODO: перенести в swagger
	_, err := uuid.Parse(string(params.ID))
	if err != nil {
		return &generated.UserGetIDGetBadRequest{}, nil
	}

	u, err := h.service.Get(ctx, string(params.ID))
	if errors.Is(err, models.ErrNotFound) {
		return &generated.UserGetIDGetNotFound{}, nil
	}
	if err != nil {
		return &generated.UserGetIDGetInternalServerError{}, err
	}

	var city string
	if u.City != nil {
		city = *u.City
	}

	var biography string
	if u.Biography != nil {
		biography = *u.Biography
	}

	return &generated.User{
		ID:         generated.NewOptUserId(generated.UserId(u.ID.String())),
		FirstName:  generated.NewOptString(u.FirstName),
		SecondName: generated.NewOptString(u.SecondName),
		Birthdate:  generated.NewOptBirthDate(generated.BirthDate(u.Birthdate)),
		Biography:  generated.NewOptString(biography),
		City:       generated.NewOptString(city),
	}, nil
}

// UserRegisterPost implements POST /user/register operation.
//
// Регистрация нового пользователя.
//
// POST /user/register
func (h *Handler) UserRegisterPost(ctx context.Context, req generated.OptUserRegisterPostReq) (generated.UserRegisterPostRes, error) {
	u, ok := req.Get()
	if !ok {
		return &generated.UserRegisterPostBadRequest{}, nil
	}

	password, ok := u.Password.Get()
	if !ok {
		return &generated.UserRegisterPostBadRequest{}, nil
	}

	birthDate, ok := u.Birthdate.Get()
	if !ok {
		return &generated.UserRegisterPostBadRequest{}, nil
	}

	var biography *string
	bio, ok := u.Biography.Get()
	if !ok {
		return &generated.UserRegisterPostBadRequest{}, nil
	}
	biography = &bio

	var city *string
	town, ok := u.City.Get()
	if !ok {
		return &generated.UserRegisterPostBadRequest{}, nil
	}
	city = &town

	params := models.CreateUserParams{
		Password:   password,
		FirstName:  u.FirstName.Value,
		SecondName: u.SecondName.Value,
		Birthdate:  time.Time(birthDate),
		Biography:  biography,
		City:       city,
	}

	id, err := h.service.Register(ctx, params)
	if err != nil {
		return &generated.UserRegisterPostInternalServerError{}, err
	}

	return &generated.UserRegisterPostOK{
		UserID: generated.NewOptString(id),
	}, nil
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

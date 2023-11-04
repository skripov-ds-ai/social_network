// Code generated by ogen, DO NOT EDIT.

package generated

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// DialogUserIDListGet implements GET /dialog/{user_id}/list operation.
	//
	// GET /dialog/{user_id}/list
	DialogUserIDListGet(ctx context.Context, params DialogUserIDListGetParams) (DialogUserIDListGetRes, error)
	// DialogUserIDSendPost implements POST /dialog/{user_id}/send operation.
	//
	// POST /dialog/{user_id}/send
	DialogUserIDSendPost(ctx context.Context, req OptDialogUserIDSendPostReq, params DialogUserIDSendPostParams) (DialogUserIDSendPostRes, error)
	// FriendDeleteUserIDPut implements PUT /friend/delete/{user_id} operation.
	//
	// PUT /friend/delete/{user_id}
	FriendDeleteUserIDPut(ctx context.Context, params FriendDeleteUserIDPutParams) (FriendDeleteUserIDPutRes, error)
	// FriendSetUserIDPut implements PUT /friend/set/{user_id} operation.
	//
	// PUT /friend/set/{user_id}
	FriendSetUserIDPut(ctx context.Context, params FriendSetUserIDPutParams) (FriendSetUserIDPutRes, error)
	// LoginPost implements POST /login operation.
	//
	// Упрощенный процесс аутентификации путем передачи
	// идентификатор пользователя и получения токена для
	// дальнейшего прохождения авторизации.
	//
	// POST /login
	LoginPost(ctx context.Context, req OptLoginPostReq) (LoginPostRes, error)
	// PostCreatePost implements POST /post/create operation.
	//
	// POST /post/create
	PostCreatePost(ctx context.Context, req OptPostCreatePostReq) (PostCreatePostRes, error)
	// PostDeleteIDPut implements PUT /post/delete/{id} operation.
	//
	// PUT /post/delete/{id}
	PostDeleteIDPut(ctx context.Context, params PostDeleteIDPutParams) (PostDeleteIDPutRes, error)
	// PostFeedGet implements GET /post/feed operation.
	//
	// GET /post/feed
	PostFeedGet(ctx context.Context, params PostFeedGetParams) (PostFeedGetRes, error)
	// PostGetIDGet implements GET /post/get/{id} operation.
	//
	// GET /post/get/{id}
	PostGetIDGet(ctx context.Context, params PostGetIDGetParams) (PostGetIDGetRes, error)
	// PostUpdatePut implements PUT /post/update operation.
	//
	// PUT /post/update
	PostUpdatePut(ctx context.Context, req OptPostUpdatePutReq) (PostUpdatePutRes, error)
	// UserGetIDGet implements GET /user/get/{id} operation.
	//
	// Получение анкеты пользователя.
	//
	// GET /user/get/{id}
	UserGetIDGet(ctx context.Context, params UserGetIDGetParams) (UserGetIDGetRes, error)
	// UserRegisterPost implements POST /user/register operation.
	//
	// Регистрация нового пользователя.
	//
	// POST /user/register
	UserRegisterPost(ctx context.Context, req OptUserRegisterPostReq) (UserRegisterPostRes, error)
	// UserSearchGet implements GET /user/search operation.
	//
	// Поиск анкет.
	//
	// GET /user/search
	UserSearchGet(ctx context.Context, params UserSearchGetParams) (UserSearchGetRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}

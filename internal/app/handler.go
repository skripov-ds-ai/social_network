package app

import (
	"context"

	"github.com/skripov-ds-ai/social_network/generated"
)

type LoginHandler interface {
	LoginPost(ctx context.Context, req generated.OptLoginPostReq) (generated.LoginPostRes, error)
}

type UserHandler interface {
	UserGetIDGet(ctx context.Context, params generated.UserGetIDGetParams) (generated.UserGetIDGetRes, error)
	UserRegisterPost(ctx context.Context, req generated.OptUserRegisterPostReq) (generated.UserRegisterPostRes, error)
	UserSearchGet(ctx context.Context, params generated.UserSearchGetParams) (generated.UserSearchGetRes, error)
}

type Handler struct {
	LoginHandler
	UserHandler
}

func NewHandler(login LoginHandler, user UserHandler) *Handler {
	return &Handler{
		LoginHandler: login,
		UserHandler:  user,
	}
}

// DialogUserIDListGet implements GET /dialog/{user_id}/list operation.
//
// GET /dialog/{user_id}/list
func (h *Handler) DialogUserIDListGet(ctx context.Context, params generated.DialogUserIDListGetParams) (generated.DialogUserIDListGetRes, error) {
	// TODO
	return &generated.DialogUserIDListGetOKApplicationJSON{}, nil
}

// DialogUserIDSendPost implements POST /dialog/{user_id}/send operation.
//
// POST /dialog/{user_id}/send
func (h *Handler) DialogUserIDSendPost(ctx context.Context, req generated.OptDialogUserIDSendPostReq, params generated.DialogUserIDSendPostParams) (generated.DialogUserIDSendPostRes, error) {
	// TODO
	return &generated.DialogUserIDSendPostOK{}, nil
}

// FriendDeleteUserIDPut implements PUT /friend/delete/{user_id} operation.
//
// PUT /friend/delete/{user_id}
func (h *Handler) FriendDeleteUserIDPut(ctx context.Context, params generated.FriendDeleteUserIDPutParams) (generated.FriendDeleteUserIDPutRes, error) {
	// TODO
	return &generated.FriendDeleteUserIDPutOK{}, nil
}

// FriendSetUserIDPut implements PUT /friend/set/{user_id} operation.
//
// PUT /friend/set/{user_id}
func (h *Handler) FriendSetUserIDPut(ctx context.Context, params generated.FriendSetUserIDPutParams) (generated.FriendSetUserIDPutRes, error) {
	// TODO
	return &generated.FriendSetUserIDPutOK{}, nil
}

// PostCreatePost implements POST /post/create operation.
//
// POST /post/create
func (h *Handler) PostCreatePost(ctx context.Context, req generated.OptPostCreatePostReq) (generated.PostCreatePostRes, error) {
	// TODO
	return &generated.PostCreatePostInternalServerError{}, nil
}

// PostDeleteIDPut implements PUT /post/delete/{id} operation.
//
// PUT /post/delete/{id}
func (h *Handler) PostDeleteIDPut(ctx context.Context, params generated.PostDeleteIDPutParams) (generated.PostDeleteIDPutRes, error) {
	// TODO
	return &generated.PostDeleteIDPutOK{}, nil
}

// PostFeedGet implements GET /post/feed operation.
//
// GET /post/feed
func (h *Handler) PostFeedGet(ctx context.Context, params generated.PostFeedGetParams) (generated.PostFeedGetRes, error) {
	// TODO
	return &generated.PostFeedGetOKApplicationJSON{}, nil
}

// PostGetIDGet implements GET /post/get/{id} operation.
//
// GET /post/get/{id}
func (h *Handler) PostGetIDGet(ctx context.Context, params generated.PostGetIDGetParams) (generated.PostGetIDGetRes, error) {
	// TODO
	return &generated.PostGetIDGetInternalServerError{}, nil
}

// PostUpdatePut implements PUT /post/update operation.
//
// PUT /post/update
func (h *Handler) PostUpdatePut(ctx context.Context, req generated.OptPostUpdatePutReq) (generated.PostUpdatePutRes, error) {
	// TODO
	return &generated.PostUpdatePutOK{}, nil
}

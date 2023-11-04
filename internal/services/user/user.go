package user

import (
	"context"
)

type User interface {
	Register(ctx context.Context)
	//Get()
}

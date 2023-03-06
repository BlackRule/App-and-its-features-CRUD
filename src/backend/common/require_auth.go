package common

import (
	"context"
	"errors"
)

func Require_auth(ctx context.Context) (uint, error) {
	userID := ctx.Value("user_id")
	if userID == nil {
		return 0, errors.New("Unauthorized: Token is invalid")
	}
	return userID.(uint), nil
}

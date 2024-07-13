package utils

import (
	"context"
	"encoding/json"
)

func GetUserIDFromContext(ctx context.Context) (userID int64, err error) {
	ret, err := ctx.Value(("userID")).(json.Number).Int64()
	if err != nil {
		return 0, err
	}
	return ret, nil
}

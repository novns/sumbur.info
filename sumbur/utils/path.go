package utils

import (
	"net/url"

	"github.com/savsgio/atreugo/v11"
)

var EmptyString = ""

func StrPathValue(ctx *atreugo.RequestCtx, key string) *string {
	value := ctx.UserValue(key)
	if value == nil {
		return &EmptyString
	}

	result, err := url.PathUnescape(value.(string))
	if err != nil {
		return &EmptyString
	}

	return &result
}

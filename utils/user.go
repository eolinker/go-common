package utils

import (
	"context"

	"github.com/gin-gonic/gin"
)

const contextUserIdKey = "user_id:context"
const contextI18nKey = "i18n:context"

func UserId(ctx context.Context) string {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetString(contextUserIdKey)
	}

	v := ctx.Value(contextUserIdKey)
	if v == nil {
		return ""
	}
	return v.(string)
}
func SetUserId(ctx context.Context, userId string) context.Context {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(contextUserIdKey, userId)
		return ginCtx
	}

	return context.WithValue(ctx, contextUserIdKey, userId)
}

func I18n(ctx context.Context) string {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetString(contextI18nKey)
	}
	v := ctx.Value(contextI18nKey)
	if v == nil {
		return ""
	}
	return v.(string)
}

func SetI18n(ctx context.Context, i18n string) context.Context {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(contextI18nKey, i18n)
		return ginCtx
	}
	return context.WithValue(ctx, contextI18nKey, i18n)
}

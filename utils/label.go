package utils

import (
	"context"

	"github.com/gin-gonic/gin"
)

const contextGatewayInvokeKey = "gateway_invoke:context"

func Label(ctx context.Context, label string) string {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		return ginCtx.GetString(label)
	}
	v := ctx.Value(label)
	if v == nil {
		return ""
	}
	return v.(string)
}

func SetLabel(ctx context.Context, label string, value string) context.Context {
	if ginCtx, ok := ctx.(*gin.Context); ok {
		ginCtx.Set(label, value)
		return ginCtx
	}
	return context.WithValue(ctx, label, value)
}

func GatewayInvoke(ctx context.Context) string {
	return Label(ctx, contextGatewayInvokeKey)
}

func SetGatewayInvoke(ctx context.Context, value string) context.Context {
	return SetLabel(ctx, contextGatewayInvokeKey, value)
}

const contextI18nKey = "i18n:context"

func I18n(ctx context.Context) string {
	return Label(ctx, contextI18nKey)
}

func SetI18n(ctx context.Context, i18n string) context.Context {
	return SetLabel(ctx, contextI18nKey, i18n)
}

package adapter

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type RPCHandlerFunc[Body any, Response any] func(ctx context.Context, in *Body, opts ...grpc.CallOption) (*Response, error)

type InExtractor[T any] func(ctx *gin.Context) (*T, error)

func WithEmpty[T any](_ *gin.Context) (*T, error) {
	return new(T), nil
}

func WithBodyExtractor[T any](ctx *gin.Context) (*T, error) {
	var in T

	if err := ctx.Bind(&in); err != nil {
		return nil, err
	}

	return &in, nil
}

func Wrap[T any, R any](h RPCHandlerFunc[T, R], inExt InExtractor[T]) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		in, err := inExt(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		}

		resp, err := h(ctx, in)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadGateway, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}

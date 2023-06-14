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

func WrapWithStatus[T any, R any](handler RPCHandlerFunc[T, R], inExtractor InExtractor[T], statusSuccess int) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		in, err := inExtractor(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
			return
		}

		resp, err := handler(ctx, in)

		if err != nil {
			status, msg := ToHTTPError(err)
			ctx.AbortWithStatusJSON(status, msg)
			return
		}

		ctx.JSON(statusSuccess, resp)
	}
}

func Wrap[T any, R any](h RPCHandlerFunc[T, R], inExt InExtractor[T]) gin.HandlerFunc {
	return WrapWithStatus[T, R](h, inExt, http.StatusOK)
}

package adapter

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ToHTTPError(err error) (int, string) {
	s := status.Convert(err)

	if s == nil {
		return http.StatusInternalServerError, "Failed to convert error"
	}

	var errCode int

	switch s.Code() {
	case codes.OK:
		errCode = http.StatusOK
	case codes.Canceled:
		errCode = 499
	case codes.Unknown:
		errCode = http.StatusInternalServerError
	case codes.InvalidArgument:
		errCode = http.StatusBadRequest
	case codes.DeadlineExceeded:
		errCode = http.StatusGatewayTimeout
	case codes.NotFound:
		errCode = http.StatusNotFound
	case codes.AlreadyExists:
		errCode = http.StatusConflict
	case codes.PermissionDenied:
		errCode = http.StatusForbidden
	case codes.Unauthenticated:
		errCode = http.StatusUnauthorized
	case codes.ResourceExhausted:
		errCode = http.StatusTooManyRequests
	case codes.FailedPrecondition:
		errCode = http.StatusBadRequest
	case codes.Aborted:
		errCode = http.StatusConflict
	case codes.OutOfRange:
		errCode = http.StatusBadRequest
	case codes.Unimplemented:
		errCode = http.StatusNotImplemented
	case codes.Internal:
		errCode = http.StatusInternalServerError
	case codes.Unavailable:
		errCode = http.StatusServiceUnavailable
	case codes.DataLoss:
		errCode = http.StatusInternalServerError
	}

	return errCode, s.Message()
}

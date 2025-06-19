package openapi

import (
	"net/http"

	"github.com/sadhasiva1984/openapi/models"
)

func ProblemDetailsSystemFailure(detail string) *models.ProblemDetails {
	return &models.ProblemDetails{
		Title:  "System failure",
		Status: http.StatusInternalServerError,
		Detail: detail,
		Cause:  "SYSTEM_FAILURE",
	}
}

func ProblemDetailsOperationNotSupported() *models.ProblemDetails {
	return &models.ProblemDetails{
		Title:  "Operation not supported",
		Status: http.StatusNotImplemented,
		Cause:  "OPERATION_NOT_SUPPORTED",
	}
}

func ProblemDetailsMalformedReqSyntax(detail string) *models.ProblemDetails {
	return &models.ProblemDetails{
		Title:  "Malformed request syntax",
		Status: http.StatusBadRequest,
		Detail: detail,
	}
}

func ProblemDetailsDataNotFound(detail string) *models.ProblemDetails {
	return &models.ProblemDetails{
		Title:  "Data not found",
		Status: http.StatusNotFound,
		Detail: detail,
	}
}

func ProblemDetailsForbidden(detail, cause string) *models.ProblemDetails {
	return &models.ProblemDetails{
		Title:  "Forbidden",
		Status: http.StatusForbidden,
		Detail: detail,
		Cause:  cause,
	}
}

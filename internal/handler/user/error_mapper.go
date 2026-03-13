package user

import (
	"errors"
	"net/http"

	domainUser "github.com/Turgho/GoFlowDesk/internal/domain/user"
	handlerpkg "github.com/Turgho/GoFlowDesk/internal/handler"
)

func HandleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, domainUser.ErrEmailAlreadyExists):
		// email já existe
		(&handlerpkg.APIResponse{
			Message: "conflict",
			Errors:  []string{err.Error()},
		}).Send(w, http.StatusConflict)

	case errors.Is(err, domainUser.ErrEmptyName),
		errors.Is(err, domainUser.ErrEmptyEmail),
		errors.Is(err, domainUser.ErrEmptyPassword),
		errors.Is(err, domainUser.ErrInvalidUserRole):

		// validação de campos
		(&handlerpkg.APIResponse{
			Message: "bad request",
			Errors:  []string{err.Error()},
		}).Send(w, http.StatusBadRequest)

	default:
		// erro inesperado
		(&handlerpkg.APIResponse{
			Message: "internal error",
			Errors:  []string{err.Error()},
		}).Send(w, http.StatusInternalServerError)
	}
}

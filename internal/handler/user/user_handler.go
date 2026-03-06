package user

import (
	"net/http"

	domainUser "github.com/Turgho/GoFlowDesk/internal/domain/user"
	"github.com/Turgho/GoFlowDesk/internal/handler/render"
	usersvc "github.com/Turgho/GoFlowDesk/internal/service/user"
)

// UserHandler handles HTTP requests related to users.
type UserHandler struct {
	userSvc *usersvc.Service
}

// NewUserHandler constructs a handler.
func NewUserHandler(svc *usersvc.Service) *UserHandler {
	return &UserHandler{userSvc: svc}
}

// createUserRequest represents input payload for POST /users
type createUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Create handles user creation.
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	req, err := render.ReadJSON[createUserRequest](w, r)
	if err != nil {
		render.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()}, nil)
		return
	}

	u, err := h.userSvc.CreateUser(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		// map domain errors to HTTP codes
		switch err {
		case domainUser.ErrEmailAlreadyExists:
			render.WriteJSON(w, http.StatusConflict, map[string]string{"error": err.Error()}, nil)
		case domainUser.ErrEmptyName, domainUser.ErrEmptyEmail, domainUser.ErrEmptyPassword, domainUser.ErrInvalidUserRole:
			render.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()}, nil)
		default:
			render.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()}, nil)
		}
		return
	}

	// avoid returning password hash
	u.PasswordHash = ""

	render.WriteJSON(w, http.StatusCreated, u, nil)
}

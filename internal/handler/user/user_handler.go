package user

import (
	"net/http"

	handlerpkg "github.com/Turgho/GoFlowDesk/internal/handler"
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
	// Leer e validar o payload JSON
	req, err := render.ReadJSON[createUserRequest](w, r)

	if err != nil {
		// payload inválido ou leitura falhou
		(&handlerpkg.APIResponse{
			Message: "bad request",
			Errors:  []string{err.Error()},
		}).Send(w, http.StatusBadRequest)
		return
	}

	u, err := h.userSvc.CreateUser(r.Context(), req.Name, req.Email, req.Password)
	if err != nil {
		// mapeamento de erros de domínio para respostas HTTP
		HandleError(w, err)
		return
	}

	// avoid returning password hash
	u.PasswordHash = ""

	(&handlerpkg.APIResponse{
		Data:    u,
		Message: "created",
	}).Send(w, http.StatusCreated)
}

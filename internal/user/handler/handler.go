package handler

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/vetrovegor/kushfinds-backend/internal/apperror"
	"github.com/vetrovegor/kushfinds-backend/internal/auth/jwt"
	"github.com/vetrovegor/kushfinds-backend/internal/handlers"
	"github.com/vetrovegor/kushfinds-backend/internal/user"
	"go.uber.org/zap"
)

type Service interface {
	GetByID(ctx context.Context, id int) (*user.User, error)
}

type handler struct {
	service        Service
	authMiddleware func(http.Handler) http.Handler
	logger         *zap.Logger
}

func New(service Service, authMiddleware func(http.Handler) http.Handler, logger *zap.Logger) handlers.Handler {
	return &handler{
		service:        service,
		authMiddleware: authMiddleware,
		logger:         logger,
	}
}

func (h *handler) Register(router chi.Router) {
	router.Group(func(privateUserRouter chi.Router) {
		privateUserRouter.Use(h.authMiddleware)

		privateUserRouter.Get("/users/me", apperror.Middleware(h.userHandler))
	})
}

//	@Tags		user
//	@Success	200		{object}	UserResponse
//	@Failure	400,500	{object}	apperror.AppError
//	@Router		/users/me [get]
func (h *handler) userHandler(w http.ResponseWriter, r *http.Request) error {
	userID := r.Context().Value(jwtauth.UserIDContextKey{}).(int)

	existingUser, err := h.service.GetByID(r.Context(), userID)
	if err != nil {
		return err
	}

	render.JSON(w, r, user.UserResponse{User: *existingUser})

	return nil
}

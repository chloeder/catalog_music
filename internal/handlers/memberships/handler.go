package memberships

import (
	"catalog-music/internal/models/memberships"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=handler.go -destination=handler_mock_test.go -package=memberships
type membershipService interface {
	SignUp(req *memberships.SignUpRequest) error
}

type Handler struct {
	*gin.Engine
	membershipService membershipService
}

func NewHandler(api *gin.Engine, membershipService membershipService) *Handler {
	return &Handler{
		api,
		membershipService,
	}
}

func (h *Handler) AuthRoute() {
	auth := h.Group("/auth")
	auth.POST("/signup", h.SignUp)
}

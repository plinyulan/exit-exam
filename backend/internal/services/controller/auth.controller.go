package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plinyulan/exit-exam/internal/conf"
	"github.com/plinyulan/exit-exam/internal/services/types"
	"github.com/plinyulan/exit-exam/internal/services/usecase"
)

type AuthController struct {
	uc usecase.AuthUsecase
}

func NewAuthController(cfg conf.Config, uc usecase.AuthUsecase) *AuthController {
	return &AuthController{uc: uc}
}

// AuthRoutes sets up the routes for authentication.
// @Summary      Auth Routes
// @Description  Sets up the routes for authentication
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        loginRequest  body      types.LoginRequest  true  "Login Request"
// @Success      200           {object}  types.LoginResponse
// @Failure      400           {object}  map[string]string
// @Failure      401           {object}  map[string]string
// @Router       /auth/login [post]
func (ac *AuthController) LoginUser(c *gin.Context) {
	var loginReq types.LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResp, err := ac.uc.LoginUserUsecase(c.Request.Context(), &loginReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loginResp)
}

func (ac *AuthController) AuthRoutes(r gin.IRouter) {
	r.POST("/login", ac.LoginUser)
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plinyulan/exit-exam/internal/services/usecase"
)

type CatController struct {
	catUsecase usecase.CatUsecase
}

func NewCatController(catUsecase usecase.CatUsecase) *CatController {
	return &CatController{
		catUsecase: catUsecase,
	}
}

// GetCats godoc
// @Summary      Get Cats
// @Description  get list of cats
// @Tags         cats
// @Accept       json
// @Produce      json
// @Success      200  {array}   string
// @Router       /cats/ [get]
func (cc *CatController) GetCats(c *gin.Context) {
	response := cc.catUsecase.GetCatsUsecase()
	c.JSON(http.StatusOK, response)
}

func (cc *CatController) CatRoutes(r gin.IRoutes) {
	r.GET("/", cc.GetCats)
}

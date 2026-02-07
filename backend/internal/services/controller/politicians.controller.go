package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/plinyulan/exit-exam/internal/services/types"
	"github.com/plinyulan/exit-exam/internal/services/usecase"
)

type PoliticiansController struct {
	polUC usecase.PoliticiansUsecase
	proUC usecase.PromisesUsecase
}

func NewPoliticiansController(polUC usecase.PoliticiansUsecase, proUC usecase.PromisesUsecase) *PoliticiansController {
	return &PoliticiansController{polUC: polUC, proUC: proUC}
}

// List godoc
// @Summary      List Politicians
// @Description  get list of politicians
// @Tags         politicians
// @Accept       json
// @Produce      json
// @Success      200  {array}   types.PoliticianResponse
// @Security     BearerAuth
// @Router       /politicians/ [get]
func (h *PoliticiansController) List(c *gin.Context) {
	items, err := h.polUC.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list politicians"})
		return
	}

	resp := make([]types.PoliticianResponse, 0, len(items))
	for _, p := range items {
		resp = append(resp, types.PoliticianResponse{
			ID:             p.ID,
			PoliticianCode: p.PoliticianCode,
			Name:           p.Name,
			Party:          p.Party,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// PromisesByPolitician godoc
// @Summary      Promises By Politician
// @Description  get list of promises by politician ID
// @Tags         politicians
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Politician ID"
// @Success      200  {array}   types.PromiseResponse
// @Failure      404  {object}  map[string]string
// @Security     BearerAuth
// @Router       /politicians/{id}/promises [get]
func (h *PoliticiansController) PromisesByPolitician(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	items, err := h.proUC.ListByPolitician(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list promises"})
		return
	}

	resp := make([]types.PromiseResponse, 0, len(items))
	for _, p := range items {
		resp = append(resp, toPromiseResponse(p))
	}
	c.JSON(http.StatusOK, resp)
}

func (h *PoliticiansController) PoliticiansRoutes(r gin.IRoutes) {
	r.GET("/", h.List)
	r.GET("/:id/promises", h.PromisesByPolitician)
}

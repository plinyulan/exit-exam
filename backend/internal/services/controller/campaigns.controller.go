package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/plinyulan/exit-exam/internal/services/types"
	"github.com/plinyulan/exit-exam/internal/services/usecase"
)

type CampaignsController struct {
	uc usecase.CampaignsUsecase
}

func NewCampaignsController(uc usecase.CampaignsUsecase) *CampaignsController {
	return &CampaignsController{uc: uc}
}

// List godoc
// @Summary      List Campaigns
// @Description  get list of campaigns
// @Tags         campaigns
// @Accept       json
// @Produce      json
// @Success      200  {array}   types.CampaignResponse
// @Security     BearerAuth
// @Router       /campaigns/ [get]
func (h *CampaignsController) List(c *gin.Context) {
	items, err := h.uc.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list campaigns"})
		return
	}

	resp := make([]types.CampaignResponse, 0, len(items))
	for _, x := range items {
		resp = append(resp, types.CampaignResponse{ID: x.ID, Year: x.Year, District: x.District})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *CampaignsController) CampaignsRoutes(r gin.IRoutes) {
	r.GET("/", h.List)
}

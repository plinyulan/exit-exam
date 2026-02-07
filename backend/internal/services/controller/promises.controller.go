package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/plinyulan/exit-exam/internal/model"
	"github.com/plinyulan/exit-exam/internal/services/types"
	"github.com/plinyulan/exit-exam/internal/services/usecase"
)

type PromisesController struct {
	uc usecase.PromisesUsecase
}

func NewPromisesController(uc usecase.PromisesUsecase) *PromisesController {
	return &PromisesController{uc: uc}
}

// ListAll godoc
// @Summary      List All Promises
// @Description  get list of all promises
// @Tags         promises
// @Accept       json
// @Produce      json
// @Success      200  {array}   types.PromiseResponse
// @Security     BearerAuth
// @Router       /promises/all [get]
func (h *PromisesController) ListAll(c *gin.Context) {
	items, err := h.uc.ListAll()
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

// Detail godoc
// @Summary      Promise Detail
// @Description  get promise detail by ID
// @Tags         promises
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Promise ID"
// @Success      200  {object}  types.PromiseDetailResponse
// @Failure      404  {object}  map[string]string
// @Security     BearerAuth
// @Router       /promises/{id} [get]
func (h *PromisesController) Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	p, err := h.uc.GetDetail(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "promise not found"})
		return
	}

	out := types.PromiseDetailResponse{
		PromiseResponse: toPromiseResponse(*p),
		Updates:         make([]types.PromiseUpdateResponse, 0, len(p.Updates)),
	}

	for _, u := range p.Updates {
		out.Updates = append(out.Updates, types.PromiseUpdateResponse{
			ID:        u.ID,
			UpdatedAt: u.UpdatedAt,
			Note:      u.Note,
		})
	}

	c.JSON(http.StatusOK, out)
}

// AddUpdate godoc
// @Summary      Add Promise Update
// @Description  add an update to a promise (admin only)
// @Tags         promises
// @Accept       json
// @Produce      json
// @Param        id                        path      int                                 true  "Promise ID"
// @Param        createPromiseUpdateRequest  body      types.CreatePromiseUpdateRequest    true  "Create Promise Update Request"
// @Success      200                       {object}  map[string]string
// @Failure      400                       {object}  map[string]string
// @Failure      403                       {object}  map[string]string
// @Security     BearerAuth
// @Router       /promises/{id}/updates [post]
func (h *PromisesController) AddUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req types.CreatePromiseUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.uc.AddUpdate(uint(id), req.UpdatedAt, req.Note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "updated",
		"redirect_to": "/promise.html?id=" + strconv.Itoa(id),
	})
}

func toPromiseResponse(p model.Promise) types.PromiseResponse {
	return types.PromiseResponse{
		ID:          p.ID,
		Detail:      p.Detail,
		AnnouncedAt: p.AnnouncedAt,
		Status:      string(p.Status),
		Politician: types.PoliticianResponse{
			ID:             p.Politician.ID,
			PoliticianCode: p.Politician.PoliticianCode,
			Name:           p.Politician.Name,
			Party:          p.Politician.Party,
		},
		Campaign: types.CampaignResponse{
			ID:       p.Campaign.ID,
			Year:     p.Campaign.Year,
			District: p.Campaign.District,
		},
	}
}

func (h *PromisesController) PromisesRoutes(r gin.IRoutes) {
	r.GET("/all", h.ListAll)
	r.GET("/:id", h.Detail)
	r.POST("/:id/updates", h.AddUpdate)
}

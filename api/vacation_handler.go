package api

import (
	"net/http"
	"strconv"

	"github.com/doutorfinancas/natural-stupidity/api/request"
	"github.com/doutorfinancas/natural-stupidity/api/response"
	"github.com/doutorfinancas/natural-stupidity/repository/model"
	"github.com/doutorfinancas/natural-stupidity/service"
	"github.com/gin-gonic/gin"
)

// RegisterVacationRoutes registers vacation endpoints under the given router group.
func RegisterVacationRoutes(rg *gin.RouterGroup, svc service.VacationService) {
	rg.POST("", createVacationHandler(svc))
	rg.GET("", listVacationsHandler(svc))
	rg.GET(":id", getVacationHandler(svc))
	rg.GET("/user/:user_id", listVacationsByUserHandler(svc))
	rg.PUT(":id", updateVacationHandler(svc))
	rg.DELETE(":id", deleteVacationHandler(svc))
}

func createVacationHandler(svc service.VacationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.CreateVacationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		vac := &model.Vacation{
			UserID:    req.UserID,
			StartDate: req.StartDate,
			EndDate:   req.EndDate,
		}
		if err := svc.Create(vac); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, response.VacationResponse{
			ID:        vac.ID,
			UserID:    vac.UserID,
			StartDate: vac.StartDate,
			EndDate:   vac.EndDate,
		})
	}
}

func listVacationsHandler(svc service.VacationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		vacs, err := svc.List()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var resp []response.VacationResponse
		for _, v := range vacs {
			resp = append(resp, response.VacationResponse{
				ID:        v.ID,
				UserID:    v.UserID,
				StartDate: v.StartDate,
				EndDate:   v.EndDate,
			})
		}
		c.JSON(http.StatusOK, resp)
	}
}

func listVacationsByUserHandler(svc service.VacationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userParam := c.Param("user_id")
		userID, err := strconv.ParseUint(userParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
			return
		}
		vacs, err := svc.ListByUser(uint(userID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var resp []response.VacationResponse
		for _, v := range vacs {
			resp = append(resp, response.VacationResponse{
				ID:        v.ID,
				UserID:    v.UserID,
				StartDate: v.StartDate,
				EndDate:   v.EndDate,
			})
		}
		c.JSON(http.StatusOK, resp)
	}
}

func getVacationHandler(svc service.VacationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		vac, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.VacationResponse{
			ID:        vac.ID,
			UserID:    vac.UserID,
			StartDate: vac.StartDate,
			EndDate:   vac.EndDate,
		})
	}
}

func updateVacationHandler(svc service.VacationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		var req request.UpdateVacationRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		vac, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if req.UserID != nil {
			vac.UserID = *req.UserID
		}
		if req.StartDate != nil {
			vac.StartDate = *req.StartDate
		}
		if req.EndDate != nil {
			vac.EndDate = *req.EndDate
		}
		if err := svc.Update(vac); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.VacationResponse{
			ID:        vac.ID,
			UserID:    vac.UserID,
			StartDate: vac.StartDate,
			EndDate:   vac.EndDate,
		})
	}
}

func deleteVacationHandler(svc service.VacationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		if err := svc.Delete(uint(id)); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.Status(http.StatusNoContent)
	}
}

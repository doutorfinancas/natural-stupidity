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

// RegisterDirectionRoutes registers direction endpoints under the given router group.
func RegisterDirectionRoutes(rg *gin.RouterGroup, svc service.DirectionService) {
	rg.POST("", createDirectionHandler(svc))
	rg.GET("", listDirectionsHandler(svc))
	rg.GET(":id", getDirectionHandler(svc))
	rg.PUT(":id", updateDirectionHandler(svc))
	rg.DELETE(":id", deleteDirectionHandler(svc))
}

func createDirectionHandler(svc service.DirectionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.CreateDirectionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dir := &model.Direction{
			Name:       req.Name,
			DirectorID: req.DirectorID,
		}
		if err := svc.Create(dir); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, response.DirectionResponse{
			ID:         dir.ID,
			Name:       dir.Name,
			DirectorID: dir.DirectorID,
		})
	}
}

func listDirectionsHandler(svc service.DirectionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		dirs, err := svc.List()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var resp []response.DirectionResponse
		for _, d := range dirs {
			resp = append(resp, response.DirectionResponse{
				ID:         d.ID,
				Name:       d.Name,
				DirectorID: d.DirectorID,
			})
		}
		c.JSON(http.StatusOK, resp)
	}
}

func getDirectionHandler(svc service.DirectionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		dir, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.DirectionResponse{
			ID:         dir.ID,
			Name:       dir.Name,
			DirectorID: dir.DirectorID,
		})
	}
}

func updateDirectionHandler(svc service.DirectionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		var req request.UpdateDirectionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		dir, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if req.Name != nil {
			dir.Name = *req.Name
		}
		if req.DirectorID != nil {
			dir.DirectorID = *req.DirectorID
		}
		if err := svc.Update(dir); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.DirectionResponse{
			ID:         dir.ID,
			Name:       dir.Name,
			DirectorID: dir.DirectorID,
		})
	}
}

func deleteDirectionHandler(svc service.DirectionService) gin.HandlerFunc {
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

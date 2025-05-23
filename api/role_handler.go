package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/doutorfinancas/natural-stupidity/api/request"
	"github.com/doutorfinancas/natural-stupidity/api/response"
	"github.com/doutorfinancas/natural-stupidity/repository/model"
	"github.com/doutorfinancas/natural-stupidity/service"
)

// RegisterRoleRoutes registers role endpoints under the given router group.
func RegisterRoleRoutes(rg *gin.RouterGroup, svc service.RoleService) {
	rg.POST("", createRoleHandler(svc))
	rg.GET("", listRolesHandler(svc))
	rg.GET(":id", getRoleHandler(svc))
	rg.PUT(":id", updateRoleHandler(svc))
	rg.DELETE(":id", deleteRoleHandler(svc))
}

func createRoleHandler(svc service.RoleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.CreateRoleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		role := &model.Role{
			Name:  req.Name,
			Level: req.Level,
		}
		if err := svc.Create(role); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, response.RoleResponse{
			ID:    role.ID,
			Name:  role.Name,
			Level: role.Level,
		})
	}
}

func listRolesHandler(svc service.RoleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, err := svc.List()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var resp []response.RoleResponse
		for _, r := range roles {
			resp = append(resp, response.RoleResponse{
				ID:    r.ID,
				Name:  r.Name,
				Level: r.Level,
			})
		}
		c.JSON(http.StatusOK, resp)
	}
}

func getRoleHandler(svc service.RoleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		role, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.RoleResponse{
			ID:    role.ID,
			Name:  role.Name,
			Level: role.Level,
		})
	}
}

func updateRoleHandler(svc service.RoleService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		var req request.UpdateRoleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		role, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if req.Name != nil {
			role.Name = *req.Name
		}
		if req.Level != nil {
			role.Level = *req.Level
		}
		if err := svc.Update(role); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.RoleResponse{
			ID:    role.ID,
			Name:  role.Name,
			Level: role.Level,
		})
	}
}

func deleteRoleHandler(svc service.RoleService) gin.HandlerFunc {
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

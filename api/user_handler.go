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

// RegisterUserRoutes registers user endpoints under the given router group.
func RegisterUserRoutes(rg *gin.RouterGroup, svc service.UserService) {
	rg.POST("", createUserHandler(svc))
	rg.GET("", listUsersHandler(svc))
	rg.GET(":id", getUserHandler(svc))
	rg.PUT(":id", updateUserHandler(svc))
	rg.DELETE(":id", deleteUserHandler(svc))
}

func createUserHandler(svc service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.CreateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user := &model.User{
			Name:   req.Name,
			Email:  req.Email,
			RoleID: req.RoleID,
			TeamID: req.TeamID,
		}
		if err := svc.Create(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, response.UserResponse{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			RoleID: user.RoleID,
			TeamID: user.TeamID,
		})
	}
}

func listUsersHandler(svc service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := svc.List()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var resp []response.UserResponse
		for _, u := range users {
			resp = append(resp, response.UserResponse{
				ID:     u.ID,
				Name:   u.Name,
				Email:  u.Email,
				RoleID: u.RoleID,
				TeamID: u.TeamID,
			})
		}
		c.JSON(http.StatusOK, resp)
	}
}

func getUserHandler(svc service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		user, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.UserResponse{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			RoleID: user.RoleID,
			TeamID: user.TeamID,
		})
	}
}

func updateUserHandler(svc service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		var req request.UpdateUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if req.Name != nil {
			user.Name = *req.Name
		}
		if req.Email != nil {
			user.Email = *req.Email
		}
		if req.RoleID != nil {
			user.RoleID = *req.RoleID
		}
		if req.TeamID != nil {
			user.TeamID = *req.TeamID
		}
		if err := svc.Update(user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.UserResponse{
			ID:     user.ID,
			Name:   user.Name,
			Email:  user.Email,
			RoleID: user.RoleID,
			TeamID: user.TeamID,
		})
	}
}

func deleteUserHandler(svc service.UserService) gin.HandlerFunc {
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

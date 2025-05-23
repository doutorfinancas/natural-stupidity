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

// RegisterTeamRoutes registers team endpoints under the given router group.
func RegisterTeamRoutes(rg *gin.RouterGroup, svc service.TeamService) {
	rg.POST("", createTeamHandler(svc))
	rg.GET("", listTeamsHandler(svc))
	rg.GET(":id", getTeamHandler(svc))
	rg.PUT(":id", updateTeamHandler(svc))
	rg.DELETE(":id", deleteTeamHandler(svc))
}

func createTeamHandler(svc service.TeamService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req request.CreateTeamRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		team := &model.Team{
			Name:        req.Name,
			LeaderID:    req.LeaderID,
			DirectionID: req.DirectionID,
		}
		if err := svc.Create(team); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, response.TeamResponse{
			ID:          team.ID,
			Name:        team.Name,
			LeaderID:    team.LeaderID,
			DirectionID: team.DirectionID,
		})
	}
}

func listTeamsHandler(svc service.TeamService) gin.HandlerFunc {
	return func(c *gin.Context) {
		teams, err := svc.List()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var resp []response.TeamResponse
		for _, t := range teams {
			resp = append(resp, response.TeamResponse{
				ID:          t.ID,
				Name:        t.Name,
				LeaderID:    t.LeaderID,
				DirectionID: t.DirectionID,
			})
		}
		c.JSON(http.StatusOK, resp)
	}
}

func getTeamHandler(svc service.TeamService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		team, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.TeamResponse{
			ID:          team.ID,
			Name:        team.Name,
			LeaderID:    team.LeaderID,
			DirectionID: team.DirectionID,
		})
	}
}

func updateTeamHandler(svc service.TeamService) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
			return
		}
		var req request.UpdateTeamRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		team, err := svc.GetByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if req.Name != nil {
			team.Name = *req.Name
		}
		if req.LeaderID != nil {
			team.LeaderID = *req.LeaderID
		}
		if req.DirectionID != nil {
			team.DirectionID = *req.DirectionID
		}
		if err := svc.Update(team); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, response.TeamResponse{
			ID:          team.ID,
			Name:        team.Name,
			LeaderID:    team.LeaderID,
			DirectionID: team.DirectionID,
		})
	}
}

func deleteTeamHandler(svc service.TeamService) gin.HandlerFunc {
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

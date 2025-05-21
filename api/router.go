package api

import (
	"github.com/doutorfinancas/natural-stupidity/service"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up all API routes under the given engine.
func RegisterRoutes(r *gin.Engine,
	userSvc service.UserService,
	roleSvc service.RoleService,
	teamSvc service.TeamService,
	dirSvc service.DirectionService,
	vacSvc service.VacationService,
) {
	// Healthcheck
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// User endpoints
	userGroup := r.Group("/users")
	RegisterUserRoutes(userGroup, userSvc)

	// Role endpoints
	roleGroup := r.Group("/roles")
	RegisterRoleRoutes(roleGroup, roleSvc)

	// Team endpoints
	teamGroup := r.Group("/teams")
	RegisterTeamRoutes(teamGroup, teamSvc)

	// Direction endpoints
	dirGroup := r.Group("/directions")
	RegisterDirectionRoutes(dirGroup, dirSvc)

	// Vacation endpoints
	vacGroup := r.Group("/vacations")
	RegisterVacationRoutes(vacGroup, vacSvc)
}

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.uber.org/dig"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/doutorfinancas/natural-stupidity/api"
	"github.com/doutorfinancas/natural-stupidity/repository"
	"github.com/doutorfinancas/natural-stupidity/repository/model"
	"github.com/doutorfinancas/natural-stupidity/service"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, proceeding with environment variables: %v", err)
	}
}

// parseUint converts string to uint (ignores errors)
func parseUint(s string) uint {
	v, _ := strconv.ParseUint(s, 10, 32)
	return uint(v)
}

// hashPassword generates a bcrypt hash of the given password
func hashPassword(pwd string) string {
	h, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(h)
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Initialize GORM/MySQL
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(
		&model.Role{},
		&model.User{},
		&model.Team{},
		&model.Direction{},
		&model.Vacation{},
	); err != nil {
		log.Fatalf("Failed to auto-migrate schema: %v", err)
	}

	// Seed initial data if empty
	var rCount int64
	if err := db.Model(&model.Role{}).Count(&rCount).Error; err != nil {
		log.Fatalf("Failed to count roles: %v", err)
	}
	if rCount == 0 {
		adminRole := model.Role{Name: "Administrator", Level: 1}
		if err := db.Create(&adminRole).Error; err != nil {
			log.Fatalf("Failed to create default role: %v", err)
		}
		log.Printf("Seeded default role: %v", adminRole)
	}
	var uCount int64
	if err := db.Model(&model.User{}).Count(&uCount).Error; err != nil {
		log.Fatalf("Failed to count users: %v", err)
	}
	if uCount == 0 {
		var defaultRole model.Role
		if err := db.Where("name = ?", "Administrator").First(&defaultRole).Error; err != nil {
			log.Fatalf("Failed to find default role: %v", err)
		}
		// Generate default password
		defaultPassword := "123456"
		hash, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Failed to hash default password: %v", err)
		}
		adminUser := model.User{Name: "Admin", Email: "admin@local.com", RoleID: defaultRole.ID, TeamID: 0, PasswordHash: string(hash)}
		if err := db.Create(&adminUser).Error; err != nil {
			log.Fatalf("Failed to create default user: %v", err)
		}
		log.Printf("Seeded default user: %v (password: %s)", adminUser.Email, defaultPassword)
	}

	// Set up DI container
	container := dig.New()

	// Provide database instance
	container.Provide(func() *gorm.DB { return db })

	// Provide repositories
	container.Provide(repository.NewRoleRepository)
	container.Provide(repository.NewUserRepository)
	container.Provide(repository.NewTeamRepository)
	container.Provide(repository.NewDirectionRepository)
	container.Provide(repository.NewVacationRepository)

	// Provide services
	container.Provide(service.NewRoleService)
	container.Provide(service.NewUserService)
	container.Provide(service.NewTeamService)
	container.Provide(service.NewDirectionService)
	container.Provide(service.NewVacationService)
	container.Provide(service.NewAuthService)

	// Provide router with handlers via DI, including HTML views
	container.Provide(func(
		userSvc service.UserService,
		roleSvc service.RoleService,
		teamSvc service.TeamService,
		dirSvc service.DirectionService,
		vacSvc service.VacationService,
		authSvc service.AuthService,
	) *gin.Engine {
		r := gin.Default()
		// Load HTML templates
		r.LoadHTMLGlob("templates/*.html")

		// Authentication routes
		api.RegisterAuthRoutes(r, authSvc)

		// Protected HTML view routes
		authGroup := r.Group("/", api.AuthRequired())
		authGroup.GET("/", func(c *gin.Context) {
			// Build events for calendar
			vacs, err := vacSvc.List()
			if err != nil {
				c.String(500, err.Error())
				return
			}
			users, _ := userSvc.List()
			userMap := make(map[uint]string)
			for _, u := range users {
				userMap[u.ID] = u.Name
			}
			type Event struct {
				Title string
				Start string
				End   string
				Color string
			}
			var events []Event
			for _, v := range vacs {
				events = append(events, Event{
					Title: userMap[v.UserID],
					Start: v.StartDate.Format("2006-01-02"),
					End:   v.EndDate.Format("2006-01-02"),
					Color: "#3788d8",
				})
			}
			c.HTML(200, "index.html", gin.H{"Events": events})
		})
		authGroup.GET("/ui/users", func(c *gin.Context) {
			users, err := userSvc.List()
			if err != nil {
				c.String(500, err.Error())
				return
			}
			roles, _ := roleSvc.List()
			teams, _ := teamSvc.List()
			roleMap := make(map[uint]string)
			for _, r := range roles {
				roleMap[r.ID] = r.Name
			}
			teamMap := make(map[uint]string)
			for _, t := range teams {
				teamMap[t.ID] = t.Name
			}
			type UserView struct {
				ID       uint
				Name     string
				Email    string
				RoleName string
				TeamName string
			}
			var userViews []UserView
			for _, u := range users {
				userViews = append(userViews, UserView{
					ID:       u.ID,
					Name:     u.Name,
					Email:    u.Email,
					RoleName: roleMap[u.RoleID],
					TeamName: teamMap[u.TeamID],
				})
			}
			c.HTML(200, "users.html", gin.H{"Users": userViews})
		})
		authGroup.GET("/ui/users/new", func(c *gin.Context) {
			roles, _ := roleSvc.List()
			teams, _ := teamSvc.List()
			c.HTML(200, "new_user.html", gin.H{"Roles": roles, "Teams": teams})
		})
		authGroup.POST("/ui/users/new", func(c *gin.Context) {
			name := c.PostForm("name")
			email := c.PostForm("email")
			password := c.PostForm("password")
			roleID := c.PostForm("role_id")
			teamID := c.PostForm("team_id")
			if err := userSvc.Create(&model.User{Name: name, Email: email, RoleID: parseUint(roleID), TeamID: parseUint(teamID), PasswordHash: hashPassword(password)}); err != nil {
				roles, _ := roleSvc.List()
				teams, _ := teamSvc.List()
				c.HTML(400, "new_user.html", gin.H{"Error": err.Error(), "Roles": roles, "Teams": teams})
				return
			}
			c.Redirect(302, "/ui/users")
		})
		authGroup.GET("/ui/users/:id/edit", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			user, _ := userSvc.GetByID(id)
			roles, _ := roleSvc.List()
			teams, _ := teamSvc.List()
			c.HTML(200, "edit_user.html", gin.H{"User": user, "Roles": roles, "Teams": teams})
		})
		authGroup.POST("/ui/users/:id/edit", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			user, _ := userSvc.GetByID(id)
			if name := c.PostForm("name"); name != "" {
				user.Name = name
			}
			if email := c.PostForm("email"); email != "" {
				user.Email = email
			}
			if pwd := c.PostForm("password"); pwd != "" {
				user.PasswordHash = hashPassword(pwd)
			}
			if rid := c.PostForm("role_id"); rid != "" {
				user.RoleID = parseUint(rid)
			}
			if tid := c.PostForm("team_id"); tid != "" {
				user.TeamID = parseUint(tid)
			}
			if err := userSvc.Update(user); err != nil {
				roles, _ := roleSvc.List()
				teams, _ := teamSvc.List()
				c.HTML(400, "edit_user.html", gin.H{"Error": err.Error(), "User": user, "Roles": roles, "Teams": teams})
				return
			}
			c.Redirect(302, "/ui/users")
		})
		authGroup.POST("/ui/users/:id/delete", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			userSvc.Delete(id)
			c.Redirect(302, "/ui/users")
		})
		authGroup.GET("/ui/roles", func(c *gin.Context) {
			roles, err := roleSvc.List()
			if err != nil {
				c.String(500, err.Error())
				return
			}
			c.HTML(200, "roles.html", gin.H{"Roles": roles})
		})
		authGroup.GET("/ui/roles/new", func(c *gin.Context) {
			c.HTML(200, "new_role.html", nil)
		})
		authGroup.POST("/ui/roles/new", func(c *gin.Context) {
			name := c.PostForm("name")
			level := int(parseUint(c.PostForm("level")))
			if err := roleSvc.Create(&model.Role{Name: name, Level: level}); err != nil {
				c.HTML(400, "new_role.html", gin.H{"Error": err.Error()})
				return
			}
			c.Redirect(302, "/ui/roles")
		})
		authGroup.GET("/ui/roles/:id/edit", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			role, _ := roleSvc.GetByID(id)
			c.HTML(200, "edit_role.html", gin.H{"Role": role})
		})
		authGroup.POST("/ui/roles/:id/edit", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			role, _ := roleSvc.GetByID(id)
			if name := c.PostForm("name"); name != "" {
				role.Name = name
			}
			if lvl := c.PostForm("level"); lvl != "" {
				role.Level = int(parseUint(lvl))
			}
			if err := roleSvc.Update(role); err != nil {
				c.HTML(400, "edit_role.html", gin.H{"Error": err.Error(), "Role": role})
				return
			}
			c.Redirect(302, "/ui/roles")
		})
		authGroup.POST("/ui/roles/:id/delete", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			roleSvc.Delete(id)
			c.Redirect(302, "/ui/roles")
		})
		authGroup.GET("/ui/teams", func(c *gin.Context) {
			teams, err := teamSvc.List()
			if err != nil {
				c.String(500, err.Error())
				return
			}
			users, _ := userSvc.List()
			dirs, _ := dirSvc.List()
			userMap := make(map[uint]string)
			for _, u := range users {
				userMap[u.ID] = u.Name
			}
			dirMap := make(map[uint]string)
			for _, d := range dirs {
				dirMap[d.ID] = d.Name
			}
			type TeamView struct {
				ID            uint
				Name          string
				LeaderName    string
				DirectionName string
			}
			var teamViews []TeamView
			for _, t := range teams {
				teamViews = append(teamViews, TeamView{
					ID:            t.ID,
					Name:          t.Name,
					LeaderName:    userMap[t.LeaderID],
					DirectionName: dirMap[t.DirectionID],
				})
			}
			c.HTML(200, "teams.html", gin.H{"Teams": teamViews})
		})
		authGroup.GET("/ui/directions", func(c *gin.Context) {
			dirs, err := dirSvc.List()
			if err != nil {
				c.String(500, err.Error())
				return
			}
			users, _ := userSvc.List()
			userMap := make(map[uint]string)
			for _, u := range users {
				userMap[u.ID] = u.Name
			}
			type DirView struct {
				ID           uint
				Name         string
				DirectorName string
			}
			var dirViews []DirView
			for _, d := range dirs {
				dirViews = append(dirViews, DirView{
					ID:           d.ID,
					Name:         d.Name,
					DirectorName: userMap[d.DirectorID],
				})
			}
			c.HTML(200, "directions.html", gin.H{"Directions": dirViews})
		})
		authGroup.GET("/ui/vacations", func(c *gin.Context) {
			vacs, err := vacSvc.List()
			if err != nil {
				c.String(500, err.Error())
				return
			}
			users, _ := userSvc.List()
			userMap := make(map[uint]string)
			for _, u := range users {
				userMap[u.ID] = u.Name
			}
			type VacView struct {
				ID        uint
				UserName  string
				StartDate time.Time
				EndDate   time.Time
			}
			var vacViews []VacView
			for _, v := range vacs {
				vacViews = append(vacViews, VacView{
					ID:        v.ID,
					UserName:  userMap[v.UserID],
					StartDate: v.StartDate,
					EndDate:   v.EndDate,
				})
			}
			c.HTML(200, "vacations.html", gin.H{"Vacations": vacViews})
		})

		// Calendar UI route
		authGroup.GET("/ui/calendar", func(c *gin.Context) {
			vacs, err := vacSvc.List()
			if err != nil {
				c.String(500, err.Error())
				return
			}
			users, _ := userSvc.List()
			userMap := make(map[uint]string)
			for _, u := range users {
				userMap[u.ID] = u.Name
			}
			type Event struct {
				Title string
				Start string
				End   string
				Color string
			}
			var events []Event
			for _, v := range vacs {
				events = append(events, Event{
					Title: userMap[v.UserID],
					Start: v.StartDate.Format("2006-01-02"),
					End:   v.EndDate.Format("2006-01-02"),
					Color: "#3788d8",
				})
			}
			c.HTML(200, "calendar.html", gin.H{"Events": events})
		})

		// Team UI routes
		authGroup.GET("/ui/teams/new", func(c *gin.Context) {
			users, _ := userSvc.List()
			dirs, _ := dirSvc.List()
			c.HTML(200, "new_team.html", gin.H{"Users": users, "Directions": dirs})
		})
		authGroup.POST("/ui/teams/new", func(c *gin.Context) {
			name := c.PostForm("name")
			leaderID := parseUint(c.PostForm("leader_id"))
			directionID := parseUint(c.PostForm("direction_id"))
			if err := teamSvc.Create(&model.Team{Name: name, LeaderID: leaderID, DirectionID: directionID}); err != nil {
				users, _ := userSvc.List()
				dirs, _ := dirSvc.List()
				c.HTML(400, "new_team.html", gin.H{"Error": err.Error(), "Users": users, "Directions": dirs})
				return
			}
			c.Redirect(302, "/ui/teams")
		})
		authGroup.GET("/ui/teams/:id/edit", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			team, _ := teamSvc.GetByID(id)
			users, _ := userSvc.List()
			dirs, _ := dirSvc.List()
			c.HTML(200, "edit_team.html", gin.H{"Team": team, "Users": users, "Directions": dirs})
		})
		authGroup.POST("/ui/teams/:id/edit", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			team, _ := teamSvc.GetByID(id)
			if name := c.PostForm("name"); name != "" {
				team.Name = name
			}
			if lid := c.PostForm("leader_id"); lid != "" {
				team.LeaderID = parseUint(lid)
			}
			if did := c.PostForm("direction_id"); did != "" {
				team.DirectionID = parseUint(did)
			}
			if err := teamSvc.Update(team); err != nil {
				users, _ := userSvc.List()
				dirs, _ := dirSvc.List()
				c.HTML(400, "edit_team.html", gin.H{"Error": err.Error(), "Team": team, "Users": users, "Directions": dirs})
				return
			}
			c.Redirect(302, "/ui/teams")
		})
		authGroup.POST("/ui/teams/:id/delete", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			teamSvc.Delete(id)
			c.Redirect(302, "/ui/teams")
		})

		// Direction UI routes
		authGroup.GET("/ui/directions/new", func(c *gin.Context) {
			users, _ := userSvc.List()
			c.HTML(200, "new_direction.html", gin.H{"Users": users})
		})
		authGroup.POST("/ui/directions/new", func(c *gin.Context) {
			name := c.PostForm("name")
			directorID := parseUint(c.PostForm("director_id"))
			if err := dirSvc.Create(&model.Direction{Name: name, DirectorID: directorID}); err != nil {
				users, _ := userSvc.List()
				c.HTML(400, "new_direction.html", gin.H{"Error": err.Error(), "Users": users})
				return
			}
			c.Redirect(302, "/ui/directions")
		})
		authGroup.GET("/ui/directions/:id/edit", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			direction, _ := dirSvc.GetByID(id)
			users, _ := userSvc.List()
			c.HTML(200, "edit_direction.html", gin.H{"Direction": direction, "Users": users})
		})
		authGroup.POST("/ui/directions/:id/edit", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			direction, _ := dirSvc.GetByID(id)
			if name := c.PostForm("name"); name != "" {
				direction.Name = name
			}
			if did := c.PostForm("director_id"); did != "" {
				direction.DirectorID = parseUint(did)
			}
			if err := dirSvc.Update(direction); err != nil {
				users, _ := userSvc.List()
				c.HTML(400, "edit_direction.html", gin.H{"Error": err.Error(), "Direction": direction, "Users": users})
				return
			}
			c.Redirect(302, "/ui/directions")
		})
		authGroup.POST("/ui/directions/:id/delete", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			dirSvc.Delete(id)
			c.Redirect(302, "/ui/directions")
		})

		// Vacation UI routes
		authGroup.GET("/ui/vacations/new", func(c *gin.Context) {
			users, _ := userSvc.List()
			c.HTML(200, "new_vacation.html", gin.H{"Users": users})
		})
		authGroup.POST("/ui/vacations/new", func(c *gin.Context) {
			userID := parseUint(c.PostForm("user_id"))
			startDate := c.PostForm("start_date")
			endDate := c.PostForm("end_date")
			sDate, _ := time.Parse("2006-01-02", startDate)
			eDate, _ := time.Parse("2006-01-02", endDate)
			if err := vacSvc.Create(&model.Vacation{UserID: userID, StartDate: sDate, EndDate: eDate}); err != nil {
				users, _ := userSvc.List()
				c.HTML(400, "new_vacation.html", gin.H{"Error": err.Error(), "Users": users})
				return
			}
			c.Redirect(302, "/ui/vacations")
		})

		authGroup.GET("/ui/vacations/:id/edit", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			vacation, _ := vacSvc.GetByID(id)
			users, _ := userSvc.List()
			c.HTML(200, "edit_vacation.html", gin.H{"Vacation": vacation, "Users": users})
		})
		authGroup.POST("/ui/vacations/:id/edit", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			vacation, _ := vacSvc.GetByID(id)
			if uid := c.PostForm("user_id"); uid != "" {
				vacation.UserID = parseUint(uid)
			}
			if sd := c.PostForm("start_date"); sd != "" {
				if sDate, err := time.Parse("2006-01-02", sd); err == nil {
					vacation.StartDate = sDate
				}
			}
			if ed := c.PostForm("end_date"); ed != "" {
				if eDate, err := time.Parse("2006-01-02", ed); err == nil {
					vacation.EndDate = eDate
				}
			}
			if err := vacSvc.Update(vacation); err != nil {
				users, _ := userSvc.List()
				c.HTML(400, "edit_vacation.html", gin.H{"Error": err.Error(), "Vacation": vacation, "Users": users})
				return
			}
			c.Redirect(302, "/ui/vacations")
		})
		authGroup.POST("/ui/vacations/:id/delete", func(c *gin.Context) {
			id := parseUint(c.Param("id"))
			vacSvc.Delete(id)
			c.Redirect(302, "/ui/vacations")
		})

		// API routes
		api.RegisterRoutes(r, userSvc, roleSvc, teamSvc, dirSvc, vacSvc)
		return r
	})

	// Start application via DI
	if err := container.Invoke(func(r *gin.Engine) error {
		fmt.Printf("Starting server on port %s\n", port)
		return r.Run(":" + port)
	}); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

package router

import (
	"SofwareGoDay3/controllers"
	"SofwareGoDay3/database"

	"github.com/gin-gonic/gin"
)

func addD(_ *gin.Context) {
	controllers.CreateDeveloper(18, 3, "Nicolas", "EPITECH", database.Db.Ctx, database.Db.Def)
}

func rD(c *gin.Context) {
	controllers.GetDevelopers(1, database.Db.Ctx, database.Db.Def)
}

func upD(c *gin.Context) {
	controllers.UpdateDeveloper(1, 19, 3, "Nicolas", "EPITECH", database.Db.Ctx, database.Db.Def)
}

func delD(c *gin.Context) {
	controllers.DeleteDeveloper(1, database.Db.Ctx, database.Db.Def)
}

func addC(c *gin.Context) {
	controllers.CreateContact(1, "nicola", "0688747870", "github", "linkedin", database.Db.Ctx, database.Db.Def)
}

func rC(c *gin.Context) {
	controllers.GetContact(1, database.Db.Ctx, database.Db.Def)
}

func upC(c *gin.Context) {
	controllers.UpdateContact(1, "nicolas.heude@epitech.eu", "0600000000", "github", "linkedin", database.Db.Ctx, database.Db.Def)
}

func delC(c *gin.Context) {
	controllers.DeleteContact(1, database.Db.Ctx, database.Db.Def)
}

func addComp(c *gin.Context) {
	controllers.CreateCompetence(1, 11, "Autre", database.Db.Ctx, database.Db.Def)
}

func upComp(c *gin.Context) {
	controllers.UpdateCompetence(1, 9, "Autre", database.Db.Ctx, database.Db.Def)
}

func delComp(c *gin.Context) {
	controllers.DeleteCompetence(1, database.Db.Ctx, database.Db.Def)
}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/ADD-DEV", addD)
	r.GET("/R-DEV", rD)
	r.GET("/UP-DEV", upD)
	r.GET("/DEL-DEV", delD)
	r.GET("/ADD-CON", addC)
	r.GET("/R-CON", rC)
	r.GET("/UP-CON", upC)
	r.GET("/DEL-CON", delC)
	r.GET("/ADD-COMP", addComp)
	r.GET("/UP-COMP", upComp)
	r.GET("/DELL-COMP", delComp)
}

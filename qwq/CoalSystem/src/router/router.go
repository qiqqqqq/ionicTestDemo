package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-contrib/cors"
	"CoalSystem/src/controller"
)

func main() {

	router := gin.Default()

	//same as
	config := cors.DefaultConfig()

	//config.AllowOrigins = [] string{"http://localhost:3000"}
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	//router.Use(cros.Default())

	//check one coal
	router.GET("/coal/:id", controller.Getonecoal)

	//post one coal
	router.POST("/coal", controller.Postcoal)

	//get all coal
	router.GET("/coals", controller.Getcoals)

	//delete one coal
	router.DELETE("/coal/:id", controller.Deletecoal)

	//put one coal
	router.PUT("/coal/:id", controller.Putonecoal)

	//一个煤的所有分析报告
	router.GET("/coal/:id/analysis/reports", controller.Coalanalysisreports)

	//修改一种煤的一种分析报告
	router.PUT("/coals/analysis/report", controller.Putonereport)

	//查看一种煤的一种分析报告
	router.GET("/coal/:id/analysis/report/:rid", controller.Getonereport)

	//删除一个煤的一种报告
	router.DELETE("/coal/:id/analysis/report/:rid", controller.Deleteonereport)

	//增加一种煤的一种分析报告
	router.POST("/coal/:cid/analysis/report", controller.Postonereport)

	//获得分析报告单的种类
	router.GET("/coals/analysis/category", controller.Allanalysisreport)

	//一个分析种类的所有属性
	router.GET("/coals/analysis/category/:id/attribute", controller.Allattribute)

	router.GET("/allreports", controller.AllReports)

	router.GET("/qiyingsheng", controller.Qiyingsheng)

	router.Run(":5000")

}

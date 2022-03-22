package main

import (
	controller "eksplorasi_gin/controller"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//Students
	router := gin.Default()
	router.POST("/student", controller.InsertStudent)
	router.GET("/student", controller.GetStudents)
	router.GET("/student/:student_id", controller.GetStudents)
	router.PUT("/student", controller.UpdateStudent)
	router.DELETE("/student/:student_id", controller.DeleteStudent)

	//Course Plan
	router.POST("/course_plan", controller.InsertCoursePlan)
	router.GET("/course_plan", controller.GetCoursePlan)
	router.GET("/course_plan/:course_plan_id", controller.GetCoursePlan)
	router.PUT("/course_plan", controller.UpdateCoursePlan)
	router.DELETE("/course_plan/:couse_plan_id", controller.DeleteCoursePlan)

	//Detail Course Plan
	router.POST("/detail_cp", controller.InsertDetailCoursePlan)
	router.DELETE("/detail_cp/:detail_cp_id", controller.DeleteDetailCoursePlan)

	router.Run(":9090")
	fmt.Println("Connected to port 9090")
}

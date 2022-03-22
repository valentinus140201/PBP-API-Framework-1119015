package controller

import (
	Model "eksplorasi_gin/model"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertDetailCoursePlan(c *gin.Context) {
	db := connect()
	defer db.Close()

	course_plan_id := c.PostForm("course_plan_id")
	student_id := c.PostForm("student_id")
	kelas := c.PostForm("kelas")

	_, errQuery := db.Exec("INSERT INTO detail_course_plan(course_plan_id, student_id, kelas) values (?,?,?)",
		course_plan_id,
		student_id,
		kelas,
	)

	fmt.Println(errQuery)

	var response Model.DetailCoursePlanResponse
	if errQuery == nil {
		response.Message = "Insert Detail Course Plan Success"
		response.Status = 200
	} else {
		response.Message = "Insert Detail Course Plan Failed"
		response.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func GetDetailCoursePlan(student_id int) []Model.DetailCoursePlan {
	db := connect()
	defer db.Close()

	id := strconv.Itoa(student_id)

	query := "SELECT a.id, a.kelas, b.nama, b.sks, b.dosen_wali FROM detail_course_plan a JOIN course_plan b ON a.course_plan_id = b.id"
	if id != "" {
		query += " where student_id = " + id + ";"
	}

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	var DetailCoursePlan Model.DetailCoursePlan
	var DetailCoursePlans []Model.DetailCoursePlan

	for rows.Next() {
		if err := rows.Scan(&DetailCoursePlan.ID, &DetailCoursePlan.Kelas, &DetailCoursePlan.CoursePlan.Nama, &DetailCoursePlan.CoursePlan.SKS, &DetailCoursePlan.CoursePlan.DosenWali); err != nil {
			log.Fatal(err.Error())
		} else {
			DetailCoursePlans = append(DetailCoursePlans, DetailCoursePlan)
		}
	}

	return DetailCoursePlans
}

func DeleteDetailCoursePlan(c *gin.Context) {
	db := connect()
	defer db.Close()

	id := c.Param("detail_cp_id")
	fmt.Println(id)
	query := "DELETE FROM detail_course_plan where id = " + id

	_, errQuery := db.Exec(query)

	if errQuery != nil {
		fmt.Println(errQuery)
	}

	var responses Model.DetailCoursePlanResponse
	if errQuery == nil {
		responses.Message = "Success Delete Detail Course Plan"
		responses.Status = 200
	} else {
		responses.Message = "Failed Delete Detail Course Plan"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

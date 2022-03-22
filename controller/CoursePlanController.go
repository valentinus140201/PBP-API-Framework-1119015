package controller

import (
	Model "eksplorasi_gin/model"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertCoursePlan(c *gin.Context) {
	db := connect()
	defer db.Close()

	nama := c.PostForm("nama")
	sks, _ := strconv.Atoi(c.PostForm("sks"))
	dosen_wali := c.PostForm("dosen_wali")

	_, errQuery := db.Exec("INSERT INTO course_plan(nama, sks, dosen_wali) values (?,?,?)",
		nama,
		sks,
		dosen_wali,
	)

	fmt.Println(errQuery)

	var response Model.CoursePlanResponse
	if errQuery == nil {
		response.Message = "Insert Student Success"
		response.Status = 200
	} else {
		response.Message = "Insert Student Failed"
		response.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func GetCoursePlan(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM course_plan"
	id := c.Param("course_plan_id")
	if id != "" {
		query += " where id = " + id + ";"
	}

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	var CoursePlan Model.CoursePlan
	var CoursePlans []Model.CoursePlan

	for rows.Next() {
		if err := rows.Scan(&CoursePlan.ID, &CoursePlan.Nama, &CoursePlan.SKS, &CoursePlan.DosenWali); err != nil {
			log.Fatal(err.Error())
		} else {
			CoursePlans = append(CoursePlans, CoursePlan)
		}
	}

	var responses Model.CoursePlanResponse
	if len(CoursePlans) > 0 {
		responses.Message = "Success Get Course Plan"
		responses.Status = 200
		responses.Data = CoursePlans
	} else {
		responses.Message = "Failed Get Course Plan"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func UpdateCoursePlan(c *gin.Context) {
	db := connect()
	defer db.Close()

	// query := "UPDATE students"

	id := c.PostForm("couse_plan_id")
	nama := c.PostForm("nama")
	sks := c.PostForm("sks")
	dosen_wali := c.PostForm("dosen_wali")

	// if nama != "" {
	// 	query += " SET nama = " + nama
	// 	if nim != "" && jurusan != "" && fakultas != "" {
	// 		query += ", SET nim = " + nim + ", SET jurusan = " + jurusan + ", SET fakultas = " + fakultas
	// 	}
	// } else {
	// 	if nim != "" {
	// 		query += " SET nim = " + nim
	// 		if jurusan != "" && fakultas != "" {
	// 			query += ", SET jurusan = " + jurusan + ", SET fakultas = " + fakultas
	// 		}
	// 	} else {
	// 		if jurusan != "" {
	// 			query += " SET jurusan = " + jurusan
	// 			if fakultas != "" {
	// 				query += ", SET fakultas = " + fakultas
	// 			}
	// 		} else {
	// 			query += " SET fakultas = " + fakultas
	// 		}
	// 	}
	// }

	_, errQuery := db.Exec("UPDATE course_plan SET nama=?, sks=?, dosen_wali=? WHERE id=?;",
		nama,
		sks,
		dosen_wali,
		id,
	)
	// query += " WHERE id = " + id + ";"
	// fmt.Println(query)
	if errQuery != nil {
		fmt.Println(errQuery)
	}

	var responses Model.CoursePlanResponse
	if errQuery == nil {
		responses.Message = "Success Update Student"
		responses.Status = 200
	} else {
		responses.Message = "Failed Update Student"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func DeleteCoursePlan(c *gin.Context) {
	db := connect()
	defer db.Close()

	id := c.Param("couse_plan_id")
	query := "DELETE FROM course_plan where id = " + id

	_, errQuery := db.Exec(query)

	if errQuery != nil {
		fmt.Println(errQuery)
	}

	var responses Model.StudentResponse
	if errQuery == nil {
		responses.Message = "Success Delete Course Plan"
		responses.Status = 200
	} else {
		responses.Message = "Failed Delete Course Plan"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func GetCoursePlanForSomething(id string) Model.CoursePlan {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM course_plan"
	if id != "" {
		query += " where id = " + id + ";"
	}

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	var CoursePlan Model.CoursePlan
	var CoursePlans []Model.CoursePlan

	for rows.Next() {
		if err := rows.Scan(&CoursePlan.ID, &CoursePlan.Nama, &CoursePlan.SKS, &CoursePlan.DosenWali); err != nil {
			log.Fatal(err.Error())
		} else {
			CoursePlans = append(CoursePlans, CoursePlan)
		}
	}
	return CoursePlans[0]
}

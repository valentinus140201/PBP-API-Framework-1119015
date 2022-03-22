package controller

import (
	Model "eksplorasi_gin/model"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertStudent(c *gin.Context) {
	db := connect()
	defer db.Close()

	nama := c.PostForm("nama")
	nim := c.PostForm("nim")
	jurusan := c.PostForm("jurusan")
	fakultas := c.PostForm("fakultas")

	_, errQuery := db.Exec("INSERT INTO students(nama, nim, jurusan, fakultas) values (?,?,?,?)",
		nama,
		nim,
		jurusan,
		fakultas,
	)

	fmt.Println(errQuery)

	var response Model.StudentResponse
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

func GetStudents(c *gin.Context) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM students"
	id := c.Param("student_id")
	if id != "" {
		query += " where id = " + id + ";"
	}

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	var Student Model.Student
	var Students []Model.Student

	for rows.Next() {
		if err := rows.Scan(&Student.ID, &Student.Nama, &Student.NIM, &Student.Jurusan, &Student.Fakultas); err != nil {
			log.Fatal(err.Error())
		} else {
			Students = append(Students, Student)
		}
	}

	for i := 0; i < len(Students); i++ {
		Students[i].DetailCoursePlan = GetDetailCoursePlan(Students[i].ID)
	}

	var responses Model.StudentResponse
	if len(Students) > 0 {
		responses.Message = "Success Get Students"
		responses.Status = 200
		responses.Data = Students
	} else {
		responses.Message = "Failed Get Students"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func UpdateStudent(c *gin.Context) {
	db := connect()
	defer db.Close()

	// query := "UPDATE students"

	id := c.PostForm("student_id")
	nama := c.PostForm("nama")
	nim := c.PostForm("nim")
	jurusan := c.PostForm("jurusan")
	fakultas := c.PostForm("fakultas")

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

	_, errQuery := db.Exec("UPDATE students SET nama=?, nim=?, jurusan=?, fakultas=? WHERE id=?;",
		nama,
		nim,
		jurusan,
		fakultas,
		id,
	)

	// query += " WHERE id = " + id + ";"
	// fmt.Println(query)
	if errQuery != nil {
		fmt.Println(errQuery)
	}

	var responses Model.StudentResponse
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

func DeleteStudent(c *gin.Context) {
	db := connect()
	defer db.Close()

	id := c.Param("student_id")
	query := "DELETE FROM students where id = " + id

	_, errQuery := db.Exec(query)

	if errQuery != nil {
		fmt.Println(errQuery)
	}

	var responses Model.StudentResponse
	if errQuery == nil {
		responses.Message = "Success Delete Student"
		responses.Status = 200
	} else {
		responses.Message = "Failed Delete Student"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

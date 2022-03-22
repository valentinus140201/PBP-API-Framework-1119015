package model

type Student struct {
	ID               int                `form:"id" json:"id"`
	Nama             string             `form:"name" json:"name"`
	NIM              string             `form:"nim" json:"nim"`
	Jurusan          string             `form:"jurusan" json:"jurusan"`
	Fakultas         string             `form:"fakultas" json:"fakultas"`
	DetailCoursePlan []DetailCoursePlan `form:"detailcourseplan" json:"detailcourseplan"`
}

type StudentResponse struct {
	Status  int       `form:"status" json:"status"`
	Message string    `form:"message" json:"message"`
	Data    []Student `form:"data" json:"data"`
}

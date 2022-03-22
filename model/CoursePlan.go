package model

type CoursePlan struct {
	ID        int    `form:"id" json:"id"`
	Nama      string `form:"nama" json:"nama"`
	SKS       int    `form:"sks" json:"sks"`
	DosenWali string `form:"dosenwali" json:"dosenwali"`
}

type CoursePlanResponse struct {
	Status  int          `form:"status" json:"status"`
	Message string       `form:"message" json:"message"`
	Data    []CoursePlan `form:"data" json:"data"`
}

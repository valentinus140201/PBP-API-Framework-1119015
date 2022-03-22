package model

type DetailCoursePlan struct {
	ID         int        `form:"id" json:"id"`
	CoursePlan CoursePlan `form:"courseplan" json:"courseplan"`
	Kelas      string     `form:"kelas" json:"kelas"`
}

type DetailCoursePlanResponse struct {
	Status  int                `form:"status" json:"status"`
	Message string             `form:"message" json:"message"`
	Data    []DetailCoursePlan `form:"data" json:"data"`
}

package dto

type Numbers struct {
	Num1 int `json:"num_1" binding:"required"`
	Num2 int `json:"num_2" binding:"required"`
}

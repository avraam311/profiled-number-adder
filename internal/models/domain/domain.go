package domain

type Sum struct {
	Result int `json:"result" binding:"required"`
}

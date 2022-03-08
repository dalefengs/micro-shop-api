package forms

type OrderForm struct {
	Id      int    `json:"id" json:"id" binding:""`
	Address string `form:"address" json:"address" binding:"required"`
	Mobile  string `form:"mobile" json:"mobile" binding:"required"`
	Name    string `form:"name" json:"name" binding:"required"`
	Post    string `form:"post" json:"post" binding:""`
}

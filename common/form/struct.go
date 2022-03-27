package form

type AddMenuForm struct {
	Pid    int    `form:"pid"`
	Type   int    `form:"type" binding:"required"`
	Name   string `form:"name" binding:"required"`
	Url    string `form:"menUrl" binding:"required"`
	Sort   int    `form:"sort" binding:"required"`
	Status int    `form:"status" binding:"required"`
	Icon   string `form:"icon" binding:"required"`
}
type EditMenuForm struct {
	Id     int    `form:"id" binding:"required"`
	Type   int    `form:"type" binding:"required"`
	Name   string `form:"name" binding:"required"`
	Url    string `form:"menUrl" binding:"required"`
	Sort   int    `form:"sort" binding:"required"`
	Status int    `form:"status" binding:"required"`
	Icon   string `form:"icon" binding:"required"`
}

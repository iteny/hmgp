package sql

type GpCategory struct {
	Id     int    `json:"id"`
	Url    string `json:"url"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Sort   int    `json:"sort"`
	Icon   string `json:"icon"`
}
type GpCategoryContent struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Pid     int    `json:"pid"`
}

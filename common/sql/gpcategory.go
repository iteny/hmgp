package sql

type GpCategory struct {
	Id     int    `json:"id"`
	Url    string `json:"url"`
	Name   string `json:"name"`
	Status int    `json:"status"`
	Sort   int    `json:"sort"`
	Icon   string `json:"icon"`
}

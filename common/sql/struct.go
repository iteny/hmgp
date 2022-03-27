package sql

type Menu struct {
	Id       int    `json:"id"`
	Url      string `json:"url"`
	Name     string `json:"name"`
	Pid      int    `json:"pid"`
	Status   int    `json:"status"`
	Sort     int    `json:"sort"`
	Icon     string `json:"icon"`
	Type     int    `json:"type"`
	Children []Menu `json:"children" gorm:"foreignKey:pid"`
}

func RecursiveMenu(arr []Menu, pid int) (ar []Menu) {
	array := make([]Menu, 0)
	for k, v := range arr {
		if pid == v.Pid {
			array = append(array, arr[k])
		}
	}
	for tk, tv := range array {
		rm := RecursiveMenu(arr, tv.Id)
		for sk := range rm {
			array[tk].Children = append(array[tk].Children, rm[sk])
		}
	}
	return array
}

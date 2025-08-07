package models

type FilterData struct {
	Page     int    `query:"page"`
	Size     int    `query:"size"`
	Search   string `query:"search"`
	Sort     string `query:"sort"`
	Status   string `query:"status"`
	Category string `query:"category"`
}

type Pagination struct {
	Page       int `json:"page"`
	Size       int `json:"size"`
	TotalCount int `json:"total_count"`
	TotalPages int `json:"total_pages"`
}

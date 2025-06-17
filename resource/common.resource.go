package resource

// PaginationQueryParam is a pagination query param
type PaginationQueryParam struct {
	Query string `form:"query" json:"query" binding:"max=1000"`
	Sort  string `form:"sort" json:"sort" binding:"max=20"`
	Order string `form:"order" json:"order" binding:"max=20"`
	Limit int    `form:"limit,default=10" json:"limit" binding:"numeric,max=1000"`
	Page  int    `form:"page,default=1" json:"page" binding:"numeric,max=1000"`
}

// Meta is a meta response
type Meta struct {
	Total       int `json:"total"`
	Limit       int `json:"limit"`
	Page        int `json:"page"`
	Offset      int `json:"offset"`
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
}

// File is a file response
type File struct {
	Path string `uri:"path" binding:"required"`
}

// NewMeta returns a meta response
func NewMeta(total, limit, page int) *Meta {
	totalPage := 1
	if limit > 0 {
		totalPage = total / limit
	}

	if total%limit > 0 {
		totalPage++
	}

	return &Meta{
		Total:       total,
		Limit:       limit,
		Page:        page,
		Offset:      (page - 1) * limit,
		CurrentPage: page,
		TotalPage:   totalPage,
	}
}

// ResponseAPIPertamina is a response api pertamina
type ResponseAPIPertamina struct {
	Success bool   `json:"success"`
	Data    string `json:"data"`
	Message string `json:"message"`
	Status  string `json:"status"`
	Code    int    `json:"code"`
}

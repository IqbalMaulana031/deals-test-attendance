package resource

// PaginationQueryParam is a pagination query param
type PaginationQueryParam struct {
	Query        string `form:"query" json:"query" binding:"max=1000"`
	Sort         string `form:"sort" json:"sort" binding:"max=20"`
	Order        string `form:"order" json:"order" binding:"max=20"`
	Limit        int    `form:"limit,default=10" json:"limit" binding:"numeric,max=1000"`
	Offset       int    `form:"offset,default=0" json:"offset" binding:"numeric,max=1000"`
	Slug         string `form:"slug" json:"slug" binding:"max=50"`
	Page         int    `form:"page,default=1" json:"page" binding:"numeric,max=1000"`
	StartDate    string `form:"start_date" json:"start_date" binding:"omitempty,datetime=2006-01-02"`
	EndDate      string `form:"end_date" json:"end_date" binding:"omitempty,datetime=2006-01-02"`
	Type         string `form:"type" json:"type" binding:"max=20"`
	CategoryID   string `form:"category_id" json:"category_id" binding:"omitempty,uuid4"`
	IsFrequently bool   `form:"is_frequently" json:"is_frequently"`
	Status       string `form:"status" json:"status"`
	Stock        string `form:"stock" json:"stock"`
	TypeMerchant string `form:"type_merchant" json:"type_merchant" binding:"max=20"`
	TypeFAQ      string `form:"type_faq" json:"type_faq" binding:"max=20"`
	PhoneNumber  string `form:"phone_number" json:"phone_number" binding:"max=20"`
	Name         string `form:"name" json:"name" binding:"max=50"`
	Payment      string `form:"payment" json:"payment" binding:"omitempty,max=20,oneof=cash my-pertamina"`
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

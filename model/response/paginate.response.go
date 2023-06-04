package response

type Paginate struct {
	Page      int64       `json:"page"`
	TotalPage int64       `json:"total_page"`
	Data      interface{} `json:"data"`
}

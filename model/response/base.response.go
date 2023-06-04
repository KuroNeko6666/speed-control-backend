package response

type Base struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

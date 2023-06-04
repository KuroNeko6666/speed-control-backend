package response

type Dashboard struct {
	Label []string    `json:"label"`
	Data  interface{} `json:"data"`
}

type DashboardCount struct {
	Count int64 `json:"count"`
}

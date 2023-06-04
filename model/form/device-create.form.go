package form

type DeviceCreate struct {
	UserID  string `json:"user_id" form:"user_id"`
	Name    string `json:"name" form:"name"`
	Model   string `json:"model" form:"model"`
	Address string `json:"address" form:"address"`
}

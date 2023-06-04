package form

type DeviceData struct {
	DeviceID string `json:"device_id" form:"device_id"`
	Speed    int    `json:"speed" form:"speed"`
	Distance int    `json:"distance" form:"distance"`
	DateTime string `json:"datetime" form:"datetime"`
}

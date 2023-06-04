package response

import "github.com/KuroNeko6666/speed-control-backend.git/database/model"

type Login struct {
	Token string     `json:"token"`
	User  model.User `json:"user"`
}

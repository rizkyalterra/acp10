package news

import (
	"acp10/models/base"
)

type NewsResponse struct {
	base.BaseResponse
	Data []News `json:"data"`
}

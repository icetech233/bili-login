package bili

var qrcode_content = "qrcode_content"

const err_qrcode_new = "qrcode new error"

type BiliApiBaseModel struct {
	Code int           `json:"code"`
	Data LoginUrlModel `json:"data"`
}

type LoginUrlModel struct {
	Url      string `json:"url"`
	OauthKey string `json:"oauthKey"`
}

package bili

var qrcode_content = "QRcode_content"

const (
	pic_size       = 256
	err_qrcode_new = "New QRcode Error"
)

// Http Base Model
type BiliApiBaseModel struct {
	Code int           `json:"code"`
	Data LoginUrlModel `json:"data"`
}

// QRcode get OauthKey
type LoginUrlModel struct {
	Url      string `json:"url"`
	OauthKey string `json:"oauthKey"`
}

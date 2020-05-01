package bili

import (
	"fmt"
	"image/color"

	"github.com/skip2/go-qrcode"
)

// hello world
func BiliHello() {
	fmt.Println("Bilibili Hello World !")
}

// qrcode demo
func demo() {

	qrcode_content = "http://www.bilibili.com"
	//var q *qrcode.QRCode
	q, err := qrcode.New(qrcode_content, qrcode.Low)
	if err != nil {
		fmt.Println(err_qrcode_new)
		//return nil, err
	}
	q.BackgroundColor = color.White
	q.ForegroundColor = color.RGBA{0, 161, 214, 255}
}

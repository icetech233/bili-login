package bili

import (
	"fmt"
)

// hello world
func BiliHello() {
	fmt.Println("Bilibili Hello World !")
}

func Demo1() {

	login_model := GetLoginUrl()
	q := GenerateQrCode(login_model)

	WriteFile(q, "login_qrcode.png")

}

// qrcode demo
func demo_qrcode() {

	// qrcode_content = "http://www.bilibili.com"
	// //var q *qrcode.QRCode
	// q, err := qrcode.New(qrcode_content, qrcode.Low)
	// if err != nil {
	// 	fmt.Println(err_qrcode_new)
	// 	//return nil, err
	// }
	// q.BackgroundColor = color.White
	// q.ForegroundColor = color.RGBA{0, 161, 214, 255}
}

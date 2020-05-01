package bili

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"os"

	"github.com/skip2/go-qrcode"
)

func GenerateQrCode(data LoginUrlModel) (q *qrcode.QRCode) {
	var err error
	//var q *qrcode.QRCode
	q, err = qrcode.New(data.Url, qrcode.Low)
	if err != nil {
		fmt.Println(err_qrcode_new)
		//return nil, err
	}
	q.BackgroundColor = color.White
	q.ForegroundColor = color.RGBA{0, 161, 214, 255}
	return q
}

func WriteFile(q *qrcode.QRCode, png_filename string) error {
	var err error
	png_bytes, err := q.PNG(pic_size)
	err = ioutil.WriteFile(png_filename, png_bytes, os.FileMode(0644))
	if err != nil {
		fmt.Println("write qrcode to file error")
	}
	return err
}

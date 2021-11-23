package makeqrcode

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"os"
)

//QR画像を生成する
func MakeQr()  {
	//バーコード用の文字列を生成
	qrCode,_ :=qr.Encode("hello world",qr.M,qr.Auto)

	qrCode,_ =barcode.Scale(qrCode,200,200)
	file,_ :=os.Create("testQr.png")
	defer file.Close()
	png.Encode(file,qrCode)
}

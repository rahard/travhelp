// generate QRcode with
// Traveler ID + Token
// Token is generated periodically
// Signed (optional)

package main
import (
   "fmt"
   qrcode "github.com/skip2/go-qrcode"
)


func main() {
   // this should be read from somewhere
   var travid = "budi"
   var token = "x123y456z"
   data := travid + " " + token

   err := qrcode.WriteFile(data, qrcode.Medium, 245, "qr.png")
   if (err != nil) {
      fmt.Println("gagal")
   }
}

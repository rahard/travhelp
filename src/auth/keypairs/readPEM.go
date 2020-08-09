// a test to open those PEM files
// @rahard
package main
import (
   "fmt"
   "os"
   "bufio"
   "encoding/pem"
   "crypto/x509"
   //"crypto/rsa"
)


func main() {
   // open private key
   privateKeyFile, err := os.Open("go_private.pem")
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   pemfileinfo, _ := privateKeyFile.Stat()
   var size int64 = pemfileinfo.Size()
   pembytes := make([]byte, size)
   buffer := bufio.NewReader(privateKeyFile)
   _, err = buffer.Read(pembytes)
   data, _ := pem.Decode([]byte(pembytes))
   privateKeyFile.Close()


   privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
   //privateKeyImported, err := rsa.ParsePKCS1PrivateKey(data.Bytes)
   if err != nil {
      fmt.Println(err)
       os.Exit(1)
   }
   fmt.Println("Private Key : ", privateKeyImported)
}

// a test to generate PEM files
// @rahard
package main
import (
   "fmt"
   "os"
   //"bufio"
   "encoding/pem"
   "crypto/x509"
   "crypto/rsa"
   "crypto/rand"
)


func main() {
   privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
   if err != nil {
      fmt.Println(err.Error)
   os.Exit(1)
   }

   publicKey := &privateKey.PublicKey
   fmt.Println("Private Key: ", privateKey)
   fmt.Println("Public key: ", publicKey)

   // save into files
   pemPrivateFile, err := os.Create("go_private.pem")
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   var pemPrivateBlock = &pem.Block{
      Type:  "RSA PRIVATE KEY",
      Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
   }
   err = pem.Encode(pemPrivateFile, pemPrivateBlock)
   if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }
   pemPrivateFile.Close()
}

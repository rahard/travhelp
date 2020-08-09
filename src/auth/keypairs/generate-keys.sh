#! /bin/sh

# from https://byparker.com/blog/2019/generating-a-golang-compatible-ssh-rsa-key-pair/

# To generate the private key, run the following:
openssl genrsa -des3 -out private.pem 4096

# To generate the *public* key from your private key, run the following:
openssl rsa -in private.pem -outform PEM -pubout -out public.pem

# To generate the ssh-rsa public key format, run the following:
ssh-keygen -f public.pem -i -mPKCS8 > id_rsa.pub

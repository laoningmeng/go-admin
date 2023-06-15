#!/bin/bash

#openssl.cnf 位置
opensslCnfPath=/private/etc/ssl/openssl.cnf
serverName="go-admin"
outPath=configs/cert

# 生成.key  私钥文件
openssl genrsa -out $outPath/ca.key 2048

# 生成.csr 证书签名请求文件
openssl req -new -key $outPath/ca.key -out $outPath/ca.csr  -subj "/C=CN/L=BJ/O=demo/CN=demo.com"

# 自签名生成.crt 证书文件
openssl req -new -x509 -days 3650 -key $outPath/ca.key -out $outPath/ca.crt  -subj "/C=CN/L=BJ/O=demo/CN=demo.com"


# 生成.key  私钥文件
openssl genrsa -out $outPath/server.key 2048

# 生成.csr 证书签名请求文件
openssl req -new -key $outPath/server.key -out $outPath/server.csr \
	-subj "/C=CN/L=BJ/O=demo/CN=CN=demo.com" \
	-reqexts SAN \
	-config <(cat $opensslCnfPath <(printf "\n[SAN]\nsubjectAltName=DNS:$serverName"))

# 签名生成.crt 证书文件
openssl x509 -req -days 3650 \
   -in $outPath/server.csr -out $outPath/server.crt \
   -CA $outPath/ca.crt -CAkey $outPath/ca.key -CAcreateserial \
   -extensions SAN \
   -extfile <(cat $opensslCnfPath <(printf "\n[SAN]\nsubjectAltName=DNS:$serverName"))


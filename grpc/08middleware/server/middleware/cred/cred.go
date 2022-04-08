package cred

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

/*
@author RandySun
@create 2022-03-29-10:49
*/
// TLSInterceptor TLS证书认证
func TLSInterceptor() grpc.ServerOption {
	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	creds, err := credentials.NewServerTLSFromFile("../pkg/tls/server_cert.pem", "../pkg/tls/server_key.pem")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	return grpc.Creds(creds)
}

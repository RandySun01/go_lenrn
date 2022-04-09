package demo

import (
	"fmt"

	"google.golang.org/grpc/metadata"
)

/*
@author RandySun
@create 2022-04-09-22:33
*/
func main() {
	// 连接服务器
	md := metadata.New(map[string]string{"go": "programming", "tour": "book"})
	fmt.Printf("%#v, \n%T \n ", md, md)
	mdPairs := metadata.Pairs(
		"go", "programming",
		"tour", "book",
		"go", "eddycjy",
	)
	fmt.Printf("%#v, \n%T \n ", mdPairs, mdPairs)

}

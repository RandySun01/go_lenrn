package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "gpc_test/tutorial"
	"io/ioutil"
	"log"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "hello word ！！",
		Email: "randy@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}
	fmt.Println(p.Name, p.GetName())
	//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./exp.proto
	//fmt.Println(tutorialpb.Name)

	//book := &pb.AddressBook{
	//	People:[]*pb.Person{
	//		{
	//			Id:    1234,
	//			Name:  "hello word ！！",
	//			Email: "randy@example.com",
	//			Phones: []*pb.Person_PhoneNumber{
	//				{Number: "555-4321", Type: pb.Person_HOME},
	//			},
	//		},
	//	},
	//}
	//// ...
	//
	//// Write the new address book back to disk.
	//out, err := proto.Marshal(book)
	//if err != nil {
	//	log.Fatalln("Failed to encode address book:", err)
	//}
	//if err := ioutil.WriteFile("grpc_file", out, 0644); err != nil {
	//	log.Fatalln("Failed to write address book:", err)
	//}


	in, err := ioutil.ReadFile("grpc_file")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Println(book.People[0])
}
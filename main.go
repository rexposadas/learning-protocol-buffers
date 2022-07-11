package main

import (
	"fmt"
	pb "github.com/rexposadas/learning-protocol-buffers/proto"
)

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   32,
			Name: "james",
		},
		MultipleDummies: []*pb.Dummy{
			{Id: 11, Name: "Andrew"},
		},
	}

}

func main() {
	fmt.Printf("%+v", doComplex())
}

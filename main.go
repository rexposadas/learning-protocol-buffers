package main

import (
	"fmt"
	pb "github.com/rexposadas/learning-protocol-buffers/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id: 33,
	}
}

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

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IDWrapper{
			"myid":  {Id: 42},
			"myid2": {Id: 43},
			"myid3": {Id: 44},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin" // the file we will write to and read from

	writeToFile(path, p)

	message := &pb.Simple{}
	readFromFile(path, message)

	fmt.Println("file: ", message)

}

func main() {
	//fmt.Println(doMap())
	doFile(doSimple())

}

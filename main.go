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

func main() {
	//fmt.Printf("%+v", doComplex())
	fmt.Println(doMap())
}

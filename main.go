package main

import (
	"fmt"
	pb "github.com/rexposadas/learning-protocol-buffers/proto"
	"google.golang.org/protobuf/proto"
	"reflect"
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

// input: json string
// output: protomessage
func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

// take a message and converts it to a JSON string
func doToJSON(p proto.Message) string {
	jsonString := toJSON(p)

	fmt.Println(jsonString)
	return jsonString
}

func main() {
	//fmt.Println(doMap())
	//doFile(doSimple())

	jsonString := doToJSON(doComplex())
	message := doFromJSON(jsonString, reflect.TypeOf(pb.Complex{}))

	fmt.Println(jsonString)
	fmt.Println(message)

}

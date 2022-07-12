package json

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
	"reflect"
)

// DoFromJSON .
// input: json string
// output: proto message
func DoFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

// take a message and converts it to a JSON string
func DoToJSON(p proto.Message) string {
	jsonString := toJSON(p)

	fmt.Println(jsonString)
	return jsonString
}

func toJSON(pb proto.Message) string {

	option := protojson.MarshalOptions{
		Multiline: true,
	}

	out, err := option.Marshal(pb)
	if err != nil {
		log.Println("cannot convert to json", err)
		return ""
	}

	return string(out)
}

func fromJSON(in string, pb proto.Message) {

	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	err := option.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatalln("cant unmarshall", err)
		return
	}
}

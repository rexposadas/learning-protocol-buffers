package main

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"log"
)

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

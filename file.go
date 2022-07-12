package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func writeToFile(name string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln(err)
	}

	if err = ioutil.WriteFile(name, out, 0644); err != nil {
		log.Fatalln("cannot write to file", err)
	}

	fmt.Println("data has been written")
}

func readFromFile(name string, pb proto.Message) {
	in, err := ioutil.ReadFile(name)

	if err != nil {
		log.Fatalln("failed to read file", in)
		return
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("failed to unmarshal ", err)
	}
}

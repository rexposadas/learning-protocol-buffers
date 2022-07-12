package addressbook

import (
	"fmt"
	pb "github.com/rexposadas/learning-protocol-buffers/proto"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

// Write the address book to file.
func Write(filename string, original, toAdd proto.Message) {

	if toAdd == nil {
		writeToFile(filename, original)
	}

	// Add the address book and write to file

}

func FirstPerson() *pb.AddressBook {
	return &pb.AddressBook{
		People: []*pb.Person{
			{Id: 22, Name: "first person"},
		},
	}

}

// Read address book from file
func Read(filename string) {
	message := &pb.AddressBook{}
	readFromFile(filename, message)

	fmt.Println("file: ", message)
}

// writes the address book, a proto message, to file.
func writeToFile(filename string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln(err)
	}

	if err = ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("cannot write to file", err)
	}

	fmt.Println("data has been written")

}

// Read from file and put in a proto message.
func readFromFile(filename string, pb proto.Message) {
	in, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln("failed to read file", in)
		return
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("failed to unmarshal ", err)
	}
}

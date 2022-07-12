package addressbook

import (
	pb "github.com/rexposadas/learning-protocol-buffers/proto"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

// Add a person to the address book.
func Add(filename string, person *pb.Person) {
	book := Read(filename)
	book.People = append(book.People, person)

	Write(filename, book)
}

// Write the address book to file.
func Write(filename string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln(err)
	}

	if err = ioutil.WriteFile(filename, out, 0644); err != nil {
		log.Fatalln("cannot write to file", err)
	}
}

func FirstPerson() *pb.AddressBook {
	return &pb.AddressBook{
		People: []*pb.Person{
			{Id: 22, Name: "first person"},
		},
	}
}

// Read address book from file
func Read(filename string) *pb.AddressBook {
	pbBook := &pb.AddressBook{}
	in, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln("failed to read file ", filename)
	}

	if err := proto.Unmarshal(in, pbBook); err != nil {
		log.Fatalln("failed to unmarshal ", err)
	}

	return pbBook
}

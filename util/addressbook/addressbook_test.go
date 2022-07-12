package addressbook

import (
	"github.com/rexposadas/learning-protocol-buffers/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Can we write and read an address book to and from a file?
func TestWriteRead(t *testing.T) {

	p := &proto.AddressBook{
		People: []*proto.Person{
			{Id: 1, Name: "first"},
			{Id: 2, Name: "second"},
		},
	}

	filename := "addressbook.bin"
	Write(filename, p)

	book := Read(filename)

	assert.Equal(t, len(book.People), 2)

	person := book.People[0]

	assert.Equal(t, int32(1), person.Id)
	assert.Equal(t, person.Name, "first")
}

func TestAdd(t *testing.T) {

	p := &proto.AddressBook{
		People: []*proto.Person{
			{Id: 1, Name: "first"},
		},
	}

	filename := "addressbook.bin"
	Write(filename, p)

	secondPerson := &proto.Person{
		Id:   2,
		Name: "second",
	}

	Add(filename, secondPerson)

	// Let's see if we added the second person
	book := Read(filename)

	assert.Equal(t, 2, len(book.People))

	second := book.People[1]

	assert.Equal(t, secondPerson.Id, second.Id)
	assert.Equal(t, secondPerson.Name, second.Name)

}

package main

import (
	"fmt"
	"path/filepath"

	"github.com/rexposadas/learning-protocol-buffers/util/addressbook"
)

func main() {
	// Step 1: Create initial address book
	addressBookName := fmt.Sprintf("generated%daddressbook.bin", filepath.Separator)
	addressbook.Write(addressBookName, addressbook.FirstPerson(), nil)
	addressbook.Read(addressBookName)

	// Step 3: add a person to address book

	// Step 4: read from address book and verify that we added correctly.
}

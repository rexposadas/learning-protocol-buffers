package main

import (
	"github.com/rexposadas/learning-protocol-buffers/util/addressbook"
	"path/filepath"
)

func main() {
	// Step 1: Create initial address book
	addressBookName := filepath.FromSlash("generated/addressbook.bin")
	addressbook.Write(addressBookName, addressbook.FirstPerson())
	// Step 3: add a person to address book

	// Step 4: read from address book and verify that we added correctly.
}

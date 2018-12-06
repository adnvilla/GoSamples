package main

import (
	fmt "fmt"
	"io/ioutil"
	"log"

	proto "github.com/golang/protobuf/proto"
)

func main() {
	fname := "test.dat"
	book := &AddressBook{
		People: []*Person{
			&Person{Name: "Tests"},
			&Person{Name: "Tests2"},
		},
	}

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book2 := &AddressBook{}
	if err := proto.Unmarshal(in, book2); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	fmt.Println(book2)
}

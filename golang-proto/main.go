package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"./src/simple"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	writeToFile(sm)
	readFromFile(sm)
}

func writeToFile(sm *personpb.Person) error {
	out, err := proto.Marshal(sm)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if err = ioutil.WriteFile("fileFromProto.bin", out, 0644); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func readFromFile(sm *personpb.Person) error {
	in, err := ioutil.ReadFile("fileFromProto.bin")
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = proto.Unmarshal(in, sm)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(sm)
	return nil
}

func doSimple() *personpb.Person {
	sm := personpb.Person{
		Age:       12,
		FirstName: "Deepak",
		LastName:  "Ahuja",
		IsVerfied: true,
	}
	return &sm
}

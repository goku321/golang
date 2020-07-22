package main

import (
	"bytes"
	"crypto/md5"
	"encoding/gob"

	"log"
)

// func createHash(args ...interface{}) {
// 	hasher := md5.New()
// 	for _, v := range args {
// 		s := fmt.Sprint(v)

// 	}
// }

func getBytes(x interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(x)
	return buf.Bytes(), err
}

type person struct {
	Name string
	Age  int
}

func main() {
	p := person{"yoda", 109}
	b, err := getBytes(p)
	if err != nil {
		log.Fatal(err)
	}

	hasher := md5.New()
	v, err := hasher.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(v)
}

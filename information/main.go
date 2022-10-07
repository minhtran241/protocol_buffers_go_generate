package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"information/protocol"

	"google.golang.org/protobuf/proto"
)

func main() {

	option := flag.String("admin", "server", "Communication between server and client")

	flag.Parse()

	switch *option {
	case "client":
		runClient()
	case "server":
		runServer()
	}
}

func runClient() {
	person := protocol.Person{
		Id:   1,
		Name: "Minh Tran",
		Age:  19,
	}

	data, err := proto.Marshal(&person)

	if err != nil {
		log.Fatal("marshal error:", err)
	}

	sendData(data)
}

func runServer() {
	listener, err := net.Listen("tcp", "127.0.0.1:8085")

	if err != nil {
		log.Fatal("listener error:", err)
	}

	for {
		connection, err := listener.Accept()

		if err != nil {
			log.Fatal("listener error:", err)
		}

		go func(c net.Conn) {
			defer c.Close()

			data, err := ioutil.ReadAll(connection)

			if err != nil {
				log.Fatal(err.Error())
			}

			person := &protocol.Person{}

			err = proto.Unmarshal(data, person)

			if err != nil {
				log.Fatal("unmarshal error:", err)
			}

			fmt.Println(person)

		}(connection)
	}

}

func sendData(data []byte) {
	connection, err := net.Dial("tcp", "127.0.0.1:8085")

	if err != nil {
		log.Fatal("connection error:", err)
	}

	defer connection.Close()

	connection.Write(data)
}

package main

import (
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"log"
	"runtime"
	"chatprotoactor/messages"
	"github.com/satori/go.uuid"
	"strings"
)

type client struct {
	Name string
}

func (legatus *client) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.Connected:
		log.Println("Connected", msg.Message)

	case *messages.UserList:
		for _, user := range msg.UserName {
			log.Println(user)
		}
	case *messages.Message:
		log.Println("New Message", msg.Message)
		//case *messages.NewMessage:
		//	legatus.notifyAll(msg)
		//case *actor.Started:
		//	log.Printf("Actor started %s", msg)
	case *actor.Terminated:
		log.Printf("Actor terminated %s", msg)
	default:
		log.Printf("Default %s", msg)

	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	server := actor.NewPID("127.0.0.1:8080", "server")
	remote.Start("127.0.0.1:0")

	id := uuid.NewV4().String()

	props := actor.FromInstance(&client{Name: id})

	client, _ := actor.SpawnNamed(props, "client")

	server.Request(&messages.Connect{Id: id}, client)
	var input string = "default"
	for input != "quit" {
		switch input, _ = console.ReadLine(); input {
		case "list":
			log.Println("List connected")
			server.Request(&messages.ListConnected{}, client)
		case "quit":
		default:
			if strings.Contains(input, "<="){
				log.Println("with destination")
				split := strings.SplitN(input, "<=", 2)
				server.Tell(&messages.Message{UserName:split[0], Message:split[1]})
			}else {
				log.Println("without destination")
				server.Tell(&messages.Message{Message:input})
			}
		}
	}
}

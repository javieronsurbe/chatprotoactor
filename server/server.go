package main

import (
	"github.com/AsynkronIT/goconsole"
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/AsynkronIT/protoactor-go/remote"
	"log"
	"runtime"
	"chatprotoactor/messages"
)

type server struct {
	connected map[string]*actor.PID
}

func (s *server) notifyAll(message interface{}) {
	for _, client := range s.connected {
		client.Tell(message)
	}
}

func (s *server) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.Connect:
		context.Sender().Tell(&messages.Connected{Message: "Wellcome " + msg.Id})
		log.Println(msg.Id+" Connected", context.Sender(), context.Sender())
		s.connected[msg.Id] = context.Sender()
	case *messages.Message:
		if len(msg.UserName) == 0 {
			s.notifyAll(msg)
		} else {
			pid := s.connected[msg.UserName]
			pid.Tell(msg)
		}
	case *messages.ListConnected:
		var users []string
		for user := range s.connected {
			users = append(users, user)
		}
		context.Sender().Tell(&messages.UserList{UserName: users})

	case *actor.Terminated:
		log.Printf("Actor terminated %s", msg)
		for key, value:= range s.connected{
			if msg.Who==value{
				log.Println("Removing dead actor", key)
				delete(s.connected, key)
			}

		}

	default:
		log.Printf("Default %s", msg)

	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	remote.Start("127.0.0.1:8080")

	props := actor.FromInstance(&server{connected: map[string]*actor.PID{}})

	actor.SpawnNamed(props, "server")
	var input string
	for input != "quit" {
		input, _ = console.ReadLine()
	}

}

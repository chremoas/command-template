package main

import (
	"fmt"
	"github.com/chremoas/services-common/config"
	"github.com/micro/go-micro"

	proto "github.com/chremoas/chremoas/proto"
	// This needs to be changed to point to the correct command location
	"github.com/chremoas/command-template/command"
)

var Version = "1.0.0"
var service micro.Service
var name = "template"

func main ( ) {
	service = config.NewService(Version, "cmd", name, initialize)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

// This function is a callback from the config.NewService function.  Read those docs
func initialize(config *config.Configuration) error {
	proto.RegisterCommandHandler(service.Server(),
		command.NewCommand(name),
	)

	return nil
}

package main

import (
	"fmt"
	"git.maurer-it.net/abaeve/services-common/config"
	"github.com/micro/go-micro"

	proto "git.maurer-it.net/abaeve/chremoas/proto"
	"git.maurer-it.net/abaeve/chremoas-command-template/command"
)

var Version string = "1.0.0"
var service micro.Service

func main ( ) {
	service = config.NewService(Version, "auth", initialize)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

// This is commented because you may or may not need/want the https://git.maurer-it.net/abaeve/services-common
//configuration related stuff.  This function is a callback from the config.NewService function.  Read those docs
func initialize(config *config.Configuration) error {
	proto.RegisterCommandHandler(service.Server(),
		command.NewCommand(config.Name),
	)

	return nil
}
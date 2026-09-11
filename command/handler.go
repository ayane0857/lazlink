package command

import "github.com/disgoorg/disgo/events"

var Handlers = map[string]func(*events.ApplicationCommandInteractionCreate){}

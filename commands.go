package main

import (
	"github.com/mitchellh/cli"
	"github.com/jakub.igla/props/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"add": func() (cli.Command, error) {
			return &command.AddCommand{
				Meta: *meta,
			}, nil
		},
		"env": func() (cli.Command, error) {
			return &command.EnvCommand{
				Meta: *meta,
			}, nil
		},
		
		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name: Name,
			}, nil
		},
	}
}

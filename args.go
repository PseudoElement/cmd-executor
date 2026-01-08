package main

type Argument struct {
	Name     string
	Required bool
}

var argsConfig = map[string][]Argument{
	"npm_i":          {Argument{Name: "package name", Required: false}},
	"npm_b":          {},
	"yarn_add":       {Argument{Name: "package name", Required: false}},
	"git_pull":       {},
	"git_push":       {},
	"git_commit":     {Argument{Name: "commit message", Required: true}},
	"git_stash_push": {Argument{Name: "stash message", Required: false}},
}

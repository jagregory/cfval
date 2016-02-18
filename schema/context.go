package schema

import "github.com/jagregory/cfval/parse"

type Context struct {
	Path        []string
	Template    *parse.Template
	Definitions ResourceDefinitions
}

func (c Context) Push(path ...string) Context {
	return Context{
		Path:        append(c.Path, path...),
		Template:    c.Template,
		Definitions: c.Definitions,
	}
}

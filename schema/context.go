package schema

import (
	"github.com/jagregory/cfval/constraints"
	"github.com/jagregory/cfval/parse"
)

type Context interface {
	Definitions() ResourceDefinitions
	Path() []string
	Template() *parse.Template
}

func ContextAdd(ctx Context, path ...string) Context {
	return fullContext{
		definitions: ctx.Definitions(),
		path:        append(ctx.Path(), path...),
		template:    ctx.Template(),
	}
}

type ResourceContext interface {
	CurrentResource() constraints.CurrentResource
	Definitions() ResourceDefinitions
	Path() []string
	Template() *parse.Template
}

func ResourceContextAdd(ctx ResourceContext, path ...string) ResourceContext {
	return fullContext{
		currentResource: ctx.CurrentResource(),
		definitions:     ctx.Definitions(),
		path:            append(ctx.Path(), path...),
		template:        ctx.Template(),
	}
}

type PropertyContext interface {
	CurrentResource() constraints.CurrentResource
	Definitions() ResourceDefinitions
	Path() []string
	Property() Schema
	Template() *parse.Template
}

func PropertyContextAdd(ctx PropertyContext, path ...string) PropertyContext {
	return fullContext{
		currentResource: ctx.CurrentResource(),
		definitions:     ctx.Definitions(),
		path:            append(ctx.Path(), path...),
		property:        ctx.Property(),
		template:        ctx.Template(),
	}
}

type fullContext struct {
	currentResource constraints.CurrentResource
	definitions     ResourceDefinitions
	path            []string
	property        Schema
	template        *parse.Template
}

func (ctx fullContext) Path() []string {
	return ctx.path
}

func (ctx fullContext) Template() *parse.Template {
	return ctx.template
}

func (ctx fullContext) Definitions() ResourceDefinitions {
	return ctx.definitions
}

func (ctx fullContext) CurrentResource() constraints.CurrentResource {
	return ctx.currentResource
}

func (ctx fullContext) Property() Schema {
	return ctx.property
}

func NewInitialContext(template *parse.Template, definitions ResourceDefinitions) Context {
	return fullContext{
		path:        make([]string, 0, 25),
		template:    template,
		definitions: definitions,
	}
}

func NewResourceContext(ctx Context, currentResource constraints.CurrentResource) ResourceContext {
	return fullContext{
		currentResource: currentResource,
		definitions:     ctx.Definitions(),
		path:            ctx.Path(),
		template:        ctx.Template(),
	}
}

func NewPropertyContext(ctx ResourceContext, property Schema) PropertyContext {
	return fullContext{
		currentResource: ctx.CurrentResource(),
		definitions:     ctx.Definitions(),
		path:            ctx.Path(),
		property:        property,
		template:        ctx.Template(),
	}
}

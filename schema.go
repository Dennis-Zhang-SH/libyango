package libyanggo

/*
#include <libyang/libyang.h>
*/
import "C"

type SchemaModule struct {
	ctx *Context
	raw *C.struct_lys_module
}

type SchemaModules struct {
	ctx   *Context
	index uint32
}

func SchemaModuleFromRaw(ctx *Context, raw *C.struct_lys_module) *SchemaModule {
	return &SchemaModule{
		ctx,
		raw,
	}
}

func NewSchemaModules(ctx *Context, skipInternal bool) *SchemaModules {
	var index uint32 = 0
	if skipInternal {
		index = ctx.InternalModuleCount()
	}
	return &SchemaModules{
		ctx,
		index,
	}
}

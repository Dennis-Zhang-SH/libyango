package libyanggo

/*
#include <libyang/libyang.h>
*/
import "C"

type SchemaModule struct {
	ctx *Context
	raw *C.struct_lys_module
}

func SchemaModuleFromRaw(ctx *Context, raw *C.struct_lys_module) *SchemaModule {
	return &SchemaModule{
		ctx,
		raw,
	}
}

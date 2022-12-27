package libyango

/*
#include <libyang/libyang.h>
*/
import "C"

type SchemaNodeKind = uint

const (
	Container SchemaNodeKind = 1 << iota
	Choice
	Leaf
	LeafList
	List
	_
	_
	Case
	Rpc
	Action
	Notification
	_
	Input
	Output
	AnyData SchemaNodeKind = 96
)

type SchemaModule struct {
	ctx *Context
	raw *C.struct_lys_module
}

type SchemaModules struct {
	ctx   *Context
	index uint32
}

type SchemaNode struct {
	ctx  *Context
	raw  *C.struct_lysc_node
	kind SchemaNodeKind
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

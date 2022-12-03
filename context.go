package libyanggo

/*
#cgo LDFLAGS: -lyang
#include <libyang/libyang.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const (
	LY_CTX_ALL_IMPLEMENTED uint16 = 1 << iota
	LY_CTX_REF_IMPLEMENTED
	LY_CTX_NO_YANGLIBRARY
	LY_CTX_DISABLE_SEARCHDIRS
	LY_CTX_DISABLE_SEARCHDIR_CWD
)

type Context struct {
	raw *C.struct_ly_ctx
}

type EmbeddedMoudleKey struct {
	mod_name string
	mod_rev *string
	submod_name *string
	submod_rev *string
}

type EmbeddedModules = map[EmbeddedMoudleKey]string

func CreateContext(options uint16) (*Context, error) {
	var ctx *C.struct_ly_ctx
	ctxp := &ctx
	ret := C.ly_ctx_new(nil, C.uint16_t(options), ctxp)
	if ret != C.LY_SUCCESS {
		return nil, fmt.Errorf("create context failed, error code: %d", ret)
	}
	return &Context{
		raw: ctx,
	}, nil
}

func (ctx *Context) SetSearchDir(path string) error {
	path_cstr := C.CString(path)
	defer C.free(unsafe.Pointer(path_cstr))
	if ret := C.ly_ctx_set_searchdir(ctx.raw, path_cstr); ret != C.LY_SUCCESS {
		return fmt.Errorf("create context failed, error code: %d", ret)
	}
	return nil
}

func (ctx *Context) UnsetSearchDirs() error {
	if ret := C.ly_ctx_unset_searchdir(ctx.raw, nil); ret != C.LY_SUCCESS {
		return fmt.Errorf("create context failed, error code: %d", ret)
	}
	return nil
}

func (ctx *Context) UnsetSearchDirLast(count uint32) error {
	if ret := C.ly_ctx_unset_searchdir_last(ctx.raw, C.uint(count)); ret != C.LY_SUCCESS {
		return fmt.Errorf("create context failed, error code: %d", ret)
	}
	return nil
}

func (ctx *Context) SetEmbeddedModules(modules EmbeddedModules) {
	
}
package libyanggo

/*
#cgo LDFLAGS: -lyang
#include <libyang/libyang.h>
LY_ERR _cgo_ly_module_import_cb(const char *mod_name, const char *mod_rev, const char *submod_name, const char *submod_rev,
        void *user_data, LYS_INFORMAT *format, const char **module_data, ly_module_imp_data_free_clb *free_module_data);

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
	mod_name    string
	mod_rev     string
	submod_name string
	submod_rev  string
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
	callback_pointer := (C.ly_module_imp_clb)(unsafe.Pointer(C._cgo_ly_module_import_cb))
	mp := unsafe.Pointer(&modules)
	C.ly_ctx_set_module_imp_clb(ctx.raw, callback_pointer, mp)
}

func (ctx *Context) UnsetEmbededModules() {
	C.ly_ctx_set_module_imp_clb(ctx.raw, nil, nil)
}

func (ctx *Context) GetOptions() uint16 {
	return uint16(C.ly_ctx_get_options(ctx.raw))
}

func (ctx *Context) SetOptions(options uint16) error {
	if ret := C.ly_ctx_set_options(ctx.raw, C.uint16_t(options)); ret != C.LY_SUCCESS {
		return fmt.Errorf("set options error, error code: %d", ret)
	}
	return nil
}

func (ctx *Context) UnsetOptions(options uint16) error {
	if ret := C.ly_ctx_unset_options(ctx.raw, C.uint16_t(options)); ret != C.LY_SUCCESS {
		return fmt.Errorf("set options error, error code: %d", ret)
	}
	return nil

}

func (ctx *Context) GetModuleSetID() uint16 {
	return uint16(C.ly_ctx_get_change_count(ctx.raw))
}

func (ctx *Context) GetModule(name, revision string) *SchemaModule {
	n := C.CString(name)
	r := C.CString(revision)
	defer C.free(unsafe.Pointer(n))
	defer C.free(unsafe.Pointer(r))
	m := C.ly_ctx_get_module(ctx.raw, n, r)
	if m == nil {
		return nil
	}
	return SchemaModuleFromRaw(ctx, m)
}

func (ctx *Context) GetModuleLatest(name string) *SchemaModule {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	m := C.ly_ctx_get_module_latest(ctx.raw, n)
	if m == nil {
		return nil
	}
	return SchemaModuleFromRaw(ctx, m)
}

func (ctx *Context) GetModuleImplemented(name string) *SchemaModule {
	n := C.CString(name)
	defer C.free(unsafe.Pointer(n))
	m := C.ly_ctx_get_module_implemented(ctx.raw, n)
	if m == nil {
		return nil
	}
	return SchemaModuleFromRaw(ctx, m)
}

func (ctx *Context) GetModuleNs(ns, revision string) *SchemaModule {
	n := C.CString(ns)
	r := C.CString(revision)
	defer C.free(unsafe.Pointer(n))
	defer C.free(unsafe.Pointer(r))
	m := C.ly_ctx_get_module_ns(ctx.raw, n, r)
	if m == nil {
		return nil
	}
	return SchemaModuleFromRaw(ctx, m)
}

func (ctx *Context) GetModuleLatestNs(ns string) *SchemaModule {
	n := C.CString(ns)
	defer C.free(unsafe.Pointer(n))
	m := C.ly_ctx_get_module_latest_ns(ctx.raw, n)
	if m == nil {
		return nil
	}
	return SchemaModuleFromRaw(ctx, m)
}

func (ctx *Context) GetModuleImplementedNs(ns string) *SchemaModule {
	n := C.CString(ns)
	defer C.free(unsafe.Pointer(n))
	m := C.ly_ctx_get_module_implemented_ns(ctx.raw, n)
	if m == nil {
		return nil
	}
	return SchemaModuleFromRaw(ctx, m)
}

func (ctx *Context) Modules(skipInternal bool) *SchemaModules {
	return NewSchemaModules(ctx, skipInternal)
}

// TODO traverse method, need find a way to implement iterator

func (ctx *Context) ResetLatests() {
	C.ly_ctx_reset_latests(ctx.raw)
}

func (ctx *Context) InternalModuleCount() uint32 {
	return uint32(C.ly_ctx_internal_modules_count(ctx.raw))
}

func (ctx *Context) LoadModule(name string, revision string, features []string) *SchemaModule {
	n := C.CString(name)
	r := C.CString(revision)
	defer C.free(unsafe.Pointer(n))
	defer C.free(unsafe.Pointer(r))
	features_ptr := make([]*C.char, 0)
	for k := range features {
		vp := C.CString(features[k])
		defer C.free(unsafe.Pointer(vp))
		features_ptr = append(features_ptr, vp)
	}
	features_ptr = append(features_ptr, nil)
	fp := (**C.char)(unsafe.Pointer(&features_ptr))
	m := C.ly_ctx_load_module(ctx.raw, n, r, fp)
	if m == nil {
		return nil
	}
	return SchemaModuleFromRaw(ctx, m)
}

// export _cgo_ly_module_import_cb
func _cgo_ly_module_import_cb(mod_name *C.char, mod_rev *C.char, submod_name *C.char, submod_rev *C.char, user_data *C.void, format C.LYS_INFORMAT, module_data *C.char, _free_module_data *C.ly_module_imp_data_free_clb) C.LY_ERR {
	mn := C.GoString(mod_name)
	mrv := C.GoString(mod_rev)
	smn := C.GoString(submod_name)
	smrv := C.GoString(submod_rev)
	m := *(*EmbeddedModules)(unsafe.Pointer(user_data))
	key := EmbeddedMoudleKey{
		mod_name:    mn,
		mod_rev:     mrv,
		submod_name: smn,
		submod_rev:  smrv,
	}
	v, find := m[key]
	if !find {
		return C.LY_ENOTFOUND
	}
	// leak the data on purpose
	data := C.CString(v)
	format = C.LYS_IN_YANG
	*module_data = *data
	return C.LY_SUCCESS
}

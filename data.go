package libyango

/*
#include <libyang/libyang.h>
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"reflect"
	"unsafe"
)

type DataTree struct {
	ctx *Context
	raw *C.struct_lyd_node
}

type Metadata struct {
	dnode *DataTree
	raw   *C.struct_lyd_meta
}

const (
	DataDiffCreate = iota
	DataDiffDelete
	DataDiffReplace
)

const (
	DataFormatXML uint = iota + 1
	DataFormatJSON
	DataFormatLYB
)

func (tree *DataTree) FindXpath(xpath string) (*Set[*DataTree, CLydNode], error) {
	xp := C.CString(xpath)
	defer C.free(unsafe.Pointer(xp))
	var set *C.struct_ly_set = nil
	if ret := C.lyd_find_xpath(tree.raw, xp, &set); ret != C.LY_SUCCESS {
		return nil, fmt.Errorf("find_xpath error, error code: %d", ret)
	}
	rnodesCount := int((*set).count)
	if rnodesCount == 0 {
		return NewSet[*DataTree, CLydNode](tree, nil), nil
	} else {
		var sp uintptr
		if BigEndian {
			sp = uintptr(binary.BigEndian.Uint64((*set).anon0[:]))
		} else {
			sp = uintptr(binary.LittleEndian.Uint64((*set).anon0[:]))
		}
		var s []CLydNode
		sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
		sh.Data = sp
		sh.Len = rnodesCount
		sh.Cap = rnodesCount
		return NewSet(tree, s), nil
	}
}

func (tree *DataTree) FindPath(path string) (*DataTree, error) {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))
	var rnode *C.struct_lyd_node
	if ret := C.lyd_find_path(tree.raw, p, 0, &rnode); ret != C.LY_SUCCESS {
		return nil, fmt.Errorf("find_xpath error, error code: %d", ret)
	}
	return &DataTree{
		ctx: tree.ctx,
		raw: rnode,
	}, nil
}

package libyango

/*
#include <libyang/libyang.h>
*/
import "C"
import "unsafe"

type SetContainer interface {
	*Context | *DataTree
}

type CLyscNode = *C.struct_lysc_node
type CLydNode = *C.struct_lyd_node
type SetSliceType interface {
	CLyscNode | CLydNode
}

type Set[C SetContainer, T SetSliceType] struct {
	container C
	slice     []T
}

func NewSet[C SetContainer, T SetSliceType](container C, slice []T) *Set[C, T] {
	return &Set[C, T]{
		container,
		slice,
	}
}

const INT_SIZE int = int(unsafe.Sizeof(0))
var BigEndian = true

func init() {
	BigEndian = getEndian()
}

func getEndian() bool {
	var i int = 0x1
	bs := (*[INT_SIZE]byte)(unsafe.Pointer(&i))
	if bs[0] == 0 {
		return true
	} else {
		return false
	}

}
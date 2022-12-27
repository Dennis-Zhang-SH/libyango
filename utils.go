package libyango

/*
#include <libyang/libyang.h>
*/
import "C"

type SetContainer interface {
	*Context
}

type CLyscNode = *C.struct_lysc_node

type SetSliceType interface {
	CLyscNode
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

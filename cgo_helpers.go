// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
//
// * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//
// * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following
// disclaimer in the documentation and/or other materials provided
// with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//

// WARNING: This file has automatically been generated on Thu, 22 Sep 2016 21:30:57 UTC.
// By https://git.io/cgogen. DO NOT EDIT.

package ykpiv

/*
#cgo LDFLAGS: -lykpiv -lcrypto -lpcsclite -lpthread
#include "ykpiv.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"runtime"
	"sync"
	"unsafe"
)

// Ref returns a reference to C object as it is.
func (x *State) Ref() *C.ykpiv_state {
	if x == nil {
		return nil
	}
	return (*C.ykpiv_state)(unsafe.Pointer(x))
}

// Free cleanups the referenced memory using C free.
func (x *State) Free() {
	if x != nil {
		C.free(unsafe.Pointer(x))
	}
}

// NewStateRef converts the C object reference into a raw struct reference without wrapping.
func NewStateRef(ref unsafe.Pointer) *State {
	return (*State)(ref)
}

// NewState allocates a new C object of this type and converts the reference into
// a raw struct reference without wrapping.
func NewState() *State {
	return (*State)(allocStateMemory(1))
}

// allocStateMemory allocates memory for type C.ykpiv_state in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStateMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStateValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStateValue = unsafe.Sizeof([1]C.ykpiv_state{})

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// PassRef returns a reference to C object as it is or allocates a new C object of this type.
func (x *State) PassRef() *C.ykpiv_state {
	if x == nil {
		x = (*State)(allocStateMemory(1))
	}
	return (*C.ykpiv_state)(unsafe.Pointer(x))
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

type stringHeader struct {
	Data uintptr
	Len  int
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
}

type sliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// allocPStateMemory allocates memory for type *C.ykpiv_state in C.
// The caller is responsible for freeing the this memory via C.free.
func allocPStateMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfPStateValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfPStateValue = unsafe.Sizeof([1]*C.ykpiv_state{})

const sizeOfPtr = unsafe.Sizeof(&struct{}{})

// unpackArgSSState transforms a sliced Go data structure into plain C format.
func unpackArgSSState(x [][]State) (unpacked **C.ykpiv_state, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	allocs = new(cgoAllocMap)
	defer runtime.SetFinalizer(&unpacked, func(***C.ykpiv_state) {
		go allocs.Free()
	})

	len0 := len(x)
	mem0 := allocPStateMemory(len0)
	allocs.Add(mem0)
	h0 := &sliceHeader{
		Data: uintptr(mem0),
		Cap:  len0,
		Len:  len0,
	}
	v0 := *(*[]*C.ykpiv_state)(unsafe.Pointer(h0))
	for i0 := range x {
		h := (*sliceHeader)(unsafe.Pointer(&x[i0]))
		v0[i0] = (*C.ykpiv_state)(unsafe.Pointer(h.Data))
	}
	h := (*sliceHeader)(unsafe.Pointer(&v0))
	unpacked = (**C.ykpiv_state)(unsafe.Pointer(h.Data))
	return
}

// packSSState reads sliced Go data structure out from plain C format.
func packSSState(v [][]State, ptr0 **C.ykpiv_state) {
	const m = 0x7fffffff
	for i0 := range v {
		ptr1 := (*(*[m / sizeOfPtr]*C.ykpiv_state)(unsafe.Pointer(ptr0)))[i0]
		hxfc4425b := (*sliceHeader)(unsafe.Pointer(&v[i0]))
		hxfc4425b.Data = uintptr(unsafe.Pointer(ptr1))
		hxfc4425b.Cap = 0x7fffffff
		// hxfc4425b.Len = ?
	}
}

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

// unpackPUcharString represents the data from Go string as *C.uchar and avoids copying.
func unpackPUcharString(str string) (*C.uchar, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.uchar)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

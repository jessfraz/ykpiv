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
import "unsafe"

// Strerror function as declared in ykpiv/ykpiv.h:63
func Strerror(err Rc) string {
	cerr, _ := (C.ykpiv_rc)(err), cgoAllocsUnknown
	__ret := C.ykpiv_strerror(cerr)
	__v := packPCharString(__ret)
	return __v
}

// StrerrorName function as declared in ykpiv/ykpiv.h:64
func StrerrorName(err Rc) string {
	cerr, _ := (C.ykpiv_rc)(err), cgoAllocsUnknown
	__ret := C.ykpiv_strerror_name(cerr)
	__v := packPCharString(__ret)
	return __v
}

// Init function as declared in ykpiv/ykpiv.h:66
func Init(state [][]State, verbose int32) Rc {
	cstate, _ := unpackArgSSState(state)
	cverbose, _ := (C.int)(verbose), cgoAllocsUnknown
	__ret := C.ykpiv_init(cstate, cverbose)
	packSSState(state, cstate)
	__v := (Rc)(__ret)
	return __v
}

// Done function as declared in ykpiv/ykpiv.h:67
func Done(state []State) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	__ret := C.ykpiv_done(cstate)
	__v := (Rc)(__ret)
	return __v
}

// Connect function as declared in ykpiv/ykpiv.h:68
func Connect(state []State, wanted string) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	cwanted, _ := unpackPCharString(wanted)
	__ret := C.ykpiv_connect(cstate, cwanted)
	__v := (Rc)(__ret)
	return __v
}

// ListReaders function as declared in ykpiv/ykpiv.h:69
func ListReaders(state []State, readers []byte, len []uint) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	creaders, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&readers)).Data)), cgoAllocsUnknown
	clen, _ := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&len)).Data)), cgoAllocsUnknown
	__ret := C.ykpiv_list_readers(cstate, creaders, clen)
	__v := (Rc)(__ret)
	return __v
}

// Disconnect function as declared in ykpiv/ykpiv.h:70
func Disconnect(state []State) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	__ret := C.ykpiv_disconnect(cstate)
	__v := (Rc)(__ret)
	return __v
}

// TransferData function as declared in ykpiv/ykpiv.h:71
func TransferData(state []State, templ string, inData string, inLen int, outData []byte, outLen []uint, sw []int32) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	ctempl, _ := unpackPUcharString(templ)
	cinData, _ := unpackPUcharString(inData)
	cinLen, _ := (C.long)(inLen), cgoAllocsUnknown
	coutData, _ := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outData)).Data)), cgoAllocsUnknown
	coutLen, _ := (*C.ulong)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outLen)).Data)), cgoAllocsUnknown
	csw, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&sw)).Data)), cgoAllocsUnknown
	__ret := C.ykpiv_transfer_data(cstate, ctempl, cinData, cinLen, coutData, coutLen, csw)
	__v := (Rc)(__ret)
	return __v
}

// Authenticate function as declared in ykpiv/ykpiv.h:74
func Authenticate(state []State, key string) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	ckey, _ := unpackPUcharString(key)
	__ret := C.ykpiv_authenticate(cstate, ckey)
	__v := (Rc)(__ret)
	return __v
}

// SetMgmkey function as declared in ykpiv/ykpiv.h:75
func SetMgmkey(state []State, newKey string) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	cnewKey, _ := unpackPUcharString(newKey)
	__ret := C.ykpiv_set_mgmkey(cstate, cnewKey)
	__v := (Rc)(__ret)
	return __v
}

// HexDecode function as declared in ykpiv/ykpiv.h:76
func HexDecode(hexIn string, inLen uint, hexOut []byte, outLen []uint) Rc {
	chexIn, _ := unpackPCharString(hexIn)
	cinLen, _ := (C.size_t)(inLen), cgoAllocsUnknown
	chexOut, _ := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&hexOut)).Data)), cgoAllocsUnknown
	coutLen, _ := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outLen)).Data)), cgoAllocsUnknown
	__ret := C.ykpiv_hex_decode(chexIn, cinLen, chexOut, coutLen)
	__v := (Rc)(__ret)
	return __v
}

// SignData function as declared in ykpiv/ykpiv.h:78
func SignData(state []State, signIn string, inLen uint, signOut []byte, outLen []uint, algorithm byte, key byte) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	csignIn, _ := unpackPUcharString(signIn)
	cinLen, _ := (C.size_t)(inLen), cgoAllocsUnknown
	csignOut, _ := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&signOut)).Data)), cgoAllocsUnknown
	coutLen, _ := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outLen)).Data)), cgoAllocsUnknown
	calgorithm, _ := (C.uchar)(algorithm), cgoAllocsUnknown
	ckey, _ := (C.uchar)(key), cgoAllocsUnknown
	__ret := C.ykpiv_sign_data(cstate, csignIn, cinLen, csignOut, coutLen, calgorithm, ckey)
	__v := (Rc)(__ret)
	return __v
}

// DecipherData function as declared in ykpiv/ykpiv.h:84
func DecipherData(state []State, encIn string, inLen uint, encOut []byte, outLen []uint, algorithm byte, key byte) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	cencIn, _ := unpackPUcharString(encIn)
	cinLen, _ := (C.size_t)(inLen), cgoAllocsUnknown
	cencOut, _ := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&encOut)).Data)), cgoAllocsUnknown
	coutLen, _ := (*C.size_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&outLen)).Data)), cgoAllocsUnknown
	calgorithm, _ := (C.uchar)(algorithm), cgoAllocsUnknown
	ckey, _ := (C.uchar)(key), cgoAllocsUnknown
	__ret := C.ykpiv_decipher_data(cstate, cencIn, cinLen, cencOut, coutLen, calgorithm, ckey)
	__v := (Rc)(__ret)
	return __v
}

// GetVersion function as declared in ykpiv/ykpiv.h:87
func GetVersion(state []State, version []byte, len uint) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	cversion, _ := (*C.char)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&version)).Data)), cgoAllocsUnknown
	clen, _ := (C.size_t)(len), cgoAllocsUnknown
	__ret := C.ykpiv_get_version(cstate, cversion, clen)
	__v := (Rc)(__ret)
	return __v
}

// Verify function as declared in ykpiv/ykpiv.h:88
func Verify(state []State, pin string, tries []int32) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	cpin, _ := unpackPCharString(pin)
	ctries, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&tries)).Data)), cgoAllocsUnknown
	__ret := C.ykpiv_verify(cstate, cpin, ctries)
	__v := (Rc)(__ret)
	return __v
}

// ChangePin function as declared in ykpiv/ykpiv.h:89
func ChangePin(state []State, currentPin string, currentPinLen uint, newPin string, newPinLen uint, tries []int32) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	ccurrentPin, _ := unpackPCharString(currentPin)
	ccurrentPinLen, _ := (C.size_t)(currentPinLen), cgoAllocsUnknown
	cnewPin, _ := unpackPCharString(newPin)
	cnewPinLen, _ := (C.size_t)(newPinLen), cgoAllocsUnknown
	ctries, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&tries)).Data)), cgoAllocsUnknown
	__ret := C.ykpiv_change_pin(cstate, ccurrentPin, ccurrentPinLen, cnewPin, cnewPinLen, ctries)
	__v := (Rc)(__ret)
	return __v
}

// ChangePuk function as declared in ykpiv/ykpiv.h:92
func ChangePuk(state []State, currentPuk string, currentPukLen uint, newPuk string, newPukLen uint, tries []int32) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	ccurrentPuk, _ := unpackPCharString(currentPuk)
	ccurrentPukLen, _ := (C.size_t)(currentPukLen), cgoAllocsUnknown
	cnewPuk, _ := unpackPCharString(newPuk)
	cnewPukLen, _ := (C.size_t)(newPukLen), cgoAllocsUnknown
	ctries, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&tries)).Data)), cgoAllocsUnknown
	__ret := C.ykpiv_change_puk(cstate, ccurrentPuk, ccurrentPukLen, cnewPuk, cnewPukLen, ctries)
	__v := (Rc)(__ret)
	return __v
}

// UnblockPin function as declared in ykpiv/ykpiv.h:95
func UnblockPin(state []State, puk string, pukLen uint, newPin string, newPinLen uint, tries []int32) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	cpuk, _ := unpackPCharString(puk)
	cpukLen, _ := (C.size_t)(pukLen), cgoAllocsUnknown
	cnewPin, _ := unpackPCharString(newPin)
	cnewPinLen, _ := (C.size_t)(newPinLen), cgoAllocsUnknown
	ctries, _ := (*C.int)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&tries)).Data)), cgoAllocsUnknown
	__ret := C.ykpiv_unblock_pin(cstate, cpuk, cpukLen, cnewPin, cnewPinLen, ctries)
	__v := (Rc)(__ret)
	return __v
}

// FetchObject function as declared in ykpiv/ykpiv.h:98
func FetchObject(state []State, objectId int32, data []byte, len []uint) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	cobjectId, _ := (C.int)(objectId), cgoAllocsUnknown
	cdata, _ := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&data)).Data)), cgoAllocsUnknown
	clen, _ := (*C.ulong)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&len)).Data)), cgoAllocsUnknown
	__ret := C.ykpiv_fetch_object(cstate, cobjectId, cdata, clen)
	__v := (Rc)(__ret)
	return __v
}

// SetMgmkey2 function as declared in ykpiv/ykpiv.h:100
func SetMgmkey2(state []State, newKey string, touch byte) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	cnewKey, _ := unpackPUcharString(newKey)
	ctouch, _ := (C.uchar)(touch), cgoAllocsUnknown
	__ret := C.ykpiv_set_mgmkey2(cstate, cnewKey, ctouch)
	__v := (Rc)(__ret)
	return __v
}

// SaveObject function as declared in ykpiv/ykpiv.h:102
func SaveObject(state []State, objectId int32, indata []byte, len uint) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	cobjectId, _ := (C.int)(objectId), cgoAllocsUnknown
	cindata, _ := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&indata)).Data)), cgoAllocsUnknown
	clen, _ := (C.size_t)(len), cgoAllocsUnknown
	__ret := C.ykpiv_save_object(cstate, cobjectId, cindata, clen)
	__v := (Rc)(__ret)
	return __v
}

// ImportPrivateKey function as declared in ykpiv/ykpiv.h:104
func ImportPrivateKey(state []State, key byte, algorithm byte, p string, pLen uint, q string, qLen uint, dp string, dpLen uint, dq string, dqLen uint, qinv string, qinvLen uint, ecData string, ecDataLen byte, pinPolicy byte, touchPolicy byte) Rc {
	cstate, _ := (*C.ykpiv_state)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&state)).Data)), cgoAllocsUnknown
	ckey, _ := (C.uchar)(key), cgoAllocsUnknown
	calgorithm, _ := (C.uchar)(algorithm), cgoAllocsUnknown
	cp, _ := unpackPUcharString(p)
	cpLen, _ := (C.size_t)(pLen), cgoAllocsUnknown
	cq, _ := unpackPUcharString(q)
	cqLen, _ := (C.size_t)(qLen), cgoAllocsUnknown
	cdp, _ := unpackPUcharString(dp)
	cdpLen, _ := (C.size_t)(dpLen), cgoAllocsUnknown
	cdq, _ := unpackPUcharString(dq)
	cdqLen, _ := (C.size_t)(dqLen), cgoAllocsUnknown
	cqinv, _ := unpackPUcharString(qinv)
	cqinvLen, _ := (C.size_t)(qinvLen), cgoAllocsUnknown
	cecData, _ := unpackPUcharString(ecData)
	cecDataLen, _ := (C.uchar)(ecDataLen), cgoAllocsUnknown
	cpinPolicy, _ := (C.uchar)(pinPolicy), cgoAllocsUnknown
	ctouchPolicy, _ := (C.uchar)(touchPolicy), cgoAllocsUnknown
	__ret := C.ykpiv_import_private_key(cstate, ckey, calgorithm, cp, cpLen, cq, cqLen, cdp, cdpLen, cdq, cdqLen, cqinv, cqinvLen, cecData, cecDataLen, cpinPolicy, ctouchPolicy)
	__v := (Rc)(__ret)
	return __v
}

// CheckVersion function as declared in lib/ykpiv-version.h:85
func CheckVersion(reqVersion string) string {
	creqVersion, _ := unpackPCharString(reqVersion)
	__ret := C.ykpiv_check_version(creqVersion)
	__v := packPCharString(__ret)
	return __v
}

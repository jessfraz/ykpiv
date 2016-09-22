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

// WARNING: This file has automatically been generated on Thu, 22 Sep 2016 16:48:47 UTC.
// By https://git.io/cgogen. DO NOT EDIT.

package ykpiv

/*
#cgo LDFLAGS: -lykpiv
#include "ykpiv.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// ALGOTAG as defined in ykpiv/ykpiv.h:113
	ALGOTAG = 0x80
	// ALGO3DES as defined in ykpiv/ykpiv.h:114
	ALGO3DES = 0x03
	// ALGORSA1024 as defined in ykpiv/ykpiv.h:115
	ALGORSA1024 = 0x06
	// ALGORSA2048 as defined in ykpiv/ykpiv.h:116
	ALGORSA2048 = 0x07
	// ALGOECCP256 as defined in ykpiv/ykpiv.h:117
	ALGOECCP256 = 0x11
	// ALGOECCP384 as defined in ykpiv/ykpiv.h:118
	ALGOECCP384 = 0x14
	// KEYAUTHENTICATION as defined in ykpiv/ykpiv.h:120
	KEYAUTHENTICATION = 0x9a
	// KEYCARDMGM as defined in ykpiv/ykpiv.h:121
	KEYCARDMGM = 0x9b
	// KEYSIGNATURE as defined in ykpiv/ykpiv.h:122
	KEYSIGNATURE = 0x9c
	// KEYKEYMGM as defined in ykpiv/ykpiv.h:123
	KEYKEYMGM = 0x9d
	// KEYCARDAUTH as defined in ykpiv/ykpiv.h:124
	KEYCARDAUTH = 0x9e
	// KEYRETIRED1 as defined in ykpiv/ykpiv.h:125
	KEYRETIRED1 = 0x82
	// KEYRETIRED2 as defined in ykpiv/ykpiv.h:126
	KEYRETIRED2 = 0x83
	// KEYRETIRED3 as defined in ykpiv/ykpiv.h:127
	KEYRETIRED3 = 0x84
	// KEYRETIRED4 as defined in ykpiv/ykpiv.h:128
	KEYRETIRED4 = 0x85
	// KEYRETIRED5 as defined in ykpiv/ykpiv.h:129
	KEYRETIRED5 = 0x86
	// KEYRETIRED6 as defined in ykpiv/ykpiv.h:130
	KEYRETIRED6 = 0x87
	// KEYRETIRED7 as defined in ykpiv/ykpiv.h:131
	KEYRETIRED7 = 0x88
	// KEYRETIRED8 as defined in ykpiv/ykpiv.h:132
	KEYRETIRED8 = 0x89
	// KEYRETIRED9 as defined in ykpiv/ykpiv.h:133
	KEYRETIRED9 = 0x8a
	// KEYRETIRED10 as defined in ykpiv/ykpiv.h:134
	KEYRETIRED10 = 0x8b
	// KEYRETIRED11 as defined in ykpiv/ykpiv.h:135
	KEYRETIRED11 = 0x8c
	// KEYRETIRED12 as defined in ykpiv/ykpiv.h:136
	KEYRETIRED12 = 0x8d
	// KEYRETIRED13 as defined in ykpiv/ykpiv.h:137
	KEYRETIRED13 = 0x8e
	// KEYRETIRED14 as defined in ykpiv/ykpiv.h:138
	KEYRETIRED14 = 0x8f
	// KEYRETIRED15 as defined in ykpiv/ykpiv.h:139
	KEYRETIRED15 = 0x90
	// KEYRETIRED16 as defined in ykpiv/ykpiv.h:140
	KEYRETIRED16 = 0x91
	// KEYRETIRED17 as defined in ykpiv/ykpiv.h:141
	KEYRETIRED17 = 0x92
	// KEYRETIRED18 as defined in ykpiv/ykpiv.h:142
	KEYRETIRED18 = 0x93
	// KEYRETIRED19 as defined in ykpiv/ykpiv.h:143
	KEYRETIRED19 = 0x94
	// KEYRETIRED20 as defined in ykpiv/ykpiv.h:144
	KEYRETIRED20 = 0x95
	// KEYATTESTATION as defined in ykpiv/ykpiv.h:145
	KEYATTESTATION = 0xf9
	// OBJCAPABILITY as defined in ykpiv/ykpiv.h:147
	OBJCAPABILITY = 0x5fc107
	// OBJCHUID as defined in ykpiv/ykpiv.h:148
	OBJCHUID = 0x5fc102
	// OBJAUTHENTICATION as defined in ykpiv/ykpiv.h:149
	OBJAUTHENTICATION = 0x5fc105
	// OBJFINGERPRINTS as defined in ykpiv/ykpiv.h:150
	OBJFINGERPRINTS = 0x5fc103
	// OBJSECURITY as defined in ykpiv/ykpiv.h:151
	OBJSECURITY = 0x5fc106
	// OBJFACIAL as defined in ykpiv/ykpiv.h:152
	OBJFACIAL = 0x5fc108
	// OBJPRINTED as defined in ykpiv/ykpiv.h:153
	OBJPRINTED = 0x5fc109
	// OBJSIGNATURE as defined in ykpiv/ykpiv.h:154
	OBJSIGNATURE = 0x5fc10a
	// OBJKEYMANAGEMENT as defined in ykpiv/ykpiv.h:155
	OBJKEYMANAGEMENT = 0x5fc10b
	// OBJCARDAUTH as defined in ykpiv/ykpiv.h:156
	OBJCARDAUTH = 0x5fc101
	// OBJDISCOVERY as defined in ykpiv/ykpiv.h:157
	OBJDISCOVERY = 0x7e
	// OBJKEYHISTORY as defined in ykpiv/ykpiv.h:158
	OBJKEYHISTORY = 0x5fc10c
	// OBJIRIS as defined in ykpiv/ykpiv.h:159
	OBJIRIS = 0x5fc121
	// OBJRETIRED1 as defined in ykpiv/ykpiv.h:161
	OBJRETIRED1 = 0x5fc10d
	// OBJRETIRED2 as defined in ykpiv/ykpiv.h:162
	OBJRETIRED2 = 0x5fc10e
	// OBJRETIRED3 as defined in ykpiv/ykpiv.h:163
	OBJRETIRED3 = 0x5fc10f
	// OBJRETIRED4 as defined in ykpiv/ykpiv.h:164
	OBJRETIRED4 = 0x5fc110
	// OBJRETIRED5 as defined in ykpiv/ykpiv.h:165
	OBJRETIRED5 = 0x5fc111
	// OBJRETIRED6 as defined in ykpiv/ykpiv.h:166
	OBJRETIRED6 = 0x5fc112
	// OBJRETIRED7 as defined in ykpiv/ykpiv.h:167
	OBJRETIRED7 = 0x5fc113
	// OBJRETIRED8 as defined in ykpiv/ykpiv.h:168
	OBJRETIRED8 = 0x5fc114
	// OBJRETIRED9 as defined in ykpiv/ykpiv.h:169
	OBJRETIRED9 = 0x5fc115
	// OBJRETIRED10 as defined in ykpiv/ykpiv.h:170
	OBJRETIRED10 = 0x5fc116
	// OBJRETIRED11 as defined in ykpiv/ykpiv.h:171
	OBJRETIRED11 = 0x5fc117
	// OBJRETIRED12 as defined in ykpiv/ykpiv.h:172
	OBJRETIRED12 = 0x5fc118
	// OBJRETIRED13 as defined in ykpiv/ykpiv.h:173
	OBJRETIRED13 = 0x5fc119
	// OBJRETIRED14 as defined in ykpiv/ykpiv.h:174
	OBJRETIRED14 = 0x5fc11a
	// OBJRETIRED15 as defined in ykpiv/ykpiv.h:175
	OBJRETIRED15 = 0x5fc11b
	// OBJRETIRED16 as defined in ykpiv/ykpiv.h:176
	OBJRETIRED16 = 0x5fc11c
	// OBJRETIRED17 as defined in ykpiv/ykpiv.h:177
	OBJRETIRED17 = 0x5fc11d
	// OBJRETIRED18 as defined in ykpiv/ykpiv.h:178
	OBJRETIRED18 = 0x5fc11e
	// OBJRETIRED19 as defined in ykpiv/ykpiv.h:179
	OBJRETIRED19 = 0x5fc11f
	// OBJRETIRED20 as defined in ykpiv/ykpiv.h:180
	OBJRETIRED20 = 0x5fc120
	// OBJATTESTATION as defined in ykpiv/ykpiv.h:182
	OBJATTESTATION = 0x5fff01
	// INSVERIFY as defined in ykpiv/ykpiv.h:184
	INSVERIFY = 0x20
	// INSCHANGEREFERENCE as defined in ykpiv/ykpiv.h:185
	INSCHANGEREFERENCE = 0x24
	// INSRESETRETRY as defined in ykpiv/ykpiv.h:186
	INSRESETRETRY = 0x2c
	// INSGENERATEASYMMETRIC as defined in ykpiv/ykpiv.h:187
	INSGENERATEASYMMETRIC = 0x47
	// INSAUTHENTICATE as defined in ykpiv/ykpiv.h:188
	INSAUTHENTICATE = 0x87
	// INSGETDATA as defined in ykpiv/ykpiv.h:189
	INSGETDATA = 0xcb
	// INSPUTDATA as defined in ykpiv/ykpiv.h:190
	INSPUTDATA = 0xdb
	// INSSETMGMKEY as defined in ykpiv/ykpiv.h:201
	INSSETMGMKEY = 0xff
	// INSIMPORTKEY as defined in ykpiv/ykpiv.h:202
	INSIMPORTKEY = 0xfe
	// INSGETVERSION as defined in ykpiv/ykpiv.h:203
	INSGETVERSION = 0xfd
	// INSRESET as defined in ykpiv/ykpiv.h:204
	INSRESET = 0xfb
	// INSSETPINRETRIES as defined in ykpiv/ykpiv.h:205
	INSSETPINRETRIES = 0xfa
	// INSATTEST as defined in ykpiv/ykpiv.h:206
	INSATTEST = 0xf9
	// PINPOLICYTAG as defined in ykpiv/ykpiv.h:208
	PINPOLICYTAG = 0xaa
	// PINPOLICYDEFAULT as defined in ykpiv/ykpiv.h:209
	PINPOLICYDEFAULT = 0
	// PINPOLICYNEVER as defined in ykpiv/ykpiv.h:210
	PINPOLICYNEVER = 1
	// PINPOLICYONCE as defined in ykpiv/ykpiv.h:211
	PINPOLICYONCE = 2
	// PINPOLICYALWAYS as defined in ykpiv/ykpiv.h:212
	PINPOLICYALWAYS = 3
	// TOUCHPOLICYTAG as defined in ykpiv/ykpiv.h:214
	TOUCHPOLICYTAG = 0xab
	// TOUCHPOLICYDEFAULT as defined in ykpiv/ykpiv.h:215
	TOUCHPOLICYDEFAULT = 0
	// TOUCHPOLICYNEVER as defined in ykpiv/ykpiv.h:216
	TOUCHPOLICYNEVER = 1
	// TOUCHPOLICYALWAYS as defined in ykpiv/ykpiv.h:217
	TOUCHPOLICYALWAYS = 2
	// TOUCHPOLICYCACHED as defined in ykpiv/ykpiv.h:218
	TOUCHPOLICYCACHED = 3
	// VERSIONSTRING as defined in lib/ykpiv-version.h:46
	VERSIONSTRING = "1.4.2"
	// VERSIONNUMBER as defined in lib/ykpiv-version.h:56
	VERSIONNUMBER = 0x010402
	// VERSIONMAJOR as defined in lib/ykpiv-version.h:65
	VERSIONMAJOR = 1
	// VERSIONMINOR as defined in lib/ykpiv-version.h:74
	VERSIONMINOR = 4
	// VERSIONPATCH as defined in lib/ykpiv-version.h:83
	VERSIONPATCH = 2
)

// Rc as declared in ykpiv/ykpiv.h:61
type Rc int32

// Rc enumeration from ykpiv/ykpiv.h:61
const (
	OK                  Rc = iota
	MEMORYERROR         Rc = -1
	PCSCERROR           Rc = -2
	SIZEERROR           Rc = -3
	APPLETERROR         Rc = -4
	AUTHENTICATIONERROR Rc = -5
	RANDOMNESSERROR     Rc = -6
	GENERICERROR        Rc = -7
	KEYERROR            Rc = -8
	PARSEERROR          Rc = -9
	WRONGPIN            Rc = -10
	INVALIDOBJECT       Rc = -11
	ALGORITHMERROR      Rc = -12
	PINLOCKED           Rc = -13
)

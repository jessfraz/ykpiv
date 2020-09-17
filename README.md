# ykpiv

[![make-all](https://github.com/jessfraz/ykpiv/workflows/make%20all/badge.svg)](https://github.com/jessfraz/ykpiv/actions?query=workflow%3A%22make+all%22)
[![make-image](https://github.com/jessfraz/ykpiv/workflows/make%20image/badge.svg)](https://github.com/jessfraz/ykpiv/actions?query=workflow%3A%22make+image%22)
[![make-test](https://github.com/jessfraz/ykpiv/workflows/make%20test/badge.svg)](https://github.com/jessfraz/ykpiv/actions?query=workflow%3A%22make+test%22)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/jessfraz/ykpiv)

Go bindings for ykpiv so you can write Go to interact with your yubikeys.
The C library lives at
[yubico/yubico-piv-tool](https://github.com/Yubico/yubico-piv-tool).

## C Libraries Required for Compilation

- libcrypto
- [libykpiv](https://github.com/Yubico/yubico-piv-tool/)
- [libpcsclite](https://pcsclite.alioth.debian.org/pcsclite.html)
- [ccid](https://pcsclite.alioth.debian.org/ccid.html)

## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/jessfraz/ykpiv"
)

func main() {
	s := ykpiv.NewState()
	defer s.Free()

	// Let's get the readers
	readers := make([]byte, 2048)
	len := []uint{2048}
	log.Println("list")
	err := ykpiv.ListReaders([]ykpiv.State{*s}, readers, len)
	if err != 0 {
		log.Fatalf("%s: %#v", ykpiv.Strerror(err), err)
	}

	fmt.Printf("readers: %s\n", string(readers))
}
```

## Starting `pcscd`

Hopefully your operating system does this for you with a nice init script but
if not here you go:

```console
$ sudo LIBCCID_ifdLogLevel=0x000F /usr/sbin/pcscd --foreground --debug --apdu --color
$ sudo /usr/sbin/pcscd --hotplug
```

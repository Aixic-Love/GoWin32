package main

import (
	"fmt"
	"syscall"
	"unsafe"
)
type SHARE_INFO_2 struct {
	shi2_netname       *uint16
	shi2_type          uint32
	shi2_remark        *uint16
	shi2_permissions   uint32
	shi2_max_uses      uint32
	shi2_current_uses  uint32
	shi2_path          *uint16
	shi2_passwd        *uint16
}

const (
	SHARE_MAX_PREFERRED_LENGTH = 0xFFFFFFFF
)

func main() {

	var (
		netapi32            = syscall.NewLazyDLL("Netapi32.dll")
		netapiShareEnumProc = netapi32.NewProc("NetShareEnum")
		usrNetApiBufferFree = netapi32.NewProc("NetApiBufferFree")
		dataPointer  uintptr
		sizeTest     SHARE_INFO_2
	)
	var (
		resume       uint32
		entriesRead  uint32
		entriesTotal uint32

	)
	r0, _, _ := netapiShareEnumProc.Call(
		0, //uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(serverName))),
		uintptr(2),
		uintptr(unsafe.Pointer(&dataPointer)),
		uintptr(uint32(SHARE_MAX_PREFERRED_LENGTH)),
		uintptr(unsafe.Pointer(&entriesRead)),
		uintptr(unsafe.Pointer(&entriesTotal)),
		uintptr(unsafe.Pointer(&resume)),
	)
	if r0 != 0 {
		return
	}
	var iter = dataPointer
	for i := uint32(0); i < entriesRead; i++ {
		var data = (*SHARE_INFO_2)(unsafe.Pointer(iter))
		ud := UTF16toString(data.shi2_netname)
		uds := UTF16toString(data.shi2_path)

		fmt.Println(ud)
		fmt.Println(uds)

		iter = uintptr(unsafe.Pointer(iter + unsafe.Sizeof(sizeTest)))
	}
	usrNetApiBufferFree.Call(dataPointer)
	return
}

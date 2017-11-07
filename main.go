package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	//"github.com/hashicorp/vault/builtin/audit/file"
)

func visit(path string, f os.FileInfo, err error) error {
	// fmt.Printf( "%s\n", path )
	file := open_file(path)
	hash, err := hash_file_md5(file)
	fmt.Println(hash)
	if err == nil {
		fmt.Println(err)
	}
	size_file(file)
	defer file.Close()
	return nil
}

func open_file(filepath string) *os.File {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Print("Error al abrir el archivo")
	}

	return file
}

func size_file(file *os.File) {
	file_statics, err := file.Stat()

	if err == nil {
		fmt.Println(err)
	}
	file_statics.Size()
	//fmt.Println("Tamaño: ", file_statics.Size())
}

func hash_file_md5(file *os.File) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Tell the program to call the following function when the current function returns
	//defer file.Close()

	//Open a new hash interface to write to
	hash := md5.New()

	//Copy the file in the hash interface and check for any error
	if _, err := io.Copy(hash, file); err != nil {
		return returnMD5String, err
	}

	//Get the 16 bytes hash
	hashInBytes := hash.Sum(nil)[:16]

	//Convert the bytes to a string
	returnMD5String = hex.EncodeToString(hashInBytes)

	return returnMD5String, nil

}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)

}

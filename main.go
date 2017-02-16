package main

import (
	"path/filepath"
	"flag"
	"fmt"
	"os"
	"crypto/md5"
	"io"
	"encoding/hex"
)

func visit( path string, f os.FileInfo, err error ) error {
	// fmt.Printf( "%s\n", path )
	hash, err := hash_file_md5(path)
	if err == nil {
		fmt.Println(hash, path)
	}
	return nil
}

func hash_file_md5(filePath string) (string, error) {
	//Initialize variable returnMD5String now in case an error has to be returned
	var returnMD5String string

	//Open the passed argument and check for any error
	file, err := os.Open(filePath)
	if err != nil {
		return returnMD5String, err
	}

	//Tell the program to call the following function when the current function returns
	defer file.Close()

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
	err := filepath.Walk( root, visit )
	fmt.Printf( "filepath.Walk() returned %v\n", err )




}

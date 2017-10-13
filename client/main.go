package main

import (
	"net"
	"os"
	"io"
	"net_test/utils"
)

func main() {
	srcFile := "./aFile"
	dstFile := "./file1"
	serverAddr := "localhost:4040"
	uploadFile(srcFile, serverAddr)
	downloadFile(dstFile, serverAddr)
}

func uploadFile(srcFile, serverAddr string) {
	// connect to server
	conn, err := net.Dial("tcp", serverAddr)
	utils.CheckError(err)
	defer conn.Close()

	// open file to upload
	fi, err := os.Open(srcFile)
	utils.CheckError(err)
	defer fi.Close()

	// upload
	_, err = io.Copy(conn, fi)
	utils.CheckError(err)
}

func downloadFile(dstFile, serverAddr string) {
	// create new file to hold response
	fo, err := os.Create(dstFile)
	utils.CheckError(err)
	defer fo.Close()

	// connect to server
	conn, err := net.Dial("tcp", serverAddr)
	utils.CheckError(err)
	defer conn.Close()

	_, err = io.Copy(fo, conn)
	utils.CheckError(err)
}

package main

import (
	"net"
	"os"
	"io"
	"net_test/utils"
)

func main() {
	dstFile := "./bFile"
	srcFile := "./file2"
	serverAddr := "localhost:4040"

	// server start listenning
	server, err := net.Listen("tcp", serverAddr)
	utils.CheckError(err)
	defer server.Close()

	recieveFile(server, dstFile)
	sendFile(server, srcFile)
}

func recieveFile(server net.Listener, dstFile string) {
	// accept connection
	conn, err := server.Accept()
	utils.CheckError(err)

	// create new file
	fo, err := os.Create(dstFile)
	utils.CheckError(err)
	defer fo.Close()

	// accept file from client & write to new file
	_, err = io.Copy(fo, conn)
	utils.CheckError(err)
}

func sendFile(server net.Listener, srcFile string) {
	// accept connection
	conn, err := server.Accept()
	utils.CheckError(err)

	// open file to send
	fi, err := os.Open(srcFile)
	utils.CheckError(err)
	defer fi.Close()

	// send file to client
	_, err = io.Copy(conn, fi)
	utils.CheckError(err)
}

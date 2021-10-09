package main

import (
	"fmt"
	"net"
	"os"
)

func availableInterfaces() {

	interfaces, err := net.Interfaces()

	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	fmt.Println("Available network interfaces on this machine : ")
	for _, i := range interfaces {
		fmt.Printf("Name : %v \n", i.Name)
	}
}

func main() {

	//if len(os.Args) != 2 {
	//	fmt.Printf("Usage : %s <interface name>\n", os.Args[0])
	//	os.Exit(0)
	//}

	//ifName := os.Args[1]

	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, v := range interfaces {
		fmt.Println(v)
	}
}
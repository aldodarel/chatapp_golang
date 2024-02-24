package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8088")	// port localhost harus berbeda dengan port yang dimiliki localhost phpmyadmin sendiri
	if err != nil {
		fmt.Println("Gagal terkoneksi ke server:", err)
		return
	}
	defer conn.Close()

	fmt.Print("Masukkan pesan: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		msg := scanner.Text()

		fmt.Fprintln(conn, msg)
		fmt.Println("Pesan terkirim:", msg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error membaca input:", err)
	}
}

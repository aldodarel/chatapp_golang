package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/chatapp_db")	// server harus disamakan dengan server phpmyadmin sendiri
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":8088")	// TCP connection, port harus sama dengan client
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()		// Closing the listener when the main function exits
	fmt.Println("Server mendengarkan pada Client...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn, db)	 // Handling each incoming connection in a separate goroutine
	}
}

func handleConnection(conn net.Conn, db *sql.DB) {
	defer conn.Close()		 // Closing the connection when the handleConnection function exits

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		msg := scanner.Text()		//  Reading a message from the connection

		_, err := db.Exec("INSERT INTO messages (message) VALUES (?)", msg)
		if err != nil {
			log.Println("Gagal menyimpan pesan:", err)		// Logging an error if failed to save the message
			return
		}

		fmt.Printf("Pesan disimpan: %s\n", msg)		// menampilkan pesan yang mengindikasikan bahwa pesan telah berhasil disimpan
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error membaca pesan:", err)	// jika pesan tidak terbaca, maka pesan tersebut akan diprint
	}
}

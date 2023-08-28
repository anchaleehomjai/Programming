package main

import "net"
import "fmt"
import "bufio"
import "os"

func main() {
	// conect to server
	conn, _ := net.Dial("tcp", "127.0.0.1:8000")
	keyboard_buffer := bufio.NewReader(os.Stdin)
	fmt.Print("Input your name: ")
	name,_ := keyboard_buffer.ReadString('\n')
	fmt.Fprintf(conn, name + "\n")
	fmt.Print("Welcome "+name)

	for{
		// create buffer to get keyboard input
		keyboard_buffer := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := keyboard_buffer.ReadString('\n')
		// send to server
		fmt.Fprintf(conn, text + "\n")
		// create buffer and recrive from conn
		server_buffer := bufio.NewReader(conn)
		massage, _ := server_buffer.ReadString('\n')
		fmt.Print("Messeage from server: "+massage)
	}
}
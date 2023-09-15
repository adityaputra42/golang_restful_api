package simple

import "fmt"

type Connection struct {
	File *File
}

func (con *Connection) Close() {
	fmt.Println("Close Connection", con.File.Name)

}

func NewConnection(file *File) (*Connection, func()) {
	connection := &Connection{File: file}
	return connection, func() {
		connection.Close()
	}

}

package main

import (
	"fmt"
	"net"
	"io"
	"strings"
)

type User struct {
	name string
	addr string
	conn net.Conn
	cReadMessage chan string
	offlineCallback func(string, string)
	recvCallback func(string, string, string)
	getAllOnlineUsersCallback func() string
	updateUserNameCallback func(string, string)
}

func NewUser(
	conn net.Conn, 
	offlineCallback func(string, string), 
	recvCallback func(string, string, string),
	getAllOnlineUsersCallback func() string,
	updateUserNameCallback func(string, string),
) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		name: userAddr,
		addr: userAddr,
		conn: conn,
		cReadMessage: make(chan string),
		offlineCallback: offlineCallback,
		recvCallback: recvCallback,
		getAllOnlineUsersCallback: getAllOnlineUsersCallback,
		updateUserNameCallback: updateUserNameCallback,
	}

	return user
}

func (this *User) getName() string {
	return this.name
}

func (this *User) setName(val string) {
	this.name = val
}

func (this *User) getAddr() string {
	return this.addr
}

func (this *User) setAddr(val string) {
	this.addr = val
}

func (this *User) Online() {
	go this.readMessage()
	go this.recv()
}

func (this *User) Write(msg string) {
	this.cReadMessage <- msg
}

func (this *User) readMessage() {
	for {
		this.conn.Write([]byte(<-this.cReadMessage + "\n")) // 写出socket
	}
}

func (this *User) recv() {
	buf := make([]byte, 4096) // 申请读窗口
	for {
		n, err := this.conn.Read(buf) // 读入socket
		if err != nil && err != io.EOF { // 若读入失败, 则返回
			fmt.Println("Conn Read err:", err)
			return
		}
		if n == 0 { // 若释放连接则下线
			this.offlineCallback(this.name, this.addr)
			return
		}
		msg := string(buf[:n - 1])
		if msg == "who" {
			this.Write(this.getAllOnlineUsersCallback())
		} else if len(msg) > 7 && msg[:7] == "rename|" {
			this.updateUserNameCallback(this.name, strings.Split(msg, "|")[1])
		} else {
			this.recvCallback(this.name, this.addr, msg)
		} 
	}
}

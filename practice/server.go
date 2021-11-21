package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	ip   string
	port int
	onlineMap map[string]*User
	mapLock sync.RWMutex
	cBroadCast chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		ip: ip,
		port: port,
		onlineMap: make(map[string]*User),
		cBroadCast: make(chan string),
	}
}

func (this *Server) Serve() {
	go this.broadCast()
	go this.accept()
}

func (this *Server) userOffline(name string, addr string) {
	this.mapLock.Lock()
	delete(this.onlineMap, name)
	this.mapLock.Unlock()
	this.cBroadCast <- "[" + addr + "@" + name + "]:offline"
}

func (this *Server) userRecv(name string, addr string, msg string) {
	this.cBroadCast <- "[" + addr + "@" + name + "]:" + msg
}

func (this *Server) getAllOnlineUsers() string {
	ret := ""
	this.mapLock.Lock()
	for _, user := range this.onlineMap {
		ret += "[" + user.getAddr() + "@" + user.getName() + "]:online...\n"
	}
	this.mapLock.Unlock()
	return ret[:len(ret) - 1]
}

func (this *Server) updateUserName(name string, newName string) {
	this.mapLock.Lock()
	user := this.onlineMap[name]
	_, ok := this.onlineMap[newName]
	if ok {
		user.Write("this name has been used!")
	} else {
		delete(this.onlineMap, name)
		this.onlineMap[newName] = user
		user.setName(newName)
		user.Write("newName has been updated:" + newName)
	}
	this.mapLock.Unlock()
}

func (this *Server) broadCast() {
	for {
		msg := <- this.cBroadCast
		this.mapLock.Lock()
		for _, user := range this.onlineMap {
			user.Write(msg)
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) accept() {
		listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.ip, this.port)) // socket listen
		if err != nil {
			fmt.Println("net.Listen err:", err)
			return
		} else {
			fmt.Printf("server run at http://%s:%d\n", this.ip, this.port)
		}
		for {
			conn, err := listener.Accept() // accept
			if err != nil {
				fmt.Println("listener accept err:", err)
				continue
			}
			go this.connectionHandler(conn) // do handler
		}
		defer listener.Close() // close listen socket
}

func (this *Server) connectionHandler(conn net.Conn) {
	fmt.Println("user connect")
	user := NewUser(conn, this.userOffline, this.userRecv, this.getAllOnlineUsers, this.updateUserName) // 用户上线, 将用户加入到onlineMap中
	user.Online()
	this.mapLock.Lock()
	this.onlineMap[user.getName()] = user
	this.mapLock.Unlock()
	this.cBroadCast <- "[" + user.getAddr() + "@" + user.getName() + "]:online"
}

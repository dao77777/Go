package main

func main() {
	server := NewServer("192.168.0.109", 8888)
	server.Serve()
	for {}
}

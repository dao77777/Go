import (
  "fmt"
  "github.com/aceld/zinx/ziface"
  "github.com/aceld/zinx/znet"
)

type PingRouter struct {
  znet.BaseRouter
}

func (this *PingRouter) {
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServer()

	s.AddRouter(0, &PingRouter{})

	s.Serve()
}

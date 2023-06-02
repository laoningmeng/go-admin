package server

//type Server interface {
//	Init()
//	Start()
//	GetRpcServer() *grpc.Server
//	GetAddr() string
//	GetPort() int
//	SetAddr(addr string)
//	SetPort(p int)
//}
//
//type MyServer struct {
//	Grpc *grpc.Server
//}
//
//func (m *MyServer) Start() {
//
//	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d"))
//	if err != nil {
//		panic(err)
//	}
//	go func() {
//		err = m.Grpc.Serve(listen)
//		if err != nil {
//			panic(err)
//		}
//	}()
//
//	// 接受终止信号
//	quit := make(chan os.Signal)
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	<-quit
//}
//func (m *MyServer) Init() *grpc.Server {
//	m.Grpc = grpc.NewServer()
//	return m.Grpc
//}
//
//
//func NewDefaultServer() Server {
//	s := &MyServer{}
//	return s
//}

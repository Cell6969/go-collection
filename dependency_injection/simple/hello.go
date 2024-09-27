package simple

// create sayhello interface and implementation
type SayHello interface {
	Hello(name string) string
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

// create HelloService from interface
type HelloService struct {
	SayHello SayHello
}

// create constructor SayHelloImpl and Hello Service
func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{SayHello: sayHello}
}

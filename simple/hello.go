package simple

// Contoh penggunaan binding interface, penggunaan data berupa interfacecd

type SayHello interface {
	Hello(name string) string
}

type HelloService struct {
	SayHello SayHello
}

type SayHelloImpl struct {
}

func (s *SayHelloImpl) Hello(name string) string {
	return "Hello " + name
}

func NewHelloService(sayHello SayHello) *HelloService {
	return &HelloService{SayHello: sayHello}
}
func NewSayHelloImpl() *SayHelloImpl {
	return &SayHelloImpl{}
}

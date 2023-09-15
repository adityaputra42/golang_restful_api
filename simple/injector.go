//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializeSimple(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializeDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMySql, NewDatabasePostgreSql, NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

var HelloSet = wire.NewSet(NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)))

func initializedHelloService() *HelloService {
	wire.Build(HelloSet, NewHelloService)
	return nil
}

// injector Salah
// func initializedHelloService() *HelloService {
// 	wire.Build(NewSayHelloImpl, NewHelloService)
// 	return nil
// }

// Contoh injector dengan struct
var FooBarSet = wire.NewSet(
	NewFoo,
	NewBar,
)

func InitializedFooBarStruct() *FooBar {
	wire.Build(FooBarSet, wire.Struct(new(FooBar), "Foo", "Bar"))
	return nil
}

var FooBarValue = wire.NewSet(
	wire.Value(&Foo{}),
	wire.Value(&Bar{}),
)

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(FooBarValue, wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}
func InitializedConfiguration() *Configuration {
	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))
	return nil
}

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}

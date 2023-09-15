//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

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

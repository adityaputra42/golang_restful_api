package simple

type Database struct {
	Name string
}

type DatabaseMySql Database
type DatabasePostgreSql Database

func NewDatabaseMySql() *DatabaseMySql {
	return (*DatabaseMySql)(&Database{Name: "MySql"})

}

func NewDatabasePostgreSql() *DatabasePostgreSql {
	return (*DatabasePostgreSql)(&Database{Name: "MySql"})

}

type DatabaseRepository struct {
	mySql      *DatabaseMySql
	postgreSql *DatabasePostgreSql
}

func NewDatabaseRepository(mySql *DatabaseMySql, postgreSql *DatabasePostgreSql) *DatabaseRepository {
	return &DatabaseRepository{
		mySql: mySql, postgreSql: postgreSql,
	}
}

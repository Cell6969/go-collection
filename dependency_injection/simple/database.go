package simple

// init database
type Database struct {
	Name string
}

// create aliasing
type DatabasePostgreSQL Database
type DatabaseMongoDB Database

func NewDatabaseMongoDB() *DatabaseMongoDB {
	return (*DatabaseMongoDB)(&Database{
		Name: "Mongodb",
	})
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{
		Name: "PostgreSQL",
	})
}

// init database repository
type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(
	databasePostgreSQL *DatabasePostgreSQL,
	databaseMongoDB *DatabaseMongoDB,
) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: databasePostgreSQL, DatabaseMongoDB: databaseMongoDB}
}

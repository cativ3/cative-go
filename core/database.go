package core

func (app *AppCore[T]) ConfigureDatabase(dbName string, mongoURI string) {
	mongoconfig := MongoConfig{}

	app.Database = mongoconfig.Configure(dbName, mongoURI).Connect()
}

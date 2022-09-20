package core

import "go.mongodb.org/mongo-driver/mongo"

type DatabaseType string

const (
	MongoDB   DatabaseType = "MongoDB"
	SqlServer DatabaseType = "SqlServer"
)

type dbtype interface {
	*mongo.Database
}

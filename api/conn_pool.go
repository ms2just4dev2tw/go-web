package api

type DatabaseConnPool interface {
	GetConn()
}

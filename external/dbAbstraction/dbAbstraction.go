package dbAbstraction

// this interface used with any db transaction from whatever the database engine
type IDbTransaction interface {
}

// this interface used with any connection from sql, or  session from nosql
type IDbConnection interface {
}

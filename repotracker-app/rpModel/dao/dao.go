package dao

type dao interface {
	selectQuery() interface{}
	insertQuery()
}

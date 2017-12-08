package base

type User struct {
	Email  string
	Alias  string
	Passwd string
}

type DBResult struct{
	Err error
	Extra interface{}
}
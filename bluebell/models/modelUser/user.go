package modelUser

/*
@author RandySun
@create 2022-01-13-8:48
*/
type User struct {
	UserId   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

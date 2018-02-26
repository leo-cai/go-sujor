package user

import (
	db "sujor.com/leo/sujor-api/database"
	"log"
	"fmt"
	"strconv"
	"time"
)

type User struct {
	Id        int    `json:"id" form:"id"`
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	SignedAt  string `json:"signed_at" form:"signed_at"`
	CreatedAt string `json:"created_at" form:"created_at"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
}

// GetUsers
func (u *User) GetUsers(limit int, page int) (users []User, err error) {
	users = make([]User, 0)
	sql := `SELECT id, username, signed_at, created_at, updated_at FROM t_user WHERE id >= (SELECT id FROM t_user ORDER BY id LIMIT ?, 1) LIMIT ?`
	stmt, err := db.SqlDB.Prepare(sql)
	defer stmt.Close()
	rows, err := stmt.Query((page-1)*limit, limit)
	// 遍历rows
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Username, &user.SignedAt, &user.CreatedAt, &user.UpdatedAt)
		users = append(users, user)
		fmt.Println(user)
	}
	rows.Close()
	return
}

// GetUserById
func (u *User) GetUserById() (user User, err error) {
	sql := `SELECT id, username, signed_at, created_at, updated_at FROM t_user WHERE id = ?`
	stmt, err := db.SqlDB.Prepare(sql)
	defer stmt.Close()

	row := stmt.QueryRow(u.Id)
	err = row.Scan(
		&user.Id, &user.Username, &user.SignedAt, &user.CreatedAt, &user.UpdatedAt,
	)
	return
}

// GetUserByName
func (u *User) GetUserByName() (user User, err error) {
	sql := `SELECT id, username, signed_at, created_at, updated_at FROM t_user WHERE username = ?`
	stmt, err := db.SqlDB.Prepare(sql)
	defer stmt.Close()

	row := stmt.QueryRow(u.Username)
	err = row.Scan(
		&user.Id, &user.Username, &user.SignedAt, &user.CreatedAt, &user.UpdatedAt,
	)
	return
}

// SignUp
func (u *User) SignUp(username string, password string) (user User, err error) {
	// insert into t_user
	sql := `INSERT INTO t_user (username, password) VALUES (?, ?)`
	stmt, err := db.SqlDB.Prepare(sql)
	defer stmt.Close()

	result, err := stmt.Exec(username, password)
	if err != nil {
		log.Println("sql error: insert into t_user table")
		return
	}
	if lastInsertId, err := result.LastInsertId(); err == nil {
		log.Println("LastInsertId: " + strconv.Itoa(int(lastInsertId)))
	}
	if rowsAffected, err := result.RowsAffected(); err == nil {
		log.Println("RowsAffected: " + strconv.Itoa(int(rowsAffected)))
	}

	// RESTFUL接口 返回整个user
	tempUser := User{Username: username}
	user, err = tempUser.GetUserByName()

	// insert into user_permission
	sql = `INSERT INTO user_role (user_id, role_id) VALUES (?, ?)`
	stmt, err = db.SqlDB.Prepare(sql)
	defer stmt.Close()

	// 如果是admin 则是管理员权限
	roleId := 2
	if username == "admin" {
		roleId = 1
	}
	result, err = stmt.Exec(user.Id, roleId)
	if err != nil {
		log.Println("sql error: insert into user_role table")
		return
	}
	if lastInsertId, err := result.LastInsertId(); err == nil {
		log.Println("LastInsertId: " + strconv.Itoa(int(lastInsertId)))
	}
	if rowsAffected, err := result.RowsAffected(); err == nil {
		log.Println("RowsAffected: " + strconv.Itoa(int(rowsAffected)))
	}
	return
}

// SignIn
func (u *User) SignIn(username string, password string) (user User, err error) {
	sql := "SELECT id, username, password, signed_at, created_at, updated_at FROM t_user WHERE username = ? AND password = ?"
	// stmt及错误处理
	stmt, err := db.SqlDB.Prepare(sql)
	defer stmt.Close()

	row := stmt.QueryRow(username, password)
	err = row.Scan(
		&user.Id, &user.Username, &user.Password, &user.SignedAt, &user.CreatedAt, &user.UpdatedAt,
	)
	return
}

// SignOut 更新上次登录时间
func (u *User) SignOut (username string) (err error) {
	sql := "UPDATE t_user SET signed_at = ? WHERE username = ?"
	// stmt及错误处理
	stmt, err := db.SqlDB.Prepare(sql)
	defer stmt.Close()
	result, err := stmt.Exec(time.Now(), username)
	if err != nil {
		log.Println("sql error: update user table")
		return
	}
	if rowsAffected, err := result.RowsAffected(); err == nil {
		log.Println("RowsAffected: " + strconv.Itoa(int(rowsAffected)))
	}
	return
}
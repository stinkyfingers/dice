package user

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stinkyfingers/dice/helpers/database"
	"github.com/stinkyfingers/dice/models/dice"
)

type User struct {
	ID       int
	Email    string
	Password string
	DiceSets dice.DiceSets
}

var (
	createUserStmt = "insert into users (email, password) values (?,?)"
	getUserStmt    = "select id, email, password from users where id = ? "
	userAuthStmt   = "select id, email, password from users where email = ? and password = ?"
	deleteUserStmt = "delete from users where id = ? "
)

func (u *User) CreateUser() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(createUserStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(u.Email, u.Password)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = int(id)
	return err
}

func (u *User) Get() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(getUserStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(u.ID).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return err
	}
	return err
}

func (u *User) Authenticate() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(userAuthStmt)
	if err != nil {
		return err
	}
	h := md5.New()
	h.Write([]byte(u.Password))
	pass := hex.EncodeToString(h.Sum(nil))

	defer stmt.Close()
	err = stmt.QueryRow(u.Email, pass).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return err
	}
	return err
}

func (u *User) Delete() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(deleteUserStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.ID)
	if err != nil {
		return err
	}
	return err
}

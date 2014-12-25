package user

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stinkyfingers/dice/helpers/database"
	"github.com/stinkyfingers/dice/models/dice"
)

type User struct {
	ID       int
	Username string
	Password string
	DiceSets dice.DiceSets
}

var (
	createUserStmt = "insert into users (username, password) values (?,?)"
	getUserStmt    = "select id, username, password from users where id = ? "
	userAuthStmt   = "select id, username, password from users where username = ? and password = ?"
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
	res, err := stmt.Exec(u.Username, u.Password)
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
	err = stmt.QueryRow(u.ID).Scan(&u.ID, &u.Username, &u.Password)
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
	defer stmt.Close()
	err = stmt.QueryRow(u.Username, u.Password).Scan(&u.ID, &u.Username, &u.Password)
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

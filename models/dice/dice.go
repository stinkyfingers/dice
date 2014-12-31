package dice

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stinkyfingers/dice/helpers/database"
	"math/rand"
)

type Die struct {
	ID        int
	DiceSetID int
	Sides     Sides
}

type Dice []Die

type Side struct {
	ID    int
	DieID int
	Value string
}
type Sides []Side

type DiceSet struct {
	ID     int
	Name   string
	Dice   Dice
	UserID int
}

type DiceSets []DiceSet

var (
	getDieStmt = `select d.id, d.diceSet_id, s.id, s.die_id, s.value from dice as d 
		left join dieSides as s on s.die_id = d.id
		where d.id = ?`
	getDiceSetStmt = `select ds.id, ds.name, ds.user_id, d.id, d.diceSet_id
		from diceSets as ds
		left join dice as d on d.diceSet_id = ds.id
		where ds.id = ?`
	getSideStmt       = `select s.id, s.die_id, s.value from dieSides as s where id = ?`
	insertDieStmt     = `insert into dice (diceSet_id) values (?)`
	updateDieStmt     = `update dice set diceSet_id = ? where id = ?`
	insertSideStmt    = `insert into dieSides(die_id, value) values(?,?)`
	insertDiceSetStmt = `insert into diceSets (name, user_id) values (?,?)`
	updateSideStmt    = `update dieSides set die_id = ? and value = ? where id = ?`
	updateDiceSetStmt = `update diceSets set name = ? where id = ?`
	deleteDieStmt     = `delete from dice where id = ?`
	deleteSideStmt    = `delete from dieSides where id = ?`
	deleteDiceSetStmt = `delete from diceSets where id = ?`
)

func (d *Die) Roll() (string, error) {
	var err error
	err = d.Get()
	if err != nil {
		return "", err
	}
	n := rand.Intn(len(d.Sides))
	return d.Sides[n].Value, err
}

func (d *Die) Create() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(insertDieStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(d.DiceSetID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	d.ID = int(id)
	return err
}

func (s *Side) Create() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(insertSideStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(s.DieID, s.Value)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	s.ID = int(id)
	return err
}

func (ds *DiceSet) Create() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(insertDiceSetStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(ds.Name, ds.UserID)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	ds.ID = int(id)
	return err
}
func (d *Die) Get() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(getDieStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Query(d.ID)
	var s Side
	for res.Next() {
		err = res.Scan(&d.ID, &d.DiceSetID, &s.ID, &s.DieID, &s.Value)
		if err != nil {
			return err
		}
		d.Sides = append(d.Sides, s)
	}
	return err
}
func (s *Side) Get() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(getSideStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Query(s.ID)
	for res.Next() {
		err = res.Scan(&s.ID, &s.DieID, &s.Value)
		if err != nil {
			return err
		}
	}
	return err
}
func (ds *DiceSet) Get() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(getDiceSetStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Query(ds.ID)
	if err != nil {
		return err
	}

	var d Die
	for res.Next() {
		err = res.Scan(&ds.ID, &ds.Name, &ds.UserID, &d.ID, &d.DiceSetID)
		if err != nil {
			return err
		}
		err = d.Get()
		if err != nil {
			return err
		}
		ds.Dice = append(ds.Dice, d)
	}
	return err
}

func (s *Side) Update() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(updateSideStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(s.DieID, s.Value, s.ID)
	if err != nil {
		return err
	}
	return err
}

func (d *Die) Update() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(updateDieStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(d.DiceSetID, d.ID)
	if err != nil {
		return err
	}
	return err
}

func (ds *DiceSet) Update() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(updateDiceSetStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(ds.UserID, ds.ID)
	if err != nil {
		return err
	}
	return err
}

func (d *Die) Delete() error {
	var err error
	for _, s := range d.Sides {
		err = s.Delete()
		if err != nil {
			return err
		}
	}
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(deleteDieStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Query(d.ID)
	if err != nil {
		return err
	}
	return err
}

func (s *Side) Delete() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(deleteSideStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Query(s.ID)
	if err != nil {
		return err
	}
	return err
}
func (ds *DiceSet) Delete() error {
	var err error
	for _, d := range ds.Dice {
		err = d.Delete()
		if err != nil {
			return err
		}
	}
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(deleteDiceSetStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Query(ds.ID)
	if err != nil {
		return err
	}
	return err
}

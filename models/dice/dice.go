package dice

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stinkyfingers/dice/helpers/database"
	"math/rand"
)

type Die struct {
	ID        int   `json:"id,omitempty"`
	DiceSetID int   `json:"diceSetId,omitempty"`
	Sides     Sides `json:"sides,omitempty"`
}

type Dice []Die

type Side struct {
	ID    int    `json:"id,omitempty"`
	DieID int    `json:"dieId,omitempty"`
	Value string `json:"value,omitempty"`
}
type Sides []Side

type DiceSet struct {
	ID     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Dice   Dice   `json:"dice,omitempty"`
	UserID int    `json:"userId,omitempty"`
	Public bool   `json:"public,omitempty"`
}

type DiceSets []DiceSet

var (
	getDieStmt = `select d.id, d.diceSet_id, s.id, s.die_id, s.value from dice as d 
		left join dieSides as s on s.die_id = d.id
		where d.id = ?`
	getDiceSetStmt = `select ds.id, ds.name, ds.user_id, ds.public
		from diceSets as ds
		where ds.id = ?`
	getPublicDiceSetsStmt = `select ds.id, ds.name, ds.user_id, ds.public
		from diceSets as ds
		where ds.public = true`
	getUserDiceSetsStmt = `select ds.id, ds.name, ds.user_id, ds.public
		from diceSets as ds
		where ds.user_id = ?`
	getDiceByDiceSetStmt = `select d.id, d.diceSet_id from dice as d 
		where d.diceSet_id = ?`
	getSideStmt       = `select s.id, s.die_id, s.value from dieSides as s where id = ?`
	getSidesByDieStmt = `select s.id, s.die_id, s.value from dieSides as s where s.die_id = ?`
	insertDieStmt     = `insert into dice (diceSet_id) values (?)`
	updateDieStmt     = `update dice set diceSet_id = ? where id = ?`
	insertSideStmt    = `insert into dieSides(die_id, value) values(?,?)`
	insertDiceSetStmt = `insert into diceSets (name, user_id, public) values (?,?,?)`
	updateSideStmt    = `update dieSides set die_id = ? and value = ? where id = ?`
	updateDiceSetStmt = `update diceSets set name = ?, public = ? where id = ?`
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
	res, err := stmt.Exec(ds.Name, ds.UserID, ds.Public)
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
	var sid, did *int
	var val *string
	for res.Next() {
		err = res.Scan(
			&d.ID,
			&d.DiceSetID,
			&sid,
			&did,
			&val,
		)
		if err != nil {
			return err
		}
		if sid != nil {
			s.ID = *sid
		}
		if did != nil {
			s.DieID = *did
		}
		if val != nil {
			s.Value = *val
		}

		d.Sides = append(d.Sides, s)
	}
	return err
}

func (ds *DiceSet) GetDiceByDiceSetID() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(getDiceByDiceSetStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Query(ds.ID)
	if err != nil {
		return err
	}
	for res.Next() {
		var d Die
		err = res.Scan(&d.ID, &d.DiceSetID)
		if err != nil {
			return err
		}
		err = d.GetSidesByDieID()
		if err != nil {
			return err
		}
		ds.Dice = append(ds.Dice, d)
	}
	return err
}

func (d *Die) GetSidesByDieID() error {
	var err error
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare(getSidesByDieStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Query(d.ID)
	if err != nil {
		return err
	}
	for res.Next() {
		var s Side
		err = res.Scan(&s.ID, &s.DieID, &s.Value)
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

	for res.Next() {
		err = res.Scan(&ds.ID, &ds.Name, &ds.UserID, &ds.Public)
		if err != nil {
			return err
		}
		err = ds.GetDiceByDiceSetID()
		if err != nil {
			return err
		}
	}
	return err
}

func GetUserDiceSets(userId int) ([]DiceSet, error) {
	var err error
	var dss []DiceSet
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return dss, err
	}
	defer db.Close()

	stmt, err := db.Prepare(getUserDiceSetsStmt)
	if err != nil {
		return dss, err
	}
	defer stmt.Close()
	res, err := stmt.Query(userId)
	if err != nil {
		return dss, err
	}

	var ds DiceSet
	for res.Next() {
		err = res.Scan(&ds.ID, &ds.Name, &ds.UserID, &ds.Public)
		if err != nil {
			return dss, err
		}
		err = ds.GetDiceByDiceSetID()
		if err != nil {
			return dss, err
		}
		dss = append(dss, ds)
	}
	return dss, err
}

func GetPublicDiceSets() ([]DiceSet, error) {
	var err error
	var dss []DiceSet
	db, err := sql.Open("mysql", database.ConnectionString())
	if err != nil {
		return dss, err
	}
	defer db.Close()

	stmt, err := db.Prepare(getPublicDiceSetsStmt)
	if err != nil {
		return dss, err
	}
	defer stmt.Close()
	res, err := stmt.Query()
	if err != nil {
		return dss, err
	}

	var ds DiceSet
	for res.Next() {
		err = res.Scan(&ds.ID, &ds.Name, &ds.UserID, &ds.Public)
		if err != nil {
			return dss, err
		}
		err = ds.GetDiceByDiceSetID()
		if err != nil {
			return dss, err
		}
		dss = append(dss, ds)
	}
	return dss, err
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
	_, err = stmt.Exec(ds.Name, ds.Public, ds.ID)
	if err != nil {
		return err
	}
	return err
}

func (d *Die) Delete() error {
	var err error
	//delete Sides
	for _, s := range d.Sides {
		err = s.Delete()
		if err != nil {
			return err
		}
	}

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
	//delete Dice
	for _, d := range ds.Dice {
		err = d.Delete()
		if err != nil {
			return err
		}
	}

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

package dice

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stinkyfingers/dice/helpers/database"
)

type Die struct {
	ID    int
	Value string
}

type Dice []Die

type DiceSet struct {
	ID   int
	Name string
	Dice Dice
}

type DiceSets []DiceSet

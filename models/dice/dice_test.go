package dice

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDice(t *testing.T) {
	Convey("Test CRUD Die", t, func() {
		var err error
		var d1, d2 Die
		var ds DiceSet
		var s1, s2, s3, s4, s5, s6 Side

		//create
		ds.Name = "two threes"
		ds.Public = true
		err = ds.Create()
		So(err, ShouldBeNil)

		d1.DiceSetID = ds.ID
		err = d1.Create()
		So(err, ShouldBeNil)
		d2.DiceSetID = ds.ID
		err = d2.Create()
		So(err, ShouldBeNil)

		s1.Value = "1"
		s1.DieID = d1.ID
		err = s1.Create()
		So(err, ShouldBeNil)
		s2.Value = "2"
		s2.DieID = d1.ID
		err = s2.Create()
		So(err, ShouldBeNil)
		s3.Value = "3"
		s3.DieID = d1.ID
		err = s3.Create()
		So(err, ShouldBeNil)
		s4.Value = "1"
		s4.DieID = d2.ID
		err = s4.Create()
		So(err, ShouldBeNil)
		s5.Value = "2"
		s5.DieID = d2.ID
		err = s5.Create()
		So(err, ShouldBeNil)
		s6.Value = "3"
		s6.DieID = d2.ID
		err = s6.Create()
		So(err, ShouldBeNil)

		//get
		err = d1.Get()
		So(err, ShouldBeNil)
		d1.Sides = nil

		err = ds.Get()
		So(err, ShouldBeNil)
		So(len(ds.Dice), ShouldEqual, 2)

		err = d1.GetSidesByDieID()
		So(len(d1.Sides), ShouldEqual, 3)
		So(err, ShouldBeNil)
		err = d2.GetSidesByDieID()
		So(len(d2.Sides), ShouldEqual, 3)
		So(err, ShouldBeNil)

		ds.Dice = nil
		err = ds.GetDiceByDiceSetID()
		So(len(ds.Dice), ShouldEqual, 2)
		So(err, ShouldBeNil)

		val, err := d1.Roll()
		So(err, ShouldBeNil)
		So(val, ShouldHaveSameTypeAs, "1")

		//update
		ds.Name = "test set new"
		err = ds.Update()
		So(err, ShouldBeNil)

		d1.DiceSetID = ds.ID
		err = d1.Update()
		So(err, ShouldBeNil)

		s6.Value = "4"
		err = s6.Update()
		So(err, ShouldBeNil)

		//delete
		err = d1.Delete()
		So(err, ShouldBeNil)

		err = s4.Delete()
		So(err, ShouldBeNil)

		err = ds.Delete()
		So(err, ShouldBeNil)
	})
}

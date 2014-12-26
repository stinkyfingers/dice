package dice

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDice(t *testing.T) {
	Convey("Test CRUD Die", t, func() {
		var err error
		var d Die
		var ds DiceSet
		var s Side

		//create
		ds.Name = "test set"
		err = ds.Create()
		So(err, ShouldBeNil)

		d.DiceSetID = ds.ID
		err = d.Create()
		So(err, ShouldBeNil)

		s.Value = "1"
		s.DieID = d.ID
		err = s.Create()
		So(err, ShouldBeNil)

		//get
		err = d.Get()
		So(err, ShouldBeNil)

		err = ds.Get()
		So(err, ShouldBeNil)
		So(len(ds.Dice), ShouldBeGreaterThan, 0)

		val, err := d.Roll()
		So(err, ShouldBeNil)
		So(val, ShouldEqual, "1")

		//update
		ds.Name = "test set new"
		err = ds.Update()
		So(err, ShouldBeNil)

		s.Value = "2"
		err = s.Update()
		So(err, ShouldBeNil)

		//delete
		err = d.Delete()
		So(err, ShouldBeNil)

		err = s.Delete()
		So(err, ShouldBeNil)

		err = ds.Delete()
		So(err, ShouldBeNil)
	})
}

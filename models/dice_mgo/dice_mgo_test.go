package dice_mgo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDiceMgo(t *testing.T) {
	Convey("Test CRUD Side", t, func() {
		var err error
		var d Die
		var ds DiceSet
		var s Side

		//create
		ds.Name = "test"
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
		err = s.Get()
		So(err, ShouldBeNil)

		err = d.Get()
		So(err, ShouldBeNil)

		err = ds.Get()
		So(err, ShouldBeNil)

		s.Value = "2"
		err = s.Update()
		So(err, ShouldBeNil)

		err = d.Update()
		So(err, ShouldBeNil)

		ds.Name = "testUpdate"
		err = ds.Update()
		So(err, ShouldBeNil)

		//delete
		err = s.Delete()
		So(err, ShouldBeNil)

		err = d.Delete()
		So(err, ShouldBeNil)

		err = ds.Delete()
		So(err, ShouldBeNil)

	})
}

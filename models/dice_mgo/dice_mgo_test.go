package dice_mgo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestDiceMgo(t *testing.T) {
	Convey("Test CRUD Side", t, func() {
		var err error
		// var d1, d2 Die
		// var ds DiceSet
		var s1 Side

		//create
		s1.Value = "1"
		s1.DieID = 1
		err = s1.Create()
		So(err, ShouldBeNil)

	})
}

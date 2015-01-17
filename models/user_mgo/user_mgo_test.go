package user_mgo

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUsers(t *testing.T) {
	Convey("Test CRUD Users", t, func() {
		var err error
		var u User

		//create
		u.Email = "test"
		u.Password = "tester"
		err = u.Create()
		So(err, ShouldBeNil)

		u.Password = "tester"
		//auth
		err = u.Authenticate()
		So(err, ShouldBeNil)

		//get
		err = u.Get()
		So(err, ShouldBeNil)

		//delete
		err = u.Delete()
		So(err, ShouldBeNil)
	})
}

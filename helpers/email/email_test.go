package email

import (
	"bytes"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"text/template"
)

func TestEmail(t *testing.T) {
	var err error
	Convey("Test Email", t, func() {
		var doc bytes.Buffer
		t := template.Must(template.New("email").Parse("HEY BUDDY"))
		t.Execute(&doc, nil)
		err = Send([]string{"john_shenk@hotmail.com"}, "Test", doc)
		So(err, ShouldBeNil)

	})
}

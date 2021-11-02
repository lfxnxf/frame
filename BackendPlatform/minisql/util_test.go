package minisql

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinisqlStart(t *testing.T) {
	Convey("Start", t, func() {
		var (
			command = "sleep"
			args    = []string{"100"}
		)
		Convey("When everything goes positive", func() {
			e, err := Start(command, args)
			Convey("Then err should be nil.e should not be nil.", func() {
				So(err, ShouldBeNil)
				So(e, ShouldNotBeNil)
			})
			e.Stop()
		})
	})
}
func TestMinisqlExec(t *testing.T) {
	Convey("Exec", t, func() {
		var (
			command = "sleep"
			args    = []string{"1"}
		)
		Convey("When everything goes positive", func() {
			e, err := Exec(command, args)
			Convey("Then err should be nil.e should not be nil.", func() {
				So(err, ShouldBeNil)
				So(e, ShouldNotBeNil)
			})
		})
	})
}

package minisql

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinisqlRun(t *testing.T) {
	Convey("Run", t, func() {
		m := New()
		m.Start()
		t.Log(m.Port)
		time.Sleep(time.Second * 2)
		m.Stop()
	})
}

func TestMinisqlgetBin(t *testing.T) {
	Convey("getBin", t, func() {
		Convey("When everything goes positive", func() {
			p1 := getBin()
			t.Log(p1)
			Convey("Then p1 should not be nil.", func() {
				So(p1, ShouldNotBeNil)
			})
		})
	})
}

package renderer

import (
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFindStringWriter(t *testing.T) {
	Convey("When writing to a FindStringWriter", t, func() {
		Convey("Writes all matches to the channel", func() {
			channel := make(chan string, 100)

			w := NewFindStringWriter(channel, regexp.MustCompile("a"), nil)
			n, err := w.Write([]byte("a a a"))

			So(n, ShouldEqual, 5)
			So(err, ShouldBeNil)

			m, ok := <-channel
			So(m, ShouldEqual, "a")
			So(ok, ShouldBeTrue)

			m, ok = <-channel
			So(m, ShouldEqual, "a")
			So(ok, ShouldBeTrue)

			m, ok = <-channel
			So(m, ShouldEqual, "a")
			So(ok, ShouldBeTrue)

			m, ok = <-channel
			So(ok, ShouldBeFalse)
		})
	})
}

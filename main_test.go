package main

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("Given a github integration", t, func() {
		So(os.Getenv("DEV_TEST"), ShouldNotBeNil)
		So(os.Getenv("DEV_TEST"), ShouldNotEqual, "")
	})
}

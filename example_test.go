package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestShowProjectName test show project name
func TestShowProjectName(t *testing.T) {
	Convey("test add", t, func() {
		a := 3
		b := 1
		c := Add(a, b)
		So(c, ShouldEqual, 4)
	})
}

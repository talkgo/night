package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

// TestBecomeAContributor test become a contributor
func TestBecomeAContributor(t *testing.T) {
	Convey("Test BecomeAContributor", t, func() {
		c := BecomeAContributor("yangwenmai")
		So(c, ShouldEqual, true)
	})
}

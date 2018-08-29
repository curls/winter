package container

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestContainer(t *testing.T) {

	Convey("Given a new container", t, func() {

		key := "first"

		Convey("When initially created", func() {

			Convey("The container should has zero element", func() {
				So(container.m, ShouldBeEmpty)
			})
		})

		Convey("When getting an element that doesn't exist", func() {

			_, err := container.Get(key)
			Convey("We will get an error", func() {
				So(err, ShouldBeError, fmt.Errorf("element '%s' does not exist", key))
			})
		})

		Convey("When add the first element", func() {

			container.Add("first", &testStruct{a: 1})
			Convey("The container should has one element", func() {
				So(len(container.m), ShouldEqual, 1)
			})
		})

		Convey("When add the second element", func() {

			container.Add("second", "a second")
			Convey("The container should has two element", func() {
				So(len(container.m), ShouldEqual, 2)
			})
		})

		Convey("When get a exist element and right type assertion", func() {

			item, _ := container.Get(key)
			value := item.(*testStruct)
			Convey("We can get the right element", func() {
				So(value.a, ShouldEqual, 1)
			})
		})

		Convey("When get a exist element and wrong type assertion", func() {

			Convey("We can get a panic", func() {
				So(func() { test() }, ShouldPanic)
			})
		})

	})
}

type testStruct struct {
	a int
}

func test() testStruct {
	item, _ := Resolve("first")
	return item.(testStruct)
}

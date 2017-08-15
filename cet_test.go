package cet_test

import (
	"testing"

	"github.com/LyricTian/cet"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuery(t *testing.T) {
	Convey("cet query not found test", t, func() {
		querier := cet.NewQuerier()

		result, err := querier.Query(nil, "370150162100108", "张三")
		So(err, ShouldEqual, cet.ErrNotFound)
		So(result, ShouldBeNil)
	})
}

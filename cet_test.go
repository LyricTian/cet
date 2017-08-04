package cet_test

import (
	"testing"

	"github.com/LyricTian/cet"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQuery(t *testing.T) {
	Convey("cet query test", t, func() {
		querier := cet.NewQuerier()

		result, err := querier.Query(nil, "370150162100708", "张新伟")
		So(err, ShouldBeNil)
		So(result, ShouldNotBeNil)
		So(result.University, ShouldEqual, "山东体育学院济南校区")
		So(result.Score, ShouldEqual, 403)
	})

	Convey("cet query not found test", t, func() {
		querier := cet.NewQuerier()

		result, err := querier.Query(nil, "370150162100718", "张伟")
		So(err, ShouldEqual, cet.ErrNotFound)
		So(result, ShouldBeNil)
	})
}

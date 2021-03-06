package connor

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestIn(t *testing.T) {
	Convey("$in", t, func() {
		cases := []struct {
			conds string
			data  string

			match  bool
			hasErr bool
		}{
			{
				`{ "x": { "$in": [1] } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "x": { "$in": [2] } }`,
				`{ "x": 1, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "x": { "$in": [1, 2, 3] } }`,
				`{ "x": 1, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "a.x": { "$in": [1] } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				true,
				false,
			},
			{
				`{ "a": { "$in": [1] } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				false,
				false,
			},
			{
				`{ "a": { "$in": 1 } }`,
				`{ "a": { "x": 1 }, "y": 2 }`,
				false,
				true,
			},
		}

		for _, c := range cases {
			Convey(fmt.Sprintf("%s & %s", c.data, c.conds), func() {
				conds := prepConds(c.conds)
				data := prepData(c.data)

				m, err := Match(conds, data)
				if c.hasErr {
					So(err, ShouldNotBeNil)
				} else {
					So(err, ShouldBeNil)
				}

				So(m, ShouldEqual, c.match)
			})
		}
	})
}

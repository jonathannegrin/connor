package connor

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGreaterEqual(t *testing.T) {
	now := time.Now()

	Convey("$ge", t, func() {
		Convey("Basic Cases", func() {
			cases := []struct {
				conds string
				data  string

				match  bool
				hasErr bool
			}{
				{
					`{ "x": { "$ge": 0 } }`,
					`{ "x": 1, "y": 2 }`,
					true,
					false,
				},
				{
					`{ "x": { "$ge": 1 } }`,
					`{ "x": 1, "y": 2 }`,
					true,
					false,
				},
				{
					`{ "x": { "$ge": 2 } }`,
					`{ "x": 1, "y": 2 }`,
					false,
					false,
				},
				{
					`{ "x": { "$ge": 0.5 } }`,
					`{ "x": 1, "y": 2 }`,
					true,
					false,
				},
				{
					`{ "a.x": { "$ge": 0 } }`,
					`{ "a": { "x": 1 }, "y": 2 }`,
					true,
					false,
				},
				{
					`{ "a": { "$ge": 2 } }`,
					`{ "a": { "x": 1 }, "y": 2 }`,
					false,
					false,
				},
				{
					`{ "x": { "$ge": "0" } }`,
					`{ "x": 1, "y": 2 }`,
					false,
					false,
				},
				{
					`{ "x": { "$ge": 0 } }`,
					`{ "x": "1", "y": 2 }`,
					false,
					false,
				},
				{
					`{ "x": { "$ge": "a" } }`,
					`{ "x": "b", "y": 2 }`,
					true,
					false,
				},
				{
					`{ "x": { "$ge": "c" } }`,
					`{ "x": "b", "y": 2 }`,
					false,
					false,
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

		Convey("Different Types", func() {
			cases := []struct {
				con  interface{}
				data interface{}

				match  bool
				hasErr bool
			}{
				{
					"abc", "def",
					true, false,
				},
				{
					"abc", "abc",
					true, false,
				},
				{
					"abc", "aaa",
					false, false,
				},
				{
					"test", 1,
					false, false,
				},
				{
					int8(10), 10,
					true, false,
				},
				{
					int16(10), 10,
					true, false,
				},
				{
					int32(10), 10,
					true, false,
				},
				{
					int64(10), 10,
					true, false,
				},
				{
					float32(10), 11,
					true, false,
				},
				{
					int64(10), float32(10),
					true, false,
				},
				{
					int64(10), "test",
					false, false,
				},
				{
					now, now,
					true, false,
				},
				{
					now, now.Add(time.Second),
					true, false,
				},
				{
					now, now.Add(-time.Second),
					false, false,
				},
				{
					now, 10,
					false, false,
				},
				{
					[]int{10}, []int{12},
					false, true,
				},
			}

			for _, c := range cases {
				conds := c.con
				data := c.data
				match := c.match
				hasErr := c.hasErr

				Convey(fmt.Sprintf("%T(%v) == %T(%v)", c.con, c.con, c.data, c.data), func() {
					m, err := Match(map[string]interface{}{
						"x": map[string]interface{}{"$ge": conds},
					}, map[string]interface{}{
						"x": data,
					})

					if hasErr {
						So(err, ShouldNotBeNil)
					} else {
						So(err, ShouldBeNil)
					}

					So(m, ShouldEqual, match)
				})
			}
		})
	})
}

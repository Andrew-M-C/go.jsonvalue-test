package usedby

import (
	"context"
	"io/ioutil"
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
)

var (
	cv = convey.Convey
	so = convey.So

	eq = convey.ShouldEqual

	isNil = convey.ShouldBeNil
)

func TestUsedBy(t *testing.T) {
	cv("测试 GetRepoCount", t, func() { testGetRepoCount(t) })
	cv("测试 GetRepoCountBadgeURL", t, func() { testGetRepoCountBadgeURL(t) })
	cv("测试 GetRepoCountBadgeImage", t, func() { testGetRepoCountBadgeImage(t) })
}

func testGetRepoCount(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cnt, err := GetRepoCount(ctx)
	so(err, isNil)
	t.Log("used by:", cnt)
}

func testGetRepoCountBadgeURL(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	badgeURL, err := GetRepoCountBadgeURL(ctx)
	so(err, isNil)
	t.Logf("badge URL: %v", badgeURL)
}

func testGetRepoCountBadgeImage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	data, ext, err := GetRepoCountBadgeImage(ctx)
	so(err, isNil)
	so(ext, eq, ".svg")

	err = ioutil.WriteFile("./.badge.svg", data, 0644)
	so(err, isNil)
}

package usedby

import (
	"path"
	"regexp"
	"sync/atomic"
	"time"
)

const (
	cacheTimeout = 30 * time.Second
)

var internal = struct {
	repoCount struct {
		cache atomic.Value
		regex regexp.Regexp
	}
	repoBadge struct {
		cache atomic.Value
	}
}{}

func init() {
	internal.repoCount.regex = *regexp.MustCompile(`\<\/svg\>[ \r\n]+(\d+)[ \r\n]+Repositories[ \r\n]+\<\/a\>`)
}

type baseCache struct {
	time time.Time
}

func (c *baseCache) isTimeout() bool {
	return time.Now().After(c.time.Add(cacheTimeout))
}

func (c *baseCache) setNowTime() {
	c.time = time.Now()
}

type countCache struct {
	baseCache
	count int
}

func newCountCache(count int) *countCache {
	c := &countCache{
		count: count,
	}
	c.setNowTime()
	return c
}

type badgeCache struct {
	baseCache
	ext  string
	data []byte
}

func newbadgeCache(data []byte, url string) *badgeCache {
	c := &badgeCache{
		ext:  path.Ext(url),
		data: data,
	}
	c.setNowTime()
	return c
}

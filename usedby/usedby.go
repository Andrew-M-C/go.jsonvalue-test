package usedby

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// GetRepoCountBadgeImage 获取 used by 徽章的图片文件
func GetRepoCountBadgeImage(ctx context.Context) (data []byte, ext string, err error) {
	if v := internal.repoBadge.cache.Load(); v != nil {
		if cache, ok := v.(*badgeCache); ok && cache != nil {
			if !cache.isTimeout() {
				return cache.data, cache.ext, nil
			}
		}
	}

	url, err := GetRepoCountBadgeURL(ctx)
	if err != nil {
		return
	}

	data, err = getHttpByURL(ctx, url)
	if err != nil {
		err = fmt.Errorf("get badge by http error: %w", err)
		return
	}

	c := newbadgeCache(data, url)
	internal.repoBadge.cache.Store(c)
	return data, c.ext, nil
}

// GetRepoCountBadgeURL 获取 used by 徽章的 URL 地址
func GetRepoCountBadgeURL(ctx context.Context) (string, error) {
	cnt, err := GetRepoCount(ctx)
	if err != nil {
		return "", nil
	}

	repoStr := "repo"
	if cnt > 1 {
		repoStr = "repos"
	}
	return fmt.Sprintf("https://img.shields.io/badge/used%%20by-%d%%20%s-blue.svg", cnt, repoStr), nil
}

// GetRepoCount 获取 used by 的 repo 数量
func GetRepoCount(ctx context.Context) (int, error) {
	if v := internal.repoCount.cache.Load(); v != nil {
		if cache, ok := v.(*countCache); ok && cache != nil {
			if !cache.isTimeout() {
				return cache.count, nil
			}
		}
	}

	// HTTP 请求
	body, err := getHttpByURL(ctx, "https://github.com/Andrew-M-C/go.jsonvalue/network/dependents?dependent_type=REPOSITORY")
	if err != nil {
		return 0, err
	}

	// 解析响应
	// fmt.Println(body)
	match := internal.repoCount.regex.FindAllStringSubmatch(string(body), -1)

	if len(match) == 0 || len(match[0]) < 2 {
		return 0, errors.New("cannot find repo count text")
	}

	s := match[0][1]
	cnt, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("parse string '%s' error: %w", s, err)
	}

	c := newCountCache(int(cnt))
	internal.repoCount.cache.Store(c)
	return int(cnt), nil
}

func getHttpByURL(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("initialize new request error: %w", err)
	}

	cli := http.DefaultClient
	resp, err := cli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read from HTTP response error: %w", err)
	}

	if resp.StatusCode != 200 {
		return body, fmt.Errorf("HTTP code error: %v", resp.StatusCode)
	}

	return body, nil
}

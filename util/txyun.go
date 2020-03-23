package util

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

func MakeTXYunUrl(urlRaw, key, us string, expireTime time.Time, rlimit int, exper int) (string, error) {
	u, err := url.Parse(urlRaw)
	if err != nil {
		return "", err
	}
	pos := strings.LastIndex(u.Path, "/")
	dir := u.Path[0 : pos+1]

	var t = expireTime.Unix()
	var tHexStr = fmt.Sprintf("%x", t)
	var rLimitStr = fmt.Sprintf("%d", rlimit)
	var experStr string
	if exper > 0 {
		experStr = fmt.Sprintf("%d", exper)
	}
	var signContent = key + dir + tHexStr + experStr + rLimitStr + us
	var sign = MD5(signContent)

	params := make(map[string]string)
	params["t"] = tHexStr
	if experStr != "" {
		params["exper"] = experStr
	}
	params["rlimit"] = rLimitStr
	params["us"] = us
	params["sign"] = sign

	var rawQuery string
	for k, v := range params {
		rawQuery += fmt.Sprintf("%s=%s&", k, v)
	}
	rawQuery = rawQuery[0 : len(rawQuery)-1]
	u.RawQuery = rawQuery
	return u.String(), nil
}

package util

import (
	"fmt"
	"testing"
	"time"
)

func Test_Txyn(t *testing.T) {
	urlRaw := "http://1253993529.vod2.myqcloud.com/7fe6eefavodtranscq1253993529/c2276bbf5285890795883087086/v.f20.mp4"
	key := "295mDpMMqrtgFQ9xF0Vz"
	us := "ewrewr"
	expireTime := time.Unix(1578389644, 0)
	rLimit := 10
	expire := 0
	url, err := MakeTXYunUrl(urlRaw, key, us, expireTime, rLimit, expire)
	if err != nil {
		t.Fail()
	}

	fmt.Println(url)
	//if url != "http://1253993529.vod2.myqcloud.com/7fe6eefavodtranscq1253993529/c2276bbf5285890795883087086/v.f20.mp4?t=5e129b26&rlimit=10&us=ewrewr&sign=855d5c40f0e78e10530c6ca76c3b391f" {
	//	t.Fail()
	//}
}

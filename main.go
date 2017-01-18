package main

import (
	"fmt"
	"regexp"

	"github.com/zouyu/resmon/parser"
)

func main() {
	p, _ := parser.CreateParser("(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S).*", 0, "v1/image/get")

	test := "2017-01-18 00:02:31 141 220.180.119.136 TCP_MISS/200 58632 GET https://res.shiqichuban.com/v1/image/get/CcMFqDhDZk3_Pwph4XRmbITlpwsYqhDAEocakXrsP6bcQWi2K-9uI8K9Ga1zmXLSmSE9gkIJ2_znlbdwlL-ENQ/m - DIRECT/kv1-pub.oss-cn-beijing-internal.aliyuncs.com image/jpeg Mozilla/5.0 (Linux; Android 4.4.4; vivo X5Max+ Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/33.0.0.0 Mobile Safari/537.36/vivo/bbk6752_lwt_kk/vivo X5Max+/19/android/com.shiqichuban.android/2.2.3/com.shiqichuban.client"
	a := p.Match(test)
	for i := 1; i < len(a); i++ {
		fmt.Printf("%s\n", a[i])
	}

	matches := make(map[string]map[string]int)
	get := "^https://res.shiqichuban.com/v1/image/get"
	reg := regexp.MustCompile(get)
	ok := reg.MatchString(a[8])
	if ok {
		fmt.Printf("matched.\n")
		e := make(map[string]int)
		e[a[1]] = 1
		matches["get"] = e
	}
}

package wap

import "fmt"

func (z *Zhengzaitv) GetAdamHeaders() map[string]string {
	headers := z.GetHeaders("adam.zhengzai.tv")
	return headers
}

func (z *Zhengzaitv) GetKylinHeaders() map[string]string {
	headers := z.GetHeaders("kylin.zhengzai.tv")
	return headers
}

func (z *Zhengzaitv) GetOrderHeaders() map[string]string {
	headers := z.GetHeaders("order.zhengzai.tv")
	return headers
}

func (z *Zhengzaitv) GetSmsLoginHeaders() map[string]string {
	headers := z.GetAdamHeaders()
	headers["version"] = "1.0.0"
	return headers
}

func (z *Zhengzaitv) GetHeaders(host string) map[string]string {
	headers := map[string]string{
		"Host":            host,
		"Connection":      "keep-alive",
		"Pragma":          "no-cache",
		"Cache-Control":   "no-cache",
		"Accept":          "application/json, text/plain, */*",
		"Content-Type":    "application/json;charset=UTF-8",
		"Authorization":   fmt.Sprintf("Bearer %v", z.Cookies),
		"source":          "H5",
		"User-Agent":      "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1",
		"version":         "1.1",
		"Origin":          "https://m.zhengzai.tv",
		"Sec-Fetch-Site":  "same-site",
		"Sec-Fetch-Mode":  "cors",
		"Sec-Fetch-Dest":  "empty",
		"Referer":         "https://m.zhengzai.tv/",
		"Accept-Encoding": "gzip, deflate, br",
		"Accept-Language": "zh-CN,zh;q=0.9",
	}

	return headers
}

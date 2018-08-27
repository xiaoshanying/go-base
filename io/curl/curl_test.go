package curl

import "testing"

const url = "http://www.baidu.com"

const destFile = "resp.txt"

func TestCurl(t *testing.T) {
	t.Log("start test curl")
	{
		code := Curl(url, destFile)
		if code == 0 {
			t.Log("operate success")
		} else {
			t.Logf("operate fail,fail reason %d", code)
		}
	}
}

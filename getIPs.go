package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// 何柳：akamai    aws    cloudflare     阿里
// 晓龙：Fastly    Level3     Limelight    Edgecast   cdnnetworks
type CNDIPs struct {
	Data []struct {
		IPAddress string `json:"IPAddress"`
		IP        string `json:"ip"`
	} `json:"data"`
}

func main() {
	cndkeyMap := map[string]string{
		"akamai":     "tKi7VvL48umjo3OK484sFBKjSx4RoD7i",
		"aws":        "iJ0hQtkd%2Bn%2FQKgyNLpwMQt8RxVB3%2B%2F3L",
		"cloudflare": "FvbMO1qsL7hvlUVTWp6AF9J6ARm37Tw%2F",
		"阿里":         "cltuPggSOiEPEZ73EuV1qQsB9U0QLbjP",
		// "Fastly":      "",
		// "Level3":      "",
		// "Limelight":   "",
		// "Edgecast":    "",
		// "cdnnetworks": "",
	}
	var cdnips CNDIPs
	client := http.Client{}
	for _, cndkey := range cndkeyMap {
		for i := 1; i < 21; i++ {
			url := "https://cdn.chinaz.com/ajax/AreaHwIP?cdnkey=" + cndkey + "&" + "pageindex=" + strconv.Itoa(i) + "&cnt=0"
			req, err := http.NewRequest(http.MethodGet, url, nil)
			if err != nil {
				fmt.Println("request err", err)
				continue
			}
			req.Header.Add("Cookie", "qHistory=aHR0cDovL2lwLnRvb2wuY2hpbmF6LmNvbV9JUC9JUHY25p+l6K+i77yM5pyN5Yqh5Zmo5Zyw5Z2A5p+l6K+i;UM_distinctid=170776d8326166-0261949a7dd434-313f68-100200-170776d8327317;bbsmax_user=79f3fa6e-02b9-4f31-b832-2ceaa46d97a8;avatarId=17af0134-e1f8-48f6-be5c-f81c4586fe68-; ASP.NET_SessionId=52nix41zzh0ak1okw2szrwa3; CNZZDATA5933533=cnzz_eid%3D1766445022-1582603946-%26ntime%3D1582616203")
			rsp, err := client.Do(req)
			if err != nil {
				fmt.Println("rsp err")
			}
			body, err := ioutil.ReadAll(rsp.Body)
			if err != nil {
				fmt.Println("read io err")
			}
			json.Unmarshal(body, &cdnips)
			fmt.Println("page:", i)
			for _, v := range cdnips.Data {
				fmt.Println(v.IP)
			}
			time.Sleep(1000)
		}

	}
}

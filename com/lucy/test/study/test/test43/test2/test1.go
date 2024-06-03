package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Line struct {
	Pct float64 `json:"pct"`
}
type Total struct {
	Line Line `json:"line"`
}
type Diff struct {
	Line Line `json:"line"`
}
type Cov struct {
	Total      Total  `json:"total"`
	Diff       Diff   `json:"diff"`
	Status     int    `json:"status"`
	Mid        string `json:"mid"`
	MergeLevel int    `json:"mergeLevel"`
}
type Body struct {
	Data []*Cov `json:"data"`
}
type Res struct {
	DiffPct  float64 `json:"diffPct"`
	TotalPct float64 `json:"totalPct"`
	Source   string  `json:"source"`
}

func main() {
	//m := make(map[string]Res)
	for {
		// 发送 HTTP GET 请求
		resp, err := http.Get("https://tcoverage.woa.com/openapi/v2/sin" +
			"gle/coverage?token=495d640d-8543-4ede-a577-15fa0e995e0c&ogid=p-310efe0d7f1745caa32e3ae86dd9b2b1.4981.0")
		if err != nil {
			fmt.Println("HTTP GET 请求失败:", err)
			return
		}
		defer resp.Body.Close()

		// 读取响应体并解析 JSON 数据
		var responseData Body
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
			fmt.Println("解析 JSON 失败:", err)
			//return
		}
		//marshal, err := json.Marshal(responseData)
		//if err != nil {
		//	return
		//}
		//str := string(marshal)
		flag := true
		if responseData.Data != nil && len(responseData.Data) != 0 {
			datas := responseData.Data
			for _, v := range datas {

				fmt.Printf("mid=%v,totalPct=%v,diffPct=%v,mergeType=%v,status=%v\n", v.Mid, v.Total.Line.Pct, v.Diff.Line.Pct, v.MergeLevel,
					v.Status)
				if !(v.Status == 401 || v.Status == 1001) {
					flag = false
				}
			}
			fmt.Println("======================================")
			if flag == true && len(datas) == 3 {
				break
			}
		}
	}
}

package tips

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func GetTip(jsonData map[string]string, key string) string {
	var result string
	for index, value := range jsonData {
		if strings.ContainsAny(key, value) {
			result = jsonData[index]
		}
	}
	return result
}

type jsonStruct struct {
	Title string `json:"title"`
	Tip   string `json:"tip"`
}

func LoadTipsFromJson() string {
	var result string
	data, err := ioutil.ReadFile("JsonFile/tips.json")
	if err != nil {
		fmt.Print(err)
	}

	var configs []jsonStruct
	err = json.Unmarshal(data, &configs)
	if err != nil {
		fmt.Println("error:", err)
	}
	result = string(data)

	//for _, v := range configs {
	//fmt.Println(v.Title, "\n", v.Tip)
	//}
	return result
}

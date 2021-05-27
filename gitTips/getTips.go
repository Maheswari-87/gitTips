package gitTips

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

func GetTip(jsonData map[string]string, key string) string {
	var result string
	for _, value := range jsonData {
		if strings.ContainsAny(key, value) {
			result = jsonData["Tip"]
			return result
		}
	}
	return ""
}

func LoadTipsFromJsonFile() string {
	var result string
	data, err := ioutil.ReadFile("JsonFile/tips.json")
	if err != nil {
		fmt.Print(err)
	}

	var configs []GitTips
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

type GitTips struct {
	Title string `json:"title"`
	Tip   string `json:"tip"`
}

//reading json file
func ReadJsonFile() ([]byte, error) {
	jsonData, err := ioutil.ReadFile("C:\\Users\\SRS\\repos\\githubProjects\\tips\\jsonFile\\tips.json")
	if err != nil {
		return []byte{}, errors.New("file issue")
	}
	return jsonData, nil
}
func LoadTipsFromJson() ([]GitTips, string) {
	p, _ := ReadJsonFile()
	var result []GitTips
	json.Unmarshal(p, &result)
	return result, string(p)
}

func GetTipRefactor(key string) string {
	data, _ := LoadTipsFromJson()
	for index, _ := range data {
		if key == data[index].Title {
			result := data[index].Tip
			return result
		}
	}
	return "key is not present"
}

func UserInput(key string) string {
	//Validation
	_, data := LoadTipsFromJson()
	if key == "" {
		return "invalid data"
	} else if strings.Contains(data, key) {
		return key
	} else {
		return ""
	}

}

func DisplayTip(tip string) map[string]string {
	data, _ := LoadTipsFromJson()
	for index, _ := range data {
		if tip == data[index].Tip {
			return map[string]string{"title": data[index].Title, "tip": data[index].Tip}
		}
	}
	return map[string]string{}

}

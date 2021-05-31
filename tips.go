package gitTips

import (
	"encoding/json"
	"errors"
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

type GitTips struct {
	Title string `json:"title"`
	Tip   string `json:"tip"`
}

//reading json data from file
func ReadJsonFile() ([]byte, error) {
	jsonData, err := ioutil.ReadFile("data\\tipsData.json")
	if err != nil {
		return []byte{}, errors.New("file issue")
	}
	return jsonData, nil
}

func LoadTipsFromJson() ([]GitTips, string) {
	p, _ := ReadJsonFile()
	var result []GitTips
	json.Unmarshal([]byte(p), &result)
	return result, string(p)
}

func GetTipRefactor(key string) string {
	data, _ := LoadTipsFromJson()
	for index := range data {
		if key == data[index].Title {
			result := data[index].Tip
			return result
		}
	}
	return "key is not present"
}

func UserInput(key string) string {
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
	for index := range data {
		if tip == data[index].Tip {
			return map[string]string{"title": data[index].Title, "tip": data[index].Tip}
		}
	}
	return map[string]string{}
}

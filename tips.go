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

//reading json file
func ReadJsonFile() ([]byte, error) {
	jsonData, err := ioutil.ReadFile("data\\tips.json")
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

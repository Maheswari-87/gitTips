package gitTips

import "io/ioutil"

func MockReadJsonFile() string {
	tips, _ := ioutil.ReadFile("data\\tips.json")
	return string(tips)
}

func MockUserInputString() string {
	return "delete"
}

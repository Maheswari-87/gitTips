package gitTips

import "io/ioutil"

func MockReadJsonFile() string {
	tips, _ := ioutil.ReadFile("C:\\Users\\SRS\\repos\\githubProjects\\tips\\jsonFile\\tips.json")

	return string(tips)
}

func MockUserInputString() string {
	return "delete"
}

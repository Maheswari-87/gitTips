package gitTips

func MockReadJsonFile() (string, error) {
	myJson :=
		`[{"title": "Everyday Git in twenty commands or so","tip": "git help everyday"},{"title": "Show helpful guides that come with Git","tip": "git help -g"},{"title": "Delete remote branch","tip": "git push origin --delete <remote_branchname>"},{"title": "Saving current state of tracked files without commiting","tip": "git stash"}]`

	return string(myJson), nil
}

func MockUserInputString() string {
	return "delete"
}

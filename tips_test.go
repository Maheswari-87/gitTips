package gitTips

import (
	"reflect"
	"testing"
)

//Test GetTip
func TestGetTip(t *testing.T) {
	jsonData := map[string]string{"Title": "Everyday Git in twenty commands or so", "Tip": "git help everyday"}
	got := GetTip(jsonData, "Tip")
	want := "git help everyday"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

//Refactor GetTip
func TestGetTipRefactor(t *testing.T) {
	t.Run("Delete", func(t *testing.T) {
		key := "Delete remote branch"
		got := GetTipRefactor(key)
		want := "git push origin --delete <remote_branchname>"
		if got != want {
			t.Errorf("got %q  want %q", got, want)
		}
	})

	t.Run("not commands", func(t *testing.T) {
		key := "hello"
		got := GetTipRefactor(key)
		want := "key is not present"
		if got != want {
			t.Errorf("got %q  want %q", got, want)
		}
	})

}

//Refactor Read json
func TestLoadTipsFromJson(t *testing.T) {
	_, got := LoadTipsFromJson()
	expected, _ := MockReadJsonFile()

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %q want %q", got, expected)
	}
}

func TestReadJsonFileNegative(t *testing.T) {
	_, err := MockReadJsonFile()
	want := "file issue"

	if reflect.DeepEqual(err, want) {
		t.Errorf("err %q want %q", err, want)
	}
}

func TestUserInput(t *testing.T) {
	key := MockUserInputString()
	got := UserInput(key)
	want := "delete"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q  want %q", got, want)
	}

}

func TestDisplayTip(t *testing.T) {
	tip := "git stash"
	got := DisplayTip(tip)
	want := map[string]string{"title": "Saving current state of tracked files without commiting", "tip": "git stash"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q  want %q", got, want)
	}

}

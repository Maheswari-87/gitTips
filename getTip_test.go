package tips

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestGetTip(t *testing.T) {
	jsonData := map[string]string{"Title": "Everyday Git in twenty commands or so", "Tip": "git help everyday"}
	got := GetTip(jsonData, "Tip")
	want := "git help everyday"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func TestLoadTipsFromJson(t *testing.T) {
	got := LoadTipsFromJson()
	data, err := ioutil.ReadFile("JsonFile/tips.json")
	if err != nil {
		fmt.Print(err)
	}
	want := string(data)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}

package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var args = []string{"find"}

func Test_man(t *testing.T) {
	if err := Exec("man", args...); err != nil {
		t.Error(err)
	}
}

/*
func Test_tldr(t *testing.T) {
	if err := Exec("tldr", args...); err != nil {
		t.Error(err)
	}
}*/

func Test_mdcat(t *testing.T) {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	workDirPath := filepath.Dir(workDir)

	mdFile := filepath.Join(workDirPath, "/../README.md")
	fmt.Println(mdFile)

	if err := Exec("mdcat", mdFile); err != nil {
		t.Error(err)
	}
}

func Test_rootDir(t *testing.T) {
	assert(t, RootDir(), "/unified-help/uh")
}

func Test_cobraCmd(t *testing.T) {
	ExecCobraCmd("intellij")
	//ExecCobraCmd("vscode")
	//ExecCobraCmd("eclipse")
}

func Test_getFiles(t *testing.T) {
	files := GetMDFiles("testdata/", "~/Bad/me/", "/Bad/Location", "https://notes.tibbes.com/notes")
	assert(t, len(files), 3)
	for _, file := range files {
		fmt.Println(file)
	}
}

func Test_getFilesFromHttp(t *testing.T) {
	//GetMDFilesFromHttp("https://github.com/powdrsoft/unified-help/tree/master/hugo/content/posts")
}

func assert(t *testing.T, result interface{}, expected interface{}) {
	fmt.Println(result)
	switch e := expected.(type) {
	case bool:
		if result.(bool) != e {
			t.Errorf("%t != %t", result, e)
		}
	case int:
		if result.(int) < e {
			t.Errorf("%d < %d", result, e)
		}
	default:
		if !strings.Contains(result.(string), expected.(string)) {
			t.Errorf("%s does not contain %s", result, e)
		}

	}
}

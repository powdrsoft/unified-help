package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

var args = []string{"find"}

func Test_man(t *testing.T) {
	if err := Exec("man", args...); err != nil {
		t.Error(err)
	}
}

func Test_tldr(t *testing.T) {
	if err := Exec("tldr", args...); err != nil {
		t.Error(err)
	}
}

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
	d := RootDir()
	fmt.Println(d)
}

func Test_cobraCmd(t *testing.T) {
	ExecCobraCmd("intellij")
	ExecCobraCmd("vscode")
	ExecCobraCmd("eclipse")
}

func Test_getFiles(t *testing.T) {
	files := getMDFiles("/Users/tericsson/dev/me/")
	for _, file := range files {
		fmt.Println(file)
	}
}

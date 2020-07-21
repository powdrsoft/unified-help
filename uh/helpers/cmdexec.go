package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

func Exec(cmd string, args ...string) error {
	path, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Println(err)
		return err
	}
	execCmd := exec.Command(cmd, args...)
	out, err := execCmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}
	fmt.Printf("From %s at %s\n", cmd, path)
	fmt.Printf("%s\n", out)
	return nil
}

func ExecCobraCmd(name string) {
	path := filepath.Join(RootDir(), "notes", name, "index.md")
	fmt.Println(path)
	Exec("mdcat", path)
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

type mdFile struct {
	path string
	name string
}

func getMDFiles(dir string) []mdFile {
	var files []mdFile
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if (filepath.Ext(path)) == ".md" {
			fmt.Println(info.Name())
			files = append(files, mdFile{path, info.Name()[:len(info.Name())-3]})
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

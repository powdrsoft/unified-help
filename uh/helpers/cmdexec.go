package helpers

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
	"strings"
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

func ExecCobraCmd(path string) {
	//path := filepath.Join(RootDir(), "notes", name, "index.md")
	fmt.Println(path)
	Exec("mdcat", path)
}

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

type mdFile struct {
	Path string
	Name string
}

func GetMDFiles(locations ...string) map[string]string {
	var files = make(map[string]string)
	for _, location := range locations {

		if strings.HasPrefix(location, "~/") {
			usr, _ := user.Current()
			location = filepath.Join(usr.HomeDir, location[2:])
		}

		err := filepath.Walk(location, func(path string, info os.FileInfo, err error) error {
			if (filepath.Ext(path)) == ".md" {
				fmt.Println(info.Name())
				files[info.Name()[:len(info.Name())-3]] = path
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	}
	return files
}

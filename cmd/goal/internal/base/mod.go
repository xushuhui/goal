package base

import (
	"bufio"
	"bytes"
	"golang.org/x/mod/modfile"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ModulePath returns go module path.
func ModulePath(filename string) (string, error) {
	modBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return modfile.ModulePath(modBytes), nil
}

// ModuleVersion returns module version.
func ModuleVersion(path string) (string, error) {
	stdout := &bytes.Buffer{}
	fd := exec.Command("go", "mod", "graph")
	fd.Stdout = stdout
	fd.Stderr = stdout
	if err := fd.Run(); err != nil {
		return "", err
	}
	rd := bufio.NewReader(stdout)
	for {
		line, _, err := rd.ReadLine()
		if err != nil {
			return "", err
		}
		str := string(line)
		i := strings.Index(str, "@")
		if strings.Contains(str, path+"@") && i != -1 {
			return path + str[i:], nil
		}
	}
}

// Mod returns goal mod.
func Mod() string {
	gopath := os.Getenv("GOPATH")
	if path, err := ModuleVersion("github.com/xushuhui/goal"); err == nil {
		// $GOPATH/pkg/mod/github.com/xushuhui/goal
		return filepath.Join(gopath, "pkg", "mod", path)
	}
	// $GOPATH/src/github.com/xushuhui/goal
	return filepath.Join(gopath, "src", "github.com", "xushuhui", "goal")
}

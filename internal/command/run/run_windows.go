//go:build windows
// +build windows

package run

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/xushuhui/goal/config"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/xushuhui/goal/internal/pkg/helper"
)

var quit = make(chan os.Signal, 1)



var excludeDir string
var includeExt string

func init() {
	CmdRun.Flags().StringVarP(&excludeDir, "excludeDir", "", excludeDir, `eg: goal run --excludeDir="tmp,vendor,.git,.idea"`)
	CmdRun.Flags().StringVarP(&includeExt, "includeExt", "", includeExt, `eg: goal run --includeExt="go,tpl,tmpl,html,yaml,yml,toml,ini,json"`)
	if excludeDir == "" {
		excludeDir = config.RunExcludeDir
	}
	if includeExt == "" {
		includeExt = config.RunIncludeExt
	}
}

var CmdRun = &cobra.Command{
	Use:     "run",
	Short:   "goal run [main.go path]",
	Long:    "goal run [main.go path]",
	Example: "goal run cmd/server",
	Run: func(cmd *cobra.Command, args []string) {
		cmdArgs, programArgs := helper.SplitArgs(cmd, args)
		var dir string
		if len(cmdArgs) > 0 {
			dir = cmdArgs[0]
		}
		base, err := os.Getwd()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
			return
		}
		if dir == "" {
			cmdPath, err := helper.FindMain(base, excludeDir)

			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
				return
			}
			switch len(cmdPath) {
			case 0:
				fmt.Fprintf(os.Stderr, "ERROR: %s\n", "The cmd directory cannot be found in the current directory")
				return
			case 1:
				for _, v := range cmdPath {
					dir = v
				}
			default:
				var cmdPaths []string
				for k := range cmdPath {
					cmdPaths = append(cmdPaths, k)
				}
				prompt := &survey.Select{
					Message:  "Which directory do you want to run?",
					Options:  cmdPaths,
					PageSize: 10,
				}
				e := survey.AskOne(prompt, &dir)
				if e != nil || dir == "" {
					return
				}
				dir = cmdPath[dir]
			}
		}
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		fmt.Printf("Goal run %s.", dir)
		fmt.Printf("Watch excludeDir %s", excludeDir)
		fmt.Printf("Watch includeExt %s", includeExt)
		watch(dir, programArgs)

	},
}

func watch(dir string, programArgs []string) {

	// Listening file path
	watchPath := "./"

	// Create a new file watcher
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer watcher.Close()

	excludeDirArr := strings.Split(excludeDir, ",")
	includeExtArr := strings.Split(includeExt, ",")
	includeExtMap := make(map[string]struct{})
	for _, s := range includeExtArr {
		includeExtMap[s] = struct{}{}
	}
	// Add files to watcher
	err = filepath.Walk(watchPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		for _, s := range excludeDirArr {
			if s == "" {
				continue
			}
			if strings.HasPrefix(path, s) {
				return nil
			}
		}
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			if _, ok := includeExtMap[strings.TrimPrefix(ext, ".")]; ok {
				err = watcher.Add(path)
				if err != nil {
					fmt.Println("Error:", err)
				}
			}

		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cmd := start(dir, programArgs)

	// Loop listening file modification
	for {
		select {
		case <-quit:
			err = killProcess(cmd)

			if err != nil {
				fmt.Printf("server exiting...")
				return
			}
			fmt.Printf("server exiting...")
			os.Exit(0)

		case event := <-watcher.Events:
			// The file has been modified or created
			if event.Op&fsnotify.Create == fsnotify.Create ||
				event.Op&fsnotify.Write == fsnotify.Write ||
				event.Op&fsnotify.Remove == fsnotify.Remove {
				fmt.Printf("file modified: %s", event.Name)
				killProcess(cmd)

				cmd = start(dir, programArgs)
			}
		case err := <-watcher.Errors:
			fmt.Println("Error:", err)
		}
	}
}

func killProcess(cmd *exec.Cmd) error {
	if cmd.Process == nil {
		return nil
	}
	// 获取进程ID
	pid := cmd.Process.Pid
	// 构造taskkill命令
	taskkill := exec.Command("taskkill", "/F", "/T", "/PID", strconv.Itoa(pid))
	err := taskkill.Run()
	if err != nil {
		return err
	}
	return nil
}
func start(dir string, programArgs []string) *exec.Cmd {
	cmd := exec.Command("go", append([]string{"run", dir}, programArgs...)...)
	// Set a new process group to kill all child processes when the program exits

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd run failed\u001B[0m")
	}
	time.Sleep(time.Second)
	fmt.Printf("running...")
	return cmd
}

package minisql

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/smartystreets/goconvey/convey"
)

type Cmd struct {
	cmd *exec.Cmd
}

func Start(command string, args []string) (e *Cmd, err error) {
	e = &Cmd{}
	e.cmd = exec.Command(command)
	e.cmd.Stdin = os.Stdin
	e.cmd.Stdout = os.Stdout
	e.cmd.Stderr = os.Stderr
	e.cmd.Args = append([]string{command}, args...)
	e.cmd.Env = os.Environ()

	fmt.Printf("%v start...\n", e.cmd.Args)
	err = e.cmd.Start()
	if err != nil {
		fmt.Errorf("cmd start err,err: %v", err)
		return
	}

	return
}
func Exec(command string, args []string) (e *Cmd, err error) {
	e = &Cmd{}
	cmd := exec.Command(command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Args = append([]string{command}, args...)
	cmd.Env = os.Environ()
	e.cmd = cmd
	fmt.Printf("%v start...\n", e.cmd.Args)
	err = cmd.Start()
	if err != nil {
		fmt.Errorf("cmd start err,err: %v", err)
		return
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Errorf("cmd exit err,err: %v", err)
	}
	return
}
func (e *Cmd) Stop() (err error) {
	fmt.Printf("%v stoped!\n", e.cmd.Args)
	err = e.cmd.Process.Kill()
	return
}

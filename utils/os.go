package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func Cls() {
	var cls *exec.Cmd
	if runtime.GOOS == "windows" {
		cls = exec.Command("cmd", "/c", "cls") //Windows example, its tested
	} else {
		cls = exec.Command("clear")
	}
	cls.Stdout = os.Stdout
	cls.Run()
}

func Wait() {
	void := ""
	fmt.Print("Press enter to continue")
	fmt.Scanf("%v\n", &void)
}

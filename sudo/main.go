package sudo

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("vim-go")
}

func Run(command string) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`do shell script "%s" with administrator privileges`, command))
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	fmt.Println(cmd)
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println(outStr, errStr)
}

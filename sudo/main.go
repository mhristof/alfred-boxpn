package sudo

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("vim-go")
}

func Run(command string) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`do shell script "%s" with administrator privileges`, command))
	_ = cmd.Run()
}

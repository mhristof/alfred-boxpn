package sudo

import (
	"fmt"
	"os/exec"
)

func Run(command string) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`do shell script "%s &> /dev/null &" with administrator privileges`, command))
	_ = cmd.Start()
}

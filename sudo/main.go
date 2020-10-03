package sudo

import (
	"fmt"
	"os/exec"
	"strings"
)

func Run(command string) {
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`do shell script "%s &> /tmp/log &" with administrator privileges`, strings.ReplaceAll(command, `"`, `\"`)))
	_ = cmd.Start()
}

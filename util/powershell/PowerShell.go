package powershell

import (
	"WindowsListener/models"
	"os/exec"
	"strings"
	"syscall"
)

func Run(command string) string {
	models.Logger.Debugf("尝试执行：%s", command)
	cmd := exec.Command("powershell", "-WindowStyle", "Hidden", "-NoProfile", "-Command", command)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.Output()
	if err != nil {
		models.Logger.Errorf("执行 %s 命令失败：%s", command, err)
	}
	models.Logger.Debugf("调用%s 成功，输出：%s", command, output)
	return strings.TrimSpace(string(output))
}

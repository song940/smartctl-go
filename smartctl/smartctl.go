package smartctl

import (
	"encoding/json"
	"os"
	"os/exec"
)

type SmartCtl struct {
	dev string
}

func Open(dev string) (smart *SmartCtl) {
	smart = &SmartCtl{dev}
	return
}

func (s *SmartCtl) Read() (info *SmartInfo, err error) {
	info = &SmartInfo{}
	cmd := exec.Command("smartctl", "-a", s.dev, "--json")
	cmd.Env = os.Environ()
	out, _ := cmd.Output()
	err = json.Unmarshal(out, info)
	return
}

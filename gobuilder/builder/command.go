package builder

import (
	"os/exec"
)

func BuildCommand(b *Builder) *exec.Cmd {
	c := exec.Command("")
	var args []string
	if b.config.Lambda != nil {
		if b.config.Lambda.Amd64 == true {
			args = append(args, "GOARCH=amd64")
		}
		if b.config.Lambda.Linux == true {
			args = append(args, "GOOS=linux")
		}
	}
	args = append(args, "go", "build", "-o", b.config.OutputName, b.config.Source)
	c.Args = args
	return c
}

package runtime

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/3JoB/unsafeConvert"

	ph "github.com/3JoB/ulib/path"
)

type Info struct {
	KernelReleaseID      string
	KernelDisplayVersion string
}

// Get the directory where the executable file is located.
func RunPath() (string, error) {
	return ph.Abs(ph.Dir(os.Args[0]))
}

func UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func UserCacheDir() (string, error) {
	return os.UserCacheDir()
}

func UserConfigDir() (string, error) {
	return os.UserConfigDir()
}

// Use pwsh
func KernelInfo() *Info {
	info := &Info{}
	info.KernelDisplayVersion = KernelDisplayVersion()
	info.KernelReleaseID = KernelReleaseID()
	return info
}

func KernelReleaseID() string {
	var execmd *exec.Cmd
	if runtime.GOOS == "windows" {
		return cmd(winPwshKernelReg("ReleaseID"))
	}
	execmd = exec.Command("uname", "-r")

	output, err := execmd.Output()
	if err != nil {
		return ""
	}
	return unsafeConvert.StringReflect(output)
}

func KernelDisplayVersion() string {
	var execmd *exec.Cmd
	if runtime.GOOS == "windows" {
		return cmd(winPwshKernelReg("DisplayVersion"))
	}
	execmd = exec.Command("uname", "-r")

	output, err := execmd.Output()
	if err != nil {
		return ""
	}
	return unsafeConvert.StringReflect(output)
}

type wsInfo struct {
	CompositionEditionID string
	CurrentBuild         string
	DisplayVersion       string
	ProductName          string
	InstallationType     string
	ReleaseID            string
}

func WindowsKernelInfo() *wsInfo {
	ws := &wsInfo{}
	if runtime.GOOS != "windows" {
		return nil
	}
	ws.CompositionEditionID = cmd(winPwshKernelReg("CompositionEditionID"))
	ws.CurrentBuild = cmd(winPwshKernelReg("CurrentBuild"))
	ws.DisplayVersion = cmd(winPwshKernelReg("DisplayVersion"))
	ws.ProductName = cmd(winPwshKernelReg("ProductName"))
	ws.ReleaseID = cmd(winPwshKernelReg("ReleaseID"))
	ws.InstallationType = cmd(winPwshKernelReg("InstallationType"))
	return ws
}

func cmd(name string, args []string) string {
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	return unsafeConvert.StringReflect(output)
}

func winPwshKernelReg(v string) (name string, args []string) {
	name = "pwsh"
	args = append(args, "-Command")
	cmd := fmt.Sprintf(`(Get-ItemProperty "HKLM:\SOFTWARE\Microsoft\Windows NT\CurrentVersion").%v`, v)
	args = append(args, cmd)
	return
}

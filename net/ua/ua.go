package ua

import (
	"bytes"

	"github.com/shirou/gopsutil/v3/host"
)

const (
	Chrome        string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
	Edge          string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.57"
	AndroidChrome string = "Mozilla/5.0 (Linux; Android 11; Jelly2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Mobile Safari/537.36"
	ULIBDefault   string = "Mozilla/5.0 (compatible; ulib/1.22.1; +https://github.com/3JoB/ulib/)"
	Nex           string = "; "
)

type Config struct {
	AddCompatible  bool
	DisableSysInfo bool
	DisableULIB    bool
	CompatibleInfo string
	CompatibleUrl  string
	SoftInfo       string
}

func GenerateUA(c Config) string {
	var buffer bytes.Buffer
	buffer.WriteString("Mozilla/5.0 (")
	h, err := host.Info()
	if err != nil {
		return ULIBDefault
	}
	if c.AddCompatible {
		buffer.WriteString("compatible")
		buffer.WriteString(Nex)
		buffer.WriteString(c.CompatibleInfo)
		buffer.WriteString(Nex)
		buffer.WriteString("+")
		buffer.WriteString(c.CompatibleUrl)
	} else {
		buffer.WriteString(h.OS)
		buffer.WriteString(Nex)
		if !c.DisableULIB {
			buffer.WriteString("ulib/1.22.1")
			buffer.WriteString(Nex)
		}
		if !c.DisableSysInfo {
			buffer.WriteString(h.Platform)
			buffer.WriteString("/")
			buffer.WriteString(h.PlatformVersion)
			buffer.WriteString(Nex)
			buffer.WriteString(h.KernelArch)
		}
	}

	buffer.WriteString(") ")

	if c.SoftInfo != "" {
		buffer.WriteString(c.SoftInfo)
	}
	return buffer.String()
}

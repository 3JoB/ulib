package ua

import (
	"strings"

	"github.com/shirou/gopsutil/v3/host"
)

const (
	Chrome        string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
	Edge          string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.57"
	AndroidChrome string = "Mozilla/5.0 (Linux; Android 11; Jelly2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Mobile Safari/537.36"
	ULIBDefault   string = "Mozilla/5.0 (compatible; ulib/1.38.0; +https://github.com/3JoB/ulib/)"
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
	var builder strings.Builder
	builder.WriteString("Mozilla/5.0 (")
	h, err := host.Info()
	if err != nil {
		return ULIBDefault
	}
	if c.AddCompatible {
		builder.WriteString("compatible")
		builder.WriteString(Nex)
		builder.WriteString(c.CompatibleInfo)
		builder.WriteString(Nex)
		builder.WriteString("+")
		builder.WriteString(c.CompatibleUrl)
	} else {
		builder.WriteString(h.OS)
		builder.WriteString(Nex)
		if !c.DisableULIB {
			builder.WriteString("ulib/1.38.0")
			builder.WriteString(Nex)
		}
		if !c.DisableSysInfo {
			builder.WriteString(h.Platform)
			builder.WriteString("/")
			builder.WriteString(h.PlatformVersion)
			builder.WriteString(Nex)
			builder.WriteString(h.KernelArch)
		}
	}

	builder.WriteString(") ")

	if c.SoftInfo != "" {
		builder.WriteString(c.SoftInfo)
	}
	return builder.String()
}

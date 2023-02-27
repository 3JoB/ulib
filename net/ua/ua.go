package ua

import (
	"github.com/shirou/gopsutil/v3/host"
)

const (
	Chrome        string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
	Edge          string = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36 Edg/110.0.1587.57"
	AndroidChrome string = "Mozilla/5.0 (Linux; Android 11; Jelly2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Mobile Safari/537.36"
	ULIBDefault   string = "Mozilla/5.0 (compatible; ulib/1.4.0; +https://core.lcag.org/ulib/)"
	Nex string = "; "
)

type Config struct {
	AddCompatible  bool
	CompatibleInfo string
	CompatibleUrl  string
	DisableSysInfo bool
	DisableULIB    bool
	SoftInfo       string
}

func GenerateUA(c Config) (ua string) {
	ua = "Mozilla/5.0 ("
	h, err := host.Info()
	if err != nil {
		return ULIBDefault
	}
	if c.AddCompatible {
		ua = ua + "compatible" + Nex + c.CompatibleInfo + Nex +"+" + c.CompatibleUrl
	} else {
		ua = ua + h.OS + Nex
		if !c.DisableULIB {
			ua = ua + "ulib/1.4.0" + Nex
		}
		if !c.DisableSysInfo {
			ua = ua + h.Platform + "/" +h.PlatformVersion + Nex + h.KernelArch
		}
	}

	ua = ua + ") "

	if c.SoftInfo != "" {
		ua = ua + c.SoftInfo
	}
	return
}

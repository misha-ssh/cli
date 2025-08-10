package version

import (
	"regexp"
	"runtime/debug"
)

var (
	regexVersion = `^v?(\d+\.\d+\.\d+)`
	emptyVersion = ""
)

func Get() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		version := info.Main.Version

		re := regexp.MustCompile(regexVersion)
		if matches := re.FindStringSubmatch(version); len(matches) > 1 {
			return matches[1]
		}
		return emptyVersion
	} else {
		return emptyVersion
	}
}

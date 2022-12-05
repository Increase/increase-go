package core

import (
	"fmt"
	"runtime"
)

func getNormalizedOS() string {
	switch runtime.GOOS {
	case "ios":
		return "iOS"
	case "android":
		return "Android"
	case "darwin":
		return "MacOS"
	case "window":
		return "Windows"
	case "freebsd":
		return "FreeBSD"
	case "openbsd":
		return "OpenBSD"
	case "linux":
		return "Linux"
	default:
		return fmt.Sprintf("Other:%s", runtime.GOOS)
	}
}

func getNormalizedArchitecture() string {
	switch runtime.GOARCH {
	case "386":
		return "x32"
	case "amd64":
		return "x64"
	case "arm":
		return "arm"
	case "arm64":
		return "arm64"
	default:
		return fmt.Sprintf("other:%s", runtime.GOARCH)
	}
}

func getPlatformProperties() map[string]string {
	return map[string]string{
		"X-Stainless-Lang":            "go",
		"X-Stainless-Package-Version": "unknown",
		"X-Stainless-OS":              getNormalizedOS(),
		"X-Stainless-Arch":            getNormalizedArchitecture(),
		"X-Stainless-Runtime":         "go",
		"X-Stainless-Runtime-Version": runtime.Version(),
	}
}

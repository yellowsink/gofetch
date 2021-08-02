package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/yellowsink/gofetch/termcontrol"
	"os/user"
	"strconv"
	"strings"
	"time"
)


func main() {
	// print ascii
	ascii := getAscii()
	asciiHeight := len(strings.Split(ascii, "\n"))
	asciiWidth := len(strings.Split(ascii, "\n")[0])
	fmt.Println(ascii + "\n")

	// query system data
	osText, uptimeText, hostText := getOS()
	cpuText := getCPU()
	memText := getMem()
	diskText := getDisk()

	// prepare to print system data
	fmt.Print(termcontrol.CursorUp(asciiHeight))
	infoCol := asciiWidth + 2
	printInfo := func(info string, color termcontrol.Color) {
		fmt.Printf("%s%s%s\n",
			termcontrol.CursorHorizontalPos(infoCol),
			termcontrol.SetOneColor(color, false, false),
			info)
	}

	// shove it all in the console
	printInfo(osText,     termcontrol.Red)
	printInfo(uptimeText, termcontrol.Magenta)
	printInfo(hostText,   termcontrol.Blue)
	printInfo(cpuText,    termcontrol.Green)
	printInfo(memText,    termcontrol.Cyan)
	printInfo(diskText,   termcontrol.Yellow)

    fmt.Print(termcontrol.SetGraphicRendition(termcontrol.Reset))
}

func getOS() (ostxt, uptimetxt, usertxt string) {
	osInfo, _ := host.Info()
	osName := osInfo.OS
	osVer := osInfo.PlatformVersion
	distro := osInfo.Platform
	osText := fmt.Sprintf("OS:     %s %s %s", distro, osName, osVer)

	uptimeRaw := osInfo.Uptime
	uptimeSeconds := uptimeRaw % 60
	uptimeMinutes := uptimeRaw / 60
	uptimeText := fmt.Sprintf("Uptime: %dm %ds", uptimeMinutes, uptimeSeconds)

	hostname := osInfo.Hostname
	userInfo, _ := user.Current()
	username := userInfo.Username
	hostText := fmt.Sprintf("Host:   %s@%s", username, hostname)

	return osText, uptimeText, hostText
}

func getCPU() string {
	cpuInfo, _ := cpu.Info()
	cpuLoads, _ := cpu.Percent(20 * time.Millisecond, false)
	cpuLoad := strconv.FormatFloat(cpuLoads[0], 'f', 2, 64)
	cpuPhysical, _ := cpu.Counts(false)
	cpuLogical, _ := cpu.Counts(true)
	return fmt.Sprintf("CPU:    %s%% %s (with %d logical processors, %d physical)", cpuLoad, cpuInfo[0].ModelName, cpuLogical, cpuPhysical)
}

func getMem() string {
	memInfo, _ := mem.VirtualMemory()
	memUsed := memInfo.Used / 1_000_000
	memTotal := memInfo.Total / 1_000_000
	memPercent := strconv.FormatFloat(memInfo.UsedPercent, 'f', 2, 64)
	memAvail := memInfo.Available / 1_000_000
	return fmt.Sprintf("Mem:    %dMB / %dMB (%s%%), %dMB Available", memUsed, memTotal, memPercent, memAvail)
}

func getDisk() string {
	usage, _ := disk.Usage("/")

	used := usage.Used / 1_000_000 // convert to MB
	usedSuffix := "MB"
	usedGb := used > 10_000 // 10GB
	if usedGb {
		used /= 1000
		usedSuffix = "GB"
	}

	total := usage.Total / 1_000_000
	totalSuffix := "MB"
	totalGb := total > 10_000
	if totalGb {
		total /= 1000
		totalSuffix = "GB"
	}

	percent := strconv.FormatFloat(usage.UsedPercent, 'f', 2, 64)

	return fmt.Sprintf("Disk:   %s%% (%d%s / %d%s)", percent, used, usedSuffix, total, totalSuffix)
}

func getAscii() string {
	osInfo, _ := host.Info()
	osName := osInfo.OS
	distro := osInfo.Platform
	if (osName == "windows") {return windows}
	if (osName == "macos") {return macos}
	if (osName == "freebsd") {return freebsd}
	if (osName == "openbsd") {return openbsd}
	dontGoAgain := false
	distroSwitch:
	switch distro {
	case "alpine":
		return alpine
	case "guix":
		return guix
	case "arch":
		return arch
	case "artix":
		return artix
	case "debian":
		return debian
	case "fedora":
		return fedora
	case "gentoo":
		return gentoo
	case "manjaro":
		return manjaro
	case "nixos":
		return nixos
	case "raspbian":
		return raspbian
	case "ubuntu":
		return ubuntu
	case "void":
		return void
	default:
		if dontGoAgain { return linux}
		dontGoAgain = true
		goto distroSwitch
	}
}

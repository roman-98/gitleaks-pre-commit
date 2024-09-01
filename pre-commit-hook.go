package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

func main() {
	homeDir, err := getUserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}

	opSystem := runtime.GOOS
	gitURL := "https://github.com/gitleaks/gitleaks.git"
	gitleaksFolder := ""
	moveCmd := ""

	if opSystem == "windows" {
		gitleaksFolder = filepath.Join(homeDir, "Downloads", "gitleaks")
		moveCmd = fmt.Sprintf("move %s\\gitleaks.exe C:\\Windows\\System32\\gitleaks.exe", gitleaksFolder)
	} else {
		gitleaksFolder = filepath.Join(homeDir, "Downloads", "gitleaks")
		moveCmd = fmt.Sprintf("sudo mv %s/gitleaks /usr/local/bin", gitleaksFolder)
	}

	fmt.Println("os :", opSystem)

	if !gitleaksEnabled() {
		fmt.Println("gitleaks disabled")
		os.Exit(0)
	}

	if !isInstalled() {
		fmt.Println("gitleaks is not installed")
		commands := []string{
			fmt.Sprintf("git clone %s %s", gitURL, gitleaksFolder),
			fmt.Sprintf("make -C %s build", gitleaksFolder),
			moveCmd,
			fmt.Sprintf("rm -rf %s", gitleaksFolder),
		}

		for _, cmd := range commands {
			res := runCommand(cmd)
			if res != nil {
				fmt.Println("error:", res, ", for cmd:", cmd)
				os.Exit(1)
			}
		}
		fmt.Println("successfully installed gitleaks")
	}

	res := runCommand("gitleaks detect --report-path=leaks-report.json --report-format=json")
	if res != nil {
		fmt.Println("error:", res)
		os.Exit(1)
	}
}

func getUserHomeDir() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return usr.HomeDir, nil
}

func gitleaksEnabled() bool {
	out, err := exec.Command("git", "config", "--bool", "hooks.gitleaks").Output()
	if err != nil {
		return false
	}
	return string(out) != "false"
}

func isInstalled() bool {
	cmd := exec.Command("gitleaks", "--version")
	err := cmd.Run()
	return err == nil
}

func runCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

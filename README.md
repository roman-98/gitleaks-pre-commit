# Gitleaks Installation and Detection Script

This Go script automates the installation and execution of [Gitleaks](https://github.com/gitleaks/gitleaks), a tool for detecting hardcoded secrets and sensitive information in your Git repositories.

## Prerequisites

Before using the script, ensure you have the following installed on your system:

- [Go](https://golang.org/doc/install)
- Git
- Make

Additionally, you need the necessary permissions to move files to system directories:
- On Linux/macOS: `/usr/local/bin`
- On Windows: `C:\Windows\System32\`

## Installation and Usage

Follow these steps to use the script:

### 1. Clone or Copy the Script

Clone this repository or copy the Go script code into a `.go` file on your machine.

### 2. Run the Script

Open a terminal or command prompt, navigate to the directory containing the Go script, and execute it using the following command:

```bash
go run pre-commit-hook.go

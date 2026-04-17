package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const version = "1.1.0"

func loadDotEnv() {
	data, err := os.ReadFile(".env")
	if err != nil {
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			os.Setenv(key, val)
		}
	}
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	if args[0] == "--version" || args[0] == "-v" {
		fmt.Printf("localshare version %s\n", version)
		return
	}

	loadDotEnv()

	token := os.Getenv("CLOUDFLARE_TUNNEL_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "Error: CLOUDFLARE_TUNNEL_TOKEN is not set.")
		fmt.Fprintln(os.Stderr, "Create a .env file with your Cloudflare tunnel token:")
		fmt.Fprintln(os.Stderr, "  CLOUDFLARE_TUNNEL_TOKEN=your_token_here")
		fmt.Fprintln(os.Stderr, "Get your token at: https://one.dash.cloudflare.com")
		os.Exit(1)
	}

	if _, err := exec.LookPath("cloudflared"); err != nil {
		fmt.Fprintln(os.Stderr, "Error: cloudflared is not installed.")
		fmt.Fprintln(os.Stderr, "Install it first:")
		fmt.Fprintln(os.Stderr, "  https://pkg.cloudflare.com/cloudflared")
		os.Exit(1)
	}

	fmt.Printf("LocalShare v%s\n", version)
	fmt.Println("-----------------------------------------")
	fmt.Println("Loading configuration from .env...")
	fmt.Println("Connecting to Cloudflare tunnel...")

	cmd := exec.Command("cloudflared", "tunnel", "run", "--token", token)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: tunnel failed: %v\n", err)
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Printf(`LocalShare v%s — Share your local app with anyone, instantly

USAGE:
  localshare

OPTIONS:
  --version, -v Show version
  --help,    -h Show this help

REQUIREMENTS:
  cloudflared must be installed: https://pkg.cloudflare.com/cloudflared
  CLOUDFLARE_TUNNEL_TOKEN must be set in a .env file

SETUP:
  1. Install cloudflared (see link above)
  2. Create a .env file with your tunnel token:
       CLOUDFLARE_TUNNEL_TOKEN=your_token_here
  3. Run: localshare

`, version)
}

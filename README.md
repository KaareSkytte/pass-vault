# Pass-Vault

A secure CLI password manager built with Go, featuring AES-256-GCM encryption and Argon2 key derivation.

## Features

- **Strong Encryption**: AES-256-GCM with Argon2 key derivation
- **Local Storage**: Encrypted JSON files stored locally
- **Cross-Platform**: Works on Linux, macOS, and Windows
- **Secure Input**: Hidden password prompts
- **Rich Entries**: Store username, password, URL, and notes
- **Fast**: Built with Go for optimal performance

## Why pass-vault?

### The Problem
- **Cloud Password Managers**: Require internet, monthly fees, and trust in third parties
- **Browser Password Managers**: Limited to browsers, sync issues, vendor lock-in
- **Existing CLI Tools**: Often complex, bloated, or lack modern encryption standards

### The Solution
pass-vault provides a **simple, secure, local-first** password management solution:

- **Privacy First**: Your passwords never leave your machine
- **Zero Cost**: No subscriptions or cloud dependencies  
- **Developer Friendly**: CLI-first design for terminal users
- **Modern Crypto**: Industry-standard encryption (AES-256-GCM + Argon2)
- **Lightweight**: Single binary, no dependencies
- **Offline**: Works without internet connection

### Perfect For
- Developers who live in the terminal
- Privacy-conscious users
- Teams needing local password storage
- Anyone wanting full control over their credentials

## Quick Start

### Option 1: Download Pre-built Binary (Recommended)

1. Go to [Releases](https://github.com/kaareskytte/pass-vault/releases)
2. Download the binary for your system:
   - **Linux**: `pass-vault-linux-amd64`
   - **macOS Intel**: `pass-vault-darwin-amd64`  
   - **macOS Apple Silicon**: `pass-vault-darwin-arm64`
   - **Windows**: `pass-vault-windows-amd64.exe`
3. Make executable (Linux/macOS): `chmod +x pass-vault-linux-amd64`
4. Run: `./pass-vault-linux-amd64 init`

### Option 2: Install with Go

```bash
go install github.com/kaareskytte/pass-vault@latest
pass-vault init
```

### Option 3: Build from Source

```bash
git clone https://github.com/kaareskytte/pass-vault.git
cd pass-vault
go install
pass-vault init
```

## Usage

```bash
# Initialize a new vault
pass-vault init

# Add a password entry
pass-vault add "Gmail"
pass-vault add "GitHub" --url "github.com" --note "Development account"

# List all entries
pass-vault list

# Get an entry (password hidden by default)
pass-vault get "Gmail"

# Show password
pass-vault get "Gmail" --show-password

# Delete an entry
pass-vault delete "Gmail"
```

## Architecture

- Encryption: AES-256-GCM authenticated encryption
- Key Derivation: Argon2 with salt for password-based key derivation
- Storage: Local encrypted JSON files with secure file permissions
- CLI Framework: Built with Cobra for professional command-line interface


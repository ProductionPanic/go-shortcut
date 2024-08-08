# Cli chortcut

This cli app is a simple alias manager for your terminal.
It allows you to create, list, delete and run aliases.

## Installation

```bash
#!/bin/bash
git clone https://github.com/ProductionPanic/go-shortcut.git ~/temp_shortcut
cd ~/temp_shortcut
go build -o shortcut
chmod +x shortcut
sudo mv shortcut /usr/local/bin
```
This will clone the repository, build the binary and move it to the /usr/local/bin directory. So you can run the command from anywhere in your terminal. using the `shortcut` command.

## Usage
```bash
shortcut [command] [alias] [command]
```

## Commands
- `add` - Add a new alias
- `list` - List all aliases
- `delete` - Delete an alias
- `run` - Run an alias

## Example
```bash
shortcut add ll "ls -la"
shortcut run ll
# or just
ll # will run the ls -la command
```



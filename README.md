# serveHere
Serves your current directory on port 8080. That's all.

## Installation

<h4> not working right now, i dont have compiled versions. ignore this plz </h4>

### Windows

Run the following commands in a CMD session:
```batch
curl -Lo "%USERPROFILE%/servehere.exe" "https://github.com/overestimate/servehere/releases/download/0.0.0/webserve-win.exe"
setx PATH "%PATH%;%USERPROFILE%\servehere.exe"
```

### Mac OS (Intel)

Install cURL if needed.
Run the following commands using your preferred terminal session:
```sh
curl -Lo "$HOME/servehere.bin" "https://github.com/overestimate/servehere/releases/download/0.0.0/webserve-mac.bin"
echo 'export PATH=$PATH:$HOME/servehere.bin' >> $HOME/.profile
```

### Linux

Install cURL if needed.
Run the following commands using your preferred terminal session:
```sh
curl -Lo "$HOME/servehere.bin" "https://github.com/overestimate/servehere/releases/download/0.0.0/webserve-linux.bin"
echo 'export PATH=$PATH:$HOME/servehere.bin' >> $HOME/.profile
```

## Usage

### Windows

Open the folder you want to serve in CMD and run `servehere.exe`. You can now open localhost:8080.

### Mac OS / Linux

Open the folder you want to serve in your shell (zsh/bash/etc.) and run `servehere.bin`. You can now open localhost:8080.
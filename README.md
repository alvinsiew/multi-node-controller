# multi-node-controller

Multi-node-controller is a tool for issuing a command to multiple ec2 instances.

## How to start using

```bash
Usage:
        To issue commands:
                mnc web -c < commandline >
                mnc app -c < commandline >

        To toggle SPCP:
                mnc app -toggle <actual|stub>

        To list IPs:
                mnc web -l
                mnc app -l

        For production
                replace web > web-prd and app > app-prd

        Production does not support -toggle and -c commandline
```

## Self-compile

```bash
# MacOS
env GOOS=darwin GOARCH=amd64 go build -o bin/mnc cmd/mnc/main.go

# Linux
env GOOS=linux GOARCH=amd64 go build -o bin/mnc cmd/mnc/main.go

# Window
env GOOS=windows GOARCH=amd64 go build -o bin/mnc cmd/mnc/main.go
```

# multi-node-controller

Multi-node-controller is a tool for issuing a command to multiple ec2 instances.

## How to start using

1) Copy conf/config.yml to HOME_DIR/.aws/config.yml
2) Update HOME_DIR/.aws/config.yml with the path to private key
3) Add repo bin folder to PATH ENV in ur ~/.bash_profile or ~/.bashrc

You now good to go!

## prerequisite

Setup your AWS_PROFILE
https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html

## Command Usage

```bash
export AWS_PROFILE=<PROFILE NAME>

Usage:
        To issue commands:
                mnc web -c < commandline >
                mnc app -c < commandline >

        To toggle SPCP:
                mnc app -toggle <actual|stub>

        To control Passenger
                mnc app -passenger start|stop|restart|status

        To control Delayed_Job
                mnc app -delay start|stop|restart|status

        To control Web Server
                mnc app -nginx start|stop|restart|status

        To list IPs:
                mnc web -l
                mnc app -l

        For production
                replace web > web-prd and app > app-prd

        Production does not support -toggle
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

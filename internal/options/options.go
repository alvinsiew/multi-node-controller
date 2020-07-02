package options

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"multi-node-controller/internal/awsinternal"
	"multi-node-controller/internal/sshcmd"
	"os"
)

// Command struct for sub command name
type Command struct {
	fs *flag.FlagSet

	name      string
	subName   bool
	subNameII string
}

// WebCommand Contain flag command for web ec2 instances
func WebCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("web", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "c", "", "Command to be executed")
	gc.fs.BoolVar(&gc.subName, "l", false, "List VMs IPs")

	return gc
}

func WebProdCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("web-prd", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "c", "", "Command to be executed")
	gc.fs.BoolVar(&gc.subName, "l", false, "List VMs IPs")

	return gc
}

// AppCommand Contain flag command for application ec2 instances
func AppCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("app", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "c", "", "Command to be executed")
	gc.fs.BoolVar(&gc.subName, "l", false, "List VMs IPs")
	gc.fs.StringVar(&gc.subNameII, "toggle", "", "actual or stub for SPCP")

	return gc
}

func AppProdCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("app-prd", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "c", "", "Command to be executed")
	gc.fs.BoolVar(&gc.subName, "l", false, "List VMs IPs")

	return gc
}

func HelpCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("help", flag.ContinueOnError),
	}

	return gc
}

// Name to return subcommand name
func (g *Command) Name() string {
	return g.fs.Name()
}

// Init get argument after subcommand
func (g *Command) Init(args []string) error {
	return g.fs.Parse(args)
}

// Parameter to decide encryption or decryption
func (g *Command) Parameter() {

	if g.Name() == "web" {
		if g.subName == true {
			ListIP("ec2*uat*web*")
		} else {
			RemoteCMD("ec2*uat*web*", g)
		} 
	} else if g.Name() == "app" {
		if g.subName == true {
			ListIP("ec2*uat*app*")
		} else {
			RemoteCMD("ec2*uat*app*", g)
		} 
	} else if g.Name() == "web-prd" {
		if g.subName == true {
			ListIP("ec2*prd*web*")
		} else {
			RemoteCMD("ec2*prd*web*", g)
		} 		
	} else if g.Name() == "app-prd" {
		if g.subName == true {
			ListIP("ec2*uat*app*")
		} else {
			RemoteCMD("ec2*uat*app*", g)
		} 
	} else if g.Name() == "help" {
		fmt.Println(help())
	} 
}

// ListIP will list VMs IPs
func ListIP(cmd string) {
	i := 0
	ipAdr, name := awsinternal.FilterInstances(cmd)

	for _, n := range name {
		fmt.Println(n)
		fmt.Println(ipAdr[i])
		i++
	}
}

// RemoteCMD will issue command to filter VMs
func RemoteCMD(cmd string, g *Command) {

	command := g.name

	if g.subNameII == "stub" {
		command = "sudo -S -u deploy /bin/bash -c '/apps/scripts/toggle_spcp_login.sh singpass Stub'"
	} else if g.subNameII == "actual" {
		command = "sudo -S -u deploy /bin/bash -c '/apps/scripts/toggle_spcp_login.sh singpass Actual'"
	} else if g.subNameII != "" {
		log.Fatalln("Invalid option for toggle")
	}

	n := 0
	ipAdr, name := awsinternal.FilterInstances(cmd)

	for _, i := range ipAdr {
		fmt.Println(name[n])
		sshcmd.RemoteCommand(i, command)
		n++
	}
}

// Runner to run function with same method
type Runner interface {
	Init([]string) error
	Name() string
	Parameter()
}

func help() string {
	helpMessage := `Usage: 
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

	Production does not support -toggle`
	return helpMessage
}

// Root Manage command line
func Root(args []string) error {
	if len(args) < 1 {
		h := help()
		fmt.Println(h)
		return errors.New("You must pass a sub-command")
	}

	cmds := []Runner{
		WebCommand(),
		WebProdCommand(),
		AppCommand(),
		AppProdCommand(),
		HelpCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			cmd.Parameter()
			return nil
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

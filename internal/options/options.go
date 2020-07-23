package options

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"multi-node-controller/internal/awsinternal"
	"multi-node-controller/internal/sshcmd"
	"multi-node-controller/internal/yamlcustom"
	"os"
)

// Command struct for sub command name
type Command struct {
	fs *flag.FlagSet

	name           string
	subName        bool
	toggleSingpass string
	passenger      string
	delayedjob     string
	web            string
	ssh            bool
	squid          string
}

// WebCommand Contain flag command for uat web ec2 instances
func WebCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("web", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "c", "", "Command to be executed")
	gc.fs.BoolVar(&gc.subName, "l", false, "List VMs IPs")
	gc.fs.StringVar(&gc.web, "nginx", "", "start|stop|restart|status")
	gc.fs.BoolVar(&gc.ssh, "ssh", false, "To connect to instance")
	return gc
}

// WebProdCommand Contain flag command for prod web ec2 instances
func WebProdCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("web-prd", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "c", "", "Command to be executed")
	gc.fs.BoolVar(&gc.subName, "l", false, "List VMs IPs")
	gc.fs.StringVar(&gc.web, "nginx", "", "start|stop|restart|status")
	gc.fs.BoolVar(&gc.ssh, "ssh", false, "To connect to instance")

	return gc
}

// AppCommand Contain flag command for application uat ec2 instances
func AppCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("app", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "c", "", "Command to be executed")
	gc.fs.BoolVar(&gc.subName, "l", false, "List VMs IPs")
	gc.fs.StringVar(&gc.toggleSingpass, "toggle", "", "actual or stub for SPCP")
	gc.fs.StringVar(&gc.passenger, "passenger", "", "start|stop|restart|status")
	gc.fs.StringVar(&gc.delayedjob, "delay", "", "start|stop|restart|status")
	gc.fs.BoolVar(&gc.ssh, "ssh", false, "To connect to instance")

	return gc
}

// AppProdCommand Contain flag command for application prod ec2 instances
func AppProdCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("app-prd", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "c", "", "Command to be executed")
	gc.fs.BoolVar(&gc.subName, "l", false, "List VMs IPs")
	gc.fs.StringVar(&gc.toggleSingpass, "toggle", "", "actual or stub for SPCP")
	gc.fs.StringVar(&gc.passenger, "passenger", "", "start|stop|restart|status")
	gc.fs.StringVar(&gc.delayedjob, "delay", "", "start|stop|restart|status")
	gc.fs.BoolVar(&gc.ssh, "ssh", false, "To connect to instance")

	return gc
}

// SquidCommand Contain flag command for squid uat ec2 instance
func SquidCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("proxy", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "c", "", "Command to be executed")
	gc.fs.BoolVar(&gc.subName, "l", false, "List VMs IPs")
	gc.fs.StringVar(&gc.squid, "squid", "", "start|stop|restart|status")
	gc.fs.BoolVar(&gc.ssh, "ssh", false, "To connect to instance")

	return gc
}

// SquidProdCommand Contain flag command for squid prod ec2 instance
func SquidProdCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("proxy-prd", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "c", "", "Command to be executed")
	gc.fs.BoolVar(&gc.subName, "l", false, "List VMs IPs")
	gc.fs.StringVar(&gc.squid, "squid", "", "start|stop|restart|status")
	gc.fs.BoolVar(&gc.ssh, "ssh", false, "To connect to instance")

	return gc
}

// HelpCommand for help paramter
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

func sudoPasswdUAT() string {
	conf := yamlcustom.ParseYAML()

	return conf.Conf[3].SudoPasswd
}

func sudoPasswdProd() string {
	conf := yamlcustom.ParseYAML()

	return conf.Conf[4].SudoPasswdProd
}

func sudoUser() string {
	conf := yamlcustom.ParseYAML()
	return conf.Conf[0].UserID
}

// Parameter to decide encryption or decryption
func (g *Command) Parameter() {
	if g.Name() == "web" {
		if g.subName == true {
			ListIP("ec2*uat*web*")
		} else if g.ssh == true {
			ConnectSSH("ec2*uat*web*", sshcmd.GetKeyUat())
		} else {
			RemoteCMD("ec2*uat*web*", g, sshcmd.GetKeyUat(), sudoPasswdUAT())
		}
	} else if g.Name() == "app" {
		if g.subName == true {
			ListIP("ec2*uat*app*")
		} else if g.ssh == true {
			ConnectSSH("ec2*uat*app*", sshcmd.GetKeyUat())
		} else {
			RemoteCMD("ec2*uat*app*", g, sshcmd.GetKeyUat(), sudoPasswdUAT())
		}
	} else if g.Name() == "proxy" {
		if g.subName == true {
			ListIP("ec2*uat*proxy*")
		} else if g.ssh == true {
			ConnectSSH("ec2*uat*proxy*", sshcmd.GetKeyUat())
		} else {
			RemoteCMD("ec2*uat*proxy*", g, sshcmd.GetKeyUat(), sudoPasswdUAT())
		}
	} else if g.Name() == "web-prd" {
		if g.subName == true {
			ListIP("ec2*prd*web*")
		} else if g.ssh == true {
			ConnectSSH("ec2*prd*web*", sshcmd.GetKeyProd())
		} else {
			RemoteCMD("ec2*prd*web*", g, sshcmd.GetKeyProd(), sudoPasswdProd())
		}
	} else if g.Name() == "app-prd" {
		if g.subName == true {
			ListIP("ec2*prd*app*")
		} else if g.ssh == true {
			ConnectSSH("ec2*prd*app*", sshcmd.GetKeyProd())
		} else {
			RemoteCMD("ec2*prd*app*", g, sshcmd.GetKeyProd(), sudoPasswdProd())
		}
	} else if g.Name() == "proxy-prd" {
		if g.subName == true {
			ListIP("ec2*prd*proxy*")
		} else if g.ssh == true {
			ConnectSSH("ec2*prd*proxy*", sshcmd.GetKeyProd())
		} else {
			RemoteCMD("ec2*prd*proxy*", g, sshcmd.GetKeyProd(), sudoPasswdProd())
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
func RemoteCMD(cmd string, g *Command, getKey []uint8, sudopass string) {
	command := g.name

	if g.toggleSingpass == "stub" {
		command = "echo " + sudopass + " | sudo -S -u deploy /bin/bash -c '/apps/scripts/toggle_spcp_login.sh singpass Stub' --stdin"
	} else if g.toggleSingpass == "actual" {
		command = "echo " + sudopass + " | sudo -S -u deploy /bin/bash -c '/apps/scripts/toggle_spcp_login.sh singpass Actual' --stdin"
	} else if g.toggleSingpass != "" {
		log.Fatalln("Invalid option for toggle")
	} else if g.passenger == "start" {
		command = "echo " + sudopass + " | sudo -S systemctl start passenger"
	} else if g.passenger == "stop" {
		command = "echo " + sudopass + " | sudo -S systemctl stop passenger"
	} else if g.passenger == "restart" {
		command = "echo " + sudopass + " | sudo -S systemctl restart passenger"
	} else if g.passenger == "status" {
		command = "echo " + sudopass + " | sudo -S systemctl status passenger"
	} else if g.passenger != "" {
		log.Fatalln("Invalid option for passenger")
	} else if g.delayedjob == "start" {
		command = "echo " + sudopass + " | sudo -S systemctl start delayed_job"
	} else if g.delayedjob == "stop" {
		command = "echo " + sudopass + " | sudo -S systemctl stop delayed_job"
	} else if g.delayedjob == "restart" {
		command = "echo " + sudopass + " | sudo -S systemctl restart delayed_job"
	} else if g.delayedjob == "status" {
		command = "echo " + sudopass + " | sudo -S systemctl status delayed_job"
	} else if g.delayedjob != "" {
		log.Fatalln("Invalid option for delay")
	} else if g.web == "start" {
		command = "echo " + sudopass + " | sudo -S systemctl start nginx"
	} else if g.web == "stop" {
		command = "echo " + sudopass + " | sudo -S systemctl stop nginx"
	} else if g.web == "restart" {
		command = "echo " + sudopass + " | sudo -S systemctl restart nginx"
	} else if g.web == "status" {
		command = "echo " + sudopass + " | sudo -S systemctl status nginx"
	} else if g.web != "" {
		log.Fatalln("Invalid option for nginx")
	} else if g.squid == "stop" {
		command = "echo " + sudopass + " | sudo -S systemctl stop squid"
	} else if g.squid == "start" {
		command = "echo " + sudopass + " | sudo -S systemctl start squid"
	} else if g.squid == "restart" {
		command = "echo " + sudopass + " | sudo -S systemctl restart squid"
	} else if g.squid == "status" {
		command = "echo " + sudopass + " | sudo -S systemctl status squid"
	} else if g.name == "" {
		log.Fatalln("Invalid option use")
	}

	n := 0
	ipAdr, name := awsinternal.FilterInstances(cmd)

	for _, i := range ipAdr {
		fmt.Println(name[n])
		sshcmd.RemoteCommand(i, command, getKey)
		n++
	}
}

// ConnectSSH will issue command to filter VMs
func ConnectSSH(cmd string, getKey []uint8) {
	var answer string
	n := 0
	ipAdr, name := awsinternal.FilterInstances(cmd)

	for _, i := range ipAdr {
		fmt.Println("Do you want to ssh " + name[n] + "? (yes/y to connect)")
		fmt.Scanf("%s", &answer)

		if answer == "yes" || answer == "y" {
			sshcmd.TerminalConn("", i, getKey)
			os.Exit(0)
		}
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
		mnc proxy -c < commandline >
	
	To toggle SPCP:
		mnc app -toggle <actual|stub>

	To control Passenger
		mnc app -passenger start|stop|restart|status

	To control Delayed_Job
		mnc app -delay start|stop|restart|status

	To control Web Server
		mnc app -nginx start|stop|restart|status

	To control Squid server
		mnc proxy -squid start|stop|restart|status

	To SSH to server
		mnc web -ssh
		mnc app -ssh
		mnc proxy -ssh
	
	To list IPs:
		mnc web -l
		mnc app -l
		mnc proxy -l
	
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
		SquidCommand(),
		SquidProdCommand(),
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

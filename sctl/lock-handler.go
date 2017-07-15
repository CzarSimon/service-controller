package main // sctl-cli

import (
	sctl "github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// UnlockCommand Command for issuing unlock requets
func (env Env) UnlockCommand() cli.Command {
	return cli.Command{
		Name:   "unlock",
		Usage:  "Unlocks the cluster",
		Action: env.RequestUnlock,
	}
}

// RequestUnlock Requests the api server to unlock the cluster
func (env Env) RequestUnlock(c *cli.Context) error {
	var emptyCommand sctl.Command
	status := env.SendCommandToNodes("unlock", emptyCommand)
	if status == util.StatusOK {
		color.Blue("Cluster unlocked!")
	} else {
		color.Red("Unlock failed")
	}
	return nil
}

// LockCommand Command for issuing lock requets
func (env Env) LockCommand() cli.Command {
	return cli.Command{
		Name:   "lock",
		Usage:  "Locks the cluster",
		Action: env.RequestLock,
	}
}

// RequestLock Requests the api server to unlock the cluster
func (env Env) RequestLock(c *cli.Context) error {
	var emptyCommand sctl.Command
	status := env.SendCommandToNodes("lock", emptyCommand)
	if status == util.StatusOK {
		color.Blue("Cluster locked!")
	} else {
		color.Red("Lock failed")
	}
	return nil
}

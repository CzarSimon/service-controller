package main // sctl-cli

import (
	"fmt"
	"path/filepath"

	sctl "github.com/CzarSimon/sctl-common"
	"github.com/CzarSimon/util"
)

// SetupNode Sends executbles and starts them on the node
func (env Env) SetupNode(node sctl.Node) {
	env.SetupMinonDB()
	SendExecutables(env.config.Folders, node)
	SendTokenDB(env.config.Folders, node)
	WaitForStartUp(env.config.Folders.Target, node)
}

// WaitForStartUp Waits unit user has started minon on node
func WaitForStartUp(targetFolder string, node sctl.Node) {
	StartDescription(targetFolder, node)
	GetInput("\nThen go back here and press ENTER")
}

// SendExecutables Sends executables to designated destination on node
func SendExecutables(folders FolderConfig, node sctl.Node) {
	execFolder := filepath.Join(folders.Exec, node.OS, "sctl-minion")
	send := node.RsyncFolderCMD(execFolder, folders.Target)
	//fmt.Println(send.ToString())
	out, err := send.Execute()
	util.CheckErrFatal(err)
	if out != "" {
		fmt.Println(out)
	}
}

// SendTokenDB Sends the token database to the designated node
func SendTokenDB(folders FolderConfig, node sctl.Node) {
	dbFile := filepath.Join(folders.Token, "token-db")
	send := node.RsyncFileCMD(dbFile, folders.Target)
	//fmt.Println(send.ToString())
	out, err := send.Execute()
	util.CheckErrFatal(err)
	if out != "" {
		fmt.Println(out)
	}
}

// StartDescription Prints the description of how to start a node
func StartDescription(targetFolder string, node sctl.Node) {
	fmt.Println("Go to", node.RsyncTarget(targetFolder), "and run:")
	fmt.Println("\nsudo sh setup-minion.sh")
}

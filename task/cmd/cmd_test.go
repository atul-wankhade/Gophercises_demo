package cmd

import "testing"

func Test_AddCommand(t *testing.T) {
	cmd := addCmd
	cmd.SetHelpTemplate("test")
	cmd.Execute()
}

func Test_DoCommand(t *testing.T) {
	cmd := doCmd
	cmd.Execute()
}

func Test_RmCommand(t *testing.T) {
	cmd := rmCmd
	cmd.Execute()
}

func Test_ListCommand(t *testing.T) {
	cmd := ListCmd
	cmd.Execute()
}

func Test_ExecuteCommand(t *testing.T) {
	cmd := RootCmd
	cmd.Execute()
}

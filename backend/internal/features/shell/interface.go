package shell

type ShellExecInteface interface {
	ExecuteCommand(command string) error
}

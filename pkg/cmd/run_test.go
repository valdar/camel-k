package cmd

import (
	"bytes"
	"context"
	"github.com/spf13/cobra"
	"testing"
)

func TestRunPropertyArgs(t *testing.T) {
	rootCmd := &cobra.Command{Use: "root", Run: emptyRun}
	options := RootCmdOptions{
		Context: context.Background(),
	}

	runCmdOptions := newCmdRun(&options)

	runCmdOptions.Command.RunE = func(c *cobra.Command, args []string) error {

		return nil
	}
	runCmdOptions.Command.Args = arbitraryArgs

	rootCmd.AddCommand(runCmdOptions.Command)

	_, err := executeCommand(rootCmd, "run", "route.java", "--property", "key1=value,othervalue", "--property", "key2=value2")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(runCmdOptions.Properties) != 2 {
		t.Errorf("Properties expected to contain: \n %v elements\nGot:\n %v elemtns\n", 2, len(runCmdOptions.Properties))
	}
	if runCmdOptions.Properties[0] != "key1=value,othervalue" || runCmdOptions.Properties[1] != "key2=value2" {
		t.Errorf("Properties expected to be: \n %v\nGot:\n %v\n", "[key1=value,othervalue key2=value2]", runCmdOptions.Properties)
	}
}

func emptyRun(*cobra.Command, []string) {}

func arbitraryArgs(cmd *cobra.Command, args []string) error {
	return nil
}

func executeCommand(root *cobra.Command, args ...string) (output string, err error) {
	_, output, err = executeCommandC(root, args...)
	return output, err
}

func executeCommandC(root *cobra.Command, args ...string) (c *cobra.Command, output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOutput(buf)
	root.SetArgs(args)

	c, err = root.ExecuteC()

	return c, buf.String(), err
}

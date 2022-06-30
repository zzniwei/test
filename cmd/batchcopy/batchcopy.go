package batchcopy

import (
	"context"
	"errors"
	"strings"

	"github.com/rclone/rclone/cmd"
	"github.com/rclone/rclone/fs/config/flags"
	"github.com/rclone/rclone/fs/sync"
	"github.com/spf13/cobra"
)

var (
	createEmptySrcDirs = false
)

func init() {
	cmd.Root.AddCommand(commandDefinition)
	cmdFlags := commandDefinition.Flags()
	flags.BoolVarP(cmdFlags, &createEmptySrcDirs, "create-empty-src-dirs", "", createEmptySrcDirs, "Create empty source dirs on destination after copy")
}

var commandDefinition = &cobra.Command{
	Use:   "batchcopy filepath/filename",
	Short: `edit job copy list to transfer dir from source to dest`,
	// Note: "|" will be replaced by backticks below
	Long: strings.ReplaceAll(`
edit txt file with 'sourcepath, destpath' line by line, batch copy the source dir to the destination.  Does not transfer files that are
identical on source and destination, testing by size and modification
time or MD5SUM.  Doesn't delete files from the destination.

Note that it is always the contents of the directory that is synced,
not the directory so when sourcpath is a directory, it's the
contents of sourcepath that are copied, not the directory name and
contents.

If destpath doesn't exist, it is created and the sourcepath contents
go there.

For example

    rclone batchcopy /home/test.txt

Let's say there are two items in file list

    sourcepath,destpath
    sourcepath1,destpath1

This copies all contents from sourcepath to destpath and sourcepath1 to destpath1

**Note**: Use the |-P|/|--progress| flag to view real-time transfer statistics.

**Note**: Use the |--dry-run| or the |--interactive|/|-i| flag to test without copying anything.
`, "|", "`"),
	Run: func(command *cobra.Command, args []string) {

		cmd.CheckArgs(1, 2, command, args)
		copyList := cmd.GetBatchCopyPairsFs(args)

		cmd.Run(true, true, command, func() error {
			restList := sync.BatchCopy(context.Background(), copyList, createEmptySrcDirs)

			if restList != nil {
				cmd.UpdateBatchCopyPairsFs(args, restList)
				return errors.New("update rest jobs list when next time need it")
			} else {
				cmd.UpdateBatchCopyPairsFs(args, nil)
				return nil
			}
		})
	},
}

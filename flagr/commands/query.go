package commands

import (
	"fmt"
	"github.com/antihax/optional"
	"github.com/checkr/goflagr"
	"github.com/urfave/cli"
)

func QuerySubCommand() cli.Command {
	return cli.Command{
		Name:    "query",
		Aliases: []string{"find"},
		Usage:   "Search for flags in Flagr",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "d, description",
				Usage: "Flags with description...",
			},
			cli.StringFlag{
				Name:  "k, key",
				Usage: "Flag with keys matching...",
			},
			cli.IntFlag{
				Name:  "i, id",
				Usage: "Flag with ID...",
			},
		},
		Action: queryRunCommand,
	}
}

func queryRunCommand(c *cli.Context) error {
	client := getFlagrClient(c)
	var flags []goflagr.Flag
	var err error

	if c.IsSet("id") {
		// For this, we use Get not Find
		flag, _, err := client.FlagApi.GetFlag(nil, int64(c.Int("id")))
		if err != nil {
			logAndDie(fmt.Sprintf("No string with ID %d", c.Int("id")))
		}

		flags = []goflagr.Flag{flag}
	}

	if c.IsSet("description"){
		desc := c.String("description")
		ffo := goflagr.FindFlagsOpts{DescriptionLike:optional.NewString(desc)}
		flags, _, err = client.FlagApi.FindFlags(nil, &ffo)
		if err != nil {
			logAndDie(fmt.Sprintf("Error finding flags %v", err))
		}
	}

	columns := []string{"ID","Description", "Key"}
	formatFlags("table", "    ", true, columns, flags)
	return nil
}

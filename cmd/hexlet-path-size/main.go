package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func PrintFileOrDirSize(ctx context.Context, cmd *cli.Command) error {
	human := cmd.Bool("human")
	all := cmd.Bool("name")
	recursive := cmd.Bool("recursive")
	if cmd.NArg() > 0 {
		path := cmd.Args().Get(0)
		res, err := code.GetPathSize(path, recursive, human, all)
		if err != nil {
			return err
		}
		fmt.Printf("%s\t%s\n", res, path)
	} else {
		fmt.Println(cmd.Usage)
	}
	return nil
}

func main() {
	cmd := &cli.Command{
		Name:   "hexlet-path-size",
		Usage:  "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		Action: PrintFileOrDirSize,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "human",
				Value:   false,
				Usage:   "human-readable sizes (auto-select unit)",
				Aliases: []string{"H"},
			},
			&cli.BoolFlag{
				Name:    "all",
				Value:   false,
				Usage:   "include hidden files and directories",
				Aliases: []string{"a"},
			},
			&cli.BoolFlag{
				Name:    "recursive",
				Value:   false,
				Usage:   "recursive size of directories",
				Aliases: []string{"r"},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

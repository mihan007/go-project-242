package main

import (
	"code"
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v3"
)

var human, all, recursive bool

func PrintFileOrDirSize(ctx context.Context, cmd *cli.Command) error {
	filename := ""
	if cmd.NArg() > 0 {
		filename = cmd.Args().Get(0)
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		res, err := code.GetPathSize(filepath.Join(dir, filename), recursive, human, all)
		if err != nil {
			return err
		}
		fmt.Println(res, filename)
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
				Name:        "human",
				Value:       false,
				Usage:       "human-readable sizes (auto-select unit)",
				Destination: &human,
				Aliases:     []string{"H"},
			},
			&cli.BoolFlag{
				Name:        "all",
				Value:       false,
				Usage:       "include hidden files and directories",
				Destination: &all,
				Aliases:     []string{"a"},
			},
			&cli.BoolFlag{
				Name:        "recursive",
				Value:       false,
				Usage:       "recursive size of directories",
				Destination: &recursive,
				Aliases:     []string{"r"},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

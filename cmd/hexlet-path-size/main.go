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

var human bool

func PrintFileOrDirSize(ctx context.Context, cmd *cli.Command) error {
	filename := ""
	if cmd.NArg() > 0 {
		filename = cmd.Args().Get(0)
		dir, err := os.Getwd()
		if err != nil {
			return err
		}
		res, err := code.GetPathSize(filepath.Join(dir, filename), false, human, false)
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
		Usage:  "print size of a file or directory",
		Action: PrintFileOrDirSize,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "human",
				Value:       false,
				Usage:       "human-readable sizes (auto-select unit)",
				Destination: &human,
				Aliases:     []string{"H"},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

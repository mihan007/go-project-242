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

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			filename := ""
			if cmd.NArg() > 0 {
				filename = cmd.Args().Get(0)
				dir, err := os.Getwd()
				if err != nil {
					return err
				}
				res, err := code.GetPathSize(filepath.Join(dir, filename), false, true, false)
				if err != nil {
					return err
				}
				fmt.Println(res, filename)
			}
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	pathsize "code"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		UsageText: "hexlet-path-size [global options]",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := "."

			if cmd.Args().Len() > 0 {
				path = cmd.Args().First()
			}

			size, ok := pathsize.GetSize(path)

			if ok != nil {
				log.Fatal(ok)
			}

			fmt.Sprintf("%d\t%s\n", size, path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

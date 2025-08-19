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
		Usage:     "print size of a file or directory; supports -r (recursive), -H (human-readable), -a (include hidden)",
		UsageText: "hexlet-path-size [global options]",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
			&cli.BoolFlag{
				Name:    "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.NArg() < 1 {
				return cli.Exit("Please provide a path", 1)
			}
			path := cmd.Args().First()
			humanReadable := cmd.Bool("human")
			withHidden := cmd.Bool("all")
			recursive := cmd.Bool("recursive")

			formattedSize, ok := pathsize.GetPathSize(path, recursive, humanReadable, withHidden)

			if ok != nil {
				log.Fatal(ok)
			}

			fmt.Printf("%s\t%s\n", formattedSize, path)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

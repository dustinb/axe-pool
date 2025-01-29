package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v3"
	"oldbute.com/axe-pool/lib"
)

func main() {
	lib.Init()

	cmd := &cli.Command{
		Name:      "axe-pool",
		Usage:     "A cli for managing Bitaxe pools",
		UsageText: `axe-pool [command] [options]`,
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "List available pools and Bitaxe",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					lib.List()
					return nil
				},
			},
			{
				Name:  "scan",
				Usage: "Scan the network for Bitaxe",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					lib.Scan()
					lib.List()
					return nil
				},
			},
			{
				Name:  "add",
				Usage: "Add a pool to the database",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					lib.AddPool()
					lib.List()
					return nil
				},
			},
			{
				Name:  "del",
				Usage: "Delete a pool from the database. axe-pool del ID",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					id := cmd.Args().Get(0)
					lib.DeletePool(id)
					lib.List()
					return nil
				},
			},
			{
				Name:  "set",
				Usage: "Update a bitaxe (or all) with primary and fallback example `pool all 3 2`",
				Action: func(ctx context.Context, cmd *cli.Command) error {
					bitaxeId := cmd.Args().Get(0)
					if bitaxeId == "" {
						return cli.Exit("Bitaxe ID is required", 1)
					}
					primaryId := cmd.Args().Get(1)
					if primaryId == "" {
						return cli.Exit("Primary pool ID is required", 1)
					}
					fallbackId := cmd.Args().Get(2)
					if fallbackId == "" {
						return cli.Exit("Fallback pool ID is required", 1)
					}
					err := lib.SetPool(bitaxeId, primaryId, fallbackId)
					if err != nil {
						return err
					}
					fmt.Println("Doing a re-scan in 20 seconds...")
					time.Sleep(20 * time.Second)
					lib.Scan()
					lib.List()
					return nil
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

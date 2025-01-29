package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/urfave/cli/v3"
)

func List() {
	OutputPools()
	fmt.Println()
	OutputBitaxes()
}

func Scan() {
	bitaxe := ScanNetwork()
	fmt.Printf("Found %d bitaxe\n\n", len(bitaxe))

	now := time.Now()

	for _, found := range bitaxe {
		exists := Bitaxe{}
		Database.Where("mac_addr = ?", found.MacAddr).First(&exists)
		if exists.ID != 0 {
			found.ID = exists.ID
		}
		Database.Save(&found)
	}

	// Clean up old Bitaxe
	Database.Unscoped().Where("updated_at < ?", now).Delete(&Bitaxe{})
}

func AddPool() {
	// get host from stdin
	fmt.Println("Enter host: ")
	reader := bufio.NewReader(os.Stdin)
	host, _ := reader.ReadString('\n')
	host = strings.TrimSpace(host)

	// get port from stdin
	fmt.Println("Enter port: ")
	port, _ := reader.ReadString('\n')
	port = strings.TrimSpace(port)

	// get user from stdin
	fmt.Println("Enter user: ")
	user, _ := reader.ReadString('\n')
	user = strings.TrimSpace(user)

	porti, _ := strconv.Atoi(port)

	pool := Pool{
		Host: host,
		Port: porti,
		User: user,
	}
	Database.Create(&pool)
}

func DeletePool(id string) cli.ExitCoder {
	Database.Unscoped().Where("id = ?", id).Delete(&Pool{})
	return nil
}

func SetPool(bitaxeId string, primaryId string, fallbackId string) cli.ExitCoder {
	var patchList []Bitaxe

	if bitaxeId == "all" {
		Database.Find(&patchList)
	} else {
		bitaxe := Bitaxe{}
		Database.Where("id = ?", bitaxeId).First(&bitaxe)
		if bitaxe.ID == 0 {
			return cli.Exit("Bitaxe with id "+bitaxeId+" not found", 1)
		}
		patchList = append(patchList, bitaxe)
	}

	primary := Pool{}
	Database.Where("id = ?", primaryId).First(&primary)
	if primary.ID == 0 {
		return cli.Exit("Primary pool not found", 1)
	}
	fallback := Pool{}
	Database.Where("id = ?", fallbackId).First(&fallback)
	if fallback.ID == 0 {
		return cli.Exit("Fallback pool not found", 1)
	}

	for _, bitaxe := range patchList {
		primaryUser := primary.User
		if !strings.Contains(primaryUser, ".") {
			primaryUser = primaryUser + "." + bitaxe.Hostname
		}
		fallbackUser := fallback.User
		if !strings.Contains(fallbackUser, ".") {
			fallbackUser = fallbackUser + "." + bitaxe.Hostname
		}
		PatchAxe(bitaxe.IP, Patch{
			StratumURL:          primary.Host,
			StratumPort:         primary.Port,
			StratumUser:         primaryUser,
			FallbackStratumURL:  fallback.Host,
			FallbackStratumPort: fallback.Port,
			FallbackStratumUser: fallbackUser,
		})
	}
	return nil
}

func OutputPools() {
	var pools []Pool
	Database.Find(&pools)

	fmt.Println("\nPools:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "ID\tHost\tUser")
	for _, pool := range pools {
		fmt.Fprintf(w, "%d\t%s:%d\t%s\n", pool.ID, pool.Host, pool.Port, pool.User)
	}
	w.Flush()
}

func OutputBitaxes() {
	var bitaxes []Bitaxe
	Database.Find(&bitaxes)

	if len(bitaxes) == 0 {
		Scan()
	}

	fmt.Println("\nBitaxe:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "ID\tHost\tStratum\tUser")
	for _, bitaxe := range bitaxes {
		fmt.Fprintf(w, "%d\t%s\t%s:%d\t%s\n", bitaxe.ID, bitaxe.Hostname, bitaxe.StratumURL, bitaxe.StratumPort, bitaxe.StratumUser)
		fmt.Fprintf(w, "%s\t%s\t%s:%d\t%s\n", "", "", bitaxe.FallbackStratumURL, bitaxe.FallbackStratumPort, bitaxe.FallbackStratumUser)
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", "", "", "", "")
	}
	w.Flush()
}

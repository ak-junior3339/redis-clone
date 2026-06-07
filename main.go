package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	store := NewStore()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Redis-clone Started.....\nEnter A command or type HELP\n\n")
	for {

		fmt.Print("(Local) > ")
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)

		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "SET":
			if len(parts) < 3 {
				fmt.Println("usage: SET <key> <value>")
				fmt.Println("Please Try Again with Correct Usage.")
				continue
			}
			store.Set(parts[1], parts[2])
			fmt.Println("OK")
		case "GET":
			if len(parts) < 2 {
				fmt.Println("usage: GET <key> ")
				fmt.Println("Please Try Again with Correct Usage.")
				continue
			}
			val, ok := store.Get(parts[1])
			if ok {
				fmt.Println(val)
			} else {
				fmt.Println("key not found")
			}
		case "DEL":
			if len(parts) < 2 {
				fmt.Println("usage: DEL <key> ")
				fmt.Println("Please Try Again with Correct Usage.")
				continue
			}
			store.Delete(parts[1])
			fmt.Println("OK")
		case "EXISTS":
			if len(parts) < 2 {
				fmt.Println("usage: EXISTS <key>")
				fmt.Println("Please Try Again with Correct Usage.")
				continue
			}
			if store.Exists(parts[1]) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}

		case "KEYS":
			keys := store.Keys()
			if len(keys) == 0 {
				fmt.Println("(empty)")
			} else {
				for i, k := range keys {
					fmt.Printf("%d) %s\n", i+1, k)
				}
			}

		case "FLUSH":
			store.Flush()
			fmt.Println("OK")
		case "EXPIRE":
			if len(parts) < 3 {
				fmt.Println("usage: EXPIRE <key> <seconds>")
				continue
			}
			seconds, err := strconv.Atoi(parts[2])
			if err != nil {
				fmt.Println("seconds must be a number")
				continue
			}
			store.Expire(parts[1], seconds)
			fmt.Println("OK")

		case "TTL":
			if len(parts) < 2 {
				fmt.Println("usage: TTL <key>")
				continue
			}
			remaining := store.TTL(parts[1])
			if remaining == -1 {
				fmt.Println("no expiry set or key expired")
			} else {
				fmt.Printf("%d seconds remaining\n", remaining)
			}
		case "HELP":
			fmt.Println("Command\tSyntax\t\t\tUsage")
			fmt.Println("-------\t------\t\t\t-----")
			fmt.Println("SET\tSET <key> <value>\tStore a key-value pair")
			fmt.Println("GET\tGET <key>\t\tRetrieve value by key")
			fmt.Println("DEL\tDEL <key>\t\tDelete a key-value pair")
			fmt.Println("EXISTS\tEXISTS <key>\t\tCheck if a key exists")
			fmt.Println("KEYS\tKEYS\t\t\tList all keys")
			fmt.Println("FLUSH\tFLUSH\t\t\tDelete all keys")
			fmt.Println("EXPIRE\tEXPIRE <key> <sec>\tSet expiry on a key")
			fmt.Println("TTL\tTTL <key>\t\tCheck remaining expiry time")
			fmt.Println("HELP\tHELP\t\t\tShow this help message")
			fmt.Println("EXIT\tEXIT\t\t\tQuit the program")
		case "EXIT":
			fmt.Println("Bye! See You Soon!")
			return
		default:
			fmt.Println("unknown command")
		}
	}
}

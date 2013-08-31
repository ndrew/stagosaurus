/**
             _
  _ __   __| |_ ____      __
 | '_ \ / _` | '__\ \ /\ / /
 | | | | (_| | |   \ V  V /
 |_| |_|\__,_|_|    \_/\_/

------------------------------

Static blog generator for http://sernyak.com

*/
package main

import "flag"
import "fmt"
import "sort"

var COMMANDS = map[string]string{
    "new": "bla", 
    "edit": "bla", 
    "publish": "bla",
    "list": "bla"}

func printHeader() {
    fmt.Println("╔════════════════════════════════╗")
    fmt.Println("╟          BLOG GENERATOR        ╢")
    fmt.Println("╚════════════════════════════════╝")
    fmt.Println("Usage: blog [--use-editor|no-editor] [--help]")
    fmt.Println("                <command> [<args>]")
    fmt.Println("")
    

}

func list_commands() {
    fmt.Println("Available commands:")

    var keys []string
    for k := range COMMANDS {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, command := range keys {
        fmt.Println("\t", command, " - ", COMMANDS[command])
    }
}

func main() {
    flag.Parse()
    var action = flag.Arg(0)

    // print list of completion rules, maybe introduce smth like --silent flag
    if (action == "autocomplete") {
        list_commands()
        return 
    }

    printHeader()

    switch {
        case action == "list": list_commands()
        default: list_commands()
    }

}

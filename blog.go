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
    "new": "<post-name> [<params>] - creates new post and opens editor", 
    "edit": "<post-name>           - opens post in editor", 
    "publish": "[<post-name>]      - renders markdown posts to html"}


func printHeader() {
    fmt.Println("╔════════════════════════════════════════╗")
    fmt.Println("╟ ░░░░░░░░░░░ BLOG GENERATOR ░░░░░░░░░░░ ╢")
    fmt.Println("╚════════════════════════════════════════╝")
    fmt.Println("Usage: blog [--config <cfg-file>] [--help]")
    fmt.Println("             <command> [<args>]")
    fmt.Println("──────────────────────────────────────────")
}


func listCommands(full_description bool) {
    if (full_description) {
        fmt.Println("Available commands:")
    }
    
    var keys []string
    for k := range COMMANDS {
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, command := range keys {
        if (full_description) {
            fmt.Println(" ■", command, COMMANDS[command])
        } else {
            // TODO: params
            fmt.Println(command)
        }
    }
}


func parseArguments(args []string) {
    // iterate through params and form a map of --properties, action and action params

}


func main() {
    flag.Parse()
    parseArguments(flag.Args())

    var action = flag.Arg(0)

    if (action == "autocomplete") {
        listCommands(false)
        return 
    }

    switch {
        case action == "new": fmt.Println("new post")
        case action == "edit": fmt.Println("edit")
        case action == "publish": fmt.Println("publish")

        default: {
            printHeader()
            listCommands(true)
        } 
    }

}

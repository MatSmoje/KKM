package main

import (
	"fmt"
	"os"
)

const version = "0.0.1"

const colorRed = "\033[0;33m"
const colorNone = "\033[0m"
const colorBlue = "\033[0;36m"


func print_help() {
	fmt.Fprintf(os.Stdout,`%skkm: Coords manager and password manager.%s

    %sUsage: kkm [-g] Type Coord Coord Coord%s

    Optional arguments
    --version         : Get version
    -h,   --help      : Print help options

    Getting Data
    -g                : Get data. First argument must be type, then coords (Max. 3)

    Adding Data
    -t    --type      : Add Type
    -a    --add       : Add register
    -k    --key       : Specify Key related to register
    -v    --value     : Specify Value related to register
    

`, colorRed , colorNone, colorBlue, colorNone)
}


func print_version() {
	fmt.Fprintf(os.Stdout, "%sKKM %s%s \n",colorRed , version, colorNone ) 
}
package main

import (
	"fmt"
	"os"
)

const version = "0.0.2"

const colorRed = "\033[0;33m"
const colorNone = "\033[0m"
const colorBlue = "\033[0;36m"


func print_help() {
	fmt.Fprintf(os.Stdout,`kkm: Coords manager. 

 %sUsage: kkm [-g] Type Coord Coord Coord%s

 %sOptional arguments%s
 --version         : Get version
 -h,   --help      : Print help options

 %sGetting Data%s
 -g                : Get data. First argument must be type, then coords (Max. 3)
 Usage:  kkm -g BANK a1 b3 c4 
 
 %sAdding Data%s
 -a    --add       : Add register
 Usage:  kkm -a BANK a1 32

 KKM %s

`, colorRed , colorNone, colorBlue, colorNone,colorBlue, colorNone,colorBlue, colorNone, version)
}


func print_version() {
	fmt.Fprintf(os.Stdout, "%sKKM %s%s \n",colorRed , version, colorNone ) 
}
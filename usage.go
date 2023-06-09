package main

import (
	"fmt"
	"os"
    //"reflect" //4type of vars
)


const colorRed = "\033[0;33m"
const colorNone = "\033[0m"
const colorBlue = "\033[0;36m"


func print_help() {
	fmt.Fprintf(os.Stdout,`%skkm: Coords manager and password manager.%s

    %sUsage: kkm [-t] [-p1] [-p2] [-p3]%s

    Optional arguments:
    -h,   --help                : Print help options
    -t,   --type                : Choose type of information to get
    -p1   -p2   -p3             : Specify which coords do you want 
    -pw   -passwd               : Set new passwd
    -a    -add                  : Add register
    -k    --key                 : Specify Key related to register
    -v    --value               : Specify Valu related to register
    --version                   : Get version

`, colorRed , colorNone, colorBlue, colorNone)
}


func print_version() {
	fmt.Fprintf(os.Stdout, "%s - [   KKM - Version 0.0.0   ] - %s \n",colorRed , colorNone ) 
}

func argumentsInterpreter(arguments []string){
    switch {
        case len(arguments) == 2:
            if arguments[1] == "-h" || arguments[1] == "--help" {
                print_help()
            } 
            if arguments[1] == "--version"{
                print_version()
            }
            if arguments[1] == "-pw" || arguments[1] == "-passwd" {
                fmt.Println("Cambio De Passwd")
            }
        case len(arguments) == 9:
            if (arguments[1] == "-t" || arguments[1] == "-type") && (arguments[3] == "-p1") && (arguments[5] == "-p2") && (arguments[7] == "-p3"){
                fmt.Println("Print coords")
            }
            
        case len(arguments) == 6:
            if (arguments[1] == "-a" || arguments[1] == "-add") && (arguments[2] == "-k") && (arguments[4] == "-v"){
                fmt.Println("Add coords")
            }            
        default:
            print_help()
            fmt.Println("default")
    }
}


func main() {
	//print_help()
	//print_version()
    argumentsInterpreter(os.Args)
}
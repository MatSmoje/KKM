package main

import (
	"fmt"
)

func print_help() {
	fmt.Println(`Usage: kkm [Options] command 
            Recognized options:
			 --help           : Print help options'
             -t               : Get type of information to get
             -p[1-3]          : Specify which coordenates do you want
             -passwd          : Change passwd
             -a               : Specify you want to add new value
             -a -k [v] -v [v] : Detail of key and value you will add
            
            '\nRecognized commands:`)
}

func main() {
	print_help()
}
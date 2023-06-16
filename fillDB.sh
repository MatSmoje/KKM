#!/bin/bash

main()
{
	coordChar=({a..j})
	coordNumber=({1..5})

	for i in ${coordChar[@]};do
		for j in ${coordNumber[@]}; do
			coord="$i$j"
			value=$(($RANDOM%98)) 
			if [[ "$value" -lt 10 ]]; then
				value="0$value"
			fi
			./PasswordManager -a -t itau -k ${coord} -v ${value}
		done;
	done
}
main "$@"

package main

import (
	"fmt"
)

func main() {
	i := fonc_ret_fonc()
	j := fonc_ret_fonc()
	fmt.Println("e1:", i(),"e2", j())
	fmt.Println("e3:", i(), "e4",i())
	fmt.Println("e5:", i(), "e6",j())
	if (i()==10){
		fmt.Println( "erreur" )
	}else{
		fmt.Println("e7:", i())
	}	
}

func fonc_ret_fonc() func() int {
	i := 0
	return func() int {
		i++
		return i + i
	}
}
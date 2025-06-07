package main

import (
	"fmt"
	"time"
)

func main(){

	i:= 2
	fmt.Println("Write", i, "as")
	switch i {
	case 1:
		fmt.Println("One")
	case 2: 
		fmt.Println("Two")
	case 3: 
		fmt.Println("Three")
		
	}

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It is a weekend")
	default:
		fmt.Println("It is a weekday")
	}


	t:= time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is afternoon")
	}

	whatAmI := func(i interface{}){
		switch t := i.(type){
		case bool:
			fmt.Println(i, "is a bool")
		case int:
			fmt.Println(i, "is an int")
		default:
			fmt.Println("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(2)
	whatAmI("a")
}
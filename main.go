package main

import "fmt"
import "reflect"
import "flag"
import "strconv"
import "github.com/asystat/project_euler/eulerutils"

type Euler struct {}

func main() {
    var e Euler
    n_problem := flag.Int("problem", 1, "problem number")
    flag.Parse()
    reflect.ValueOf(&e).MethodByName("Problem_" + strconv.Itoa(*n_problem)).Call([]reflect.Value{})
}


func (e *Euler) Problem_1() {
    sum := 0
    for i:=1; i<1000; i++{
	if i % 3 == 0 || i % 5 == 0{
	    sum += i
	}
    }
    fmt.Println("Solution: ", sum)
}


func (e *Euler) Problem_2() {
    sum := 0
    last := 1
    lastlast := 0
    for last < 4000000{
	fib := last + lastlast
	lastlast = last
	last = fib
	if fib % 2 == 0{
	    sum += fib
	} 
    }
    fmt.Println("Solution: ", sum)
}

func (e *Euler) Problem_3() {
    var num int64
    num = 600851475143
    for i := eulerutils.Sqrt(num); i>0; i--{
	if num % i == 0 && eulerutils.IsPrime(i){
	    fmt.Println("Solution: ", i)
	    return
	}
    }
    fmt.Println("something bad happened")
}

func (e *Euler) Problem_4() {
    m1 := 999
    m2 := 999
    max := 0

    for {
	mult := m1 * m2
        if eulerutils.IsPalindrome(mult){
	    if max < mult{
	        max = mult
	    }
	}

        m1 --
        if m1 < 100 {
	    m2-- 
	    if m2<100{
	        break
	    }
	    m1 = m2
	}
    }
    fmt.Println("Solution: ", max)
}

func (e *Euler) Problem_5() {
    num := 20
    for{
	for i:=1; i<=20; i++{
	    if num % i != 0{
	        num++
		break
	    }
	    if i == 20{
	        fmt.Println("Solution: ", num)
		return
	    }
	}
    }
}

func (e *Euler) Problem_6() {
    
}

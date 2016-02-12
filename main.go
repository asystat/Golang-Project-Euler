package main

import "fmt"
import "reflect"
import "flag"
import "strconv"
import "math"
import "github.com/asystat/project_euler/eulerutils"

type Euler struct{}

func main() {
	var e Euler
	n_problem := flag.Int("problem", 1, "problem number")
	flag.Parse()
	reflect.ValueOf(&e).MethodByName("Problem_" + strconv.Itoa(*n_problem)).Call([]reflect.Value{})
}

func (e *Euler) Problem_1() {
	sum := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("Solution: ", sum)
}

func (e *Euler) Problem_2() {
	sum := 0
	last := 1
	lastlast := 0
	for last < 4000000 {
		fib := last + lastlast
		lastlast = last
		last = fib
		if fib%2 == 0 {
			sum += fib
		}
	}
	fmt.Println("Solution: ", sum)
}

func (e *Euler) Problem_3() {
	var num int64
	num = 600851475143
	for i := eulerutils.Sqrt(num); i > 0; i-- {
		if num%i == 0 && eulerutils.IsPrime(i) {
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
		if eulerutils.IsPalindrome(mult) {
			if max < mult {
				max = mult
			}
		}

		m1--
		if m1 < 100 {
			m2--
			if m2 < 100 {
				break
			}
			m1 = m2
		}
	}
	fmt.Println("Solution: ", max)
}

func (e *Euler) Problem_5() {
	num := 20
	for {
		for i := 1; i <= 20; i++ {
			if num%i != 0 {
				num++
				break
			}
			if i == 20 {
				fmt.Println("Solution: ", num)
				return
			}
		}
	}
}

func (e *Euler) Problem_6() {
	var squares_sum float64 = 0
	var sum_squared int64 = 0
	for count := 0; count <= 100; count++ {
		squares_sum += math.Pow(float64(count), 2)
		sum_squared += int64(count)
	}
	result := math.Pow(float64(sum_squared), 2) - squares_sum
	fmt.Println("El resultado es: ", int64(result))
}

func (e *Euler) Problem_7() {
	primeCount := 0
	var index int64 = 0
	for {
		if eulerutils.IsPrime(index) {
			primeCount++
			if primeCount == 10001 {
				break
			}
		}
		index++
	}
	fmt.Println("The 10001 prime is ", index)
}

func (e *Euler) Problem_8() {
	//solution for problem 8
}
func (e *Euler) Problem_9() {

}
func (e *Euler) Problem_10() {

}
func (e *Euler) Problem_11() {

}
func (e *Euler) Problem_12() {

}
func (e *Euler) Problem_13() {

}
func (e *Euler) Problem_14() {

}

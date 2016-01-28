package main

import "fmt"
import "reflect"
import "flag"
import "strconv"
import "github.com/asystat/project_euler/eulerutils"
import "math"

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
    var sumsquare, squaresum float64
    sumsquare = math.Pow(((100 * 101)/2), 2)
    fmt.Println("(1 + 2 + 3 + ... + 100)^2 = ", sumsquare)
    squaresum = 0
    index := 0
    for index < 100 {
	index++
	squaresum += math.Pow(float64(index), 2)
    }
    fmt.Println("1^2 + 2^2 + 3^2 + ... + 100^2 = ", squaresum)
    fmt.Println("Solution: ", int64(sumsquare - squaresum))
}

func (e *Euler) Problem_7(){
    var index int64 = 0
    nprimes := 0
    for {
	if eulerutils.IsPrime(index){
	    nprimes++
	    if nprimes == 10001{
	        break
	    }
	}
        index++
    }
    fmt.Println("The 10001th prime number is ", index)
}

func (e *Euler) Problem_8(){
    danumber := "7316717653133062491922511967442657474235534919493496983520312774506326239578318016984801869478851843858615607891129494954595017379583319528532088055111254069874715852386305071569329096329522744304355766896648950445244523161731856403098711121722383113622298934233803081353362766142828064444866452387493035890729629049156044077239071381051585930796086670172427121883998797908792274921901699720888093776657273330010533678812202354218097512545405947522435258490771167055601360483958644670632441572215539753697817977846174064955149290862569321978468622482839722413756570560574902614079729686524145351004748216637048440319989000889524345065854122758866688116427171479924442928230863465674813919123162824586178664583591245665294765456828489128831426076900422421902267105562632111110937054421750694165896040807198403850962455444362981230987879927244284909188845801561660979191338754992005240636899125607176060588611646710940507754100225698315520005593572972571636269561882670428252483600823257530420752963450"

    max := 0

    for i := 0; i < len(danumber)-13; i++ {
        pivot := int(danumber[i] - '0')
	for j := i+1; j < i+13; j++ {
	    pivot = int(danumber[j] - '0') * pivot
	    fmt.Print(danumber[j] - '0', " * ")
	}
	if pivot > max{
	    fmt.Println("MAX at position: ", i)
	    max = pivot
	}
	fmt.Println(" = ", pivot)
    }

    fmt.Println("the max product is ", max)
}

func (e *Euler) Problem_9(){
    for num_a := 1; num_a < 1000; num_a++{
        for num_b := num_a; num_b < 1000; num_b++{
	    num_c := math.Sqrt(math.Pow(float64(num_a), 2) + math.Pow(float64(num_b), 2))
	    //Check if whole number:
	    if float64(num_a) + float64(num_b) + num_c == 1000{
		fmt.Println(int64(float64(num_a) * float64(num_b) * num_c))
		return
	    }
	}	
    }
}

func (e *Euler) Problem_10(){
    var num, sum int64 = 2, 0
    
    for num = 2; num < 2000000; num++{
	if eulerutils.IsPrime(num){
	    sum += num
	}
    }
    fmt.Println(sum)
}

func (e *Euler) Problem_11(){
    values := [][]int{}
    values = append(values, []int{8, 2, 22, 97, 38, 15, 0, 40, 0, 75, 4, 5, 7, 78, 52, 12, 50, 77, 91 ,8})
    values = append(values, []int{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0})
    values = append(values, []int{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 4, 56, 62, 0}) 
    values = append(values, []int{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 3, 49, 13, 36, 65}) 
    values = append(values, []int{52, 70, 95, 23, 4, 60, 11, 42, 69, 24, 68, 56, 1, 32, 56, 71, 37, 2, 36, 91}) 
    values = append(values, []int{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80}) 
    values = append(values, []int{24, 47, 32, 60, 99, 03, 45, 2, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50}) 
    values = append(values, []int{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70}) 
    values = append(values, []int{67, 26, 20, 68, 2, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21}) 
    values = append(values, []int{24, 55, 58, 5, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72}) 
    values = append(values, []int{21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95}) 
    values = append(values, []int{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 3, 80, 4, 62, 16, 14, 9, 53, 56, 92}) 
    values = append(values, []int{16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57}) 
    values = append(values, []int{86, 56, 0, 48, 35, 71, 89, 07, 5, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58}) 
    values = append(values, []int{19, 80, 81, 68, 5, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 4, 89, 55, 40}) 
    values = append(values, []int{4, 52, 8, 83, 97, 35, 99, 16, 7, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66}) 
    values = append(values, []int{88, 36, 68, 87, 57, 62, 20, 72, 3, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69}) 
    values = append(values, []int{4, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36}) 
    values = append(values, []int{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 4, 36, 16}) 
    values = append(values, []int{20, 73, 35, 29, 78, 31, 90, 1, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 5, 54}) 
    values = append(values, []int{1, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 1, 89, 19, 67, 48}) 

    var max int64 = 0
    for i:=0; i<20; i++{
	for j:=0; j<20; j++{
	    //horizontal:
	    hor := 0
	    if j < 17{
            	hor = values[i][j]*values[i][j+1]*values[i][j+2]*values[i][j+3]
	    }
	    //vertical
	    ver := 0
	    if i < 17{
	        ver = values[i][j]*values[i+1][j]*values[i+2][j]*values[i+3][j]
	    }
	    //diagonalDownRight
	    ddr := 0
            if j < 17 && i<17{
		ddr = values[i][j]*values[i+1][j+1]*values[i+2][j+2]*values[i+3][j+3]
	    }
	    //diagonalUpRight
	    dur := 0
	    if j < 17 && i > 2{
		dur = values[i][j]*values[i-1][j+1]*values[i-2][j+2]*values[i-3][j+3]
	    }

	    if int64(ver) > max{
		max = int64(ver)
	    }
	    if int64(hor) > max{
		max = int64(hor)
	    }
	    if int64(ddr) > max{
		max = int64(ddr)
	    }
	    if int64(dur) > max{
		max = int64(dur)
	    }
	}
    }
    fmt.Println("Solution: ", max)
}


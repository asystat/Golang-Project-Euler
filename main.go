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

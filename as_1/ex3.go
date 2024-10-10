package main
import "fmt"

func main() {
  inp := 0
  
  if inp > 0 {
      fmt.Println("it is positive")
  } else if inp == 0 {
      fmt.Println("it is zero")
  } else {
      fmt.Println("it is negative")
  }
  
  sum := 0
  
  for i := 1; i <= 10; i++ {
      sum += i
  }
  fmt.Println("The sum is", sum)
  
  day := 5 
  
    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    case 4:
        fmt.Println("Thursday")
    case 5:
        fmt.Println("Friday")
    case 6:
        fmt.Println("Saturday")
    case 7:
        fmt.Println("Sunday")
    default:
        fmt.Println("Invalid day")
    }
}

// in Go the condition in if statement does not need parentheses, while in Java it is required. Whereas in Python we use colon

// basic
// for i := 1; i <= 10; i++ {
    // fmt.Println(i)
// }

// similar to while
// i := 0
// for i < 10 {
//     fmt.Println(i)
//     i++
// }

// infinite loop
// for {
    // fmt.Println("aaaasassasaasasasa")
// }

// every case in Go switch statement breaks automatically unlike in Python and C where you have to break the case manually


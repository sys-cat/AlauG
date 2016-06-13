package main
import (
  "github.com/dsparling/go-apache-log-parser"
  "fmt"
  "log"
  "regexp"
)

type All struct {
  v40 int
  v41 int
  v42 int
  v43 int
  v44 int
  v50 int
  v51 int
  v60 int
}

func main() {
  v60 := regexp.MustCompile(`Android 6.0.\d`)
  v51 := regexp.MustCompile(`Android 5.1.\d`)
  v50 := regexp.MustCompile(`Android 5.0.\d`)
  v44 := regexp.MustCompile(`Android 4.4.\d`)
  v43 := regexp.MustCompile(`Android 4.3.\d`)
  v42 := regexp.MustCompile(`Android 4.2.\d`)
  v41 := regexp.MustCompile(`Android 4.1.\d`)
  v40 := regexp.MustCompile(`Android 4.0.\d`)
  all := All{}
  lines, err := apachelogparser.Parse("./log/test.log")
  if err != nil {log.Fatal(err)}
  for _, line := range lines {
    if v40.MatchString(line.UserAgent) {
      //fmt.Println("Android 4.0.x")
      all.v40 += 1
    } else if v41.MatchString(line.UserAgent) {
      //fmt.Println("Android 4.1.x")
      all.v41 += 1
    } else if v42.MatchString(line.UserAgent) {
      //fmt.Println("Android 4.2.x")
      all.v42 += 1
    } else if v43.MatchString(line.UserAgent) {
      //fmt.Println("Android 4.3.x")
      all.v43 += 1
    } else if v44.MatchString(line.UserAgent) {
      //fmt.Println("Android 4.4.x")
      all.v44 += 1
    } else if v50.MatchString(line.UserAgent) {
      //fmt.Println("Android 5.0.x")
      all.v50 += 1
    } else if v51.MatchString(line.UserAgent) {
      //fmt.Println("Android 5.1.x")
      all.v51 += 1
    } else if v60.MatchString(line.UserAgent) {
      //fmt.Println("Android 6.0.x")
      all.v60 += 1
    }
  }
  fmt.Println("Android View Users")
  fmt.Printf("4.0.x: %d\n", all.v40)
  fmt.Printf("4.1.x: %d\n", all.v41)
  fmt.Printf("4.2.x: %d\n", all.v42)
  fmt.Printf("4.3.x: %d\n", all.v43)
  fmt.Printf("4.4.x: %d\n", all.v44)
  fmt.Printf("5.0.x: %d\n", all.v50)
  fmt.Printf("5.1.x: %d\n", all.v51)
  fmt.Printf("6.0.x: %d\n", all.v60)
}

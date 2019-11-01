package main 

import( "fmt"
"strings")

func findin(str string, h string) int {
  return strings.Index(h, str)
}
//next thing to work on
func findReplaceText(str string, h string) string {
  return "the text insert"
}

func main() {

var html = "jfkasdhfkjhsdafkj***insert***ldkfsajsdlkfj"
  fmt.Println(findin("***", html))

}

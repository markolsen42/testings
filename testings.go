package main 

import( "fmt"
"strings")

func findin(str string, h string) int {
  return strings.Index(h, str)
}
//next thing to work on
func findReplaceText(str string, h string) string {
	var splits = strings.Split(h, str);
	fmt.Print(splits)
	return splits[1]
}


func main() {

var html = "jfkasdhfkjhsdafkj***insert***ldkfsajsdlkfj***insert2***"
  fmt.Print(findReplaceText("***",html))

}

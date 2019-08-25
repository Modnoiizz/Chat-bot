package main
import "fmt"

func main(){
  //fmt.Println("Hello Word")
  str_first_var := "Success"
  str_second_var := "Fail"

  fmt.Println("=====================>Befor Swap<=====================")
  fmt.Println(str_first_var)
  fmt.Println(str_second_var)

  str_first_var = str_first_var + str_second_var
  str_second_var = string(str_first_var[0:len(str_first_var)-len(str_second_var)])
  str_first_var = string(str_first_var[len(str_second_var):len(str_first_var)])

  fmt.Println("=====================>After Swap<=====================")
  fmt.Println(str_first_var)
  fmt.Println(str_second_var)
}

package main 
import "fmt"
import "github.com/saad189/golangexamples"

func main(){
	fmt.Printf(golangexamples.EZGreetings("Hi\n"))
	arr := [] byte {'H','I','N','T','!'}
	fmt.Printf(golangexamples.ConcatSlice(arr)+"\n")
	fmt.Printf(golangexamples.ConcatSlice(golangexamples.Encrypt(arr,3)))
	fmt.Printf("\n")
	fmt.Printf(golangexamples.EZGreetings("Saad")+"\n")
	
}

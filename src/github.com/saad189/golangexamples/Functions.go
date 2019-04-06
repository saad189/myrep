package golangexamples
import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat []byte) string {
	var x string =""
	for i:=0; i< len(sliceToConcat); i++ {
		x+= string(sliceToConcat[i])+"-";
	}
	return x
}
func Encrypt(sliceToEncrypt []byte, ceaserCount int ) []byte{
	x := sliceToEncrypt	
	for i:=0; i<len(sliceToEncrypt); i++ {
		x[i]=byte((int(sliceToEncrypt[i]-'0')+ceaserCount)+'0');
	}
	return x
}
func EZGreetings(name string) string {
	return greetings.PrintGreetings(name)
}

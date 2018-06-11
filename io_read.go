package main
import	(
	"io/ioutil"
	"fmt"
	"time"
	"bufio"
	"os"
)
func main(){
	filedata,err := ioutil.ReadFile("./hola.txt")
	if err != nil {
		fmt.Println("no erro")
	}

	fmt.Println(string(filedata))

	otro,error := os.Open("./hola.txt")
	if error != nil {
		fmt.Println("no erroeeeerrrrr")
	}
	scanner := bufio.NewScanner(otro)
	for scanner.Scan(){
		linea := scanner.Text()
		time.Sleep(2 * time.Second)
		fmt.Println(linea)
	}
}

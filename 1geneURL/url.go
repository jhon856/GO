package main
import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
)

func createURL() string {
	u,err := url.Parse("/params")/// para url estaticas poner de una http://localhost:8000/params
	if err != nil{
		panic(err)
	}
	u.Host = "localhost:8000"
	u.Scheme = "http"

	//para pasar parametors
	//generar un map
	query := u.Query()
	query.Add("nombre","valor")
	u.RawQuery = query.Encode() // esto es para }
	return u.String()
}
func main() {
	url := createURL()
	fmt.Println("la url final es :"+ url)
	request,err := http.NewRequest("GET",url,nil)
	if err != nil{
		panic(err)
	}
	client :=&http.Client {}// revisar esto.. es el valor del puntero
	response,err1 := client.Do(request)
	if err1 != nil{
		panic(err1)
	}
	body,err1 := ioutil.ReadAll(response.Body)
	if err1 != nil{
		panic(err1)
	}
	fmt.Println("la url final es ", string(body))
	
	fmt.Println("la url final es :status", response.Status)
	fmt.Println("la url final es :status", response.Header)
	
}


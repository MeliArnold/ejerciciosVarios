package variables

import "fmt"

var temperatura int = 25
var humedad int = 36
var presion int = 12

func Clima() {

	respuesta := map[string]int{
		"temperatura": temperatura,
		"humedad":     humedad,
		"presion":     presion,
	}
	/*	fmt.Println("la temperatura de hoy es: ", temperatura)
		fmt.Println("la humedad de hoy es: ", humedad)
		fmt.Println("la presion de hoy es: ", presion)*/
	fmt.Println(respuesta)

}

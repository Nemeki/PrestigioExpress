package camiones
import "math/rand"

type Camion struct {
	Tipo int
	Estado int
	Pendientes int
	IDPaquetes [2]int
	Valorpaquetes [2]int
	Direcciones [2]string
	var Historial []Registro
	Ultimo int
}

func (c *Camion) elegir_paquete(i int){
	var ID = 0
	if i == 0{
		if (*c).Valorpaquetes[0] > (*c).Valorpaquetes[1]{
			ID = (*c).Valorpaquetes[0]
			(*c).Ultimo = 1
		}
		else{
			ID = (*c).Valorpaquetes[1]
			(*c).Ultimo = 2
		}
	}
	else{
		if (*c).Ultimo == 1{
			ID = (*c).Valorpaquetes[1]
			(*c).Ultimo = 2
		}
		else{
			ID = (*c).Valorpaquetes[0]
			(*c).Ultimo = 1
		}
	}
	(*c).entregar(ID)
}

func (c *Camion) entregar(ID int){
	var r = rand.Intn(100)
	entregado = false
	if (r >=20){
		entregado = true
		(*c)Pendientes = (*c)Pendientes-1
		(*c).Ultimo = 0
	}
	//mensaje(paquete, entregado)
}

func (c *Camion) rutina():
	switch (*c).pendiente{
	case 0:
		(*c).Estado = 0
		//Solicitar paquetes
		//Esperar un paquete inicial indefinidamente
	case 1:
		if (*c).Estado == 0{
		//Solicitar paquete
		//Esperar un paquete inicial un tiempo definido
		}
		else{
			(*c).elegir_paquete(0)
		}
	case 2:
		(*c).elegir_paquete((*c).Ultimo)
	}
}
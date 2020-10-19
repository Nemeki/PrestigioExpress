package paquetes

type Paquete struct {
	IDPaquete int
	Seguimiento int
	Tipo int
	Valor int
	Intentos int
	Estado int
}

func (p *Paquete) Actulizar_intentos(){
	//Funcion que actuliza los intentos
	(*p).Intentos = Intentos + 1
	max_intentos = (*p).intentos_maximos()
	if (*p).Intentos == max_intentos{
		Terminar_paquete(0)
	}
}

func (p *Paquete) Terminar_paquete(entregado bool){
	actualizar_carga(camion, entregado)
	ganancia = calcular_costo(entregado)
	if entregado{
		Estado = 2
	}
	else{
		Estado = 3
	}
}

func (p Paquete) intentos_maximos(){
	var intentos_permitidos = 0
	if 30 > p.Valor{
		intentos = p.Valor%10
	}
	else{
		intentos = 3
	}
	return intentos
}
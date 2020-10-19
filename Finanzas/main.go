package main
import registro

func calcular_costo(entregado bool, Valor int, tipo int){
	var ganancia int
	//tipo 1-retail, 2-normal, 3-prioritario
	if tipo == 1{
		ganancia = Valor
	}
	else if entregado & tipo == 3{
		ganancia = 1.3*Valor
	}
	else if tipo == 2{
		ganancia = 0
	}
	else{
		ganancia = 0.3*Valor
	}
	return ganancia
}

func actualizar_balance(ganancia int, Intentos int){
	Perdidas = Perdidas + Intentos*10
	Ganancias = Gancancias + ganancia
}

func registrar(int id, int intentos, int Valor, bool entregado, int Tipo){
	ganancia = calcular_costo(entregado, Valor, tipo)
	actualizar_balance(ganancia, intentos)
	r := registro{
		IDPaquete: id,
		Intentos: intentos,
		Ganancia: ganancia,
		Entregado: entregado,
	}
}

func main(){
	//esperando recibir mensajes
	//mensaje con información de un paquete
	registrar(id, intentos, Valor, entregado, Tipo)
}
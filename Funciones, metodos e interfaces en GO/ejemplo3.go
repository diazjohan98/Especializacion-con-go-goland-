package main

import (
	"fmt"
	"time"
)

// Esta es la tarea lenta (la cocina)
func cocinarHamburguesa(mesa string) {
	fmt.Println("👨‍🍳 Empezando pedido para", mesa)
	// time.Sleep simula que la tarea demora 2 segundos
	time.Sleep(2 * time.Second)
	fmt.Println("✅ ¡Lista la hamburguesa de la", mesa, "!")
}

func main() {
	fmt.Println("🚪 Restaurante abierto. Mesero listo.")

	// Usamos 'go' para mandar las tareas al fondo sin bloquear al mesero
	go cocinarHamburguesa("Mesa 1")
	go cocinarHamburguesa("Mesa 2")
	go cocinarHamburguesa("Mesa 3")

	fmt.Println("📝 El mesero ya tomó los 3 pedidos y está libre al instante.")

	// Le decimos al programa principal que no se cierre todavía,
	// para darle tiempo a la cocina de terminar.
	time.Sleep(3 * time.Second)
	fmt.Println("🚪 Restaurante cerrado.")
}

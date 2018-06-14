// main
package main

import (
	"fmt"
)

func Шапка_Печать() {
	fmt.Println("\n\033[1m   =[ Программа учёта \"РАЙАВТОДОР\" ]=\n")

}

func main() {
	var цДейств int = -1
мЦикл:
	for {
		Шапка_Печать()
		fmt.Println("\nДействие:")
		fmt.Println("   0) Выход из программы")
		fmt.Println("   1) Создать новую базу")
		fmt.Println("   2) Открыть существующую базу")
		fmt.Print("Действие > ")
		fmt.Scanf("%d", &цДейств)
		switch цДейств {
		case 0:
			fmt.Println("Конец программы")
			break мЦикл
		case 1:
			fmt.Println("Создание новой базы")
		case 2:
			fmt.Println("Открытие новой базы")
		default:
			fmt.Println("Неизвестное значение")
		}
	}

}

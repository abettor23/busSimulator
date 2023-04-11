package main

import (
	"fmt"
)

func main() {
	fmt.Println("Перед тобой симулятор маршрутки. v.1.0")

	busFullness := 0
	wallet := 0
	ticket := 20
	var passengers int

	fmt.Println("Прибыли на остановку \"Почта\".")
	fmt.Println("В салоне", busFullness, "человек.")
	fmt.Println("Сколько человек зашло на этой остановке?")
	fmt.Scan(&passengers)
	busFullness += passengers
	wallet += passengers * ticket
	fmt.Println("Отправляемся с остановки \"Почта\".")

	fmt.Println("Прибыли на остановку \"Проспект мира\".")
	fmt.Println("В салоне", busFullness, "человек.")
	fmt.Println("Сколько человек вышло на этой остановке?")
	fmt.Scan(&passengers)
	busFullness -= passengers
	fmt.Println("Сколько человек зашло на этой остановке?")
	fmt.Scan(&passengers)
	busFullness += passengers
	wallet += passengers * ticket
	fmt.Println("Отправляемся с остановки \"Проспект мира\".")

	fmt.Println("Прибыли на остановку \"Снегири\".")
	fmt.Println("В салоне", busFullness, "человек.")
	fmt.Println("Сколько человек вышло на этой остановке?")
	fmt.Scan(&passengers)
	busFullness -= passengers
	fmt.Println("Сколько человек зашло на этой остановке?")
	fmt.Scan(&passengers)
	busFullness += passengers
	wallet += passengers * ticket
	fmt.Println("Отправляемся с остановки \"Снегири\".")

	fmt.Println("Прибыли на конечную остановку \"Тупик\".")
	fmt.Println("В салоне", busFullness, "человек.")
	fmt.Println("Конечная, все пассажиры вышли из салона.")
	//busFullness := passengers

	driverSalory := wallet / 4
	fuelBill := wallet / 5
	taxes := wallet / 5
	busRepair := wallet / 5
	profit := wallet - (driverSalory + fuelBill + taxes + busRepair)

	fmt.Println("Всего заработано:", wallet)
	fmt.Println("Зарплата водителю:", driverSalory)
	fmt.Println("Расходы на топливо:", fuelBill)
	fmt.Println("Налоги:", taxes)
	fmt.Println("Ремонт маршрутки:", busRepair)
	fmt.Println("Итого доход:", profit)

}

/* Замечания ChatGPT

    Код выглядит вполне рабочим и функциональным, но есть несколько замечаний:

    В строке 8 нужно убрать лишний пробел перед "busFullness".

    В строке 47 есть закомментированный код, который можно удалить.

    Нет проверки на то, что введенное пользователем значение является корректным.

	Например, что бы произошло, если пользователь введет отрицательное число?

    Нет возможности для пользователя увидеть, сколько стоит билет.

	Можно добавить информацию о цене билета в начало программы.

    Можно добавить больше комментариев, чтобы код был более понятен для других разработчиков.

    Разделение переменных на passengersIn и passengersOut может упростить код и повысить его читабельность,

	так как это позволяет явно указать количество пассажиров, которые зашли и вышли на каждой остановке.

	Однако, в данном случае, можно обойтись без такого разделения, так как пассажиры, которые выходят,

	могут быть вычтены из общего количества пассажиров на маршрутке (busFullness) с помощью оператора вычитания.

	Таким образом, в данном коде принцип подсчета пассажиров (passengers) реализован достаточно корректно, и разделение

	переменных не является обязательным. */

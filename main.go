package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var input string //буфер клавиатурного ввода

type card struct { //Тип данных "Карта": имеет имя, стоимость в очках (value) и статус: 0 - в колоде, 1 - в руке игрока, 2 - в руке дилера
	name          string //имя
	status, value int    //статус, стоимость в очках
}

type gameStatus struct { //Содержит состояние игры: счёт, расположение карт и т.д.
	mu         sync.Mutex
	plScore    int          //начальный счет игрока
	cpScore    int          //начальный счет дилера
	plAces     int          //количество тузов у игрока
	cpAces     int          //количество тузов у дилера
	cardNumber int          //случайный номер для выбора карты из колоды
	cardName   string       //он же в виде текстовой строки для обращения к map
	ace        bool         //по умолчанию выбранная карта - не туз
	plFinished bool         //ни одно из условий для завершения хода игрока ещё не наступило
	cpFinished bool         //ни одно из условий для завершения хода дилера ещё не наступило
	Deck       map[int]card // Колода именованных карт. Имена цифровые, чтобы можно было выбрать карту по генератору случайных чисел
}

func (g *gameStatus) Pick(score int) (new_score, cardNumber int, ace bool) { //функция выбирает из колоды случайную карту, не находящуюся ни в чьей руке и добавляет её стоимость к счёту. Также функция собщает, является ли карта тузом
	status := 3       //по умолчанию, статус "ошибка"
	value := 0        //стоимость по умолчанию "ошибка"
	ace = false       // по умолчанию выбранная карта - не туз
	for status != 0 { //до выбора карты, не находящейся ни в чьей руке
		cardNumber = rand.Intn(52) //случайный номер карты
		g.mu.Lock()
		value = g.Deck[cardNumber].value   //стоимость карты
		status = g.Deck[cardNumber].status //статус карты
		g.mu.Unlock()
		if value == 11 { //если карта является тузом
			ace = true //возвращаем туз == ИСТИНА
		}
	}
	fmt.Println(g.Deck[cardNumber].name) //выводим на экран выбранную карту
	new_score = score + value            //прибавляем стоимость карты к счёту выбранного игрока
	return new_score, cardNumber, ace    // возвращает обновлённый счёт
}

func (g *gameStatus) Gameover() string { //определяет результат игры, возвращая его в виде текстовой строки
	if g.plScore == g.cpScore { //счёт равный -> ничья
		return "draw"
	}
	if g.cpScore > 21 { //перебор у дилера -> игрок победил
		return "player win"
	}
	if g.cpScore > g.plScore && g.cpScore < 22 { //счёт дилера выше чем у игрока, но ниже 22 -> компьютер победил
		return "computer win"
	}
	if g.cpScore < g.plScore && g.plScore < 22 { //счёт игрока выше чем у дилера, но ниже 22 -> игрок победил
		return "player win"
	}
	return "computer win" // перебор у игрока -> компьютер победил
}

func init() {
	rand.Seed(time.Now().UnixNano()) //карты будут вытаскиваться в соответствии со сгенерированной случайной последовательностью
}

func main() {

	g := gameStatus{ //Состояние игры, с объявлением начальных значений - счёта, количества тузов у игроков, состояния колоды и т.д.
		mu:         sync.Mutex{},
		plScore:    0,     //начальный счет игрока
		cpScore:    0,     //начальный счет дилера
		plAces:     0,     //количество тузов у игрока
		cpAces:     0,     //количество тузов у дилера
		cardNumber: 0,     //случайный номер для выбора карты из колоды
		cardName:   "",    //он же в виде текстовой строки для обращения к map
		ace:        false, //по умолчанию выбранная карта - не туз
		plFinished: false, //ни одно из условий для завершения хода игрока ещё не наступило
		cpFinished: false, //ни одно из условий для завершения хода дилера ещё не наступило
		Deck: map[int]card{ // Колода именованных карт. Имена цифровые, чтобы можно было выбрать карту по генератору случайных чисел
			1: {"hearts02", 0, 2}, 2: {"hearts03", 0, 3}, 3: {"hearts04", 0, 4}, 4: {"hearts05", 0, 5}, 5: {"hearts06", 0, 6}, 6: {"hearts07", 0, 7}, 7: {"hearts08", 0, 8}, 8: {"hearts09", 0, 9}, 9: {"hearts10", 0, 10}, 10: {"heartsJack", 0, 10}, 11: {"heartsQueen", 0, 10}, 12: {"heartsKing", 0, 10}, 13: {"heartsAce", 0, 11}, 14: {"spades02", 0, 2}, 15: {"spades03", 0, 3}, 16: {"spades04", 0, 4}, 17: {"spades05", 0, 5}, 18: {"spades06", 0, 6}, 19: {"spades07", 0, 7}, 20: {"spades08", 0, 8}, 21: {"spades09", 0, 9}, 22: {"spades10", 0, 10}, 23: {"spadesJack", 0, 10}, 24: {"spadesQueen", 0, 10}, 25: {"spadesKing", 0, 10}, 26: {"spadesAce", 0, 11}, 27: {"diamonds02", 0, 2}, 28: {"diamonds03", 0, 3}, 29: {"diamonds04", 0, 4}, 30: {"diamonds05", 0, 5}, 31: {"diamonds06", 0, 6}, 32: {"diamonds07", 0, 7}, 33: {"diamonds08", 0, 8}, 34: {"diamonds09", 0, 9}, 35: {"diamonds10", 0, 10}, 36: {"diamondsJack", 0, 10}, 37: {"diamondsQueen", 0, 10}, 38: {"diamondsKing", 0, 10}, 39: {"diamondsAce", 0, 11}, 40: {"clubs02", 0, 2}, 41: {"clubs03", 0, 3}, 42: {"clubs04", 0, 4}, 43: {"clubs05", 0, 5}, 44: {"clubs06", 0, 6}, 45: {"clubs07", 0, 7}, 46: {"clubs08", 0, 8}, 47: {"clubs09", 0, 9}, 48: {"clubs10", 0, 10}, 49: {"clubsJack", 0, 10}, 50: {"clubsQueen", 0, 10}, 51: {"clubsKing", 0, 10}, 0: {"clubsAce", 0, 11}},
	}

	fmt.Println("Type hit or stand") //ход игрока. Игроку доступны две команды: ЕЩЁ (hit) и СЕБЕ (stand)
	for !g.plFinished {              //игрок завершит ход при вводе команды "СЕБЕ" ИЛИ в случае перебора (<21 очков)
		//g.mu.Lock()
		fmt.Scanf("%s\n", &input) //ожидаем действие игрока
		if input == "hit" {       //игрок ввёл команду "ещё"
			fmt.Println("You picking the card:")
			g.plScore, g.cardNumber, g.ace = g.Pick(g.plScore) //новый счёт, карта, является ли карта тузом = Pick(счёт игрока на данный момент)
			g.mu.Lock()
			g.Deck[g.cardNumber] = card{g.cardName, 1, 0} //присваиваем карте статус "в руке игрока" чтобы больше она не вытаскивалась
			g.mu.Unlock()
			if g.ace { //если карта оказалась тузом:
				g.plAces++ // запомним, что у игрока +1 туз
			}
			if g.plScore > 21 && g.plAces > 0 { //ситуация перебора: если у игрока остались тузы, которые можно объявить единицей, то
				g.plScore -= 10 //уменьшаем очки на 10 и продолжаем
				g.plAces--      //мы использовали один туз для погашения перебора
			}
			fmt.Println("Your score is:", g.plScore)
			if g.plScore > 20 { //при счёте более 20 дальнейшее взятие карт невозможно
				g.plFinished = true //автоматический переход хода
			}
			//g.mu.Unlock()
		}
		if input == "stand" { //игрок ввёл команду "себе"
			g.plFinished = true //переход хода по команде игрока
		}
	}

	fmt.Println("Ход дилера") //ход дилера

	for !g.cpFinished { //дилер завершит ход по достижении 18 очков ИЛИ в случае перебора (<21 очков)
		fmt.Println("I picking the card:")
		g.cpScore, g.cardNumber, g.ace = g.Pick(g.cpScore) //новый счёт, карта, является ли карта тузом = Pick(счёт дилера на данный момент)
		g.mu.Lock()
		g.Deck[g.cardNumber] = card{g.cardName, 2, 0} //присваиваем карте статус "в руке дилера" чтобы она более не вытаскивалась
		g.mu.Unlock()
		if g.ace { //если карта оказалась тузом:
			g.cpAces++ // запомним, что у дилера +1 туз
		}
		if g.cpScore > 21 && g.cpAces > 0 { //ситуация перебора: если у дилера остались тузы, которые можно объявить единицей, то
			g.cpScore -= 10 //уменьшаем очки на 10 и продолжаем
			g.cpAces--      //мы использовали один туз для погашения перебора
		}
		fmt.Println("My score is:", g.cpScore)
		if g.cpScore > 16 { //счёт дилера достиг 17
			g.cpFinished = true //завершение игры
		}

	}
	fmt.Println(g.Gameover()) //вывод результата игры
}

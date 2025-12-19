package console

import (
"bufio"
"fmt"
"os"
"strconv"
"strings"

"moscow-zoo/domain"
"moscow-zoo/services"
)

// App - консольное приложение
type App struct {
	zoo    services.IZoo
	reader *bufio.Reader
}

// NewApp создает приложение
func NewApp(zoo services.IZoo) *App {
	return &App{
		zoo:    zoo,
		reader: bufio.NewReader(os.Stdin),
	}
}

// Run запускает приложение
func (a *App) Run() {
	fmt.Println("========================================")
	fmt.Println("  ERP-система Московского Зоопарка")
	fmt.Println("========================================")
	fmt.Println()

	for {
		a.showMenu()
		choice := a.readInput("Выберите пункт меню: ")

		switch choice {
		case "1":
			a.addAnimal()
		case "2":
			a.addThing()
		case "3":
			a.showAnimalsReport()
		case "4":
			a.showFoodConsumption()
		case "5":
			a.showContactZooAnimals()
		case "6":
			a.showInventory()
		case "0":
			fmt.Println("\nДо свидания!")
			return
		default:
			fmt.Println("\nНеверный выбор. Попробуйте снова.")
		}
		fmt.Println()
	}
}

func (a *App) showMenu() {
	fmt.Println("--- ГЛАВНОЕ МЕНЮ ---")
	fmt.Println("1. Добавить животное")
	fmt.Println("2. Добавить вещь на баланс")
	fmt.Println("3. Отчет по животным")
	fmt.Println("4. Общее потребление еды")
	fmt.Println("5. Животные для контактного зоопарка")
	fmt.Println("6. Инвентаризационная ведомость")
	fmt.Println("0. Выход")
}

func (a *App) addAnimal() {
	fmt.Println("\n=== Добавление животного ===")
	fmt.Println("Выберите тип:")
	fmt.Println("1. Обезьяна (травоядное)")
	fmt.Println("2. Кролик (травоядное)")
	fmt.Println("3. Тигр (хищник)")
	fmt.Println("4. Волк (хищник)")

	typeChoice := a.readInput("Тип: ")

	foodStr := a.readInput("Потребление еды (кг/сутки): ")
	food, err := strconv.Atoi(foodStr)
	if err != nil || food < 0 {
		fmt.Println("Ошибка: некорректное значение")
		return
	}

	numberStr := a.readInput("Инвентарный номер: ")
	number, err := strconv.Atoi(numberStr)
	if err != nil || number < 0 {
		fmt.Println("Ошибка: некорректный номер")
		return
	}

	healthyStr := a.readInput("Здорово? (да/нет): ")
	healthy := strings.ToLower(healthyStr) == "да"

	var animal domain.IAnimal

	switch typeChoice {
	case "1":
		kindness := a.readKindness()
		animal = domain.NewMonkey(food, number, healthy, kindness)
	case "2":
		kindness := a.readKindness()
		animal = domain.NewRabbit(food, number, healthy, kindness)
	case "3":
		animal = domain.NewTiger(food, number, healthy)
	case "4":
		animal = domain.NewWolf(food, number, healthy)
	default:
		fmt.Println("Неверный тип")
		return
	}

	if a.zoo.AddAnimal(animal) {
		fmt.Printf("Животное '%s' добавлено!\n", animal.GetName())
	} else {
		fmt.Printf("Животное '%s' не прошло ветосмотр.\n", animal.GetName())
	}
}

func (a *App) readKindness() int {
	kindnessStr := a.readInput("Уровень доброты (0-10): ")
	kindness, err := strconv.Atoi(kindnessStr)
	if err != nil {
		return 0
	}
	return kindness
}

func (a *App) addThing() {
	fmt.Println("\n=== Добавление вещи ===")
	fmt.Println("1. Стол")
	fmt.Println("2. Компьютер")

	typeChoice := a.readInput("Тип: ")

	numberStr := a.readInput("Инвентарный номер: ")
	number, err := strconv.Atoi(numberStr)
	if err != nil || number < 0 {
		fmt.Println("Ошибка: некорректный номер")
		return
	}

	var thing domain.IInventory

	switch typeChoice {
	case "1":
		thing = domain.NewTable(number)
	case "2":
		thing = domain.NewComputer(number)
	default:
		fmt.Println("Неверный тип")
		return
	}

	a.zoo.AddThing(thing)
	fmt.Printf("Вещь '%s' добавлена!\n", thing.GetName())
}

func (a *App) showAnimalsReport() {
	fmt.Println("\n=== ОТЧЕТ ПО ЖИВОТНЫМ ===")

	animals := a.zoo.GetAnimals()
	if len(animals) == 0 {
		fmt.Println("В зоопарке нет животных.")
		return
	}

	fmt.Printf("Всего животных: %d\n\n", len(animals))
	fmt.Printf("%-4s %-15s %-12s %-15s\n", "N", "Название", "Инв.номер", "Еда(кг/сутки)")
	fmt.Println(strings.Repeat("-", 50))

	for i, animal := range animals {
		fmt.Printf("%-4d %-15s %-12d %-15d\n",
i+1, animal.GetName(), animal.GetNumber(), animal.GetFood())
	}
}

func (a *App) showFoodConsumption() {
	fmt.Println("\n=== ПОТРЕБЛЕНИЕ ЕДЫ ===")
	total := a.zoo.GetTotalFoodConsumption()
	count := a.zoo.GetAnimalsCount()

	fmt.Printf("Количество животных: %d\n", count)
	fmt.Printf("Общее потребление: %d кг/сутки\n", total)

	if count > 0 {
		fmt.Printf("Среднее: %.2f кг/сутки\n", float64(total)/float64(count))
	}
}

func (a *App) showContactZooAnimals() {
	fmt.Println("\n=== КОНТАКТНЫЙ ЗООПАРК ===")
	fmt.Println("(Травоядные с добротой > 5)")

	animals := a.zoo.GetContactZooAnimals()
	if len(animals) == 0 {
		fmt.Println("Нет подходящих животных.")
		return
	}

	fmt.Printf("\nНайдено: %d\n\n", len(animals))
	fmt.Printf("%-4s %-15s %-12s %-10s\n", "N", "Название", "Инв.номер", "Доброта")
	fmt.Println(strings.Repeat("-", 45))

	for i, animal := range animals {
		fmt.Printf("%-4d %-15s %-12d %-10d\n",
i+1, animal.GetName(), animal.GetNumber(), animal.GetKindnessLevel())
	}
}

func (a *App) showInventory() {
	fmt.Println("\n=== ИНВЕНТАРИЗАЦИЯ ===")

	inventory := a.zoo.GetAllInventory()
	if len(inventory) == 0 {
		fmt.Println("На балансе нет объектов.")
		return
	}

	fmt.Printf("Всего объектов: %d\n\n", len(inventory))
	fmt.Printf("%-4s %-20s %-12s\n", "N", "Наименование", "Инв.номер")
	fmt.Println(strings.Repeat("-", 40))

	for i, item := range inventory {
		fmt.Printf("%-4d %-20s %-12d\n",
i+1, item.GetName(), item.GetNumber())
	}
}

func (a *App) readInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := a.reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Правила іменування
// Поганий приклад (використовуються беззмістовні імена)
var d int
func GetThem(theList [][]int) [][]int { /* ... */ return nil }

// Гарний приклад
var elapsedTimeInDays int
func GetFlaggedCells(gameBoard [][]int) [][]int { /* ... */ return nil }


// Правила написання функцій
// Поганий приклад (функція виконує кілька операцій на різних рівнях абстракції)
func ProcessOrder(order Order) {
    if order.Quantity <= 0 {
        return
    }
    total := float64(order.Quantity) * 15.5
    db.Save(order.ID, total)
}

// Гарний приклад 
func ProcessOrder(order Order) {
    if !isValid(order) {
        return
    }
    saveToDatabase(order.ID, calculateTotal(order))
}


// Правила аргументів-прапорців
// Поганий приклад (Використовується прапорець, виклик виглядає як render(true))
func Render(pageData PageData, isSuite bool) {
    if isSuite {
        // логіка підготовки для набору тестів
    } else {
        // логіка підготовки для одного тесту
    }
}

// Гарний приклад
func RenderForSuite(pageData PageData) {
    // логіка підготовки для набору тестів
}

func RenderForSingleTest(pageData PageData) {
    // логіка підготовки для одного тесту
}
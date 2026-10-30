// Виділення функції
// Код до рефакторингу
func Total(qty int, price float64) float64 {
    subtotal := float64(qty) * price
    return subtotal * 1.2
}

// Код після рефакторингу
func Subtotal(qty int, price float64) float64 {
    return float64(qty) * price
}

func Total(qty int, price float64) float64 {
    return Subtotal(qty, price) * 1.2
}


// Усунення дублювання (DRY)
// Код до рефакторингу
func AdultPrice(p float64) float64 {
    return p – p * 0.1
}

func StudentPrice(p float64) float64 {
    return p – p * 0.2
}

// Код після рефакторингу
func WithDiscount(p, rate float64) float64 {
    return p – p * rate
}

func AdultPrice(p float64) float64 {
    return WithDiscount(p, 0.1)
}

func StudentPrice(p float64) float64 {
    return WithDiscount(p, 0.2)
}


// Принцип розділення команд і запитів
// Код до рефакторингу
func (c *Counter) Next() int {
    c.value++
    return c.value
}

// Код після рефакторингу
type Counter struct { value int }
func (c *Counter) Increment() {
    c.value++
}

func (c *Counter) Value() int {
    return c.value
}
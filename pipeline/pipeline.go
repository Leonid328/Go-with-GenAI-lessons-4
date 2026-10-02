// Package pipeline реалізує Частину 1 домашньої роботи: простий
// послідовний конвеєр обробки даних за моделлю CSP.
//
// Generate виробляє випадкові числа, Filter пропускає лише парні.
// Обидві функції використовують спрямовані типи каналів у
// сигнатурах (Завдання 1.2) — це саме той стиль API, який ми
// проходили на занятті.
package pipeline

import "math/rand"

// Generate запускає горутину, що генерує рівно n випадкових цілих
// чисел у діапазоні [1, 100] і надсилає їх у повернутий канал.
// Після надсилання n-го числа канал має бути закритий.
func Generate(n int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 0; i < n; i++ {
			out <- rand.Intn(100) + 1
		}
	}()
	return out
}

// Filter приймає числа з in і пропускає в повернутий канал лише
// парні значення. Коли in закривається і вичерпується, Filter
// закриває свій вихідний канал.
func Filter(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if v%2 == 0 {
				out <- v
			}
		}
	}()
	return out
}

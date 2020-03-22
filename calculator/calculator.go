package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"unicode"
)

// (2*300)+30/7*1-(16*3)
// 2 300 * 30 7 / 1 * + 16 3 * -

// 3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3
// 3 4 2 * 1 5 - 2 3 ^ ^ / +

// 3.1415926

// 2^8

// Stack .
type Stack struct {
	list list.List
}

// Len .
func (s *Stack) Len() int {
	return s.list.Len()
}

// PushFront .
func (s *Stack) PushFront(str string) {
	s.list.PushFront(str)
}

// PushBack .
func (s *Stack) PushBack(str string) {
	s.list.PushBack(str)
}

// Show .
func (s *Stack) Show() {
	fmt.Print("stack: ")
	for e := s.list.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value.(string), " ")
	}
	fmt.Println()
}

// PeekBack .
func (s *Stack) PeekBack() *list.Element {
	return s.list.Back()
}

// PeekFront .
func (s *Stack) PeekFront() *list.Element {
	return s.list.Front()
}

// PopFront .
func (s *Stack) PopFront() *list.Element {
	e := s.list.Front()
	if e != nil {
		s.list.Remove(e)
	}
	return e
}

// PopBack .
func (s *Stack) PopBack() *list.Element {
	e := s.list.Back()
	if e != nil {
		s.list.Remove(e)
	}
	return e
}

// Remove .
func (s *Stack) Remove(e *list.Element) {
	s.list.Remove(e)
}

// Clear .
func (s *Stack) Clear() {
	s.list.Init()
}

// GetSymbolOrder .
func GetSymbolOrder(str string) int {
	priority := -1
	switch str {
	case "(":
		priority = 0

	case "+":
		priority = 1
	case "-":
		priority = 1

	case "*":
		priority = 2
	case "/":
		priority = 2
	case "%":
		priority = 2

	case "^":
		priority = 3

	}
	return priority
}

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", log.Lshortfile)
}

func main() {

	r := bufio.NewReader(os.Stdin)
	stack, postfix := Stack{}, Stack{}

	for {
		str, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}

		if str == "exit\n" {
			break
		}

		number := ""
		for _, r := range str {
			char := string(r)
			logger.Println("char       ", char)
			postfix.Show()
			stack.Show()
			println()

			if char == "\n" {
				if number != "" {
					postfix.PushBack(number)
					number = ""
				}
				break
			}

			if char == " " {
				if number != "" {
					postfix.PushBack(number)
					number = ""
				}
				continue
			}

			if unicode.IsDigit(r) == true || char == "." {
				number += char
				continue
			}

			if unicode.IsGraphic(r) {
				if number != "" {
					postfix.PushBack(number)
					number = ""
				}

				if char == "(" {

					stack.PushBack(char)
					continue
				}

				if char == ")" {
					for {
						e := stack.PopBack()
						if e != nil {
							symbol := e.Value.(string)
							if symbol != "(" {
								postfix.PushBack(symbol)
							} else {
								break
							}
						} else {
							break
						}
					}
					continue
				}

				for stack.Len() != 0 {
					charPriority := GetSymbolOrder(char)
					symbolPriority := GetSymbolOrder(stack.PeekBack().Value.(string))
					if char != "^" && charPriority <= symbolPriority || char == "^" && charPriority < symbolPriority {
						symbol := stack.PopBack().Value.(string)
						postfix.PushBack(symbol)
					} else {
						break
					}
				}
				stack.PushBack(char)
				continue
			}
		}

		for stack.Len() != 0 {
			symbol := stack.PopBack().Value.(string)
			postfix.PushBack(symbol)
		}

		if number != "" {
			postfix.PushBack(number)
			number = ""
		}

		postfix.Show()

		// 3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3
		// 3 4 2 * 1 5 - 2 3 ^ ^ / +

		// prevElements := []*list.Element{}
		// for e := postfix.PopFront(); e != nil; e = e.Next() {
		// 	temp := e.Value.(string)
		// 	switch temp {
		// 	case "+":
		// 		l := len(prevElements)
		// 		e1, e2 := prevElements[l-2], prevElements[l-1]
		// 		n1, _ := strconv.ParseFloat(e1.Value.(string), 64)
		// 		n2, _ := strconv.ParseFloat(e2.Value.(string), 64)

		// 		prevElements = prevElements[:l-2]

		// 		var e3 *list.Element
		// 		e3.Value = fmt.Sprintf("%f", n1+n2)
		// 		prevElements = append(prevElements, e)

		// 	case "-":
		// 		l := len(prevElements)
		// 		e1, e2 := prevElements[l-2], prevElements[l-1]
		// 		n1, _ := strconv.ParseFloat(e1.Value.(string), 64)
		// 		n2, _ := strconv.ParseFloat(e2.Value.(string), 64)

		// 		prevElements = prevElements[:l-2]

		// 		var e3 *list.Element
		// 		e3.Value = fmt.Sprintf("%f", n1-n2)
		// 		prevElements = append(prevElements, e)

		// 	case "^":
		// 		l := len(prevElements)
		// 		e1, e2 := prevElements[l-2], prevElements[l-1]
		// 		n1, _ := strconv.ParseFloat(e1.Value.(string), 64)
		// 		n2, _ := strconv.ParseFloat(e2.Value.(string), 64)

		// 		prevElements = prevElements[:l-2]

		// 		var e3 *list.Element
		// 		e3.Value = fmt.Sprintf("%f", math.Pow(n1, n2))
		// 		prevElements = append(prevElements, e)

		// 	case "*":
		// 		l := len(prevElements)
		// 		e1, e2 := prevElements[l-2], prevElements[l-1]
		// 		n1, _ := strconv.ParseFloat(e1.Value.(string), 64)
		// 		n2, _ := strconv.ParseFloat(e2.Value.(string), 64)

		// 		prevElements = prevElements[:l-2]

		// 		var e3 *list.Element
		// 		e3.Value = fmt.Sprintf("%f", n1*n2)
		// 		prevElements = append(prevElements, e)

		// 	case "/":
		// 		l := len(prevElements)
		// 		e1, e2 := prevElements[l-2], prevElements[l-1]
		// 		n1, _ := strconv.ParseFloat(e1.Value.(string), 64)
		// 		n2, _ := strconv.ParseFloat(e2.Value.(string), 64)

		// 		prevElements = prevElements[:l-2]

		// 		var e3 *list.Element
		// 		e3.Value = fmt.Sprintf("%f", n1/n2)
		// 		prevElements = append(prevElements, e)

		// 	default:
		// 		prevElements = append(prevElements, e)
		// 	}
		// }
		// fmt.Println("answer:", prevElements[0].Value)

		stack.Clear()
		postfix.Clear()
	}

	// fmt.Println("unicode.IsGraphic(+)", unicode.IsGraphic('+'))
	// fmt.Println("unicode.IsGraphic(-)", unicode.IsGraphic('-'))
	// fmt.Println("unicode.IsGraphic(*)", unicode.IsGraphic('*'))
	// fmt.Println("unicode.IsGraphic(/)", unicode.IsGraphic('/'))
	// fmt.Println("unicode.IsGraphic(^)", unicode.IsGraphic('/'))

	// fmt.Println("unicode.IsSymbol(+)", unicode.IsSymbol('+'))
	// fmt.Println("unicode.IsSymbol(-)", unicode.IsSymbol('-'))
	// fmt.Println("unicode.IsSymbol(*)", unicode.IsSymbol('*'))
	// fmt.Println("unicode.IsSymbol(/)", unicode.IsSymbol('/'))
	// fmt.Println("unicode.IsSymbol(^)", unicode.IsSymbol('/'))
}

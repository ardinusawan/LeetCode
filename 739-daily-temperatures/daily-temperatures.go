import (
    "errors"
    "fmt"
)

type Stack interface {
    Push(int)
    Pop() (int, error)
    Size() int
    Head() int
}

type stack struct {
    elements []int
}

func NewStack() Stack {
    return &stack{}
}

func (s *stack) Push(data int) {
    s.elements = append(s.elements, data)
}

func (s *stack) Pop() (int, error) {
    if s.Size() == 0 {
        return 0, errors.New("Stack is empty")
    }
    result := s.elements[len(s.elements)-1]
    s.elements = s.elements[:len(s.elements)-1]
    return result, nil
}

func (s *stack) Size() int {
    return len(s.elements)
}

func (s *stack) Head() int {
    return s.elements[s.Size()-1]
}

func dailyTemperatures(temperatures []int) []int {
    result := make([]int, len(temperatures))
    stack := NewStack()

    for i, temp := range temperatures {
        for stack.Size() > 0 {
            tempHead := temperatures[stack.Head()]
            if temp > tempHead {
                result[stack.Head()] = i - stack.Head()
                stack.Pop()
            } else {
                break
            }
        }
        stack.Push(i)
    }

    return result
}

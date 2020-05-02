package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) Run() {
	fmt.Printf("%s, I am running\n", h.name)
}

func (h *Human) Say() {
	fmt.Printf("%s, I am %d year old \n", h.name, h.age)
}

type Singer struct {
	Human
	collection string
}

func (s *Singer) sing() {
	fmt.Printf("%s\n", s.collection)
}

type Man interface {
	Run()
	Say()
}

type IPv4 []byte

func (a IPv4) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", a[0], a[1], a[2], a[3])
}

type NewHuman struct {
	name string
	age  int
}

func main() {
	foo := Human{
		name: "foo",
		age:  21,
	}
	// foo.Run()
	// foo.Say()

	singer := Singer{
		Human: Human{
			name: "bar",
			age:  32,
		},
		collection: "test",
	}

	test := []Man{}

	test = append(test, &foo, &singer)
	for _, v := range test {
		v.Run()
		v.Say()
	}

	ipv4 := map[string]IPv4{
		"localhost": {127, 0, 0, 1},
		"google":    {8, 8, 8, 8},
	}

	for i, v := range ipv4 {
		fmt.Printf("%v: %v\n", i, v)
	}

	temp := []interface{}{}
	a := 1
	b := "foo"
	c := NewHuman{
		name: "bar",
		age:  21,
	}
	d := 100.1
	temp = append(temp, a, b, c, d)
	for i, value := range temp {
		if v, ok := value.(int); ok {
			fmt.Printf("%d, value: %d\n", i, v)
		} else if v, ok := value.(string); ok {
			fmt.Printf("%d, string value: %s\n", i, v)
		} else if v, ok := value.(NewHuman); ok {
			fmt.Printf("%d, Human value: %s\n", i, v.name)
		} else {
			fmt.Printf("%d, Unknown\n", i)
		}
	}

	for i, value := range temp {
		switch v := value.(type) {
		case int:
			fmt.Printf("%d, value: %d\n", i, v)
		case string:
			fmt.Printf("%d, string value: %s\n", i, v)
		case NewHuman:
			fmt.Printf("%d, Human value: %s\n", i, v.name)
		default:
			fmt.Printf("%d, Unknown\n", i)
		}
	}
}

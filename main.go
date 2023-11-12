// операции удаления и подмассивы
package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	/*debugSlice(s)
	//subslices -подмассивы - это срез в определенных границах среза
	d := s[1:4] //правая граница не входит в итоговый срез,т.о. cap будет 5, а не 6
	debugSlice(d)

	e := s[3:5] // cap будет 3
	debugSlice(e)

	e[0] = 10
	fmt.Println(s)
	fmt.Println(d) //s и d теперь вместо 4 имеют 10
	//изменения в любом subslice повлечет изменения во всех других и исходном слайсе*/
	//чтобы этого не происходило используют копирование:
	//copy()
	/*d := make([]int, 2, 2)
	copy(s[2:4], d) //copy(source, destination)
	d[0] = 10
	fmt.Println(s)
	fmt.Println(d)*/
	/*d := make([]int, 2)
	copy(d, s[2:4])
	d[0] = 10
	fmt.Println(s)
	fmt.Println(d)*/

	//в го из слайс просто так удалить нельзя, он не поддерживает такое операции, для этого пишется обёртка
	//рассмотрим несколько операций удаления из слайса
	s = remove(s, 2)
	debugSlice(s)
	s = anotherremove(s, 3)
	debugSlice(s)
}

// функция обертка для удаления из слайс
func remove(slice []int, index int) []int {

	return append(slice[:index], slice[index+1:]...) //если опускаем левую границу, отсчёт идёт с нуля, если правую len-1
	//... - потому что аппенд принимает эдлементы, мы не можем ей передать массив, поэтому спрединг

}

// рассмотрим еще один вариант
func anotherremove(slice []int, index int) []int {
	slice[len(slice)-1], slice[index] = slice[index], slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func debugSlice(input []int) {
	fmt.Printf("data: %+v\n", input)
	fmt.Printf("len: %+v\n", len(input))
	fmt.Printf("cap: %+v\n", cap(input))
}

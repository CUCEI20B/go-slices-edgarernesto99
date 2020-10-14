package main

import "fmt"

func main()  {
	var n, aux int
	result := 0

	fmt.Scan(&n);
	s := make([]int, 0, n)

	for i:=0; i<n; i++ {
		fmt.Scan(&aux);
		s = append(s, aux);
	}
	for i:=0; i<len(s); i++ {//Es más rapido en el for anterior
		result += s[i];
	}
	fmt.Println(result)
}
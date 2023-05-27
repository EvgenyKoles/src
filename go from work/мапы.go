map[0:-1 1:0 2:1 3:2 4:5 5:6]

3 1 5 2 3 5 3 0 3 4
2 0 6 2+2 6+2 -1 2 5
2 0 6 1 2 6 2 -1 2 5

2 0 6 1 2+6+2+-12+5

PS C:\Program Files\Go\src\stepic>
map[3:2 12:2 13:5]
map[1:0 2:1 3:2 5:6 12:2 13:5]
map[1:0 2:1 3:2 5:6 12:2 13:5]
map[0:-1 1:0 2:1 3:2 5:6 12:2 13:5]
map[0:-1 1:0 2:1 3:2 5:6 12:2 13:5]
map[0:-1 1:0 2:1 3:2 4:5 5:6 12:2 13:5]

	m := make(map[int]int)
	var numbers int

	for i := 0; i < 10; i++{
		fmt.Scan(&numbers)
		m[numbers] = work(numbers)

		for key, value := range m {
			if numbers == key{
				fmt.Print(value, " ") // 10	
		}

			}

	}
	
	fmt.Println(m[numbers])
	
	--------------------------
	
	
			for key, _ := range m {
			if numbers == key {
				fmt.Print(m[numbers], "+")
				break
			}else {
				m[numbers] = work(numbers)
				fmt.Print(m[numbers])
				break
			}
		}
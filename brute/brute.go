package brute

// import "fmt"

func Force() [][]string {
	island := [][]string{}
	first_row := []string{"1","0"}
	second_row := []string{"3","2"}
	island = append(island, first_row, second_row)


	return island
}
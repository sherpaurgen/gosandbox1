package main
import (
	"fmt"
	"log"
	"os"
	"sync"
	"bufio"
	"strings"
	"strconv"
)
func sort(tmpsli []int,mainslice *[][]int,wg *sync.WaitGroup) {
	for idx:=0;idx<len(tmpsli);idx++{
		for cidx:=1;cidx<len(tmpsli);cidx++{
			tempvar:=tmpsli[cidx-1]
			if tmpsli[cidx]<tmpsli[cidx-1]{
				tmpsli[cidx-1]=tmpsli[cidx]
				tmpsli[cidx]=tempvar
			}
		}
	}
	*mainslice=append(*mainslice,tmpsli)
	// closing gooutine
	fmt.Println("sub goroutine o/p:",tmpsli)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	// part 1 start user input & converting to slice of int
	fmt.Println("Enter series of numbers delimited by single whitespace")
	fmt.Println("example: 12 4 9 2 7")
	fmt.Printf(">")
	reader := bufio.NewReader(os.Stdin)
	userinput, _ := reader.ReadString('\n')
	userinput = strings.TrimSpace(userinput)
	tmpstr := strings.Split(userinput, " ")
	if len(tmpstr) < 1 {
		log.Fatal("Input error")
	}
	arr := []int{}
	if len(tmpstr) > 0 {
		for _, val := range tmpstr {
			tmp, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			arr = append(arr,tmp)
		}
		//arr := [21]int{4,14,29,2,42,25,21,11,41,16,20,31,34,30,43,2,40,4,23,6,37}
		rem := len(arr) % 4
		fact := len(arr)/4
		y := make([][]int, 0)
		arr2:=arr[0:len(arr)-rem]
		if fact > 0 && len(arr) > 0 {
			x := []int{}
			func(f int, v *[]int, mainslice *[][]int) *int {
				//8 9 16 37 4 14 31 43 1

				for i := 0; i < len(arr2) ; i=i+fact {
					tmpsli := arr2[i : i+fact]
					fmt.Println("tmpsli first:",tmpsli)
					//implementing bubble sort algorithm
					wg.Add(1)
					go sort(tmpsli, &y, &wg)
					//*mainslice = append(*mainslice, tmpsli)
				}
				return nil
			}(fact, &x, &y)
		}

		wg.Wait()
		final := []int{}

		fmt.Printf("divided array:%v \n", y)
		for i := 0; i < len(y); i++ {
			func(t []int) {
				for _, val := range t {
					final = append(final, val)
				}
			}(y[i])
		}
		if rem > 0 {
			lastsli := arr[len(arr)-rem:] //arr[N:]
			final = append(final, lastsli...)
		}
		fmt.Printf("Combined Slice: %v \n",final)
		//final sorting the merged slice
		for e:=0;e<len(final);e++ {
			for f:=1;f<len(final);f++ {
				t:=final[f-1]
				if final[f]<final[f-1]{
					final[f-1] = final[f]
					final[f] = t
				}
			}
		}
		fmt.Printf("Sorted Combined Slice/Array: %v \n",final)

	}
}

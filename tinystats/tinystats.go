package main


import ("fmt"
		"sort"
		"math"
	)

func mean(data []float64) float64 {
	var sum float64
	for _, num := range data {
		sum += num
	}
	return float64(sum) / float64(len(data)) 
}


func median(data []float64) float64 {
	sort.Float64s(data) 
	if len(data) % 2 == 0{
		return float64(data[len(data)/2]+data[len(data)/2-1]) / 2
	} else {
		return float64(data[len(data)/2])
	}
}


func quartile(data []float64, p float64) float64{
	var pth int 
	pth = int(math.Ceil(p * float64(len(data) - 1)))
	sort.Float64s(data) 
	return data[pth]
}


func mode(data []float64) int {
	freq := map[float64]int{}
	for _, val := range data{
		freq[val]++
	}

	mode := data[0]
	maxCount := 0

	for key, val := range freq{
		if val > maxCount{
			mode = key
			maxCount = val
		}
	}

	return int(mode)
}


func data_range(data []float64) float64{
	var max, min float64


	if len(data) == 0 {
		return 0
	}

	max = data[0]
	min = data[0]

	for _, val := range data{
		if val > max{
			max = val
		}
		
		if val < min{
			min = val
		}
	}
	return max - min
}										


func de_mean(data []float64, f func([]float64) float64) []float64{
	avg := f(data)
	var temp []float64
	for _, val := range data{
		temp = append(temp, (val - avg))
	}
	return temp
}

// func variance(data []float64, f func([]float64, func([]float64) float64) float64) float64{
// 	var n int

// 	n := len(data)
// 	deviation := f(data)



// }

func main(){

	nums := [] float64 {100,49,41,40,25,21,21,19,19,18,18,16,15,15,15,15,14,14,
						13,13,13,13,12,12,11,10,10,10,10,10,10,10,10,10,10,10,10,
						10,10,10,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,9,8,8,8,8,8,8,
						8,8,8,8,8,8,8,7,7,7,7,7,7,7,7,7,7,7,7,7,7,7,6,6,6,6,6,6,6,
						6,6,6,6,6,6,6,6,6,6,6,6,6,6,6,5,5,5,5,5,5,5,5,5,5,5,5,5,5,
						5,5,5,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,4,3,3,3,3,3,3,
						3,3,3,3,3,3,3,3,3,3,3,3,3,3,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,
						2,2,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}

	fmt.Println(data_range(nums))
}
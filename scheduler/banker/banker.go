package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

const (
	customerQuantity int = 5
	resourceQuantity int = 4
)

var (

	//the total amount of each resource
	total [resourceQuantity]int

	//the available amount of each resource
	available [resourceQuantity]int

	//the maximum demand of each customer
	maximum [customerQuantity][resourceQuantity]int

	//the amount currently allocated to each customer
	allocation [customerQuantity][resourceQuantity]int

	//the remaining need of each customer
	need [customerQuantity][resourceQuantity]int

	request [customerQuantity][resourceQuantity]int

	//the virtualAvailable simulate available amount of each resource
	virtualAvailable [resourceQuantity]int

	//the finish simulate finish amount of each task
	finish [customerQuantity]bool

	canFinish [customerQuantity]bool

	safeSequence []int

	executeSequence []int

	wg sync.WaitGroup

	rw sync.RWMutex
)

func banker(customerID int) bool {
	canRun := true

	for i := range request[customerID] {
		if request[customerID][i] > need[customerID][i] {
			canRun = false
		}
	}

	for i := range request[customerID] {
		if request[customerID][i] > available[i] {
			canRun = false
		}
	}

	if canRun == true {
		for i := range available {
			available[i] -= request[customerID][i]
			allocation[customerID][i] += request[customerID][i]
			need[customerID][i] -= request[customerID][i]
		}

		safe := checkSafe()
		if safe == false {
			canRun = false
			for i := range available {
				available[i] += request[customerID][i]
				allocation[customerID][i] -= request[customerID][i]
				need[customerID][i] += request[customerID][i]
			}
		}
	}
	return canRun
}

func checkSafe() bool {
	safeFinishQuantity := 0
	executeSequence = executeSequence[:0]

	for i := 0; i < resourceQuantity; i++ {
		virtualAvailable[i] = available[i]
	}

	for i := 0; i < customerQuantity; i++ {
		canFinish[i] = false
	}

	for i := 0; i < customerQuantity; i++ {
		j := 0
		//扣除所需大於現有資源的
		for j = 0; j < resourceQuantity; j++ {
			if need[i][j] > virtualAvailable[j] {
				break
			}
		}
		// 未完成 且 現有資源能滿足需求
		if canFinish[i] == false && j == resourceQuantity {
			canFinish[i] = true
			for k := 0; k < resourceQuantity; k++ {
				virtualAvailable[k] += allocation[i][k]
			}
			executeSequence = append(executeSequence, i)
			safeFinishQuantity++
			i = -1 // 令 loop 結束 i 又重置為 0，從頭開始找其他能放入序列的值
		} else {
			continue // 現存資源無法滿足需求 繼續找機會
		}
		if safeFinishQuantity == customerQuantity { // 能找到安全的執行順序
			return true
		}
	}
	//fmt.Println("unsafe")
	return false
}

func requestResources(customerID int) bool {
	requestSuccess := false
	requestSuccess = banker(customerID)
	return requestSuccess
}

func releaseResources(customerID int) {
	for i := 0; i < resourceQuantity; i++ {
		available[i] += allocation[customerID][i]
		maximum[customerID][i] = 0
		need[customerID][i] = 0
	}
}

func executeWrite(customerID int) bool {
	executeSuccess := false
	rw.Lock()
	file, _ := os.OpenFile("test.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	needWrite := "customerID: " + strconv.Itoa(customerID)
	needWrite += "\n"
	byteSlice := []byte(needWrite)
	file.Write(byteSlice)
	file.Close()
	rw.Unlock()
	executeSuccess = true
	return executeSuccess
}

func answerA() {
	available = [4]int{3, 3, 2, 1}

	maximum[0] = [4]int{4, 2, 1, 2}
	maximum[1] = [4]int{5, 2, 5, 2}
	maximum[2] = [4]int{2, 3, 1, 6}
	maximum[3] = [4]int{1, 4, 2, 4}
	maximum[4] = [4]int{3, 6, 6, 5}

	allocation[0] = [4]int{2, 0, 0, 1}
	allocation[1] = [4]int{3, 1, 2, 1}
	allocation[2] = [4]int{2, 1, 0, 3}
	allocation[3] = [4]int{1, 3, 1, 2}
	allocation[4] = [4]int{1, 4, 3, 2}

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			need[i][j] = maximum[i][j] - allocation[i][j]
		}
	}

	for i := range request {
		request[i] = [4]int{0, 0, 0, 0}
	}
	checkSafe()
	fmt.Println("a. execute sequence:", executeSequence)
}

func answerB() {
	available = [4]int{3, 3, 2, 1}

	maximum[0] = [4]int{4, 2, 1, 2}
	maximum[1] = [4]int{5, 2, 5, 2}
	maximum[2] = [4]int{2, 3, 1, 6}
	maximum[3] = [4]int{1, 4, 2, 4}
	maximum[4] = [4]int{3, 6, 6, 5}

	allocation[0] = [4]int{2, 0, 0, 1}
	allocation[1] = [4]int{3, 1, 2, 1}
	allocation[2] = [4]int{2, 1, 0, 3}
	allocation[3] = [4]int{1, 3, 1, 2}
	allocation[4] = [4]int{1, 4, 3, 2}

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			need[i][j] = maximum[i][j] - allocation[i][j]
		}
	}

	for i := range request {
		request[i] = [4]int{0, 0, 0, 0}
	}

	request[1] = [4]int{1, 1, 0, 0}

	can := banker(1)
	fmt.Println("b. p1 1, 1, 0, 0 can the request be granted immediately:", can)

}

func answerC() {
	available = [4]int{3, 3, 2, 1}

	maximum[0] = [4]int{4, 2, 1, 2}
	maximum[1] = [4]int{5, 2, 5, 2}
	maximum[2] = [4]int{2, 3, 1, 6}
	maximum[3] = [4]int{1, 4, 2, 4}
	maximum[4] = [4]int{3, 6, 6, 5}

	allocation[0] = [4]int{2, 0, 0, 1}
	allocation[1] = [4]int{3, 1, 2, 1}
	allocation[2] = [4]int{2, 1, 0, 3}
	allocation[3] = [4]int{1, 3, 1, 2}
	allocation[4] = [4]int{1, 4, 3, 2}

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			need[i][j] = maximum[i][j] - allocation[i][j]
		}
	}

	for i := range request {
		request[i] = [4]int{0, 0, 0, 0}
	}

	request[4] = [4]int{0, 0, 2, 0}

	can := banker(4)
	fmt.Println("c. p4 0, 0, 2, 0 can the request be granted immediately:", can)
}

func main() {
	answerA()
	answerB()
	answerC()
	/*
		// total 12, 12, 8, 10
		available = [4]int{3, 3, 2, 1}

		maximum[0] = [4]int{4, 2, 1, 2}
		maximum[1] = [4]int{5, 2, 5, 2}
		maximum[2] = [4]int{2, 3, 1, 6}
		maximum[3] = [4]int{1, 4, 2, 4}
		maximum[4] = [4]int{3, 6, 6, 5}

		allocation[0] = [4]int{2, 0, 0, 1}
		allocation[1] = [4]int{3, 1, 2, 1}
		allocation[2] = [4]int{2, 1, 0, 3}
		allocation[3] = [4]int{1, 3, 1, 2}
		allocation[4] = [4]int{1, 4, 3, 2}

		for i := 0; i < 5; i++ {
			for j := 0; j < 4; j++ {
				need[i][j] = maximum[i][j] - allocation[i][j]
			}
		}

		for i := range request {
			request[i] = [4]int{0, 0, 0, 0}
		}

		//request[1] = [4]int{1, 1, 0, 0}
		//request[4] = [4]int{0, 0, 2, 0}


		//total[0], _ = strconv.Atoi(os.Args[1])
		//total[1], _ = strconv.Atoi(os.Args[2])
		//total[2], _ = strconv.Atoi(os.Args[3])
		//total[3], _ = strconv.Atoi(os.Args[4])

		wg.Add(5)
		for i := 0; i < 5; i++ {
			go func(customerID int) {
				for {
					requestSuccess := requestResources(customerID)
					if requestSuccess == true {
						ok := executeWrite(customerID)
						if ok == true {
							releaseResources(customerID)
							finish[customerID] = true
							break
						}
					}
				}
				wg.Done()
			}(i)
		}
		wg.Wait()

		fmt.Println(executeSequence)
	*/
}

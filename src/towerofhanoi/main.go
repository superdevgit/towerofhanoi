package main
  
import (
    "fmt"
    "math"
)

var modulus = math.Pow(10, 9);

type HanoiPath struct {
    From    int
    To      int
}

var fullPath []HanoiPath

//main function
func main() {

    var cost int

    cost = e(2, 5, 1, 3, 5)
    fmt.Printf("Expected cost is 60 and obtained cost of E(2,5,1,3,5) is ")
    fmt.Println(cost)

    cost = e(3, 20, 4, 9, 17)
    fmt.Printf("Expected cost is 2358 and obtained cost of E(3, 20, 4, 9, 17) is ")
    fmt.Println(cost)
    
    //this will exceed memory limit and cause error
    summationCost()
}

//move of disks
func moveDisk(n int, from int, to int, aux int) {

    if n == 1 {
        if (len(fullPath) == 0) {
            if to < aux {
                fullPath = append(fullPath, HanoiPath{to, from})
            } else {
                fullPath = append(fullPath, HanoiPath{aux, from})
            }
        } else {
            var last = fullPath[len(fullPath) - 1]
            fullPath = append(fullPath, HanoiPath{last.To, from})
        }
        fullPath = append(fullPath, HanoiPath{from, to})
        return;
    }

    moveDisk(n-1, from, aux, to)

    if (len(fullPath) == 0) {
        fullPath = append(fullPath, HanoiPath{aux, from})
    } else {
        var last = fullPath[len(fullPath) - 1]
        fullPath = append(fullPath, HanoiPath{last.To, from})
    }

    fullPath = append(fullPath, HanoiPath{from, to})

    moveDisk(n-1, aux, to, from)
}

//get each move cost
func moveCast(tiles int, i int, j int) int {
    
    var cost int
    if i < j {
        cost = int(math.Pow((float64(j) - 1), 2) - math.Pow((float64(i) - 1), 2))
    } else {
        cost = int(math.Pow((float64(tiles) - float64(j)), 2) - math.Pow((float64(tiles) - float64(i)), 2))
    }
    return int(cost % int(modulus))
}

func e(numberOfDisk int, numberOfTiles int, towerA int, towerB int, towerC int) int {
    
    fullPath = []HanoiPath{}

    moveDisk(numberOfDisk, towerA, towerC, towerB)

    var pathCost = 0

    for _, element := range fullPath {
        pathCost += moveCast(numberOfTiles, element.From, element.To)
    }

    pathCost = pathCost % int(modulus)

    return pathCost

}

func summationCost() {
    var summationCost = 0
    var i int

    for i <= 1500 {
        var numOfDisk = i % int(modulus)
        var numOfTiles = (numOfDisk * 10) % int(modulus)
        var towerA = (numOfDisk * 3) % int(modulus)
        var towerB = (numOfDisk * 6) % int(modulus)
        var towerC = (numOfDisk * 9) % int(modulus)
        fullPath = []HanoiPath{}

        moveDisk(numOfDisk, towerA, towerC, towerB)

        var pathCost =0

        for _, element := range fullPath {
            pathCost += moveCast(numOfTiles, element.From, element.To)
        }
        summationCost += pathCost
        summationCost = summationCost % int(modulus)
        i = i + 1
    }

    fmt.Printf("Summation Cost is ")
    fmt.Println(summationCost)
}
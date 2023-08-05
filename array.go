package main

import ("fmt"
	"log"
	"os/exec"
	"os"
)

func main() {

	var chr = "8"

	var fuck [10][10]string

	fuck = newMap()
	
	printMap(fuck)

	var choice string
	var x int
	var y int
	
	for {
		
		fmt.Println("type command up")
		fmt.Scanln(&choice)	
		
		switch choice {

			case "up", "w":
				x--
			case "down", "s":
				x++
			case "left", "a":
				y--
			case "right", "d":
				y++

			}
	clear()
	fuck = newMap()
	

	fuck[3][7] = "="
	fuck[3][8] = "="
        fuck[3][9] = "="
	fuck[5][7] = "H"
	fuck[5][8] = "n" 
	fuck[5][9] = "H"
	fuck[4][7] = "H"
        fuck[4][8] = "H"
        fuck[4][9] = "H"


	fuck[5][4] = "E"
	
	fuck[x][y] = chr
	
	if fuck[5][4] == chr {
                log.Fatal()
        }

	if fuck[5][8] == chr {
		levelTwo()
	}

	printMap(fuck)

	}
	


}

func levelTwo() {
var chr = "8"
        var fuck [10][10]string

        fuck = newMap()

        printMap(fuck)

        var choice string
        var x int
        var y int

        for {
		
                fmt.Println("type command up")
                fmt.Scanln(&choice)

                switch choice {

                        case "up", "w":
                                x--
                        case "down", "s":
                                x++
                        case "left", "a":
                                y--
                        case "right", "d":
                                y++

                        }
	clear()
        fuck = newMap()


        fuck[1][5] = "%"
        fuck[1][6] = "%"
        fuck[1][7] = "%"
        fuck[3][5] = "H"
        fuck[3][6] = "n"
        fuck[3][7] = "H"
        fuck[2][5] = "H"
        fuck[2][6] = "H"
        fuck[2][7] = "H"


        fuck[5][4] = "E"

        fuck[x][y] = chr

        if fuck[5][4] == chr {
                log.Fatal()
        }

        if fuck[5][8] == chr {

        }

        printMap(fuck)

        }


}

func clear() {

        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
}


//prints 2d array
func printMap(fuck [10][10]string) {

	for d := 0; d <= 9; d++ {

                fmt.Println("", fuck[d], "\n")

        }

}

//returns 2d array of all Xs
func newMap() [10][10]string  {

	var fuck [10][10]string

        for j := 0; j <= 9; j++ {
                for i := 0; i <= 9; i++ {

                        fuck[j][i] = "x"

                }
        }

	for d := 1; d <= 8; d++ {
		for f := 1; f <= 8; f++ {
			fuck[d][f] = "0"
		}
        }



return fuck

}



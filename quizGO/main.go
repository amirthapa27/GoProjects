package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type problem struct {
	q string //quetsion
	a string //answer
}

func problemPuller(fileName string) ([]problem, error) {
	//read all the problems from the file
	//open the file
	if fObj, err := os.Open(fileName); err == nil {
		//read the file
		csvR := csv.NewReader(fObj)
		if clines, err := csvR.ReadAll(); err == nil {
			return parseProblem(clines), nil
		} else {
			return nil, fmt.Errorf("error in reading data in csv"+"format from %s file ; %s", fileName, err)

		}
	} else {
		return nil, fmt.Errorf("error in opening %s file ; %s", fileName, err)
	}

	//call
}

func main() {
	//input the name of the file
	fName := flag.String("f", "quiz.csv", "path of csv file")
	//timer
	timer := flag.Int("t", 30, "time for the quiz")
	flag.Parse()
	//pull the problems from the file (calliing our problem puller function)
	problems, err := problemPuller(*fName)
	//handle error
	if err != nil {
		panic(err)
	}
	//create a variable to count th ecorrect answers
	correctANS := 0
	//intitalize a timer
	tObj := time.NewTimer(time.Duration(*timer) * time.Second)
	ansC := make(chan string)
	//loop through problems
problemLoop:
	for i, p := range problems {
		var answer string
		fmt.Printf("Problems %d: %s=", i+1, p.q)

		go func() {
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()
		select {
		case <-tObj.C:
			fmt.Println()
			break problemLoop

		case iAns := <-ansC:
			if iAns == p.a {
				correctANS++

			}
			if i == len(problems)-1 {
				close(ansC)
			}
		}
	}
	//print the results
	fmt.Printf("Youre result is %d out of %d \n", correctANS, len(problems))
	fmt.Println("Press enter to exit")
	<-ansC
}

func parseProblem(lines [][]string) []problem {
	//go over the lines and parse them
	r := make([]problem, len(lines))
	for i := 0; i < len(lines); i++ {
		r[i] = problem{q: lines[i][0], a: lines[i][1]}
	}
	return r
}

func exit(msg string) {
	fmt.Println("msg")
	os.Exit(1)
}

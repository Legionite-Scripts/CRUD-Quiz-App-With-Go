package main ///main package

import (
	"fmt"  //fmt package : printing and scanning
	"time" // time package : implementation of  timer function for code
)

func main() { //Main Method
	fmt.Print("WELCOME TO LEGIONITE'S TECH QUIZ\nPlease input your name to start playing : ")
	var nameInput string   //username variable
	fmt.Scanln(&nameInput) // Scanner to collect user name string input
	fmt.Println("Okay " + nameInput + "\nSelect what mode you want to play : ")

	modeArray := [...]string{"1. Web Quiz", "2. Mobile Quiz"}
	for _, val := range modeArray {
		fmt.Printf("%v\n", val)
	}

	var modeInput int
	fmt.Scanln(&modeInput)

	//switch case that checks the quiz mode which the user wants to play
	switch modeInput {
	//CASE ONE OF MODE INPUT
	case 1:
		fmt.Println("You will get a maximum time of 30 Seconds for this quiz\n There are a total of 10 questions")
		fmt.Print("Press 1 to start : ")
		var pressOne int
		fmt.Scan(&pressOne)

		//INNER PRESS ONE SWITCH CASE
		switch pressOne {
		case 1:
			timeDelay := 30 * time.Second

			fmt.Println("You have chosen the `WEB` quiz\n\n")
			//Question one for web
			fmt.Println("1. What is the full meaning of Html ? \n(a)Helix Tuple Mark Language\n(b)Hyper Text Markup Language\n(c)Hyper Text Mark Language\n(d)Huge Tuple Management Language")
			var questionOneInput string
			fmt.Scan(&questionOneInput)

			//Question Two for web
			fmt.Println("2. What year was JavaScript invented?\n(a)1992\n(b) 1993\n(c)1995\n(d)2006")
			var questionTwoInput string
			fmt.Scan(&questionTwoInput)

			//Question Three for web
			fmt.Println("3. All of these is an IDE except : \n(a)Apache Netbeans\n(b) Notepad\n(c)IntelliJ Environment\n(d)Pycharm")
			var questionThreeInput string
			fmt.Scan(&questionThreeInput)

			//Question four for web
			fmt.Println("4. Who Invented Python? \n(a) Guido van Rossum \n(b)Adolf Samuel Hitler\n(c)Ludwig Van Beethoven\n(d)Jimmy Q. Fallon")
			var questionFourInput string
			fmt.Scan(&questionFourInput)

			//Question five for web
			fmt.Println("5. What is the extension for a javascript file?\n(a) .jv\n(b) .jsc\n(c).js\n(d)  .jsp")
			var questionFiveInput string
			fmt.Scan(&questionFiveInput)

			//Question six for web
			fmt.Println("6. What is the full meaning of CSS?\n(a) Cascading Style Sheet\n(b) Calamanja Style Sheet\n(c) Candace Style Shirt\n(d) Canadian Smooth Style")
			var questionSixInput string
			fmt.Scan(&questionSixInput)

			//Question seven for web
			fmt.Println("7. What is the full meaning of API? \n(a) Application Permission Interface\n(b) Application Printing Interface\n(c)Application Programming Interface\n(d) Application Picturesque Interface")
			var questionSevenInput string
			fmt.Scan(&questionSevenInput)

			//Question eight for web
			fmt.Println("8. All of these are databases except : \n(a) MongoDB\n(b)CouchDB\n(c)NOSql\n(d)BenchDB")
			var questionEightInput string
			fmt.Scan(&questionEightInput)

			//Question nine for web
			fmt.Println("9. Which of these are alll JavaScript Frameworks\n(a) ReactJs, Step.Js, NodeJs\n(b) StepJs, AngularJs, BrightJs\n(c) AngularJs,ReactJs,NodeJs\n(d)BrightJs, NodeJs, NextJs")
			var questionNineInput string
			fmt.Scan(&questionNineInput)

			//Question 10 for web
			fmt.Println("10. Identify the fastest language among the list : \n(a) C#\n(b) Python\n(c) Golang\n(d) Java")
			var questionTenInput string
			fmt.Scan(&questionTenInput)

			//	WEB QUESTIONS VALIDATION

			var score int32 = 0

			//Inner if statements to validate answers
			if questionOneInput == "b" || questionOneInput == "B" {
				fmt.Println("You got question one right")
				score++
			} else {
				fmt.Println("You got question one wrong")
			}

			if questionTwoInput == "c" || questionTwoInput == "C" {
				fmt.Println("You got question two right")
				score++
			} else {
				fmt.Println("You got question two wrong")
			}

			if questionThreeInput == "b" || questionThreeInput == "B" {
				fmt.Println("You got question three right")
				score++
			} else {
				fmt.Println("You got question three wrong")
			}

			if questionFourInput == "a" || questionFourInput == "A" {
				fmt.Println("You got question four right")
				score++

			} else {
				fmt.Println("You got question four wrong")
			}

			if questionFiveInput == "c" || questionFiveInput == "C" {
				fmt.Println("You got question five right")
				score++
			} else {
				fmt.Println("You got question five wrong")
			}

			if questionSixInput == "a" || questionSixInput == "A" {
				fmt.Println("You got question six right")
				score++
			} else {
				fmt.Println("You got question six wrong")
			}

			if questionSevenInput == "c" || questionSevenInput == "C" {
				fmt.Println("You got question seven right")
				score++
			} else {
				fmt.Println("You got question seven wrong")
			}

			if questionEightInput == "d" || questionEightInput == "D" {
				fmt.Println("You got question Eight right")
				score++
			} else {
				fmt.Println("You got question Eight wrong")
			}

			if questionNineInput == "c" || questionNineInput == "C" {
				fmt.Println("You got question Nine right")
				score++
			} else {
				fmt.Println("You got question Nine wrong")
			}

			if questionTenInput == "c" || questionTenInput == "C" {
				fmt.Println("You got question Ten right")
				score++
			} else {
				fmt.Println("You got question Ten wrong")
			}

			if score > 5 {
				fmt.Println("\n\nGood job "+nameInput, "\nYour score is ", score)
			} else {
				fmt.Println("\n\nAww💔 "+nameInput+" You could do better\nYour score is ", score)
			}

			<-time.After(timeDelay) // time function : check if user has exceeded specified time
			fmt.Println("Time's up!")

		default:
			fmt.Println("Wrong Input")
			main()
		}

		//CASE 2 OF MODE INPUT GOES HERE
	case 2:
		timeDelay := 10 * time.Second
		fmt.Println("You have chosen the `MOBILE` quiz")

		//Question one for mobile
		fmt.Println("1. Which of these is a mobile programming language?\n(a) Python\n(b) CSS\n(c) Java\n(d) Binglet")
		var mobileOneInput string
		fmt.Scanln(&mobileOneInput)

		//Question two for mobile
		fmt.Println("2. What is the official logo of Swift?\n(a) A Cheetah\n(b) A frog\n(c) A bird\n(d)  A monkey")
		var mobileTwoInput string
		fmt.Scanln(&mobileTwoInput)

		//Question three for mobile
		fmt.Println("3. What is function recursion? \n(a) a method reversing itself\n(b) a method recalling itself\n(c) a variable calling a function\n(d) a method calling a variable")
		var mobileThreeInput string
		fmt.Scanln(&mobileThreeInput)

		//Question four for mobile
		fmt.Println("4. Which of these is not a mobile programming technology? \n(a) Java Swing\n(b) Linux\n(c) Flutter\n(d) Amaterasu")
		var mobileFourInput string
		fmt.Scanln(&mobileFourInput)

		//Question five for mobile
		fmt.Println("5. Which of these is a basic requirement to run flutter? \n(a) Kotlin\n(b) Dart\n(c) Java\n(d) Golang")
		var mobileFiveInput string
		fmt.Scanln(&mobileFiveInput)

		//Question six for mobile
		fmt.Println("6. Which of these is not a java framework?\n(a) Swing\n(b) Outline\n(c) Hibernate\n(d) Spring")
		var mobileSixInput string
		fmt.Scanln(&mobileSixInput)

		//Question seven for mobile
		fmt.Println("7. Which of these is needed by every single mobile app?\n(a) Multiply file\n(b) Manifest file\n(c) Minizip file\n(d) Mundane file")
		var mobileSevenInput string
		fmt.Scanln(&mobileSevenInput)

		//Question eight for mobile
		fmt.Println("8. Which is an alternative to if/else?\n(a) Snitch case\n(b) Spring case\n(c) Speech case\n(d) Switch case")
		var mobileEightInput string
		fmt.Scanln(&mobileEightInput)

		//Question nine for mobile
		fmt.Println("9. Which is used for java SQL connection?\n(a) JGBA\n(b) JDBC\n(c) JSQL\n(d) JDCB")
		var mobileNineInput string
		fmt.Scanln(&mobileNineInput)

		//Question ten for mobile
		fmt.Println("10. Which of these is a  Java Framework?\n(a) Spring\n(b) Grasp\n(c) Flask\n(d) Kettle.Java")
		var mobileTenInput string
		fmt.Scanln(&mobileTenInput)

		//MOBILE QUESTIONS VALIDATION

		var score int32 = 0

		//Inner if statements to validate answers
		if mobileOneInput == "c" || mobileOneInput == "C" {
			fmt.Println("You got question one right")
			score++
		} else {
			fmt.Println("You got question one wrong")
		}

		if mobileTwoInput == "c" || mobileTwoInput == "C" {
			fmt.Println("You got question two right")
			score++
		} else {
			fmt.Println("You got question two wrong")
		}

		if mobileThreeInput == "b" || mobileThreeInput == "B" {
			fmt.Println("You got question three right")
			score++
		} else {
			fmt.Println("You got question three wrong")
		}

		if mobileFourInput == "a" || mobileFourInput == "A" {
			fmt.Println("You got question four right")
			score++

		} else {
			fmt.Println("You got question four wrong")
		}

		if mobileFiveInput == "b" || mobileFiveInput == "B" {
			fmt.Println("You got question five right")
			score++
		} else {
			fmt.Println("You got question five wrong")
		}

		if mobileSixInput == "b" || mobileSixInput == "B" {
			fmt.Println("You got question six right")
			score++
		} else {
			fmt.Println("You got question six wrong")
		}

		if mobileSevenInput == "b" || mobileSevenInput == "B" {
			fmt.Println("You got question seven right")
			score++
		} else {
			fmt.Println("You got question seven wrong")
		}

		if mobileEightInput == "d" || mobileEightInput == "D" {
			fmt.Println("You got question Eight right")
			score++
		} else {
			fmt.Println("You got question Eight wrong")
		}

		if mobileNineInput == "b" || mobileNineInput == "B" {
			fmt.Println("You got question Nine right")
			score++
		} else {
			fmt.Println("You got question Nine wrong")
		}

		if mobileTenInput == "a" || mobileTenInput == "A" {
			fmt.Println("You got question Ten right")
			score++
		} else {
			fmt.Println("You got question Ten wrong")
		}

		if score > 5 {
			fmt.Println("\n\nGood job "+nameInput, "\nYour score is ", score)
		} else {
			fmt.Println("\n\nAww💔 "+nameInput+" You could do better\nYour score is ", score)
		}

		<-time.After(timeDelay) // time function : check if user has exceeded specified time
		fmt.Println("Time's up!")
	//DEFAULT CASE OF MODE INPUT GOES HERE
	default:
		fmt.Println("Wrong input")

	}
}

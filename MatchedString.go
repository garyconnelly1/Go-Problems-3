package main


//imports
import (
	"math/rand"
	"time"
	"regexp"
	"fmt"
	"bufio"
	"os"
	"strings"
	
)



func ElizaResponse(inputStr string) string{

	input :=inputStr
	
//back ticks instead of quotes to make sure it doesnt leave the characters first
	if matched,_ := regexp.MatchString(`(?i). *\bfather\b.*`,input); matched {
		//return the string below
		return "Why dont you tell me more about your father?"
	}

//capture I am
	re := regexp.MustCompile("[Ii] (?:A|a)(?:M|m) ([^.!?]*)[.!?]?")
	

	if re.MatchString(input){	
		return re.ReplaceAllString(input, "How do you know you are $1?") 

	}

	

	answers := []string{
		"I'm not sure what you are trying to say. Could you explain it to me?",
		"How does that make you feel?",
		 "Why do you say that?",
	}
	//returning a single string response
		return answers[rand.Intn(len(answers))]

}


//part 5
func refPro(inputStr string) string{
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(inputStr, -1)
	
	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`you`, `I`},
		{`my`, `your`},
		{`your`, `my`},
		{`I am`,`you are`},
	}
	
	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				
				break
			}
		}
	}
	
	// Put the tokens back together.
	return strings.Join(tokens, ``)
}
	//end part 5


func main(){
	tester :=refPro("I am not sure that you understand the effect your questions are having on me.")
	fmt.Println(tester)
	rand.Seed(time.Now().UTC().UnixNano())//get a random number
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please message me")

	input, _ := reader.ReadString('\n')
	fmt.Scanf("%s", &input)
	fmt.Println(ElizaResponse(input))

}

//useful resource 
// https://shapeshed.com/golang-regexp/#what-is-the-regex-package-in-go


//I am not sure that you understand the effect your questions are having on me.
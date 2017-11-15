package main


//imports
import (
	"math/rand"
	"time"
	"regexp"
	"fmt"
	"bufio"
	"os"
	
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

func main(){
	rand.Seed(time.Now().UTC().UnixNano())//get a random number
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please message me")

	input, _ := reader.ReadString('\n')
	fmt.Scanf("%s", &input)
	fmt.Println(ElizaResponse(input))

}
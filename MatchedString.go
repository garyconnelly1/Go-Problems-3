package main


//imports
import (
	"math/rand"
	"time"
	"regexp"
	"fmt"
)

func ElizaResponse(input string) string{

	var input string

	fmt.println("Please message me")
	fmt.Scanf("%s", &input)
//back ticks instead of quotes to make sure it doesnt leave the characters first
	if matched,_ := regexp.MatchString(`(?i). *\bfather\b.*`,input); matched {
		//return the string below
		return "Why dont you tell me more about your father?"
	}

//capture I am
	re := regexp.MustCompile("I am ([^.!?]*)[.!?]?")
	if re.MatchedString(input){
		return re.ReplaceAllString(input, "How do you know you are $1?")) 

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
	ElizaResponse()

}
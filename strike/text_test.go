package strike_test

import (
	"fmt"
	"testing"

	strike "github.com/strike-official/go-sdk/strike"
)

func TestTextOutput(t *testing.T) {

	strikeObj := strike.Create("location_handler", "https://abc.com") //This creates the bare minimum of the response

	question1 := strikeObj.Question("Key1").
		QuestionText().
		SetTextToQuestion("How are you?", "Text Description, getting used for testing purpose.")

	question1.Answer(true).
		LocationInput("Select the property location")

	fmt.Println(string(strikeObj.ToJson())) //ToJson() just converts the struct to json value
}

package strike_test

import (
	"testing"
	//"encoding/json"
	"fmt"
	"github.com/strike-official/go-sdk/strike"
)

func TestCreate(t *testing.T){

	strike := strike.Create("first_handler","https://abc.com")	//This creates the bare minimum of the response
	
	question1 := strike.Question("Key1"). 						//Creates a question object with context as Key1
		QuestionCard(strike).									//inserts a card object in the question. Other interfaces to be used here
		SetHeaderToQuestion(strike, 2, "HALF").					//Sets the header for the card object. Always next to QuestionCard function. HALF FULL
		AddGraphicRowToQuestion(strike, "pic_row", "https://xyz.com").		//Add a photo or video to the card. pic_row video_row
		AddTextRowToQuestion(strike, "h3", "News for the young!", "#e456tc",true).		//Add a text row to the card
		AddTextRowToQuestion(strike, "h4", "Hey there this the top news for you", "black",false)	//Add a text row to the card
	
	_ = question1.Answer(strike,true).							//Creates an answer object with multiselect true or false. Must be after Question and not before	
		AnswerCardArray(strike,"VERTICAL").						//Add an array of Card with orientation VERTICAL HORIZONTAL
		AnswerCard(strike).										//A card in the above added array
		SetHeaderToAnswer(strike, 2, "FULL").					//Sets the header for the card object
		AddGraphicRowToAnswer(strike, "pic_row", "https://abc.com").		//Add a pic or video row to the card. pic_row video_row
		AddTextRowToAnswer(strike, "h3", "wolla!", "#e456tc",true).						
		AnswerCard(strike).										//Another card in the array	
		SetHeaderToAnswer(strike, 2, "FULL").					//Header for the above card
		AddTextRowToAnswer(strike, "h3", "next card!", "#e456tc",true)

	_ = strike.Question("Key2").								//Another question with context as key2	
	    QuestionCard(strike).									//This is again card type of question						
		SetHeaderToQuestion(strike, 2, "HALF").					//Observe Question without answer is allowed but Answer without Question is not
		AddTextRowToQuestion(strike, "h3", "This is the last question!", "#e456tc",true)
	
		
	    fmt.Println(string(strike.ToJson()))						//ToJson() just converts the struct to json value



}


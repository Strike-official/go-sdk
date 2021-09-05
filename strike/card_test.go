package strike_test

import (
	"testing"

	//"encoding/json"
	"fmt"

	"github.com/strike-official/go-sdk/strike"
)

func TestCreate(t *testing.T) {

	strike := strike.Create("first_handler", "https://abc.com") //This creates the bare minimum of the response

	question1 := strike.Question("Key1"). //Creates a question object with context as Key1
						QuestionCard().                                                                   //inserts a card object in the question. Other interfaces to be used here
						SetHeaderToQuestion(2, "HALF").                                                   //Sets the header for the card object. Always next to QuestionCard function. HALF FULL
						AddGraphicRowToQuestion("pic_row", "https://xyz.com").                            //Add a photo or video to the card. pic_row video_row
						AddTextRowToQuestion("h3", "News for the young!", "#e456tc", true).               //Add a text row to the card
						AddTextRowToQuestion("h4", "Hey there this the top news for you", "black", false) //Add a text row to the card

	_ = question1.Answer(true). //Creates an answer object with multiselect true or false. Must be after Question and not before
					AnswerCardArray("VERTICAL").                         //Add an array of Card with orientation VERTICAL HORIZONTAL
					AnswerCard().                                        //A card in the above added array
					SetHeaderToAnswer(2, "FULL").                        //Sets the header for the card object
					AddGraphicRowToAnswer("pic_row", "https://abc.com"). //Add a pic or video row to the card. pic_row video_row
					AddTextRowToAnswer("h3", "wolla!", "#e456tc", true).
					AnswerCard().                 //Another card in the array
					SetHeaderToAnswer(2, "FULL"). //Header for the above card
					AddTextRowToAnswer("h3", "next card!", "#e456tc", true)

	_ = strike.Question("Key2"). //Another question with context as key2
					QuestionCard().                 //This is again card type of question
					SetHeaderToQuestion(2, "HALF"). //Observe Question without answer is allowed but Answer without Question is not
					AddTextRowToQuestion("h3", "This is the last question!", "#e456tc", true)

	// You need to bind this strike struct in wrapper() to make it strike complient

	// response := strike.Wrapper()

	fmt.Println(string(strike.ToJson())) //ToJson() just converts the struct to json value

}

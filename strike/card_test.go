package strike_test

import (
	"testing"
	"fmt"
	app "github.com/strike-official/go-sdk/strike"
)

func TestCreate(t *testing.T) {

	strike := app.Create("first_handler", "https://abc.com") //This creates the bare minimum of the response

	question1 := strike.Question("Key1"). //Creates a question object with context as Key1
						QuestionCard().                                                                   //inserts a card object in the question. Other interfaces to be used here
						SetHeaderToQuestion(2, app.HALF).                                                   //Sets the header for the card object. Always next to QuestionCard function. HALF FULL
						AddGraphicRowToQuestion(app.PICTURE_ROW, []string{"https://abc.com"}).                            //Add a photo or video to the card. pic_row video_row
						AddTextRowToQuestion(app.H3, "News for the young!", "#e456tc", true).               //Add a text row to the card
						AddTextRowToQuestion(app.H4, "Hey there this the top news for you", "black", false) //Add a text row to the card

	_ = question1.Answer(true). //Creates an answer object with multiselect true or false. Must be after Question and not before
					AnswerCardArray(app.VERTICAL).                         //Add an array of Card with orientation VERTICAL HORIZONTAL
					AnswerCard().                                        //A card in the above added array
					SetHeaderToAnswer(2, app.FULL).                        //Sets the header for the card object
					AddGraphicRowToAnswer(app.PICTURE_ROW, []string{"https://abc.com","https://xyz.com"}). //Add a pic or video row to the card. pic_row video_row
					AddTextRowToAnswer(app.H3, "wolla!", "#e456tc", true).
					AnswerCard().                 //Another card in the array
					SetHeaderToAnswer(2, app.FULL). //Header for the above card
					AddTextRowToAnswer(app.H3, "next card!", "#e456tc", true)

	_ = strike.Question("Key2"). //Another question with context as key2
					QuestionCard().                 //This is again card type of question
					SetHeaderToQuestion(2, app.HALF). //Observe Question without answer is allowed but Answer without Question is not
					AddTextRowToQuestion(app.H3, "This is the last question!", "#e456tc", true)

	fmt.Println(string(strike.ToJson())) //ToJson() just converts the struct to json value
}

package strike_test

import (
	"testing"
	"fmt"
	strike "github.com/strike-official/go-sdk/strike"
)

func TestCard(t *testing.T) {

	strikeObj := strike.Create("first_handler", "https://abc.com") //This creates the bare minimum of the response

	question1 := strikeObj.Question("Key1"). //Creates a question object with context as Key1
						QuestionCard().                                                                   //inserts a card object in the question. Other interfaces to be used here
						SetHeaderToQuestion(2, strike.HALF_WIDTH).                                                   //Sets the header for the card object. Always next to QuestionCard function. HALF FULL
						AddGraphicRowToQuestion(strike.PICTURE_ROW, []string{"https://abc.com"}).                            //Add a photo or video to the card. pic_row video_row
						AddTextRowToQuestion(strike.H3, "News for the young!", "#e456tc", true).               //Add a text row to the card
						AddTextRowToQuestion(strike.H4, "Hey there this the top news for you", "black", false) //Add a text row to the card

	_ = question1.Answer(true). //Creates an answer object with multiselect true or false. Must be after Question and not before
					AnswerCardArray(strike.VERTICAL_ORIENTATION).                         //Add an array of Card with orientation VERTICAL HORIZONTAL
					AnswerCard().                                        //A card in the above added array
					SetHeaderToAnswer(2, strike.FULL_WIDTH).                        //Sets the header for the card object
					AddGraphicRowToAnswer(strike.PICTURE_ROW, []string{"https://abc.com","https://xyz.com"}). //Add a pic or video row to the card. pic_row video_row
					AddTextRowToAnswer(strike.H3, "wolla!", "#e456tc", true).
					AnswerCard().                 //Another card in the array
					SetHeaderToAnswer(2, strike.FULL_WIDTH). //Header for the above card
					AddTextRowToAnswer(strike.H3, "next card!", "#e456tc", true)

	_ = strikeObj.Question("Key2"). //Another question with context as key2
					QuestionCard().                 //This is again card type of question
					SetHeaderToQuestion(2, strike.HALF_WIDTH). //Observe Question without answer is allowed but Answer without Question is not
					AddTextRowToQuestion(strike.H3, "This is the last question!", "#e456tc", true)

	fmt.Println(string(strikeObj.ToJson())) //ToJson() just converts the struct to json value
}

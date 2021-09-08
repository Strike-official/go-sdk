package strike_test

import (
	"testing"
	"fmt"
	strike "github.com/strike-official/go-sdk/strike"
)

func TestLocationInput(t *testing.T) {

	strikeObj := strike.Create("location_handler", "https://abc.com") //This creates the bare minimum of the response

	question1 := strikeObj.Question("Key1"). //Creates a question object with context as Key1
						QuestionCard().                                                                   //inserts a card object in the question. Other interfaces to be used here
						SetHeaderToQuestion(2, strike.HALF_WIDTH).                                                   //Sets the header for the card object. Always next to QuestionCard function. HALF FULL
						AddGraphicRowToQuestion(strike.PICTURE_ROW, []string{"https://abc.com"}).                            //Add a photo or video to the card. pic_row video_row
						AddTextRowToQuestion(strike.H3, "News for the young!", "#e456tc", true).               //Add a text row to the card
						AddTextRowToQuestion(strike.H4, "Hey there this the top news for you", "black", false) //Add a text row to the card

	_ = question1.Answer(true). //Creates an answer object with multiselect true or false. Must be after Question and not before
				LocationInput("Select the property location")


	fmt.Println(string(strikeObj.ToJson())) //ToJson() . It just converts the struct to json value
}

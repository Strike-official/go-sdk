package strike_test

import (
	"testing"
	"fmt"
	app "github.com/strike-official/go-sdk/strike"
)

func TestLocationInput(t *testing.T) {

	strike := app.Create("location_handler", "https://abc.com") //This creates the bare minimum of the response

	question1 := strike.Question("Key1"). //Creates a question object with context as Key1
						QuestionCard().                                                                   //inserts a card object in the question. Other interfaces to be used here
						SetHeaderToQuestion(2, app.HALFCARD).                                                   //Sets the header for the card object. Always next to QuestionCard function. HALF FULL
						AddGraphicRowToQuestion(app.PICTURE_ROW, []string{"https://abc.com"}).                            //Add a photo or video to the card. pic_row video_row
						AddTextRowToQuestion(app.H3, "News for the young!", "#e456tc", true).               //Add a text row to the card
						AddTextRowToQuestion(app.H4, "Hey there this the top news for you", "black", false) //Add a text row to the card

	_ = question1.Answer(true). //Creates an answer object with multiselect true or false. Must be after Question and not before
				LocationInput()


	fmt.Println(string(strike.ToJson())) //ToJson() just converts the struct to json value
}

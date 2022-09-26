package strike_test

import (
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/go-playground/assert"
	"github.com/strike-official/go-sdk/strike"
)

func TestDo(t *testing.T) {
	inputs := []struct {
		testName     string
		notification strike.SendNotificationRequest
		wantResponse strike.SendNotificationResponse
	}{
		{
			testName: "All fields are empty",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr(""),
				AppID:  stringToPtr(""),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr(""),
					PicURL: stringToPtr(""),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("User ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "UserID is not provided",
			notification: strike.SendNotificationRequest{
				AppID: stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr("It's a test notification"),
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("User ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "AppID is not provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr("It's a test notification"),
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("App ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "Content is not provided and pic url is empty",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					PicURL: stringToPtr(""),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Either Content or PictureURL must be set to send push notification"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "Content is not provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Notification sent successfully"),
				Status: stringToPtr("sent"),
			},
		},
		{
			testName: "All things populated",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr("It's a test notification"),
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Notification sent successfully"),
				Status: stringToPtr("sent"),
			},
		},
	}

	for _, input := range inputs {
		name := input.testName
		t.Run(name, func(t *testing.T) {
			actualResponse := input.notification.Do()
			if !assert.IsEqual(*input.wantResponse.Result, *actualResponse.Result) {
				t.Fail()
			}
		})
	}
}

func TestNotificationWithPicURL(t *testing.T) {
	inputs := []struct {
		testName     string
		notification strike.SendNotificationRequest
		wantResponse strike.SendNotificationResponse
	}{
		{
			testName: "All fields are empty",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr(""),
				AppID:  stringToPtr(""),
				PushNotification: &strike.PushNotificationRequest{
					PicURL: stringToPtr(""),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("User ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "UserID is not provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr(""),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("User ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "AppId is not provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr(""),
				PushNotification: &strike.PushNotificationRequest{
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("App ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "pic url is empty",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					PicURL: stringToPtr(""),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Either Content or PictureURL must be set to send push notification"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "Only pic URL is provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Notification sent successfully"),
				Status: stringToPtr("sent"),
			},
		},
	}

	for _, input := range inputs {
		name := input.testName
		t.Run(name, func(t *testing.T) {
			actualResponse := strike.Notification(*input.notification.UserID, *input.notification.AppID).SetPictureURL(*input.notification.PushNotification.PicURL).Do()
			if !assert.IsEqual(*input.wantResponse.Result, *actualResponse.Result) {
				t.Fail()
			}
		})
	}
}

func TestNotificationWithContent(t *testing.T) {
	inputs := []struct {
		testName     string
		notification strike.SendNotificationRequest
		wantResponse strike.SendNotificationResponse
	}{
		{
			testName: "All fields are empty",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr(""),
				AppID:  stringToPtr(""),
				PushNotification: &strike.PushNotificationRequest{
					Story: stringToPtr(""),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("User ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "UserID is not provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr(""),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					Story: stringToPtr("Testing strike's go-sdk"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("User ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "AppId is not provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr(""),
				PushNotification: &strike.PushNotificationRequest{
					Story: stringToPtr("Testing strike's go-sdk"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("App ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "Content is empty",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					Story: stringToPtr(""),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Either Content or PictureURL must be set to send push notification"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "Only content is provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					Story: stringToPtr("Testing strike's go-sdk"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Notification sent successfully"),
				Status: stringToPtr("sent"),
			},
		},
	}

	for _, input := range inputs {
		name := input.testName
		t.Run(name, func(t *testing.T) {
			actualResponse := strike.Notification(*input.notification.UserID, *input.notification.AppID).SetContent(*input.notification.PushNotification.Story).Do()
			if !assert.IsEqual(*input.wantResponse.Result, *actualResponse.Result) {
				t.Fail()
			}
		})
	}
}

func TestNotificationWithContentAndPicture(t *testing.T) {
	inputs := []struct {
		testName     string
		notification strike.SendNotificationRequest
		wantResponse strike.SendNotificationResponse
	}{
		{
			testName: "All fields are empty",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr(""),
				AppID:  stringToPtr(""),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr(""),
					PicURL: stringToPtr(""),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("User ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "UserID is not provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr(""),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr("It's a test notification"),
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("User ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "AppID is not provided",
			notification: strike.SendNotificationRequest{
				AppID:  stringToPtr(""),
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr("It's a test notification"),
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("App ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "Both content and pic url is empty",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					PicURL: stringToPtr(""),
					Story:  stringToPtr(""),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Either Content or PictureURL must be set to send push notification"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "Content is not provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
					Story:  stringToPtr(""),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Notification sent successfully"),
				Status: stringToPtr("sent"),
			},
		},
		{
			testName: "All things populated",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr("It's a test notification"),
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Notification sent successfully"),
				Status: stringToPtr("sent"),
			},
		},
	}

	for _, input := range inputs {
		name := input.testName
		t.Run(name, func(t *testing.T) {
			actualResponse := strike.Notification(*input.notification.UserID, *input.notification.AppID).SetContent(*input.notification.PushNotification.Story).SetPictureURL(*input.notification.PushNotification.PicURL).Do()
			if !assert.IsEqual(*input.wantResponse.Result, *actualResponse.Result) {
				t.Fail()
			}
		})
	}
}

func TestNotificationWithTargetDateAndTime(t *testing.T) {

	tm := time.Now()
	utc := tm.UTC()
	timePlusTwoMin := utc.Add(120000000001)
	y, m, d := timePlusTwoMin.Date()

	customTargetDate := normalizedString(y) + "-" + normalizedString(int(m)) + "-" + normalizedString(d)
	h, min, sec := timePlusTwoMin.Clock()
	customTargetTime := normalizedString(h) + ":" + normalizedString(min) + ":" + normalizedString(sec)
	log.Println(customTargetDate, customTargetTime)

	inputs := []struct {
		testName     string
		notification strike.SendNotificationRequest
		wantResponse strike.SendNotificationResponse
	}{
		{
			testName: "All fields are empty",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr(""),
				AppID:  stringToPtr(""),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr(""),
					PicURL: stringToPtr(""),
				},
				TargetTime: stringToPtr(""),
				TargetDate: stringToPtr(""),
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("User ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "UserID is not provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr(""),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr("It's a test notification"),
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
				TargetTime: stringToPtr("12:45:00"),
				TargetDate: stringToPtr("2022-09-24"),
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("User ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "AppID is not provided",
			notification: strike.SendNotificationRequest{
				AppID:  stringToPtr(""),
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr("It's a test notification"),
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
				TargetTime: stringToPtr("12:45:00"),
				TargetDate: stringToPtr("2022-09-24"),
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("App ID not provided"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "Both content and pic url is empty",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					PicURL: stringToPtr(""),
					Story:  stringToPtr(""),
				},
				TargetTime: stringToPtr("12:45:00"),
				TargetDate: stringToPtr("2022-09-24"),
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Either Content or PictureURL must be set to send push notification"),
				Status: stringToPtr("failed_to_send"),
			},
		},
		{
			testName: "Content is not provided",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
					Story:  stringToPtr(""),
				},
				TargetTime: stringToPtr(customTargetTime),
				TargetDate: stringToPtr(customTargetDate),
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Notification scheduled at " + customTargetDate + " " + customTargetTime),
				Status: stringToPtr("sent"),
			},
		},
		{
			testName: "All things populated",
			notification: strike.SendNotificationRequest{
				UserID: stringToPtr("623efb9195ba637fe92fb07b"),
				AppID:  stringToPtr("629c800e2c5817b22d7b3066"),
				PushNotification: &strike.PushNotificationRequest{
					Story:  stringToPtr("It's a test notification"),
					PicURL: stringToPtr("https://cdn.pixabay.com/photo/2016/02/19/15/46/labrador-retriever-1210559__480.jpg"),
				},
				TargetTime: stringToPtr(customTargetTime),
				TargetDate: stringToPtr(customTargetDate),
			},
			wantResponse: strike.SendNotificationResponse{
				Result: stringToPtr("Notification scheduled at " + customTargetDate + " " + customTargetTime),
				Status: stringToPtr("sent"),
			},
		},
	}

	for _, input := range inputs {
		name := input.testName
		t.Run(name, func(t *testing.T) {
			actualResponse := strike.Notification(*input.notification.UserID, *input.notification.AppID).SetContent(*input.notification.PushNotification.Story).SetPictureURL(*input.notification.PushNotification.PicURL).SetTargetDateUTC(*input.notification.TargetDate).SetTargetTimeUTC(*input.notification.TargetTime).Do()
			log.Println("actualResponse", *actualResponse.Result)
			if !assert.IsEqual(*input.wantResponse.Result, *actualResponse.Result) {
				t.Fail()
			}
		})
	}
}

func stringToPtr(s string) *string {
	return &s
}

func normalizedString(i int) string {
	s := strconv.Itoa(i)
	if i/10 == 0 {
		s = "0" + strconv.Itoa(i)
	}
	return s
}

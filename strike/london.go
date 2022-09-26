package strike

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type SendNotificationRequest struct {
	UserID           *string                  `json:"user_id"`
	AppID            *string                  `json:"app_id"`
	PushNotification *PushNotificationRequest `json:"push_notification"`
	TargetTime       *string                  `json:"target_time"`
	TargetDate       *string                  `json:"target_date"`
}

var notificationRequest *SendNotificationRequest

type PushNotificationRequest struct {
	Story  *string `json:"story"`
	PicURL *string `json:"pic_url"`
}

type SendNotificationResponse struct {
	NotificationID *string `json:"notification_id"`
	Result         *string `json:"result"`
	Status         *string `json:"status"`
}

func Notification(userID, appID string) *SendNotificationRequest {
	emptyString := ""
	pushNotification := &PushNotificationRequest{
		Story:  &emptyString,
		PicURL: &emptyString,
	}

	notificationRequest = &SendNotificationRequest{
		UserID:           &userID,
		AppID:            &appID,
		PushNotification: pushNotification,
	}
	return notificationRequest
}

func (payload *SendNotificationRequest) SetContent(text string) *SendNotificationRequest {
	if payload != nil && payload.PushNotification != nil {
		if payload.PushNotification.Story != nil {
			payload.PushNotification.Story = &text
		}
	}
	return payload
}

func (payload *SendNotificationRequest) SetPictureURL(url string) *SendNotificationRequest {
	if payload != nil && payload.PushNotification != nil {
		if payload.PushNotification.PicURL != nil {
			payload.PushNotification.PicURL = &url
		}
	}
	return payload
}

func (payload *SendNotificationRequest) SetTargetTimeUTC(targetTime string) *SendNotificationRequest {
	if payload != nil {
		payload.TargetTime = &targetTime
	}
	return payload
}

func (payload *SendNotificationRequest) SetTargetDateUTC(targetDate string) *SendNotificationRequest {
	if payload != nil {
		payload.TargetDate = &targetDate
	}
	return payload
}

func (payload *SendNotificationRequest) Do() *SendNotificationResponse {

	return validatePayload(payload)

}

func validatePayload(payload *SendNotificationRequest) *SendNotificationResponse {
	if payload == nil {
		return setResponseValue("", "malformed builder", "failed_to_send")
	}
	if payload.UserID == nil || (payload.UserID != nil && *payload.UserID == "") {
		return setResponseValue("", "User ID not provided", "failed_to_send")
	}
	if payload.AppID == nil || (payload.AppID != nil && *payload.AppID == "") {
		return setResponseValue("", "App ID not provided", "failed_to_send")
	}
	if payload.PushNotification == nil {
		return setResponseValue("", "malformed builder, notification not provided", "failed_to_send")
	}
	if (payload.PushNotification.Story == nil || (payload.PushNotification.Story != nil && *payload.PushNotification.Story == "")) && (payload.PushNotification.PicURL == nil || (payload.PushNotification.PicURL != nil && *payload.PushNotification.PicURL == "")) {
		return setResponseValue("", "Either Content or PictureURL must be set to send push notification", "failed_to_send")
	}

	if payload.TargetDate != nil {
		if payload.TargetTime == nil {
			return setResponseValue("", "Must provide target time if target date is being used", "failed_to_send")
		}
	}
	if payload.TargetTime != nil {
		if payload.TargetDate == nil {
			return setResponseValue("", "Must provide target date if target time is being used", "failed_to_send")
		}
	}

	return sendHTTPPostRequest(payload)
}

func sendHTTPPostRequest(payload *SendNotificationRequest) *SendNotificationResponse {
	// Hit send notification API
	json_data, err := json.Marshal(payload)
	if err != nil {
		return setResponseValue("", "Internal server error: "+err.Error(), "failed_to_send")
	}

	resp, err := http.Post("https://shashank:prakash@london.bybrisk.com/notification/send/push", "application/json",
		bytes.NewBuffer(json_data))

	if err != nil {
		return setResponseValue("", "Internal server error: "+err.Error(), "failed_to_send")
	}

	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	if res != nil {
		return setResponseValue(res["notification_id"].(string), res["result"].(string), res["status"].(string))
	}
	return setResponseValue("", "Unknown error", "failed_to_send")
}

func setResponseValue(notificationID, result, status string) *SendNotificationResponse {
	response := &SendNotificationResponse{
		NotificationID: &notificationID,
		Result:         &result,
		Status:         &status,
	}
	return response
}

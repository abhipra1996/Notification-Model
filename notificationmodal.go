package notificationmodal

import "gorm.io/gorm"

// Function to create Notification Object
func CreateNotification(Db *gorm.DB, params map[string]interface{}) (Notification, bool, string, string) {
	var notification Notification

	res, errorCode, errorDesc := ValidatePresenceOfParams(params, "entity_type", "type", "request_data")

	if res {

		fields := []string{"entity_type", "type", "request_data", "entity_id"}
		createParams := PermitParams(fields, params)
		notification.EntityType = createParams["entity_type"].(string)
		notification.NotificationType = createParams["type"].(string)
		notification.RequestData = createParams["request_data"].(map[string]interface{})

		if createParams["entity_id"] != nil {
			notification.EntityID = createParams["entity_id"].(string)
		}
		err := Db.Create(&notification).Error
		if err != nil {
			errorCode = "E002"
			errorDesc = err.Error()
			res = false
		}
	}

	return notification, res, errorCode, errorDesc
}

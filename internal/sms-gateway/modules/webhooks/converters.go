package webhooks

import (
	"github.com/BaskBoomy/client-go/smsgateway"
)

func webhookToDTO(model *Webhook) smsgateway.Webhook {
	return smsgateway.Webhook{
		ID:       model.ExtID,
		DeviceID: model.DeviceID,
		URL:      model.URL,
		Event:    model.Event,
	}
}

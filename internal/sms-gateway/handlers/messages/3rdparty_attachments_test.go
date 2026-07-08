package messages

import (
	"strings"
	"testing"

	"github.com/android-sms-gateway/client-go/smsgateway"
)

func TestValidateAttachmentSizes(t *testing.T) {
	// base64 inflates ~4/3. Decoded limit is 300*1024 bytes ≈ 409600 base64 chars.
	underLimit := strings.Repeat("A", 400*1024) // decoded ~300KB, under limit
	overLimit := strings.Repeat("A", 420*1024)  // decoded ~315KB, over limit

	tests := []struct {
		name    string
		atts    []smsgateway.Attachment
		wantErr bool
	}{
		{"nil ok", nil, false},
		{"under limit ok", []smsgateway.Attachment{{ContentType: "image/jpeg", Data: underLimit}}, false},
		{"over limit rejected", []smsgateway.Attachment{{ContentType: "image/jpeg", Data: overLimit}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateAttachmentSizes(tt.atts)
			if (err != nil) != tt.wantErr {
				t.Fatalf("validateAttachmentSizes() err=%v wantErr=%v", err, tt.wantErr)
			}
		})
	}
}

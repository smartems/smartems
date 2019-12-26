package notifications

import (
	m "github.com/smartems/smartems/pkg/models"
	"github.com/smartems/smartems/pkg/setting"
)

// AttachedFile is struct representating email attached files
type AttachedFile struct {
	Name    string
	Content []byte
}

// Message is representation of the email message
type Message struct {
	To            []string
	From          string
	Subject       string
	Body          string
	Info          string
	ReplyTo       []string
	EmbededFiles  []string
	AttachedFiles []*AttachedFile
}

func setDefaultTemplateData(data map[string]interface{}, u *m.User) {
	data["AppUrl"] = setting.AppUrl
	data["BuildVersion"] = setting.BuildVersion
	data["BuildStamp"] = setting.BuildStamp
	data["EmailCodeValidHours"] = setting.EmailCodeValidMinutes / 60
	data["Subject"] = map[string]interface{}{}
	if u != nil {
		data["Name"] = u.NameOrFallback()
	}
}

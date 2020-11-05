package amqp

import "github.com/gofrs/uuid"

type AqmpPayload struct {
	Language string `json:"language"`
	TraceId uuid.UUID `json:"traceId"`
	Sender string `json:"sender"`
	ObjectType string `json:"objectType"`
	Payload interface{} `json:"payload"`
	ResponseQueue *string `json:"responseQueue"`
	EventName *string `json:"eventName"`
}

func (a AqmpPayload) Init()  {
	if len(a.Language) == 0 {
		a.Language = "ru"
	}
	if len(a.TraceId) == 0 {
		uid,err := uuid.NewV4()
		if err != nil{
			panic(err)
		}
		a.TraceId = uid
	}
}

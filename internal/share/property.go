package share

type (
	PropertyName string
)

type ComponentName string

const (
	ComponentVCalendar ComponentName = "VCALENDAR"
	ComponentVEvent    ComponentName = "VEVENT"
	ComponentVAlarm    ComponentName = "VALARM"
)

type RELATED struct {
	UID     string
	RelType *string // PARENT / CHILD / SIBLING
}

func NewRELATED(uid string) RELATED {
	return RELATED{UID: uid}
}

type RequestStatus struct {
	Code        string // 2.0
	Description string
	Extra       *string
}

type ATTACH struct {
	URI  *string // http://...
	Data []byte  // future: base64

	FmtType *string // FMTTYPE
}

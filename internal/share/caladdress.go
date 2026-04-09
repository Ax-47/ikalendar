package share

type CALADDRESS struct {
	Address string // mailto:xxx

	CN       *string // Common Name
	Role     *string // REQ-PARTICIPANT
	PartStat *string // ACCEPTED / DECLINED
	RSVP     *bool

	CUTYPE        *string // INDIVIDUAL / GROUP
	DelegatedTo   []string
	DelegatedFrom []string
}

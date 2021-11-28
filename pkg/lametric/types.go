package lametric

// NotificationPriority is a priority for a notification.
type NotificationPriority string

// Notification Priorities
const (
	NotificationPriorityInfo     = "info"
	NotificationPriorityWarning  = "warning"
	NotificationPriorityCritical = "critical"
)

// IconType is a type for a icons.
type IconType string

// Icon Types
const (
	IconTypeNone  = "none"
	IconTypeInfo  = "info"
	IconTypeAlert = "alert"
)

// Notification is the input for .
type Notification struct {
	Priority NotificationPriority `json:"priority,omitempty"`
	IconType IconType             `json:"iconType,omitempty"`
	Lifetime int                  `json:"lifetime,omitempty"`
	Model    NotificationModel    `json:"model"`
}

// NotificationModel is a component.
type NotificationModel struct {
	Frames []Frame `json:"frames,omitempty"`
	Sound  *Sound  `json:"sound,omitempty"`
	Cycles int     `json:"cycles,omitempty"`
}

// Frame is a component.
type Frame struct {
	Icon      string    `json:"icon,omitempty"`
	Text      string    `json:"text,omitempty"`
	GoalData  *GoalData `json:"goalData,omitempty"`
	ChartData []int     `json:"chartData,omitempty"`
}

// GoalData is a gauge for a notification.
type GoalData struct {
	Start   float64 `json:"start,omitempty"`
	Current float64 `json:"current,omitempty"`
	End     float64 `json:"end,omitempty"`
	Unit    string  `json:"unit,omitempty"`
}

// SoundCategory is a category for sounds.
type SoundCategory string

// Sound Categories
const (
	SoundCategoryAlarms        = "alarms"
	SoundCategoryNotifications = "notifications"
)

// Sound is options for sound notifications.
type Sound struct {
	Category SoundCategory `json:"category,omitempty"`
	ID       SoundID       `json:"id,omitempty"`
	Repeat   int           `json:"repeat,omitempty"`
}

// SoundID is a sound identifier
type SoundID string

// Sound IDs
const (
	SoundNotificationBicycle     = "bicycle"
	SoundNotificationCar         = "car"
	SoundNotificationCash        = "cash"
	SoundNotificationCat         = "cat"
	SoundNotificationDog         = "dog"
	SoundNotificationDog2        = "dog2"
	SoundNotificationEnergy      = "energy"
	SoundNotificationKnockKnock  = "knock-knock"
	SoundNotificationLetterEmail = "letter_email"
	SoundNotificationLose1       = "lose1"
	SoundNotificationLose2       = "lose2"
	SoundNotificationNegative1   = "negative1"
	SoundNotificationNegative2   = "negative2"
	SoundNotificationNegative3   = "negative3"
	SoundNotificationNegative4   = "negative4"
	SoundNotificationNegative5   = "negative5"
	SoundNotification            = "notification"
	SoundNotification2           = "notification2"
	SoundNotification3           = "notification3"
	SoundNotification4           = "notification4"
	SoundNotificationOpenDoor    = "open_door"
	SoundNotificationPositive1   = "positive1"
	SoundNotificationPositive2   = "positive2"
	SoundNotificationPositive3   = "positive3"
	SoundNotificationPositive4   = "positive4"
	SoundNotificationPositive5   = "positive5"
	SoundNotificationPositive6   = "positive6"
	SoundNotificationStatistic   = "statistic"
	SoundNotificationThunder     = "thunder"
	SoundNotificationWater1      = "water1"
	SoundNotificationWater2      = "water2"
	SoundNotificationWin         = "win"
	SoundNotificationWin2        = "win2"
	SoundNotificationWind        = "wind"
	SoundNotificationWindShort   = "wind_short"

	SoundAlarm1  = "alarm1"
	SoundAlarm2  = "alarm2"
	SoundAlarm3  = "alarm3"
	SoundAlarm4  = "alarm4"
	SoundAlarm5  = "alarm5"
	SoundAlarm6  = "alarm6"
	SoundAlarm7  = "alarm7"
	SoundAlarm8  = "alarm8"
	SoundAlarm9  = "alarm9"
	SoundAlarm10 = "alarm10"
	SoundAlarm11 = "alarm11"
	SoundAlarm12 = "alarm12"
	SoundAlarm13 = "alarm13"
)

// CreateNotificationOutput is the output for CreateNotification.
type CreateNotificationOutput struct {
	Success Identifier `json:"success"`
}

// Identifier is an identifier.
type Identifier struct {
	ID string `json:"id"`
}

// Icon is a constant for an icon.
type Icon string

// Icon constants
const (
	IconAppleLogo   = "i37"
	IconAttention   = "i555"
	IconBeach       = "i386"
	IconCalendar    = "i66"
	IconClock       = "i82"
	IconDoge        = "i6219"
	IconDollar      = "i34"
	IconFacebook    = "i128"
	IconFacebookAlt = "i28817"
	IconGmail       = "i43"
	IconHeart       = "i230"
	IconInstagram   = "i3741"
	IconMario       = "i3061"
	IconMatrix      = "i653"
	IconPoop        = "i8520"
	IconRSS         = "i85"
	IconSmile       = "i87"
	IconTool        = "i93"
	IconTwitch      = "i549"
	IconTwitter     = "i70"
	IconUSA         = "i413"
	IconYoutube     = "i974"
)

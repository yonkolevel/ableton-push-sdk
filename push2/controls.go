package push2

// ControlType - defines the type of a control
type ControlType string

const (
	// Cc - control of type CC
	Cc ControlType = "CC"
	// Note - control of type Note
	Note ControlType = "NOTE"
)

// Control - holds the midi control representation of a button/knob
type Control struct {
	Name string
	Type ControlType
	/**
	 * Number of the note or control
	 */
	ID uint
}

// ControlID - the id of the control
type ControlID uint8

// Soft Buttons
const (
	SoftButton1  ControlID = 20
	SoftButton2  ControlID = 21
	SoftButton3  ControlID = 22
	SoftButton4  ControlID = 23
	SoftButton5  ControlID = 24
	SoftButton6  ControlID = 25
	SoftButton7  ControlID = 26
	SoftButton8  ControlID = 27
	SoftButton9  ControlID = 102
	SoftButton10 ControlID = 103
	SoftButton11 ControlID = 104
	SoftButton12 ControlID = 105
	SoftButton13 ControlID = 106
	SoftButton14 ControlID = 107
	SoftButton15 ControlID = 108
	SoftButton16 ControlID = 109
)

//   Left hand side buttons
const (
	TapTempo    ControlID = 3
	Metronome   ControlID = 9
	Delete      ControlID = 118
	Undo        ControlID = 119
	Mute        ControlID = 60
	Solo        ControlID = 61
	Stop        ControlID = 29
	Convert     ControlID = 35
	DoubleLoop  ControlID = 117
	Quantize    ControlID = 116
	Duplicate   ControlID = 88
	NewButton   ControlID = 87
	FixedLength ControlID = 90
	Automate    ControlID = 89
	Record      ControlID = 86
	Play        ControlID = 85
	TouchStrip  ControlID = 12
)

// Pads
const (
	Pad1  ControlID = 36
	Pad2  ControlID = 37
	Pad3  ControlID = 38
	Pad4  ControlID = 39
	Pad5  ControlID = 40
	Pad6  ControlID = 41
	Pad7  ControlID = 42
	Pad8  ControlID = 43
	Pad9  ControlID = 44
	Pad10 ControlID = 45
	Pad11 ControlID = 46
	Pad12 ControlID = 47
	Pad13 ControlID = 48
	Pad14 ControlID = 49
	Pad15 ControlID = 50
	Pad16 ControlID = 51
	Pad17 ControlID = 52
	Pad18 ControlID = 53
	Pad19 ControlID = 54
	Pad20 ControlID = 55
	Pad21 ControlID = 56
	Pad22 ControlID = 57
	Pad23 ControlID = 58
	Pad24 ControlID = 59
	Pad25 ControlID = 60
	Pad26 ControlID = 61
	Pad27 ControlID = 62
	Pad28 ControlID = 63
	Pad29 ControlID = 64
	Pad30 ControlID = 65
	Pad31 ControlID = 66
	Pad32 ControlID = 67
	Pad33 ControlID = 68
	Pad34 ControlID = 69
	Pad35 ControlID = 70
	Pad36 ControlID = 71
	Pad37 ControlID = 72
	Pad38 ControlID = 73
	Pad39 ControlID = 74
	Pad40 ControlID = 75
	Pad41 ControlID = 76
	Pad42 ControlID = 77
	Pad43 ControlID = 78
	Pad44 ControlID = 79
	Pad45 ControlID = 80
	Pad46 ControlID = 81
	Pad47 ControlID = 82
	Pad48 ControlID = 83
	Pad49 ControlID = 84
	Pad50 ControlID = 85
	Pad51 ControlID = 86
	Pad52 ControlID = 87
	Pad53 ControlID = 88
	Pad54 ControlID = 89
	Pad55 ControlID = 90
	Pad56 ControlID = 91
	Pad57 ControlID = 92
	Pad58 ControlID = 93
	Pad59 ControlID = 94
	Pad60 ControlID = 95
	Pad61 ControlID = 96
	Pad62 ControlID = 97
	Pad63 ControlID = 98
	Pad64 ControlID = 99
)

// Knobs
const (
	TempoKob   ControlID = 14
	SwingKnob  ControlID = 15
	MasterKnob ControlID = 79
	Knob1      ControlID = 71
	Knob2      ControlID = 72
	Knob3      ControlID = 73
	Knob4      ControlID = 74
	Knob5      ControlID = 75
	Knob6      ControlID = 76
	Knob7      ControlID = 77
	Knob8      ControlID = 78
)

// Right hand side
const (
	SETUP         ControlID = 30
	USER          ControlID = 59
	AddDevice     ControlID = 52
	AddTrack      ControlID = 54
	DeviceBtn     ControlID = 110
	Browse        ControlID = 111
	Mix           ControlID = 112
	Clip          ControlID = 113
	MasterBtn     ControlID = 28
	Repeat        ControlID = 56
	Accent        ControlID = 57
	Scale         ControlID = 58
	Layout        ControlID = 31
	NoteBtn       ControlID = 50
	SessionBtn    ControlID = 51
	ArrowUp       ControlID = 46
	ArrowDown     ControlID = 47
	ArrowLeft     ControlID = 44
	ArrowRight    ControlID = 45
	OctaveUp      ControlID = 55
	OctaveDown    ControlID = 54
	PageLeft      ControlID = 62
	PageRight     ControlID = 63
	Shift         ControlID = 49
	Select        ControlID = 48
	GridDivision1 ControlID = 36
	GridDivision2 ControlID = 37
	GridDivision3 ControlID = 38
	GridDivision4 ControlID = 39
	GridDivision5 ControlID = 40
	GridDivision6 ControlID = 41
	GridDivision7 ControlID = 42
	GridDivision8 ControlID = 43
)

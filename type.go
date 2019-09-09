package mxl

import (
	"encoding/xml"
)

type Accidental struct {
	Text string `xml:",chardata"`
}

type Alter struct {
	Text string `xml:",chardata"`
}

// Attributes represents
type Attributes struct {
	Clef      []*Clef    `xml:"clef"`
	Divisions *Divisions `xml:"divisions"`
	Key       *Key       `xml:"key"`
	Staves    *Staves    `xml:"staves"`
	Time      *Time      `xml:"time"`
}

type Backup struct {
	Duration *Duration `xml:"duration"`
}

type BarStyle struct {
	Text string `xml:",chardata"`
}

type Barline struct {
	AttrLocation string    `xml:"location,attr"`
	BarStyle     *BarStyle `xml:"bar-style"`
}

type Beam struct {
	AttrNumber string `xml:"number,attr"`
	Text       string `xml:",chardata"`
}

type BeatType struct {
	Text string `xml:",chardata"`
}

type Beats struct {
	Text string `xml:",chardata"`
}

type BottomMargin struct {
	Text string `xml:",chardata"`
}

type Chord struct {
}

// Clef represents a clef change
type Clef struct {
	AttrNumber string `xml:"number,attr"`
	Line       *Line  `xml:"line"`
	Sign       *Sign  `xml:"sign"`
}

type Credit struct {
	AttrPage    string       `xml:"page,attr"`
	CreditWords *CreditWords `xml:"credit-words"`
}

type CreditWords struct {
	AttrDefaultX string `xml:"default-x,attr"`
	AttrDefaultY string `xml:"default-y,attr"`
	AttrFontSize string `xml:"font-size,attr"`
	AttrJustify  string `xml:"justify,attr"`
	AttrValign   string `xml:"valign,attr"`
	Text         string `xml:",chardata"`
}

type Defaults struct {
	PageLayout *PageLayout `xml:"page-layout"`
	Scaling    *Scaling    `xml:"scaling"`
}

type DisplayOctave struct {
	Text string `xml:",chardata"`
}

type DisplayStep struct {
	Text string `xml:",chardata"`
}

type Divisions struct {
	Text string `xml:",chardata"`
}

type Dot struct {
}

type Duration struct {
	Text string `xml:",chardata"`
}

// Encoding holds encoding info
type Encoding struct {
	Date     *EncodingDate `xml:"encoding-date"`
	Software *Software     `xml:"software"`
}

type EncodingDate struct {
	Text string `xml:",chardata"`
}

type Fifths struct {
	Text string `xml:",chardata"`
}

type GroupSymbol struct {
	Text string `xml:",chardata"`
}

// Identification holds all of the ident information for a music xml file
type Identification struct {
	Composer string    `xml:"creator"`
	Encoding *Encoding `xml:"encoding"`
	Rights   string    `xml:"rights"`
	Source   string    `xml:"source"`
	Title    string    `xml:"movement-title"`
}

type InstrumentName struct {
}

// Key represents a key signature change
type Key struct {
	Fifths *Fifths `xml:"fifths"`
	Mode   *Mode   `xml:"mode"`
}

type LeftMargin struct {
	Text string `xml:",chardata"`
}

type Line struct {
	Text string `xml:",chardata"`
}

// Measure represents a measure in a piece of music
type Measure struct {
	AttrWidth string        `xml:"width,attr"`
	Atters    []*Attributes `xml:"attributes"`
	Backup    []*Backup     `xml:"backup"`
	Barline   *Barline      `xml:"barline"`
	Notes     []*Note       `xml:"note"`
	Number    int           `xml:"number,attr"`
	Print     *Print        `xml:"print"`
}

type MidiChannel struct {
	Text string `xml:",chardata"`
}

type MidiInstrument struct {
	AttrId      string       `xml:"id,attr"`
	MidiChannel *MidiChannel `xml:"midi-channel"`
	MidiProgram *MidiProgram `xml:"midi-program"`
}

type MidiProgram struct {
	Text string `xml:",chardata"`
}

type Millimeters struct {
	Text string `xml:",chardata"`
}

type Mode struct {
	Text string `xml:",chardata"`
}

// MXLDoc holds all data for a music xml file
type MXLDoc struct {
	Score          xml.Name `xml:"score-partwise"`
	Identification `xml:"identification"`
	Parts          []Part `xml:"part"`
}

type Notations struct {
	Tied []*Tied `xml:"tied"`
}

// Note represents a note in a measure
type Note struct {
	AttrDefaultX string      `xml:"default-x,attr"`
	AttrDefaultY string      `xml:"default-y,attr"`
	Accidental   *Accidental `xml:"accidental"`
	Beam         []*Beam     `xml:"beam"`
	Chord        *Chord      `xml:"chord"`
	Dot          *Dot        `xml:"dot"`
	Duration     *Duration   `xml:"duration"`
	Notations    *Notations  `xml:"notations"`
	Pitch        *Pitch      `xml:"pitch"`
	Rest         *Rest       `xml:"rest"`
	Staff        *Staff      `xml:"staff"`
	Stem         *Stem       `xml:"stem"`
	Tie          []*Tie      `xml:"tie"`
	Type         string      `xml:"type"`
	Voice        *Voice      `xml:"voice"`
}

type Octave struct {
	Text string `xml:",chardata"`
}

type PageHeight struct {
	Text string `xml:",chardata"`
}

type PageLayout struct {
	PageHeight  *PageHeight  `xml:"page-height"`
	PageMargins *PageMargins `xml:"page-margins"`
	PageWidth   *PageWidth   `xml:"page-width"`
}

type PageMargins struct {
	AttrType     string        `xml:"type,attr"`
	BottomMargin *BottomMargin `xml:"bottom-margin"`
	LeftMargin   *LeftMargin   `xml:"left-margin"`
	RightMargin  *RightMargin  `xml:"right-margin"`
	TopMargin    *TopMargin    `xml:"top-margin"`
}

type PageWidth struct {
	Text string `xml:",chardata"`
}

// Part represents a part in a piece of music
type Part struct {
	AttrId  string     `xml:"id,attr"`
	Measure []*Measure `xml:"measure"`
}

type PartGroup struct {
	AttrNumber  string       `xml:"number,attr"`
	AttrType    string       `xml:"type,attr"`
	GroupSymbol *GroupSymbol `xml:"group-symbol"`
}

type PartList struct {
	PartGroup []*PartGroup `xml:"part-group"`
	ScorePart *ScorePart   `xml:"score-part"`
}

type PartName struct {
}

// Pitch represents the pitch of a note
type Pitch struct {
	Alter  *Alter  `xml:"alter"`
	Octave *Octave `xml:"octave"`
	Step   *Step   `xml:"step"`
}

type Print struct {
	AttrNewPage   string         `xml:"new-page,attr"`
	AttrNewSystem string         `xml:"new-system,attr"`
	StaffLayout   []*StaffLayout `xml:"staff-layout"`
	SystemLayout  *SystemLayout  `xml:"system-layout"`
}

type Rest struct {
	DisplayOctave *DisplayOctave `xml:"display-octave"`
	DisplayStep   *DisplayStep   `xml:"display-step"`
}

type RightMargin struct {
	Text string `xml:",chardata"`
}

type Root struct {
	ScorePartwise *ScorePartwise `xml:"score-partwise"`
}

type Scaling struct {
	Millimeters *Millimeters `xml:"millimeters"`
	Tenths      *Tenths      `xml:"tenths"`
}

type ScoreInstrument struct {
	AttrId         string          `xml:"id,attr"`
	InstrumentName *InstrumentName `xml:"instrument-name"`
}

type ScorePart struct {
	AttrId          string           `xml:"id,attr"`
	MidiInstrument  *MidiInstrument  `xml:"midi-instrument"`
	PartName        *PartName        `xml:"part-name"`
	ScoreInstrument *ScoreInstrument `xml:"score-instrument"`
}

type ScorePartwise struct {
	Credit         []*Credit       `xml:"credit"`
	Defaults       *Defaults       `xml:"defaults"`
	Identification *Identification `xml:"identification"`
	Part           *Part           `xml:"part"`
	PartList       *PartList       `xml:"part-list"`
}

type Sign struct {
	Text string `xml:",chardata"`
}

type Software struct {
	Text string `xml:",chardata"`
}

type Staff struct {
	Text string `xml:",chardata"`
}

type StaffDistance struct {
	Text string `xml:",chardata"`
}

type StaffLayout struct {
	AttrNumber    string         `xml:"number,attr"`
	StaffDistance *StaffDistance `xml:"staff-distance"`
}

type Staves struct {
	Text string `xml:",chardata"`
}

type Stem struct {
	Text string `xml:",chardata"`
}

type Step struct {
	Text string `xml:",chardata"`
}

type SystemDistance struct {
	Text string `xml:",chardata"`
}

type SystemLayout struct {
	SystemDistance    *SystemDistance    `xml:"system-distance"`
	SystemMargins     *SystemMargins     `xml:"system-margins"`
	TopSystemDistance *TopSystemDistance `xml:"top-system-distance"`
}

type SystemMargins struct {
	LeftMargin  *LeftMargin  `xml:"left-margin"`
	RightMargin *RightMargin `xml:"right-margin"`
}

type Tenths struct {
	Text string `xml:",chardata"`
}

// Tie represents whether or not a note is tied.
type Tie struct {
	AttrType string `xml:"type,attr"`
}

type Tied struct {
	AttrType string `xml:"type,attr"`
}

// Time represents a time signature change
type Time struct {
	BeatType *BeatType `xml:"beat-type"`
	Beats    *Beats    `xml:" beats"`
}

type TopMargin struct {
	Text string `xml:",chardata"`
}

type TopSystemDistance struct {
	Text string `xml:",chardata"`
}

type Voice struct {
	Text string `xml:",chardata"`
}

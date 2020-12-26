package push2

import (
	"gitlab.com/gomidi/midi/writer"
)

// SendCCOn - Sends a MIDI note on message
func (p *Device) SendCCOn(controller uint8) error {
	err := writer.CcOn(p.LivePortWriter, 60)
	return err
}

// SendCC - Sends a MIDI note on message
func (p *Device) SendCC(controller, value uint8) error {
	err := writer.ControlChange(p.LivePortWriter, controller, value)
	return err
}

// SendCCOff - Sends a MIDI note off message
func (p *Device) SendCCOff(controller uint8) error {
	err := writer.CcOff(p.LivePortWriter, controller)
	return err
}

// SendPitchBend - Sends a MIDI note off message
func (p *Device) SendPitchBend(value int16) error {
	err := writer.Pitchbend(p.LivePortWriter, value)
	return err
}

// MidiStart - Start writes the start realtime message
func (p *Device) MidiStart() error {
	wr := writer.New(p.LivePortOut)
	err := writer.RTStart(wr)
	return err
}

// MidiStop - Stop writes the start realtime message
func (p *Device) MidiStop() error {
	wr := writer.New(p.LivePortOut)
	err := writer.RTStop(wr)
	return err
}

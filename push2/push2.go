package push2

import (
	"time"

	"gitlab.com/gomidi/midi/writer"

	// "encoding/hex"
	"encoding/hex"
	"fmt"
	"log"

	"gitlab.com/gomidi/midi"

	"gitlab.com/gomidi/midi/reader"
	driver "gitlab.com/gomidi/rtmididrv"
)

// Push - the interface for Ableton Push
type Push interface {
	SendCCOn(controller uint8) error
	SendCC(controller uint8, value uint8) error
	SendCCOff(controller uint8) error
	CloseSession() error
	OpenSession() error
	MidiStart() error
	MidiStop() error
	WritePixels() error
}

type noteEventFunc func(p *reader.Position, channel, key, vel uint8)
type sysexEventFunc func(p *reader.Position, data []byte)

type notesState map[uint8]uint8

var iacDriver = "IAC Driver Live Port"

var push2LivePort = "Ableton Push 2 Live Port"
var push2UserPort = "Ableton Push 2 Live Port"

var pushIDReply = []byte{0x7E, 0x01, 0x06, 0x02, 0x00, 0x21, 0x1D, 0x67, 0x32, 0x02, 0x00, 0x01, 0x00, 0x47, 0x00, 0x7E, 0x1C, 0x3B, 0x08, 0x00, 0x01}
var deviceStatsReq = []byte{0x00, 0x21, 0x1D, 0x01, 0x01, 0x1A}
var deviceIDReq = []byte{0x7E, 0x7F, 0x06, 0x01}

var setLiveMode = []byte{0x00, 0x21, 0x1D, 0x01, 01, 0x0A, 0x00}
var padSensitivityReq = []byte{0x00, 0x21, 0x1D, 0x01, 0x01, 0x28, 0x00, 0x00, 0x01}

// Device - holds the ableton push device state
type Device struct {
	LivePortIn     midi.In
	LivePortOut    midi.Out
	UserPortIn     midi.In
	UserPortOut    midi.Out
	Driver         midi.Driver
	NotesState     notesState
	LivePortWriter writer.ChannelWriter
	UserPortWriter writer.ChannelWriter
	Reader         midi.Reader
	Display        Display
}

func must(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// New - returns a new instance of Device
func New() *Device {
	d := NewPush2Display()
	d.Open()
	return &Device{
		Display: d,
	}
}

// CloseSession - Closes current midi driver session
func (push *Device) CloseSession() error {
	push.LivePortIn.Close()
	push.LivePortOut.Close()
	push.UserPortIn.Close()
	push.UserPortOut.Close()
	return push.Driver.Close()
}

func (p *Device) WritePixels(pixels []byte) {
	p.Display.WritePixels(pixels)
}

// OpenSession - Opens a new midi driver session
func (push *Device) OpenSession() error {

	drv, err := driver.New()
	outs, _ := drv.Outs()

	log.Println(outs)

	must(err)
	push.Driver = drv

	livePortIn, err := midi.OpenIn(drv, -1, push2LivePort)
	livePortOut, err := midi.OpenOut(drv, -1, push2LivePort)

	userPortIn, err := midi.OpenIn(drv, -1, push2UserPort)
	userPortOut, err := midi.OpenOut(drv, -1, push2UserPort)

	if err != nil {
		panic(err)
	}

	must(livePortIn.Open())
	must(livePortOut.Open())
	must(userPortIn.Open())
	must(userPortOut.Open())

	lwr := writer.New(livePortOut)
	uwr := writer.New(userPortOut)

	push.LivePortWriter = lwr
	push.UserPortWriter = uwr
	push.LivePortIn = livePortIn
	push.LivePortOut = livePortOut
	push.UserPortIn = userPortIn
	push.UserPortOut = userPortOut

	rd := reader.New(
		reader.NoLogger(),
		reader.Device(func(_ reader.Position, name string) {
			log.Println(name)
		}),
		reader.NoteOn(func(p *reader.Position, channel, key, velocity uint8) {
			log.Println(key)
			//  TODO: handle events
		}),
		reader.NoteOff(func(p *reader.Position, channel, key, velocity uint8) {
			//  TODO: handle events
		}),
		reader.RTStart(func() {
			log.Println("Start msg")
			//  TODO: handle events
		}),
		reader.RTStop(func() {
			log.Println("Stop msg")
			//  TODO: handle events
		}),
		reader.SysEx(func(r *reader.Position, data []byte) {
			fmt.Sprint("Sysex got %s", string(data))
			if testEq(data, deviceIDReq) {
				log.Println("Device id requested")

				msg := pushIDReply

				time.Sleep(time.Millisecond * 500)
				err = writer.SysEx(push.LivePortWriter, msg)

				if err != nil {
					log.Println(err)
				}

				log.Println("Sent")
				log.Println(msg)
			}

			if testEq(data, deviceStatsReq) {
				log.Println("Device stats requested")
			}

			if testEq(data, setLiveMode) {
				log.Println("Set live mode requested")

				err = writer.SysEx(push.LivePortWriter, setLiveMode)

				if err != nil {
					log.Println(err)
				}

				err = writer.SysEx(push.UserPortWriter, setLiveMode)

				if err != nil {
					log.Println(err)
				}
			}
		}),
	)

	push.Reader = rd

	rd.ListenTo(livePortIn)
	rd.ListenTo(userPortIn)

	// push.activeSense(time.Millisecond * 300)

	// err = writer.SysEx(push.LivePortWriter, pushIdReply)

	// if err != nil {
	// 	log.Println(err)
	// }

	// err = writer.SysEx(push.UserPortWriter, pushIdReply)

	// if err != nil {
	// 	log.Println(err)
	// }

	// go (func (){
	// 	time.Sleep(time.Millisecond * 3000)
	// 	writer.SysEx(push.LivePortWriter, setLiveMode)
	// })()

	// go (func (){
	// 	time.Sleep(time.Millisecond * 3000)
	// 	writer.SysEx(push.UserPortWriter, setLiveMode)
	// })()

	return nil
}

func (push *Device) activeSense(tick time.Duration) {
	ticker := time.NewTicker(tick)

	wr := writer.New(push.LivePortOut)

	done := make(chan bool)
	var ticks int = 0
	go func() {
		for {
			select {
			case <-ticker.C:
				ticks++
				writer.RTActivesense(wr)
			case <-done:
				ticker.Stop()
				return
			}
		}
	}()
}

func testEq(a, b []byte) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func getDecodedHexString(src []byte) string {
	o := hex.EncodeToString(src)

	return o
}

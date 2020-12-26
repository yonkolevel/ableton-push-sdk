package push2

import (
	"sync"
	"time"

	"context"
	"errors"
	"log"

	"github.com/google/gousb"
	"github.com/google/gousb/usbid"

	"flag"
	"fmt"
)

var abletonVendorID gousb.ID = 0x2982
var pushProductID gousb.ID = 0x1967
var frameHeader = []byte{0xff, 0xcc, 0xaa, 0x88,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00}

var (
	debug = flag.Int("debug", 0, "libusb debug level (0..3)")
)

// DisplayInterface - describes the methods for the Push 2 Display
type DisplayInterface interface {
	Open() error
	Close() error
}

// Display - Ableton Push 2 Display Interface
type Display struct {
	pixels []byte
	device *gousb.Device
	ctx    *gousb.Context
	intf   *gousb.Interface
}

// NewPush2Display - returns a new instance of Push2Display
func NewPush2Display() Display {
	return Display{}
}

// Close - Closes the connection with the Push2Display
func (d *Display) Close() error {
	if d.device == nil || d.ctx == nil {
		return nil
	}

	err := d.device.Close()

	err = d.ctx.Close()

	return err
}

// Open - Opens the Ableton Push 2 usb device
func (d *Display) Open() error {
	// Only one context should be needed for an application.  It should always be closed.
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Debugging can be turned on; this shows some of the inner workings of the libusb package.
	ctx.Debug(*debug)

	device, err := ctx.OpenDeviceWithVIDPID(abletonVendorID, pushProductID)

	// OpenDevices can occasionally fail, so be sure to check its return value.
	if err != nil {
		log.Fatalf("list: %s", err)
		return err
	}

	if device == nil {
		return errors.New("Unable to open device")
	}

	log.Println("Device opened")

	d.device = device

	fmt.Printf("%03d.%03d %s:%s %s\n", device.Desc.Bus, device.Desc.Address, device.Desc.Vendor, device.Desc.Product, usbid.Describe(device.Desc))
	fmt.Printf("  Protocol: %s\n", usbid.Classify(device.Desc))

	interF, _, _ := d.device.DefaultInterface()

	d.intf = interF

	return nil
}

type contextReader interface {
	ReadContext(context.Context, []byte) (int, error)
}

// WritePixels - writes incoming pixels to the Display
func (d *Display) WritePixels(pixelschan chan []byte) error {

	outEp, _ := d.intf.OutEndpoint(1)
	ctx := context.Background()
	var wg sync.WaitGroup
	wg.Add(1)

	select {
	case pixels := <-pixelschan:
		wg.Done()
		fmt.Println("received message", pixels)
		go func(pxs []byte) {
			for {
				go outEp.WriteContext(ctx, frameHeader)
				go outEp.WriteContext(ctx, pxs)
				time.Sleep(1000)
			}
		}(pixels)
		wg.Wait()
	default:
		fmt.Println("no message received")
	}

	return nil
}

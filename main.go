package main

import (
	_ "embed"
	"machine"
	"time"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/drivers/st7789"
)

//go:embed img.raw
var data []byte

func main() {
	machine.Serial.Configure(machine.UARTConfig{
		BaudRate: 115200,
	})

	time.Sleep(2 * time.Second)

	machine.SPI1.Configure(machine.SPIConfig{
		Frequency: 62_500_000, // 62.5MHz
		SCK:       machine.GP10,
		SDO:       machine.GP11,
		SDI:       machine.GP12, // ダミー - ST7789は受信しないが、SPIドライバが要求
		Mode:      0,
	})

	display := st7789.New(
		machine.SPI1,
		machine.GP12, // RST
		machine.GP8,  // DC
		machine.GP9,  // CS
		machine.GP13, // BL
	)

	display.Configure(st7789.Config{
		Width:        240,
		Height:       320,
		Rotation:     st7789.ROTATION_90, // 90度回転で横向き表示
		RowOffset:    0,
		ColumnOffset: 0,
	})

	img := pixel.NewImageFromBytes[pixel.RGB565BE](320, 240, data)

	display.DrawBitmap(0, 0, img)
}

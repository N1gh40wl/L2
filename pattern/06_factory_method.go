package pattern

import "fmt"

type iDevice interface {
	setName(name string)
	setOS(name string)
	getName() string
	getOS() string
}

type device struct {
	name string
	OS   string
}

func (d *device) setName(name string) {
	d.name = name
}

func (d *device) setOS(OS string) {
	d.OS = OS
}

func (d *device) getName() string {
	return d.name
}

func (d *device) getOS() string {
	return d.OS
}

type iphone struct {
	device
}

func newIphone() iDevice {
	return &iphone{
		device: device{
			name: "iPhone",
			OS:   "iOS",
		},
	}
}

type PC struct {
	device
}

func newPC() iDevice {
	return &PC{
		device: device{
			name: "PC",
			OS:   "Windows 10",
		},
	}
}

func getDevice(deviceType string) (iDevice, error) {
	if deviceType == "iPhone" {
		return newIphone(), nil
	}
	if deviceType == "PC" {
		return newPC(), nil
	}
	return nil, fmt.Errorf("Wrong device type")
}

func RunFactory() {
	iphone, _ := getDevice("iPhone")
	PC, _ := getDevice("PC")

	printDevices(iphone)
	printDevices(PC)
}

func printDevices(d iDevice) {
	fmt.Println("Name:", d.getName())
	fmt.Println()
	fmt.Println("OS:", d.getOS())
	fmt.Println()
}

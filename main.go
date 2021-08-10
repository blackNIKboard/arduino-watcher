package main

import (
	"flag"
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/tarm/serial"
	_ "github.com/tklauser/go-sysconf"
	_ "golang.org/x/sys/unix"
	"log"
	"math"
	"os/user"
	"strconv"
	"time"
)

var device = flag.String("device", "", "device to call")

func main() {
	if current, _ := user.Current(); current.Username != "root" {
		//log.Fatalf("App should be started with elevated permissions, got: %s", current.Username)
	}

	flag.Parse()

	c := &serial.Config{
		Name: *device,
		Baud: 9600,
		//		ReadTimeout: time.Millisecond * 5,
	}

	port, err := serial.OpenPort(c)
	if err != nil {
		fmt.Println(err)
	}
	defer port.Close() //todo catch sigs

	for {
		data, err := parseToSend()
		if err != nil {
			log.Fatal(err)
		}
		n, err := port.Write([]byte(data + "\n"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Sent %v bytes of info: %s\n", n, data)

		time.Sleep(time.Millisecond * 1000)
	}
}

func getCpuTemp() (string, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return "NA", err
	}
	fmt.Println(cpuInfo[0].ModelName)

	sensors, err := host.SensorsTemperatures()
	if err != nil {
		return "NA", err
	}

	for _, sensor := range sensors {
		if sensor.SensorKey == "TC0P" {
			return strconv.Itoa(int(sensor.Temperature)), nil
		}
	}

	return "NA", nil
}

func getMaxTemp() (string, error) {
	sensors, err := host.SensorsTemperatures()
	if err != nil {
		return "NA", err
	}

	var maxTemp float64

	for _, sensor := range sensors {
		if maxTemp < sensor.Temperature {
			maxTemp = sensor.Temperature
		}
	}

	return strconv.Itoa(int(maxTemp)), err
}

func getFanRPM() (string, error) {
	//out, err := exec.Command("/usr/bin/powermetrics -i 200 -n1 --samplers smc").Output()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Printf("The date is %s\n", out)

	return "NA", nil
}

func getMemLoad() (string, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return "NA", err
	}

	return fmt.Sprintf("%.2f", convertToGB(v.Used)) + "/" + fmt.Sprintf("%.0f", convertToGB(v.Total)) +
		" " + fmt.Sprintf("%.2f%%", v.UsedPercent), nil
}

func getCpuLoad() (string, error) {
	info, err := cpu.Percent(0, false)
	if err != nil {
		return " NA", nil
	}

	return fmt.Sprintf("%.2d", int(info[0])) + "%", nil
}

func convertToGB(bytes uint64) float64 {
	return float64(bytes) / (math.Pow(1024, 3))
}

func parseToSend() (string, error) {
	maxTemp, err := getCpuTemp()
	if err != nil {
		return "NA", err
	}

	memory, err := getMemLoad()
	if err != nil {
		return "NA", err
	}

	cpuLoad, err := getCpuLoad()
	if err != nil {
		return "NA", err
	}

	return fmt.Sprintf("%.16s%.16s", "cpu temp: "+maxTemp+" "+cpuLoad, memory), nil
}

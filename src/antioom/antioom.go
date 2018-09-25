package antioom

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron"
)

type Instance struct {
	memoryThreshold     int
	checkMemoryDuration int
	bashCommand         []string
}

// New create antioom instance
func New(memoryThreshold int, checkMemoryDuration int) (instance *Instance) {
	instance = &Instance{
		memoryThreshold:     memoryThreshold,
		checkMemoryDuration: checkMemoryDuration,
		bashCommand:         []string{},
	}
	if instance.checkMemoryDuration <= 0 {
		instance.checkMemoryDuration = 1
	}

	return
}

// Run run the antioom instance
func (i *Instance) Run() {
	log.Println("Run Anti OOM")
	onCheck := false

	cronClient := cron.New()
	cronClient.AddFunc("@every "+strconv.Itoa(i.checkMemoryDuration)+"s", func() {
		if !onCheck {
			onCheck = true
			freeMemory, err := getCurrentFreeMemory()
			if err != nil {
				log.Println("[error]", err)
			} else if freeMemory <= i.memoryThreshold {
				log.Println("Anti OOM is triggered (current free memory ="+strconv.Itoa(freeMemory)+"): ", time.Now().Format("2006-01-02 15:04:05"))

				for _, command := range i.bashCommand {
					log.Println("bash command called:", command)
					c1 := exec.Command("/bin/bash", "-c", command)
					c1Output, errOutput := c1.Output()
					if errOutput != nil {
						log.Println("error:", errOutput)
					} else {
						log.Println("result:", string(c1Output))
					}
				}
				log.Println("------------------------------------------------------------------------")
			}
			onCheck = false
		}
	})
	cronClient.Start()
}

// getCurrentFreeMemory check current server memory (ubuntu)
func getCurrentFreeMemory() (freeMemory int, err error) {
	cmd := exec.Command("/bin/bash", "-c", "cat /proc/meminfo | grep MemAvailable | awk '{print $2}'")
	output, errOutput := cmd.Output()
	if errOutput != nil {
		err = errOutput
		freeMemory = -1
		return
	}
	freeMemoryStr := strings.Trim(string(output), "\n")
	freeMemory, err = strconv.Atoi(freeMemoryStr)
	if err != nil {
		freeMemory = -1
	}
	return
}

// AddBashCommand add bash command when memory got oom
func (i *Instance) AddBashCommand(command string) {
	i.bashCommand = append(i.bashCommand, command)
}

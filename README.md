# antioom
First Aid for Out of Memory Problem

## How to Use

Run the Service:
```
./antioom -M=500000 -C="service memleaker1 restart" -C="service memleaker1 restart"
M = memory size threshold. if memory size under the threshold, commands (-C) will be trigerred
C = bash command which will triggered when memory under the threshold
```

Using the Library:
```
go get github.com/herryg91/antioom
import "github.com/herryg91/antioom/src/antioom"

antiOOMInstance, _ := antioom.New(memthreshold, 1)
antiOOMInstance.AddBashCommand("service memleaker1 restart")
antiOOMInstance.Run()
```

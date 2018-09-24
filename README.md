# antioom
First Aid for Out of Memory Problem

## How to Use

Run the Service:
```
./antioom -M=500000 -C="service memleaker1 restart" -C="service memleaker1 restart"

args:
-M or --memory = memory threshold. if memory size under the threshold, commands (-C) will be trigerred
-C or --commands = bash command which will triggered when memory under the threshold
```

Using the Library:
```
go get github.com/herryg91/antioom

import "github.com/herryg91/antioom/src/antioom"

aoInstance, _ := antioom.New(memthreshold, 1)
aoInstance.AddBashCommand("service memleaker1 restart")
aoInstance.AddBashCommand("service memleaker2 restart")
aoInstance.Run()
```

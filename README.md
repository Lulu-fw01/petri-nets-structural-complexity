# petri-nets-structural-complexity


### Run util app

```bash 
cd cmd/util-app
go run main.go --settings /path/to/settings.json --net /path/to/net.xml
```
**flags**
* --metrics - option to choose which metric to use. "all" is default value print all versions of metrics. "v1", "v2" first and second version of metric;
* --settings - path to settings file;
* --net - path to xml file which describe net.

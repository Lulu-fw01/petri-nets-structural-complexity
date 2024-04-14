# petri-nets-structural-complexity


### Run util app

```bash 
cd cmd/util-app
go run main.go --settings /path/to/settings.json --net /path/to/net.xml
```
**flags**
* --metrics - option to choose which metric to use. "all" is default value print all versions of metrics. "v1", "v2" first and second version of metric;
* --settings-type - type of settings standard simple settings or regexp settings (default "simple", "regexp" for regexp settings);
* --settings - path to settings file;
* --net - path to xml file which describe net.


## Settings

### Simple settings

```json
{
  "agentsToTransitions": {
    "a1": [
      "t1",
      "t2",
      "t3"
    ],
    "a2": [
      "q1",
      "q2",
      "q3"
    ]
  },
  "silentTransitions": [
    "t1",
    "q2"
  ]
}
```

### Regexp settings

```json
{
  "agentToTransitionRegexp": {
    "a1": "^a1-",
    "a2": "^(a2-.*|t1|t2)$"
  },
  "silentTransitionRegexp": "^(s1|s2|s3)$"
}
```
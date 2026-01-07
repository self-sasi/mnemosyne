# This document stores the plan for how the implementation and behaviour of the tool will be.

This prints the general 'dictionary' informing the user of mnemo's features
```
mnemo
```

This backs up a database. the specifications for the database will be in the config file. 
This config file will hold database name, host, type, outputDir, etc. 
```
mnemo backup --configPath="/some/folder/config.json"
```
For now, it is unclear whether we want the config to be json or yaml

These configurations can also be passed via flags
```
mnemo backup --type="postgres" --port=1234 
```

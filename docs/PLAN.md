# This document stores the plan for how the implementation and behaviour of the tool will be.

This prints the general 'dictionary' informing the user of mnemo's features
```zsh
mnemo
```

This backs up a database. the specifications for the database will be in the config file. 
This config file will hold database name, host, type, outputDir, etc. 
```zsh
mnemo backup --configPath /some/folder/config.json
```
For now, it is unclear whether we want the config to be json or yaml

These configurations can also be passed via flags
```zsh
mnemo backup --type postgres --port 1234 
```

This restores the database. The specification regarding the database to restore and where to get the backup from the config file itself. 
```zsh
mnemo restore --configPath /some/folder/ 
# or
mnemo restore --type postgres --backupPath /backup/ --port 1234
```
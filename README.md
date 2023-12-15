Convert xshell sessions config to WindTerm sessions config.

## Build and Usage

> Note: Please Backup origin user.sessions file at first

```bash
# build
go build xshell2windterm.go
# show help
./xshell2windterm -h
# convert
./xshell2windterm -i "C:\Users\administrator\Documents\NetSarang Computer\6\Xshell\Sessions"

# Note: Please Backup origin user.sessions file at first
# Backup origin user.sessions
cp ${ProfilesDirectory}/.wind/profiles/default.v10/terminal/user.sessions user.sessions.backup_$(date "+%Y%m%d")
# import to windterm
cp -f user.sessions ${ProfilesDirectory}/.wind/profiles/default.v10/terminal/user.sessions
```


> Note

    * Please update the ssh key path manually in user.sessions.

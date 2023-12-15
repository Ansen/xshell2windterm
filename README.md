Convert xshell sessions config to WindTerm sessions config.

## Build and Usage

```bash
# build
go build xshell2windterm.go
# show help
./xshell2windterm -h
# convert
./xshell2windterm -i "C:\Users\administrator\Documents\NetSarang Computer\6\Xshell\Sessions"
# import to windterm
mv `user.sessions` `${ProfilesDirectory}/.wind/profiles/default.v10/terminal/
```


> Note

    * Please update the ssh key path manually in user.sessions.

# ArdWatcher
This is the small utility that displays needed information on I2C display using Arduino

## Installation
```bigquery
unzip arduino-watcher.zip
cp arduino-watcher /usr/local/bin
cp com.blacknikboard.ardwatcher.plist ~/Library/LaunchAgents
launchctl load ~/Library/LaunchAgents/com.blacknikboard.ardwatcher.plist
```
WIP about make
## Building
`go build -o arduino-watcher .`
## Hardware setup
![Screenshot 2021-08-10 at 11 07 58 AM](https://user-images.githubusercontent.com/18029685/128831636-baca5083-e059-4a23-9290-4f76a737ca06.png)

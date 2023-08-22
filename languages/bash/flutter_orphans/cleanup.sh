kill -9 $(ps | grep -w "/Applications/Xcode.app/Contents/Developer/usr/bin/xcdevice observe" | awk '{print $1}')
# kill -9 "$(ps aux | grep -w "/Applications/Xcode.app/Contents/Developer/usr/bin/xcdevice observe --both" | grep -v "grep" | awk '{print $2}')"

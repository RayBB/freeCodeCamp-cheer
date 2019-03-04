#!/bin/sh
# This is a script to notify when a user gets more FreeCodeCamp points. Then you can cheer them on!

USER=$1
function getPoints(){
    echo $(curl -v --silent https://www.freecodecamp.org/api/users/get-public-profile?username=$USER 2>&1 | grep -o "points\":\w\d\{1,\}" | grep -o "\d\{1,\}")
}
function notify(){
    /usr/bin/osascript -e 'display notification "'$1'" with title "FreeCodeCamp"'
}
OLDPOINTS=-1

while true; do
    CURPOINTS=$(getPoints)
    if [ "$CURPOINTS" -gt "$OLDPOINTS" ]; then
        echo "$(date)\t$CURPOINTS"
        notify "points...$CURPOINTS"
        OLDPOINTS=$CURPOINTS
    fi
    sleep 11
done;

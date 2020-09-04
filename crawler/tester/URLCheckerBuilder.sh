#!/usr/bin/env bash
DIR=$(dirname $(realpath $0))
PDIR=$(dirname $DIR)
TARGET=$PDIR/bin/URLChecker
LOCATION=$PDIR/src/URLChecker
FIN=$DIR/URLChecker

if [ -f $FIN ]; then
    echo "File already exists ... Done"
    exit 0
fi

if [ -f $TARGET ]; then
    echo "Check if file exists ... Done"
else
    echo "File not exists ..."
    echo "Build executable file ..."
    if [ -d $LOCATION ]; then
        cd $LOCATION
        go install
    else
        echo "Build directory doesn't exist ..."
        exit 0
    fi
fi

cp $TARGET $DIR
chmod 755 $FIN
echo "Move file ... Done"
#!/usr/bin/env bash

log() {
  printf "  \033[36m%10s\033[0m : \e[2m%s\e[22m\033[0m\n" $1 $2
}

abort() {
  printf "\n  \033[31mError: $@\033[0m\n\n" && exit 1
}

# ensure setx exists
command -v setx >/dev/null 2>&1 || { echo >&2 "setx not exists，install first!"; exit 1; }

command -v curl >/dev/null 2>&1 || { echo >&2 "curl not exists，install first!"; exit 1; }

# ensure HOME exist in env
if ! test -d $HOME; then
  abort "HOME env not exist!"
fi

DEST=$HOME/.gost
URL="http://congz.pw/gost-cli/gost.exe"
NAME="gost.exe"

# makedir DEST
if test -d $DEST; then
  log $DEST exists
else
  log mkdir $DEST
  mkdir $DEST
fi

if test -e $DEST/$NAME; then
  printf "\n $DEST/$NAME exists, will replace it\n"
fi

cd $DEST
printf "\ndownload start $URL\n\n"
curl $URL -o $NAME

if ! test -e $DEST/$NAME; then
  abort "download error, please try again later"
fi

command -v gost >/dev/null 2>&1 && log path "already set" || {
  log set $DEST to PATH
  setx PATH "%PATH%;$DEST"
 }

log install finished
printf "\nopen a new terminal and try it.\n"

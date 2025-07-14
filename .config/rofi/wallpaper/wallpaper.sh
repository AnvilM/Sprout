#!/bin/bash

DIR="$HOME/.config/rice/wallpapers"

if [ -z "$1" ]; then
  for a in "$DIR"/*; do
    echo -en "$(basename "$a")\0icon\x1f$a\n"
  done
else
  swww img "$DIR/$1" -t wipe --transition-duration 2 --transition-fps 60
fi
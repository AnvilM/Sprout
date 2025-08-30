#!/bin/bash

prompt='Cliphist'

cliphist list | \
    rofi -dmenu \
    -display-columns 2 \
    -p $prompt \
    -theme ~/.rice/rofi/cliphist/cliphist.rasi | \
    cliphist decode | \
    wl-copy
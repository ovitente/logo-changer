#!/bin/bash

# This script resizes images and attaches to it border with text img under it.
# You have to install imagemagick to be able to use it.
# Ignore errors, just look if you have result.

# convert -scale 40% pravda-zhizni.png pravda-zhizni-rdy.png && 
# convert pravda-zhizni-rdy.png -bordercolor black -border 1 pravda-zhizni-rdy2.png && 
# convert -append pravda-zhizni-rdy2.png under-line.png pravda-zhizni-rdy.png


SOURCE_DIR=../img-source
TARGET_DIR=../img-ready

for i in $(ls); do convert -scale 40% $SOURCE_DIR/$i $TARGET_DIR/$i; done;
for i in $(ls); do convert $SOURCE_DIR/$i -bordercolor black -border 1 $TARGET_DIR/$i; done;
for i in $(ls); do convert -append $SOURCE_DIR/$i ../under-line.png $TARGET_DIR/$i; done;
echo "Done."

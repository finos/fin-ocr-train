#!/bin/bash
# Copyright (c) 2024 Discover Financial Services

if [ $# -ne 1 ]; then
   echo "Usage: train <numImages>"
   exit 1
fi


#rm "nohup.out"

# The number of images to generate and to train over
NUM_IMAGES=$1

# One epoch means to iterate once over all of the images
# Setting it to 5 is just a guess
NUM_EPOCHS=2

echo "----------------------------------"

start_date_time=$(date)
echo "Start date and time: $start_date_time"
echo "Num Images: $NUM_IMAGES"
echo "Num Epochs: $NUM_EPOCHS"

echo "Removing previous training data"
rm -rf out
rm -rf data/micr*
export OVERLAP=50

echo "Generating ground truth"
GT_DIR=./data/micr_FT-ground-truth
python3 ./gen_ground_truth_trdg.py $NUM_IMAGES $GT_DIR

delete_start_date_time=$(date)
echo "Delete start date and time: $delete_start_date_time"
echo "Removing zero size box files"
find . -maxdepth 7 -type f -exec bash -c 'if [ ! -s "$1" ]; then BASE=${1%.*}; rm -f $BASE.*; fi' _ {} \;

echo "Start training"
gmake training MODEL_NAME=micr_FT START_MODEL=micr PSM=13 EPOCHS=$NUM_EPOCHS TARGET_ERROR_RATE=-1 TESSDATA=tessdata

end_date_time=$(date)
echo "End date and time: $end_date_time"

echo "----------------------------------"

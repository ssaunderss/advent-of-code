#!/bin/bash

echo "Starting Download Script...";

while getopts d:c: flag
do
    case "${flag}" in
        d) day=${OPTARG};;
        c) cookie=${OPTARG};;
    esac
done

echo "Day: $day";
echo "Cookie: $cookie";

curl "https://adventofcode.com/2022/day/$day/input" --cookie "session=$cookie" >> "$day/input-1.txt";

echo "Finished Downloading, input file located at $day/input-1.txt";
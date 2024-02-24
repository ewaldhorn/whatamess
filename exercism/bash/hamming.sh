#!/usr/bin/env bash
if (( $# != 2 )) 
then
    echo "Usage: hamming.sh <string1> <string2>"
    exit 1
fi
line1=$1
line2=$2
len1=${#line1}
len2=${#line2}
if (( $len1 != $len2 ))
then
    echo "strands must be of equal length"
    exit 1
fi
if (( $len1 == 0 ))
then
    echo 0
    exit 0
fi
diff=0
for ((i = 0; i < ${#line1}; i++)); do
    if [ "${line1:i:1}" != "${line2:i:1}" ]
    then
        diff=$((diff+=1))
    fi
done
echo $diff


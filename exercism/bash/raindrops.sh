#!/usr/bin/env bash
done=false
if [[ $(expr $1 % 3) == 0 ]] 
then
    echo -n "Pling"
    done=true
fi
if [[ $(expr $1 % 5) == 0 ]] 
then
    echo -n "Plang"
    done=true
fi
if [[ $(expr $1 % 7) == 0 ]] 
then
    echo -n "Plong"
    done=true
fi

if [[ $done == false ]] 
then
    echo $1
fi



#!/bin/bash

if [ `alias go 2> /dev/null | wc -l` == 0 ]; then
	# get absolute path to current script
	abs=$( cd `dirname "$BASH_SOURCE"` ; pwd -P)
	alias go=". $abs/go go"
fi

# ------------------------------------------------------
help() {
# ------------------------------------------------------
# Affichage de l'aide
#-------------------------------------------------------
cat << 'END'
go <rep1/rep2/...> :
        Allows to jump into a place by providing only the first characters of each intermediate directory.
        If there is only one path matching, then it jumps directly into it.
        Otherwise, it lists possible paths and ask user for choice.
END
}


#=======================================================
# help
#=======================================================
if [ "$2" = "-h" ]
then
  help
  return
fi

#=======================================================
# go
#=======================================================
if [ "$1" = "go" ]; then
  if [ "$2" = "" ]; then
    echo "-> $HOME"
    cd $HOME
    return
  fi
  gp=`echo "$2" | sed -e 's/\([^\/]\)$/\1*/' -e 's/\//\*\//g'`
  lst=`/bin/ls -d $gp 2> /dev/null`
  nb=0
  for val in $lst
  do
    if [ -d $val ]; then
      nb=$((nb+1))
      arrP[$nb]=$val
    fi
  done
  if [ $nb -eq 0 ]; then
    echo "-> Not found !"
  elif [ $nb -eq 1 ]; then
    echo "-> ${arrP[1]}"
    cd ${arrP[1]}
  else
    choice=
    while [ "$choice" = "" ]
    do
      i=1
      while [ $i -le $nb ]
      do
        echo "$i ${arrP[$i]}/"
        i=$((i+1))
        if [ $i -eq 9 ]; then
          echo "..."
          break
        fi
      done
      printf "> "
      read num
      choice=${arrP[$num]}
    done
    echo "-> $choice"
    cd $choice
  fi
fi
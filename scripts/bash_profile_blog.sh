#!/bin/bash
#
# copypaste contents of this file to ~/.bash_profile

BLOG_HOME="/Users/ndrw/blog"
export BLOG_HOME

# TODO: use $BLOG_HOME? or generate this file  
alias blog='$BLOG_HOME/./blog'

_blog() 
{
local cur goals

  COMPREPLY=()
  cur=${COMP_WORDS[COMP_CWORD]}
  # todo exection blog --list
  goals="$($BLOG_HOME/./blog autocomplete)"
  
  echo goals
  cur=`echo $cur`
  COMPREPLY=($(compgen -W "${goals}" ${cur}))
}

complete -F _blog blog 
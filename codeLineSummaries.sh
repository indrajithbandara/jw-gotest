#!/bin/bash
#author : cuijun 20161223
#pwd
seldate=`date -d "1 weeks ago" +%Y-%m-%d`
echo " the current date is ->$seldate"
cd /E/IDEidea/IdeaProjects
projectname=jw-source
sourceextession=.java
echo " the project name is ->$projectname-$sourceextession"
pwd
annonumbner=`find $projectname -name "*$sourceextession" |xargs cat | grep -v ^$|wc -l`    #Output:36335
# exclude the lines begin with //
#noannonumbner=find . -name "*.java"|xargs cat|grep -v -e ^$ -e ^\s*\/\/.*$|wc -l    #Output:36064
noannonumbner=`find $projectname -name *$sourceextession | xargs cat | grep -v -e ^$ -e ^\s*\/\/.*$ | wc -l`
echo " the project has annotation codeline is ->$annonumbner"
echo " the project has no annotation codeline is ->$noannonumbner"

percent_1=$(printf "%d%%" $((noannonumbner*100/annonumbner)))
# no 0.1
percent_2=`awk 'BEGIN{printf "%.1f%%\n",('$noannonumbner'/'$annonumbner')*100}'`
# save 0.1
echo $percent_1
echo $percent_2

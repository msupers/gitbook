#!bash

#step 1 
#clear output dir and output.tar.gz
rm -rf output 

rm -rf output.tar.gz

# step 2
# copy files to output dir
mkdir  -p ./output

for i in `ls|grep -v build.sh`

do

cp -r $i ./output  

done

# step 3
#got output.tar.gz
tar -zcvf output.tar.gz output/
rm -rf ./output
# step 4 
# submit output.tar.gz  to http server 
#cp -r output.tar.gz /home/files/

#step 5
# submit to docker registry

#!/bin/bash
### 写脚本实现，可以用shell、perl等。在目录/tmp下找到100个以abc开头的文件，然后把这些文件的第一行保存到文件new中。 



for i in `find /tmp -type  f -name "abc*"|head -n 100`
do
    sed '1p' $i>>new
done

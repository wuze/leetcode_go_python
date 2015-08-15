#!/bin/bash python
#encoding:utf-8

### 主要思想是采用队列
### queue


def displayRet(num,total):
    print "Total: ", total ," Num:",num 
    print  range(1, total+1)

    if total==1:
        print 1
    arr=range(1,total+1)

    while len(arr)>0:
        for i in range(num-1):
            arr.append(arr.pop(0))
        print arr.pop(0), ", ",


if __name__=="__main__":
    total=10
    num=5
    displayRet(num,total)

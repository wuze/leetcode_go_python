#!/usr/bin/env python
#encoding:utf-8


def displayData(line):
    arr=[([0]*line)  for i in range(line)]
    arr[0][0]=1

    for i in range(1,line):
        for j in range(i+1):
            if j==0:
                arr[i][0]=1
                continue

            if j==i:
                arr[i][j]=1
            else:
                arr[i][j] = arr[i-1][j-1]+arr[i-1][j]



    for i in range(line):
        for j in range(line-i):
            print " ",
        for j in range(len(arr[i])):
            if arr[i][j]>0:
                print arr[i][j], " ",
        print "\n"

    
    return





if __name__=="__main__":
    line=7
    displayData(line)

    


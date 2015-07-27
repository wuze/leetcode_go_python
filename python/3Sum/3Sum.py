#!/usr/bin/env python
#encoding:utf-8


def mySum(*arr):
    lenth=len(arr)
    sorted(arr)
    ret =[]

    for i in range(lenth-2): 
        node=arr[i]
        low=i+1;
        high=lenth-1

        while low<high:
            a=arr[low]
            b=arr[high]

            if node+a+b==0:
                ret.append([node, a, b])
                while low<lenth and arr[low] ==arr[low+1]:
                    low+=1
                while high>0 and arr[high] ==arr[high-1]:
                    high-=1

                low+=1
                high-=1
            elif node+b+a>0:
                while high>0 and arr[high]==arr[high-1]:
                    high-=1
                high-=1
            else:
                while low<lenth and arr[low]==arr[low+1]:
                    low+=1
                low+=1
        i+=1
    return ret       


if __name__=="__main__":
    print [-4,-1,0,1,2]
    ret=mySum(-4 , -1 , 0 , 1 , 2)

    print "result:"
    print ret
    


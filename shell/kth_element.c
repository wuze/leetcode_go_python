#include <iostream>
#include <cstdio>
#include <cstdlib>

using namespace std;

int kth_elem(int a[],int low, int high,int k)
{
    int pivot     = a[low];/*{{{*/
    int low_temp  = low;
    int high_temp = high;

    while(low<high)
    {
        while(low<high&&a[high]>=pivot)     --high;
        a[low]=a[high];

        while(low<high&&a[low]<pivot) ++low;
        a[high]=a[low];
    }

    a[low]=pivot;
    
    //以下就是主要思想中所述的内容
    if(low==k-1) 
    {
        return a[low];
    }
    else if(low>k-1) 
    {
        return kth_elem(a,low_temp,low-1,k);
    }
    else 
    {
        return kth_elem(a,low+1,high_temp,k);
    }
/*}}}*/
}

void Swap(int *a, int *b)
{
    int c;
    c = *a;
    *a = *b;
    *b = c;
}

int divide(int *vi, int low, int up)  
{  
    int pivot = vi[up];  
    int i = low-1;  
    for (int j = low; j < up; j++)  
    {  
        if(vi[j] <= pivot)  
        {  
            i++;  
            Swap(&vi[i], &vi[j]);  
        }  
    }  
    Swap(&vi[i+1], &vi[up]);  
    return i+1;  
}



void quick(int *a, int l, int h)
{
/*{{{*/
    if(l>=h) return;

    int mid = divide(a,l,h);

    quick(a, l,mid-1);
    quick(a, mid+1, h);
/*}}}*/
}

void print(int a[], int l)
{
/*{{{*/
    for(int i=0;i<l;i++)
    {
        printf("%d\t", a[i]);
    }

    printf("\n");/*}}}*/
}

int St(int *a, int length)
{
    quick(a, 0, length-1);
    print(a,length);
}

int main()
{
    int len=10;
    int a[len];/*{{{*/
    int k;

    for(int i=0;i<len;i++)
    {
        a[i] = rand()%100;    //随机地生成序列中的数
        printf("%3d  ",a[i]);
        if(i%5==4) 
        cout << endl;
    }

    printf("\nAfter Sort\n");

    St(a,len);

    cin >> k;
    cout << "the" << k << "th number is:" << kth_elem(a,0,len-1,k)<<endl;

    getchar();
    return 0;/*}}}*/
}

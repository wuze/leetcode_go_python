#include<stdio.h>

struct LinkNode{
    int val;
    LinkNode * next;
};


LinkNode *reverseLinkedList3(LinkNode* head)
{  
    if(head==NULL||head->next==NULL)  
        return  head;  

    LinkNode* p    = head; //指向head
    LinkNode* r    = head->next; //指向待搬运的节点，即依次指向从第2个节点到最后一个节点的所有节点
    LinkNode* m    = NULL; //充当搬运工作用的节点
    LinkNode* tail = head->next;

    while(r!=NULL)
    { 
        m       = r;
        r       = r->next;
        m->next = p->next;
        p->next = m;
    }  

    head       = p->next;
    tail->next = p;
    p->next    = NULL;
    tail       = p;

    return head;
} 


int main()
{

    int a[] = {1,2,3,4,5,6,7};
    int len = sizeof(a)/sizeof(a[0]);

    LinkNode * head;
    LinkNode *t  = new LinkNode;
    t->val     = 1;
    t->next    = NULL;
    head->next = t;

    for(int i=0;i<len;i++)
    {
        LinkNode *c= new LinkNode;
        c->val  = i;
        c->next = NULL;

        return 0;
        t->next = c;
        t       = t->next;
    }

    LinkNode *m = head;

    while(m->next)
    {
        printf("%d\t",m->val);
        m = m->next;
    }

    printf("Over\n");
    return 0;
}

/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
class Solution {
public:
    ListNode* oddEvenList(ListNode* head) {
        ListNode *o=NULL;
        ListNode *e=NULL;
        ListNode *oh=NULL;
        ListNode *eh=NULL;
        ListNode *tmp=NULL;
        i=0;
        while(head){
            i++;
            tmp=head->next;
            head->next=NULL;
            if(i%2==0){
                // even
                if(e==NULL){
                    e=head;
                    eh=e;
                }else{
                    e->next=head;
                    e=e->next;
                }
            }else{
                // odd
                if(o==NULL){
                    o=head;
                    oh=head;
                }else{
                    o->next=head;
                    o=o->next;
                }
            }
            head=tmp;
        }
        if(o){
            o->next=eh;
            return oh;
        }
        return eh;
    }
};


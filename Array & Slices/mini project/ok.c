//Array Reverse (Loop + Array) taking 5 user input then reverse
#include<stdio.h>
int main(){
    int a[5];
    printf("enter 5 number:\n");
    for (int i = 0; i < 5; i++)
    {
    scanf("%d",&a[i]);
    printf("write :%d\n",a[i]);
    }
    for (int i = 0; i < 5; i++)
    {
        printf("you entered %d\n",a[i]);
    }

    
printf("now reverse time\n");
for (int  i = 4; i>=0; i--)
{
    printf("%d\n",a[i]);
}
return 0;

}
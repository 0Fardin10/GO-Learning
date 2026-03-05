#include <stdio.h>
int findMax(int a[5]){
    int max = a[0];
    for(int i=1;i<5;i++){
        if(a[i]>max){
            max=a[i];
        }
        return max;
    }
    
}
int findsum(int a[5])
{
    int sum = 0;
    for (int i = 0; i<5; i++){
        sum += a [i];
    }
    return sum;
}
int main()
{
    int num[5];
    printf ("Enter 5 number : ");
    for (int i = 0; i<5; i++)
    {
        scanf ("%d", &num[i]);
    }
    for (int i = 0; i<5; i++)
        {
        printf ("Your 5 number is : %d\n", num[i]);
    }
    int total = findsum(num);
    printf ("Total number: %d\n",total);
    int maximum = findMax(num);
    printf("Maximum number is %d\n",maximum);
    return 0;
}

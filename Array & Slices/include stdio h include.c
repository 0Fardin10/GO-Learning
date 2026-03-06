#include <stdio.h>
#include <string.h>
#include <ctype.h>
void processText(char text [],char result[]){
    int len = strlen(text);
    int j =0;
    for(int i =len - 1;i>=0;i--){
        result[j]=toupper(text[i]);
        j++;
        
    }
result[j]='\0';
    }

int main()
{
    char text[100];
    char result[100];
    printf("Enter text: ");
    scanf("%s",&text);
processText(text,result);
 printf("result is : %s\n",result);
    return 0;
}

#include <stdio.h>
#include <string.h>

int main(int argc, char **argv)
{
    char src[40];
   char dest[40];
  
   memset(dest, '\0', sizeof(dest));
   strcpy(src, "This is tutorialspoint.com");
   int len = strlen(src);
   strncpy(dest, src + 5, len - 5);

   printf("Final copied string : %s\n", dest);
    return(0);
}

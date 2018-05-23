#include <stdio.h>
#define MAX_LINE 4096

/*
 * Leer de a una línea por vez (hasta MAX_LINE - 1 caracteres)
 * y escribir en buffer
 */

static void cat(FILE *fp)
{
    char buffer[MAX_LINE];
    while (fgets(buffer, sizeof(buffer), fp) != 0)
         fputs(buffer, stdout);
}

int main(int argc, char **argv)
{
    FILE *fp;
    const char *file;
    while ((file = *++argv) != 0)
    {
        if ((fp = fopen(file, "r")) != 0)
        {
            cat(fp);
            fclose(fp);
        }
    }
    return(0);
}

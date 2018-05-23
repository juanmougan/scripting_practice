#include <stdio.h>
#define MAX_LINE 4096
#define LATE 1
#define NOLA 0

/*
 * Leer de a una línea por vez (hasta MAX_LINE - 1 caracteres)
 * y escribir en buffer
 */

static void cat(FILE *fp, FILE *output_stream)
{
    char buffer[MAX_LINE];
    while (fgets(buffer, sizeof(buffer), fp) != 0)
         fputs(buffer, output_stream);
}

int main(int argc, char **argv)
{
    FILE *fp;
    FILE *output_stream;
    const char *file;
    if (argc >= 3)  // dos nombres de archivos + program name me da 3
    {
        printf("%s %s\n", argv[1], argv[2]);
        char* input_name = "cromos_input.txt";
        char* output_name = "cromos_output.txt";
        fp = fopen(input_name, "r");
        output_stream = fopen(output_name, "w");
        printf("%d %d\n", fp, output_stream);
        if (fp != 0 && output_stream != 0)
        {
            printf("Pude abrir los dos\n");
            cat(fp, output_stream);
            fclose(fp);
            fclose(output_stream);
        }
    }
    return(0);
}

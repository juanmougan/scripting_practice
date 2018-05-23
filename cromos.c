#include <stdio.h>
#define MAX_LINE 4096
#define LATE 1
#define NOLA 0
#define MAX_NOMBRE 35
#define MAX_PAIS 20
#define LIMIT_NUMERO 3
#define LIMIT_NOMBRE 39

typedef struct Cromo {
    int numero;
    char nombre[MAX_NOMBRE];
    char pais[MAX_PAIS];
} Cromo;

/*
 * Leer de a una línea por vez (hasta MAX_LINE - 1 caracteres)
 * y armar el struct
 */
static void cat(FILE *fp, FILE *output_stream);

int main(int argc, char **argv)
{
    FILE *fp;
    FILE *output_stream;
    const char *file;
    printf("%s %s\n", argv[1], argv[2]);
    //char* input_name = "cromos_input.txt";
    char* input_name = "cromos.txt";
    char* output_name = "cromos_output.txt";
    fp = fopen(input_name, "r");
    output_stream = fopen(output_name, "w");
    printf("%d %d\n", fp, output_stream);
    if (fp != 0 && output_stream != 0)
    {
        cat(fp, output_stream);
        fclose(fp);
        fclose(output_stream);
    }
    return(0);
}

static void cat(FILE *fp, FILE *output_stream) {
    //struct Cromo cada_cromo;
    char buffer[MAX_LINE];
    char numero[4];
    char nombre[100];
    char pais[20];
    int len = 0;

    memset(numero, '\0', sizeof(numero));
    memset(nombre, '\0', sizeof(nombre));
    memset(pais, '\0', sizeof(pais));
    
    while (fgets(buffer, sizeof(buffer), fp) != 0) {
         //fputs(buffer, output_stream);

        // Armo mi struct, una por línea
        len = strlen(buffer);
        printf("len: %d\n", len);

        // Para armar un substring
        // strncpy(dest, src + beginIndex, endIndex - beginIndex);

        // Para el número
        strncpy(numero, buffer, LIMIT_NUMERO);
        printf("Número: %s\n", numero);

        // Para el nombre
        strncpy(nombre, buffer + 4, 36 - 4);
        printf("Nombre: %s\n", nombre);

        // Para el pais
        strncpy(pais, buffer + 35 + 1, len);
        printf("País: %s\n", pais);
    }
}

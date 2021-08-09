#include "../libft.h"

char *ft_strcpy(char *dest, const char *src)
{
    char *temp_src = (char *) src;
    while ((*dest++ = *temp_src++) != '\0')
        ;
    return dest;
}

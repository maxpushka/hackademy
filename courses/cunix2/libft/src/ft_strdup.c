#include "../libft.h"

char *ft_strdup(const char *s)
{
    char *new_s = malloc(ft_strlen(s) + 1);
    ft_strcpy(new_s, s);
    return new_s;
}

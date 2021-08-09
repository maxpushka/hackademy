#include "../libft.h"

size_t ft_strlen(const char *s)
{
    size_t res = 0;
    for (size_t i = 0; s[i] != '\0'; i++, res++)
        ;
    return res;
}

#include "../libft.h"

void *ft_memccpy(void *dest, const void *src, int c, size_t n)
{
    while (n-- > 0)
    {
        if ((*(unsigned char *)(dest++) = *(unsigned char *)(src++)) == c)
        {
            return dest;
        }
    }
    return NULL;
}

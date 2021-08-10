#include "../libft.h"

void *ft_memcpy(void *dest, const void *src, size_t n)
{
    char *ch_dest = (char *) dest;
    const char *ch_src = (const char *) src;

    for (size_t i = 0; i < n; i++)
    {
        ch_dest[i] = ch_src[i];
    }

    return dest;
}

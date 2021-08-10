#include "../libft.h"

void *ft_memccpy(void *dest, const void *src, int c, size_t n)
{
    char *ch_dest = (char *) dest;
    const char *ch_src = (const char *) src;

    for (size_t i = 0; i < n; i++)
    {
        if (ch_src[i] == (char) c)
        {
            break;
        }

        ch_dest[i] = ch_src[i];
    }

    return dest;
}

#include "../libft.h"

void *ft_memmove(void *dest, const void *src, size_t n)
{
    char *ch_dest = (char *) dest;
    const char *ch_src = (const char *) src;

    if (ch_src < ch_dest)
    {
        for (int i = n; i >= 0; i--)
        {
            ch_dest[i] = ch_src[i];
        }
    }
    else if (ch_src > ch_dest)
    {
        for (size_t i = 0; i < n; i++)
        {
            ch_dest[i] = ch_src[i];
        }
    }

    return dest;  // + case ch_src == ch_dest
}

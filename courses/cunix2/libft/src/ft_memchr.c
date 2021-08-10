#include "../libft.h"

void *ft_memchr(const void *s, unsigned char c, size_t n)
{
    const unsigned char *ss = (const unsigned char *) s;

    for (size_t i = 0; i < n; i++)
    {
        if (ss[i] == c)
        {
            return (void *)(s + i);
        }
    }

    return NULL;
}

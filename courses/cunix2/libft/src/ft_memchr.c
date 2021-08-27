#include "../libft.h"

void *ft_memchr(const void *s, char c, size_t n)
{
    const char *ss = (const char *) s;

    for (size_t i = 0; i < n; i++)
    {
        if (ss[i] == c)
        {
            return (void *)(s + i);
        }
    }

    return NULL;
}

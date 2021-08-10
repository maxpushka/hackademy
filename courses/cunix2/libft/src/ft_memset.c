#include "../libft.h"

void *ft_memset(void *s, int c, size_t n)
{
    char *ss = s;
    for (size_t i = 0; i < n; i++)
    {
        ss[i] = (char) c;
    }

    return s;
}

#include "../libft.h"

void ft_bzero(void *s, size_t n)
{
    if (s)
    {
        char * ss = (char *) s;

        for (size_t i = 0; i < n; i++)
        {
            ss[i] = '\0';
        }
    }
}

#include "../libft.h"

int ft_memcmp(const void *s1, const void *s2, size_t n)
{
    const unsigned char *ss1 = (const unsigned char *) s1;
    const unsigned char *ss2 = (const unsigned char *) s2;

    for (size_t i = 0; i < n; i++)
    {
        if (ss1[i] > ss2[i])
        {
            return 1;
        }
        else if (ss1[i] < ss2[i])
        {
            return -1;
        }
    }

    return 0;
}

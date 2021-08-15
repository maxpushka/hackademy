#include "../libft.h"

void *ft_memmove(void *dest, const void *src, size_t n)
{
    char *arr = (char *) malloc(n);

    size_t i = 0;
    for (; i < n; i++)
    {
        arr[i] = *(char *)(src + i);
    }

    for (i = 0; i < n; i++)
    {
        *(char *)(dest + i) = arr[i];
    }

    free(arr);
    return dest;
}


#include "../libft.h"

int ft_toupper(int c)
{
    if ('a' <= c && c <= 'z')
    {
        return c - ('a' - 'A');
    }

    return c;
}

#include "../libft.h"

int ft_isalpha(int c)
{
    int is_upper = ('A' <= c && c <= 'Z');
    int is_lower = ('a' <= c && c <= 'z');

    if (is_upper || is_lower)
    {
        return 1;
    }
    else
    {
        return 0;
    }
}

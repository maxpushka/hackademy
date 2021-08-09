#include "../libft.h"

int ft_isdigit(int c)
{
    if ('0' <= c && c <= '9')
    {
        return 1;
    }
    else
    {
        return 0;
    }
}

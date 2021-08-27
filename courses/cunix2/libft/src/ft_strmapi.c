#include "../libft.h"

char *ft_strmap(const char *s, char (*f)(char))
{
    const size_t s_size = ft_strlen(s);
    char *new_s = malloc(s_size + 1);
    new_s[s_size] = '\0';

    for (size_t i = 0; i < s_size; i++)
    {
        new_s[i] = (*f)(s[i]);
    }

    return new_s;
}

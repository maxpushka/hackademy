#include "../libft.h"

char *ft_strchr(const char *s, int c)
{
    const size_t s_size = ft_strlen(s);

    for (size_t i = 0; i <= s_size; i++)
    {
        if (s[i] == (char) c)
        {
            return (char *)(&s[i]);
        }
    }

    return NULL;
}

#include "../libft.h"

char *ft_strrchr(const char *s, int c)
{
    for (int i = ft_strlen(s); i >= 0; i--)
    {
        if (s[i] == (char) c)
        {
            return (char *)(&s[i]);
        }
    }

    return NULL;
}

#include "../libft.h"

char *ft_strsub(const char *s, unsigned int start, size_t len)
{
    const size_t s_len = ft_strlen(s);

    if (start > s_len)
    {
        char *fail_s = (char *) malloc(1);
        fail_s[0] = '\0';
        return fail_s;
    }

    len = s_len - start > len ? len : s_len - start;

    char *sub_s = (char *) malloc(len + 1);
    sub_s[len] = '\0';

    for (size_t i = 0; i < len; i++)
    {
        sub_s[i] = s[i + start];
    }

    return sub_s;
}


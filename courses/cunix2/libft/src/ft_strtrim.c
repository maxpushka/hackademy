#include "../libft.h"

char *ft_strtrim(const char *s)
{
    const char* s_start = s;
    const char* s_end = s + ft_strlen(s) - 1;

    while (*s_start == ' ' || *s_start == '\n' || *s_start == '\t')
    {
        s_start++;
    }

    if (s_start == s_end + 1)
    {
        char *empty_s = (char *) malloc(1);
        empty_s[0] = '\0';
        return empty_s;
    }

    while (*s_end == ' ' || *s_end == '\n' || *s_end == '\t')
    {
        s_end--;
    }

    const size_t s_trimmed_size = ++s_end - s_start;
    char *s_trimmed = (char *) malloc(s_trimmed_size + 1);
    s_trimmed[s_trimmed_size] = '\0';

    for (size_t i = 0; s_start != s_end && i < s_trimmed_size; i++)
    {
        s_trimmed[i] = *s_start++;
    }

    return s_trimmed;
}

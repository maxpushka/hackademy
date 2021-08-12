#include "../libft.h"

char *ft_strjoin(const char *s1, const char *s2)
{
    const size_t s1_size = ft_strlen(s1);
    const size_t s2_size = ft_strlen(s2);

    char *new_s = (char *) malloc(s1_size + s2_size + 1);
    ft_strcpy(new_s, s1);
    ft_strcpy(new_s + s1_size, s2);
    new_s[s1_size + s2_size] = '\0';

    return new_s;
}

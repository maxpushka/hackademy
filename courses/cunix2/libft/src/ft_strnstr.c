#include "../libft.h"

char *ft_strnstr(const char *haystack, const char *needle, size_t search_len)
{
    char needle_ch, hay_ch;

    if ((needle_ch = *needle++) != '\0')
    {
        size_t needle_len = ft_strlen(needle);

        do 
        {
            do
            {
                if (search_len-- < 1 || (hay_ch = *haystack++) == '\0')
                {
                    return NULL;
                }
            } 
            while (hay_ch != needle_ch);

            if (needle_len > search_len)
            {
                return NULL;
            }
        }
        while (ft_strncmp(haystack, needle, needle_len) != 0);
        haystack--;
    }

    return (char *) haystack;
}

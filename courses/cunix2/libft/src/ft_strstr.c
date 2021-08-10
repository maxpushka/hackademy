#include "../libft.h"

char *ft_strstr(const char *haystack, const char *needle)
{
    char needle_ch, hay_ch;

    if ((needle_ch = *needle++) != 0) 
    {
        size_t needle_len = ft_strlen(needle);

        do 
        {
            do 
            {
                if ((hay_ch = *haystack++) == 0)
                {
                    return NULL;
                }
            }
            while (hay_ch != needle_ch);
        }
        while (ft_strncmp(haystack, needle, needle_len) != 0);
        haystack--;
    }

    return (char *) haystack; 
}

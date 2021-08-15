#include "../libft.h"

char *find_substr_start(char const *s, char c)
{
    size_t i = 0;
    while (s[i] == c || s[i] == '\0')
    {
        if (s[i] == '\0')
        {
            return 0;
        }
        i++;
    }
    return (char *)(s + i);
}

size_t find_substr_end_index(char const *s, char c)
{
    size_t i = 0;
    while (s[i] != c && s[i] != '\0')
    {
        i++;
    }
    return i;
}

size_t substr_length(char const *s, char c)
{
    size_t i = 0;
    while (s[i] != c && s[i] != '\0')
    {
        i++;
    }

    return i;
}


size_t next_word_offset(char const *s, char const *subs, char c)
{
    return (subs - s) + find_substr_end_index(subs, c);
}

char **ft_strsplit(char const *s, char c)
{
    char *subs;
    size_t s_iter = 0;

    // count words
    size_t words = 0;
    while ((subs = find_substr_start((s + s_iter), c)))
    {
        words++;
        s_iter += next_word_offset(s + s_iter, subs, c); 
    }

    // create array of words
    char **ss;
    if (!(ss = (char **) malloc(sizeof(char *) * (words + 1))))
    {
        return NULL;
    }

    // iterate through words
    s_iter = 0;
    size_t word_length;
    size_t i = 0;
    size_t j;

    while ((subs = find_substr_start((s + s_iter), c)))
    {
        word_length = substr_length(subs, c);
        if (!(ss[i] = (char *)malloc(sizeof(char) * (word_length + 1))))
        {
            return NULL;
        }


        j = 0;
        while (j < word_length)
        {
            ss[i][j] = subs[j];
            j++;
        }
        ss[i][j] = '\0';

        i++;
        s_iter += next_word_offset(s + s_iter, subs, c);
    }
    
    ss[i] = NULL;

    return ss;
}

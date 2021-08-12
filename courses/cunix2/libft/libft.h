#ifndef LIBFT_H_
#define LIBFT_H_

#include <stddef.h>
#include <malloc.h>

void ft_bzero(void *s, size_t n);

size_t ft_strlen(const char *s);
char *ft_strcpy(char *dest, const char *src);
char *ft_strdup(const char *s);

int ft_strncmp(const char *s1, const char *s2, size_t n);

char *ft_strchr(const char *s, int c);
char *ft_strrchr(const char *s, int c);

int ft_isalpha(int c);
int ft_isdigit(int c);
int ft_isascii(int c);

int ft_toupper(int c);
int ft_tolower(int c);

int ft_abs(int j);

typedef struct div
{
    int quot;
    int rem;
}
_div_t;

_div_t ft_div(int numerator, int denominator);

char *ft_strstr(const char *haystack, const char *needle);
char *ft_strnstr(const char *haystack, const char *needle, size_t search_len);

void *ft_memset(void *s, int c, size_t n);
void *ft_memcpy(void *dest, const void *src, size_t n);
void *ft_memccpy(void *dest, const void *src, int c, size_t n);
void *ft_memmove(void *dest, const void *src, size_t n);
void *ft_memchr(const void *s, unsigned char c, size_t n);
int ft_memcmp(const void *s1, const void *s2, size_t n);

void ft_striteri(char *s, void (*f)(unsigned int, char *));
char *ft_strmapi(const char *s, char (*f)(unsigned int, char));
char *ft_strsub(const char * s, unsigned int start, size_t len);

#endif

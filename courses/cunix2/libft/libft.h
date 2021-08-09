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

#endif

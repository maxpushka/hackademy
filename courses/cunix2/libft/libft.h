#ifndef LIBFT_H_
#define LIBFT_H_

#include <stddef.h>
#include <malloc.h>

void ft_bzero(void *s, size_t n);

size_t ft_strlen(const char *s);
char *ft_strcpy(char *dest, const char *src);
char *ft_strdup(const char *s);

#endif

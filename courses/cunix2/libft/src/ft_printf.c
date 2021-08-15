#include "../libft.h"
#include <stddef.h>
#include <stdlib.h>
#include <stdarg.h>
#include <unistd.h>

#include <stdio.h>

int ft_to_string(int val, char *buf);
void print_filler(char *filler, int amount);
void print(char *string);
void reset_format(int *is_formated, int *f_min_width, int *left_side, int *is_sign);


int ft_printf(char *format, ...)
{
    va_list arg_list;
    int number_of_chars = 0;

    int int_value;
    char *string_value;
    char char_value;
    char buffer[13];

    int is_formated = 0;
    int f_min_width = 0;
    int is_sign = 0;   // 0 - default, 1 - sign is obl., 2 - space is no sign
    int left_side = 0; // 0 -default, 1 - allign. left, 2 - fill with zeros
    int size_of_string;

    va_start(arg_list, format);

    for (int i = 0; format[i] != 0; i++)
    {
        if (format[i] != '%' && is_formated == 0)
        {
            write(1, &format[i], sizeof(char));
            number_of_chars++;
            continue;
        }
        i = is_formated ? i : i + 1;

        switch (format[i])
        {
            case 'd':
            case 'i':
                int_value = va_arg(arg_list, int);
                size_of_string = ft_to_string(int_value, buffer);

                char *sign;
                sign = int_value < 0 ? "-" : (is_sign == 1 ? "+" : (is_sign == 2 ? " " : ""));
                size_of_string = sign[0] == 0 ? size_of_string : size_of_string + 1;

                string_value = buffer;
                if (f_min_width <= size_of_string)
                {
                    write(1, sign, sizeof(char));
                    print(string_value);
                }
                else if (left_side == 0)
                {
                    print_filler(" ", f_min_width - size_of_string);
                    write(1, sign, sizeof(char));
                    print(string_value);
                }
                else if (left_side == 1)
                {
                    write(1, sign, sizeof(char));
                    print(string_value);
                    print_filler(" ", f_min_width - size_of_string);
                }
                else if (left_side == 2)
                {
                    write(1, sign, sizeof(char));
                    print_filler("0", f_min_width - size_of_string);
                    print(string_value);
                }
                reset_format(&is_formated, &f_min_width, &left_side, &is_sign);
                break;
            case 'c':
                char_value = va_arg(arg_list, int);
                string_value = &char_value;
                write(1, string_value, sizeof(char));
                reset_format(&is_formated, &f_min_width, &left_side, &is_sign);
                break;
            case 's':
                string_value = va_arg(arg_list, char *);
                size_of_string = 0;
                while (string_value[size_of_string++])
                    ;
                size_of_string--;

                if (f_min_width <= size_of_string)
                {
                    print(string_value);
                }
                else if (left_side == 0)
                {
                    print_filler(" ", f_min_width - size_of_string);
                    print(string_value);
                }
                else if (left_side == 1)
                {
                    print(string_value);
                    print_filler(" ", f_min_width - size_of_string);
                }
                else if (left_side == 2)
                {
                    print_filler("0", f_min_width - size_of_string);
                    print(string_value);
                }
                reset_format(&is_formated, &f_min_width, &left_side, &is_sign);
                break;
            case '%':
                write(1, "%", sizeof(char));
                reset_format(&is_formated, &f_min_width, &left_side, &is_sign);
                break;
            default:
                is_formated = 1;
                f_min_width = 0;

                for (; format[i] != 's' && format[i] != 'c' && format[i] != 'i' && format[i] != 'd'; i++)
                {
                    if (format[i] >= '0' && format[i] <= '9')
                    {
                        f_min_width *= 10;
                        f_min_width += (format[i] - '0');
                    }

                    if (f_min_width > 0)
                    {
                        continue;
                    }

                    is_sign = is_sign != 0 ? is_sign : format[i] == '+' ? 1
                              : format[i] == ' '   ? 2
                              : 0;

                    left_side = left_side != 0 ? left_side : format[i] == '-' ? 1
                                : format[i] == '0'   ? 2
                                : 0;
                }
                i--;
        }
    }

    va_end(arg_list);

    return number_of_chars;
}

void print_filler(char *filler, int amount)
{
    for (int i = 0; i < amount; i++)
    {
        write(1, filler, sizeof(char));
    }
}

void print(char *string)
{
    for (; *string; string++)
    {
        write(1, string, sizeof(char));
    }
}

int ft_to_string(int val, char *buf)
{
    int string_size = 0;
    long digit = 1;
    int nmb2 = val;

    if (val == 0)
    {
        buf[0] = '0';
        buf[1] = '\0';
        return 1;
    }

    for (int i = 0; nmb2 != 0; i++)
    {
        nmb2 /= 10;
        digit *= 10;
        string_size++;
    }

    digit /= 10;

    int i = 0;
    if (val < 0)
    {
        val = -val;
    }
    for (; i < string_size; i++)
    {
        buf[i] = (val / digit) + '0';
        val %= digit;
        digit /= 10;
    }

    buf[i] = 0;
    return string_size;
}

void reset_format(int *is_formated, int *f_min_width, int *left_side, int *is_sign)
{
    *is_formated = 0;
    *f_min_width = 0;
    *is_sign = 0;
    *left_side = 0;
}


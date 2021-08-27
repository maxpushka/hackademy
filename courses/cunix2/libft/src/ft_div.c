#include "../libft.h"

div_t ft_div(int numerator, int denominator)
{
    div_t result;

    result.quot = numerator / denominator;
    result.rem = numerator % denominator;

    if (numerator >= 0 && result.rem < 0)
    {
        result.quot++;
        result.rem -= denominator;
    }

    return result;
}


CC := gcc
RM := rm -rf
CFLAGS += -Iinclude -Wall -Wextra -Werror -g -std=c99

SRCS := $(wildcard src/ft_*.c)
OBJS := $(SRCS:.c=.o)
NAME := libft.a

all: $(NAME) clean

$(NAME): $(OBJS)
	ar rcs $(NAME) $(OBJS)

clean:
	$(RM) $(OBJS)

fclean: clean
	$(RM) $(NAME)


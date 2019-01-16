NAME=expert_system
SRC_PATH=.
GO=go
BUILD=build
DEBUG=debug

SRC_NAME=main.go \
flags.go \
parser.go 

all: $(NAME)
SRC = $(addprefix $(SRC_PATH)/, $(SRC_NAME))

$(NAME): $(SRC)
	go build -o $(NAME) $(SRC)

fclean:
	rm -rf $(NAME)

re: fclean all

debug:
	go build -gcflags "-m -m -l" -o $(DEBUG) $(SRC) 

fclean_debug:
	@rm -rfv $(DEBUG)

re_debug: fclean_debug debug

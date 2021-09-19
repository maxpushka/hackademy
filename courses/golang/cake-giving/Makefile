launch_server:
	go run -race ./

register:
	@echo "Registration\n"
	@curl -X POST localhost:8000/user/register --data '{"email":"hackademy","password":"fdj232dlf","favorite_cake":"cheesecake"}'
	@echo "\n----\n"

right_login:
	@echo "Right login\n"
	@curl -X POST localhost:8000/user/jwt --data '{"email":"hackademy","password":"fdj232dlf"}' 
	@echo "\n----\n"

wrong_login:
	@echo "Wrong login\n"
	@curl -X POST localhost:8000/user/jwt --data '{"email":"hackademy","password":"aaaaaaa"}'
	@echo "\n----\n"

get_cake_error:
	@echo "Get cake unauthorized\n"
	@curl -X GET localhost:8000/cake
	@echo "\n----\n"

get_cake_ok:
	@echo "Get cake authorized (doesn`t work)\n"	
	@export JWT=$(curl -X POST localhost:8000/user/jwt --data '{"email":"hackademy","password":"fdj232dlf"}')
	@echo ${JWT}
	@curl -X GET localhost:8000/cake -H "Authorization: Bearer ${JWT}"
	@echo "\n----\n"

test: register wrong_login right_login get_cake_error get_cake_ok
	@echo "Tests complete"
	
echo "This is a quick command to facilitate application development."
echo "If you found any bugs please report them on this Github. Thanks for reading <3."
echo "Github: https://github.com/RabbitHouseCorp/tokamak"
echo "Issues: https://github.com/RabbitHouseCorp/tokamak/issues"
echo "RabbitHouseCorp: https://github.com/RabbitHouseCorp"
echo 
echo 
cd ./src && echo "[1/3] Running GO GET" && go get  && echo "[2/3] Install..." && go install && echo "[3/3] Tokamak ready!"   && go run main.go

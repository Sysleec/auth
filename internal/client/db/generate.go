package db

//go:generate sh -c "rm -rf mocks && mkdir -p mocks"
//go:generate minimock -i Client -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i TxManager -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i Transactor -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i SQLExecer -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i NamedExecer  -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i QueryExecer -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i Pinger -o ./mocks/ -s "_minimock.go"
//go:generate minimock -i DB -o ./mocks/ -s "_minimock.go"

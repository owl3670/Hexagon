# Hexagon

## Run

1. go 설치 https://go.dev/
2. docker 설치 https://docs.docker.com/engine/install/
3. docker-compose 설치 https://docs.docker.com/compose/install/
4. go migrate CLI 설치 https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
5. docker container 시작
   1. 터미널 실행 
   2. 프로젝트 루트 디렉토리로 이동
   3. docker-compose up -d 실행
6. Go module 다운로드
   1. 터미널 실행
   2. 프로젝트 루트 디렉토리로 이동
   3. go mod download 실행
7. migrate 실행
   1. 터미널 실행
   2. 프로젝트 루트 디렉토리로 이동
   3. migrate -path ./migrations -database postgres://postgres:testpasswd@localhost:5432/postgres?sslmode=disable up 실행
8. 서버 실행
   1. 터미널 실행
   2. 프로젝트 루트 디렉토리로 이동
   3. go run ./cmd/api 실행

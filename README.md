# Hexagon

## Skill

#### Language
Go
#### Database
PostgreSQL  
Redis
#### Environment
Docker

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

## Architecture

![img](./readme/Hexagonal_Architecture.png)

헥사고널 아케텍쳐를 구현 하고자 하며,  
각각의 domain 별로 core 로써 domains, dto, ports 를 명세하고,  
명세된 ports 에 대응 되는 adapter 를 구현하여 구현 관심사의 분리를 하고자 한다.

### core
#### domains 
domains 에는 외부에 의존성이 없는 domain model 의 변경등과 같은 비지니스 로직이 구현된다.

#### ports
domain 으로의 접근 혹은 domain 으로 부터 infra 로의 접근을 위한 기능을 인터페이스로 명세한다.

#### dto
port 에 전달되는 data의 구조를 명세한다.

### handlers
server 의 router 로 부터 호출되는 handler 를 구현한다.

### infra
DB 등의 외부 infra 에 접근하기 위한 기능을 구현하는 것으로 ports 의 repository 를 구현한다.

### service
외부 에서 domain 에 접근하기 위한 기능을 구현하는 것으로 ports 의 service 를 구현한다.

## API Specification

![img](./readme/API.jpeg)

https://documenter.getpostman.com/view/20455344/UVyysYA2

## Feature Specification

### 회원 가입

1. SMS 인증 코드 및 토큰 발급 (sign up) (TEST 를 위해 response 에 code 포함)
2. SMS 인증
3. 유저 회원 가입 (email, nickname, phone_number 중복 허용하지 않음)

### 유저 로그인

1. 유저 로그인   
email, nickname, phone_number 중 택 1 하여 로그인 가능
2. 유저 로그인 response 에 포함된 jwt_token 의 access_token 으로 세션을 대체

### 내 정보 조회

1. 유저 로그인
2. 유저 로그인 JWT 저장
3. 내 정보 조회

### 비밀번호 재설정

1. SMS 인증 코드 및 토큰 발급 (reset password) (TEST 를 위해 response 에 code 포함)
2. SMS 인증
3. 유저 비밀 번호 변경 

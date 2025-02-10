Go의 구조체(struct)

1. 기본 구조체 정의

```go
type Person struct {
    Name    string
    Age     int
    Email   string
}

// 중첩 구조체
type Address struct {
    Street  string
    City    string
    Country string
}

type Employee struct {
    Person      // 구조체 임베딩
    Address     // 구조체 임베딩
    Company     string
    Department  string
}
``` 


2. 구조체 초기화 방법들

```go 
    // 방법 1: 필드 순서대로
    person1 := Person{"John", 30, "john@email.com"}
    
    // 방법 2: 필드명 지정 (권장)
    person2 := Person{
        Name: "Alice",
        Age: 25,
        Email: "alice@email.com",
    }
    
    // 방법 3: new 키워드 사용
    person3 := new(Person)
    person3.Name = "Bob"
    
    // 방법 4: 포인터로 생성
    person4 := &Person{
        Name: "Charlie",
        Age: 35,
    }
```

3. 메서드 정의

```go
// 값 리시버
func (p Person) GetInfo() string {
    return fmt.Sprintf("%s is %d years old", p.Name, p.Age)
}

// 포인터 리시버 (상태 변경 가능)
func (p *Person) UpdateAge(newAge int) {
    p.Age = newAge
}

// 사용 예시
func main() {
    person := Person{Name: "John", Age: 30}
    fmt.Println(person.GetInfo())
    person.UpdateAge(31)
}
``` 

4. 태그 사용

```go
type User struct {
    ID        int       `json:"id" db:"user_id"`
    Username  string    `json:"username" validate:"required"`
    Password  string    `json:"-" validate:"required"`
    CreatedAt time.Time `json:"created_at"`
}
```

5. 프라이빗/퍼블릭 필드

```go 
type Account struct {
    Username    string  // 퍼블릭 (대문자)
    password    string  // 프라이빗 (소문자)
}
```

6. 구조체 컴포지션 (상속과 유사)

```go
type Animal struct {
    Name string
    Age  int
}

func (a Animal) Sleep() {
    fmt.Printf("%s is sleeping\n", a.Name)
}

type Dog struct {
    Animal      // 구조체 임베딩
    Breed string
}

// 사용 예시
dog := Dog{
    Animal: Animal{Name: "Rex", Age: 3},
    Breed: "German Shepherd",
}
dog.Sleep() // Animal의 메서드 사용 가능
```

7. 인터페이스 구현

```go
type Speaker interface {
    Speak() string
}

type Human struct {
    Name string
}

// Human이 Speaker 인터페이스 구현
func (h Human) Speak() string {
    return fmt.Sprintf("Hello, I'm %s", h.Name)
}
```

8. 제로 값과 생성자 패턴

```go
type Config struct {
    Host    string
    Port    int
    Timeout time.Duration
}

// 생성자 함수
func NewConfig() *Config {
    return &Config{
        Host:    "localhost",
        Port:    8080,
        Timeout: 30 * time.Second,
    }
}
```

9. 구조체 메모리 최적화
```go
// 좋은 예 (메모리 효율적)
type Optimized struct {
    flag    bool    // 1 byte
    pad     byte    // 1 byte
    count   uint16  // 2 bytes
    id      uint32  // 4 bytes
}

// 나쁜 예 (패딩으로 인한 메모리 낭비)
type Unoptimized struct {
    id      uint32  // 4 bytes
    flag    bool    // 1 byte + 3 bytes padding
    count   uint16  // 2 bytes + 2 bytes padding
}
```

10. 구조체 비교

```go 
type Compare struct {
    Name string
    Age  int
}

func main() {
a := Compare{"John", 30}
b := Compare{"John", 30}

    // 직접 비교 가능
    fmt.Println(a == b) // true
    
    // 깊은 비교
    fmt.Println(reflect.DeepEqual(a, b)) // true
}
``` 

11. 실제 활용 예시:

```go
type Server struct {
    config  *Config
    db      *sql.DB
    router  *mux.Router
    logger  *log.Logger
}

func NewServer(cfg *Config) (*Server, error) {
    db, err := sql.Open("postgres", cfg.DBUrl)
    if err != nil {
        return nil, err
    }

    return &Server{
		
        config:  cfg,
        db:      db,
        router:  mux.NewRouter(),
        logger:  log.New(os.Stdout, "", log.LstdFlags),
    }, nil
}

func (s *Server) Start() error {
    return http.ListenAndServe(s.config.Port, s.router)
}
```
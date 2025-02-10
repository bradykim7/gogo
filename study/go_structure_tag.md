Go struct tags and their common uses. Tags are string literals attached to struct fields that provide metadata.

1. Common Tag Formats and Uses:
```go
type User struct {
    // JSON tags
    Name  string `json:"name,omitempty"`
    Email string `json:"email" validate:"required,email"`

    // Database tags
    ID        int64  `db:"id" gorm:"primaryKey"`
    CreatedAt time.Time `db:"created_at" gorm:"autoCreateTime"`

    // Form/validation tags
    Password string `form:"password" validate:"required,min=8"`

    // XML tags
    Address string `xml:"address,attr"`

    // YAML tags
    Config string `yaml:"configs"`
}
```

2. Common Tag Types and Their Purposes:

   1. json tags:
    ```go 
    type Example struct {
        // "name" - specifies JSON field name
        // "omitempty" - skips field if empty
        Name string `json:"name,omitempty"`
    
        // "-" - ignores field in JSON
        Internal string `json:"-"`
    }
    ```
    2. Database tags:
    ```go 
    type Product struct {
        // gorm tags for ORM
        ID      int64  `gorm:"primaryKey"`
        Price   float64 `gorm:"type:decimal(10,2)"`
        
        // sql tags
        Status string `db:"status" sql:"type:ENUM('active','inactive')"`
    }
    ```
   3. Validation tags:
    ```go
    type Registration struct {
        // validate tags for input validation
        Email    string `validate:"required,email"`
        Password string `validate:"required,min=8,max=32"`
        Age      int    `validate:"gte=18,lte=120"`
    }
   ```
    4. Form tags:
    ```go
    type LoginForm struct {
        // form tags for HTTP form binding
        Username string `form:"username" binding:"required"`
        Password string `form:"password" binding:"required"`
    }
    ```

3. Using Tags in Code:
```go 
import "reflect"

type Person struct {
    Name string `json:"name" validate:"required"`
}

func main() {
    t := reflect.TypeOf(Person{})
    field := t.Field(0)

    // Get specific tag
    jsonTag := field.Tag.Get("json")     // "name"
    validateTag := field.Tag.Get("validate") // "required"
}
```

4. Common Tag Conventions:
   - Key-value format: key:"value"
   - Multiple options: json:"name,omitempty"
   - Multiple tags: json:"name" validate:"required"
   - Case sensitivity: Tags are case-sensitive
   - No spaces around colons or commas

5. 주요 사용 사례:

   1. JSON 직렬화/역직렬화:
   ```go
   type Product struct {    
        ID          int     `json:"id"`
        Name        string  `json:"name"`
        Price       float64 `json:"price,omitempty"`   // 값이 비어있으면 생략
        SKU         string  `json:"-"`                 // JSON 변환 시 제외
        InternalID  string  `json:"-,"`                // 하이픈을 필드명으로 사용
   }

   // 사용 예시
   product := Product{
       ID: 1,
       Name: "Phone",
       Price: 999.99,
   }
   
   // JSON으로 변환
   jsonData, _ := json.Marshal(product)
   ```
   2. 데이터베이스 매핑:
   ```go 
   type User struct {
        ID        int       `db:"user_id"`      // 데이터베이스 컬럼명 매핑
        Username  string    `db:"username"`
        CreatedAt time.Time `db:"created_at"`
   }
   ```
   3. 폼 데이터 바인딩:
   ```go
      type LoginForm struct {
         Username string `form:"username" binding:"required"`
         Password string `form:"password" binding:"required,min=6"`
      }
   ```

   4. 태그 접근과 파싱:
   ```go
   import "reflect"

   type Person struct {
        Name string `validate:"required" description:"사용자 이름"`
   }
   
   func main() {
        t := reflect.TypeOf(Person{})
        field := t.Field(0) // Name 필드 정보 가져오기
   
       // 특정 태그 값 가져오기
       validateTag := field.Tag.Get("validate")    // "required"
       descTag := field.Tag.Get("description")     // "사용자 이름"
   }

   ```

   5. 커스텀 태그 검증:

      ```go 
      type Custom struct {
           Value string `custom:"key=value,option=true"`
      }
   
      func parseCustomTag(tag string) map[string]string {
         result := make(map[string]string)
         pairs := strings.Split(tag, ",")
   
          for _, pair := range pairs {
              kv := strings.Split(pair, "=")
              if len(kv) == 2 {
                  result[kv[0]] = kv[1]
              }
          }
          return result
      }
      ```

6. 실제 활용 예시:

   검증(Validation) 라이브러리와 함께 사용:
   ```go 
   type Order struct {
      ID      int     `validate:"required"`
      UserID  int     `validate:"required"`
      Amount  float64 `validate:"required,gte=0"`
      Status  string  `validate:"required,oneof=pending approved rejected"`
   }
   
   
   func validateOrder(order Order) error {
      validate := validator.New()
      return validate.Struct(order)
   }
   ```

   GORM(ORM)과 함께 사용:
   ```go
   type Product struct {
      ID        uint      `gorm:"primaryKey"`
      Name      string    `gorm:"size:255;not null"`
      Price     float64   `gorm:"type:decimal(10,2)"`
      CreatedAt time.Time `gorm:"autoCreateTime"`
      DeletedAt time.Time `gorm:"index"`
   }
   ```

7. 자주 사용되는 태그 규칙:
```go
type Example struct {
    // JSON 태그
    Field1 string `json:"field_1,omitempty"`

    // 데이터베이스 태그
    Field2 string `db:"field_2" gorm:"column:field_2"`
    
    // 검증 태그
    Field3 string `validate:"required,min=3,max=32"`
    
    // XML 태그
    Field4 string `xml:"field_4,attr"`
    
    // YAML 태그
    Field5 string `yaml:"field_5"`
    
    // 여러 태그 조합
    Field6 string `json:"field_6" validate:"required" db:"field_6"`
}
```
8. 태그를 사용할 때 주의사항:

   태그는 런타임에 파싱되므로 성능에 영향을 줄 수 있음
   태그 문법이 정확해야 함 (쌍따옴표, 띄어쓰기 등)
   필요한 태그만 사용하여 가독성 유지
   태그의 의미를 문서화하여 유지보수성 향상
   
   이런 태그들은 주로 프레임워크나 라이브러리와 함께 사용될 때 매우 유용하며, 코드의 선언적 특성을 높여줍니다.
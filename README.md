## Encoder() of a healthy person

Accepts json and returns url-encoded body

### Usage

```
import "github.com/fruitgum/URLEncoder"

func main(){
  postFields := `{
    "field1": 0,
    "filed2": "test",
    }`

  postData, err := URLEncoder.EncodeBody(postFields) // Will return field1=0&field2=test
  r, err := http.NewRequest("POST", "https://expample.com", strings.NewReader(postData))
}
```

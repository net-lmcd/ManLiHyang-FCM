FCM, ManLiHyang Application
=========================

### firebase-admin sdk
```bash
$go get firebase.google.com/go
```

### gin framework
```bash
$go get -u github.com/gin-gonic/gin
```

### go-mysql driver 
```bash
$go get -u github.com/go-sql-driver/mysql
```

### usage 
```go
import (
    	"net/http"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
```

### run server
```bash
$go build
$./ManLiHyang 
```

# How to stub mongo connection with test container

## Required

- Docker

## Repository

- ตัวอย่าง Repository สำหรับจำลอง Integration test

```golang
package custrepo

import (
  "context"

  "github.com/pkg/errors"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type Customer struct {
  ID   primitive.ObjectID `bson:"_id"`
  Name string             `bson:"name"`
}

type ICustomerRepository interface {
  CreateCustomer(ctx context.Context, data Customer) error
}

type CustomerRepository struct {
  Collection mong.ICollection
}

func NewCustomerRepository(dbconn mong.IConnect) ICustomerRepository {
  return &CustomerRepository{
    Collection: dbconn.Database().Collection("customers"),
  }
}

func (repo CustomerRepository) CreateCustomer(ctx context.Context, data Customer) error {
  if _, err := repo.Collection.InsertOne(ctx, data); err != nil {
    return errors.Wrap(err, "Create customer")
  }

  return nil
}
```

## Setting Vistual studio code

- ตั้งค่า Vistual studio code ให้สามารถอ่านค่า build tags ได้จาก gopls
- แก้ไขใน .vscode/settings.json เพิ่มค่าด้านล่าง

```json
{
  "go.buildFlags": ["-tags=integration"],
  "go.testTags": "integration"
}
```

## Test file

- ด้านบนของ File เมื่อเป็น integration test ให้เพิ่ม build tags เพื่อแยก test case

```golang
//go:build integration
// +build integration
```

- ข้อมูลของ test case สำหรับ function CreateCustomer

```golang
//go:build integration
// +build integration

package custrepo_test

import (
  "context"
  "os"
  "strings"
  "testing"

  "github.com/stretchr/testify/assert"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

var (
  collection mong.ICollection
  repo       custrepo.ICustomerRepository
)

func TestMain(m *testing.M) {
  container, err := mongstub.Connect("demos")
  if err != nil {
    panic(err)
  }

  defer func(ctx context.Context) {
    container.Client.Close()
    container.Terminate(ctx)

  }(context.Background())

  // Keep collection
  collection = container.Client.Database().Collection("customers")

  // New repository
  repo = custrepo.NewCustomerRepository(container.Client)

  os.Exit(m.Run())
}

func TestCreateCustomer(t *testing.T) {
  // Define variables
  var (
    ctx    context.Context
    result custrepo.Customer
  )

  beforeEach := func() {
    ctx = context.Background()
    result = custrepo.Customer{}
  }

  afterEach := func() {
    collection.Drop(ctx)
  }

  t.Run("Should be insert customer to database", func(t *testing.T) {
    beforeEach()
    defer afterEach()

    err := repo.CreateCustomer(ctx, custrepo.Customer{Name: "Tester"})
    assert.NoError(t, err)

    // Check database
    err = collection.FindOne(ctx, &result, bson.M{})
    assert.NoError(t, err)

    assert.Equal(t, "Tester", result.Name)
  })

  t.Run("Should be return err when insert customer", func(t *testing.T) {
    beforeEach()
    defer afterEach()

    id, _ := primitive.ObjectIDFromHex(strings.Repeat("0", 23) + "1")

    err := repo.CreateCustomer(ctx, custrepo.Customer{ID: id, Name: "Tester"})
    assert.NoError(t, err)

    // Create duplicate id
    err = repo.CreateCustomer(ctx, custrepo.Customer{ID: id, Name: "Tester"})
    assert.ErrorContains(t, err, "E11000 duplicate key error collection")
  })
}
```

## Run test

```shell
go test ./... -tags=integration
```

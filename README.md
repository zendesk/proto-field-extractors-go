# What

Field extractors provides a mechanism to transparently extract arbitrarily nested fields from protobuf messages in Go, only requiring the path to the desired value.

It can be used for abstracting protobuf message traversal, and for creating arbitrary transformations on protobuf encoded data without unnecessary boilerplate.
## Basic usage

Simple value extraction

```protobuf
message Account {
    google.protobuf.Int64Value   id = 1;
    google.protobuf.BoolValue    is_active = 2;
    google.protobuf.StringValue  name = 3;
    google.protobuf.StringValue  subdomain = 4;
    google.protobuf.StringValue  time_zone = 5;
    google.protobuf.Timestamp    created_at = 6;
    google.protobuf.Timestamp    updated_at = 7;
}
```

```go
type Account struct {
  ID   int64
  Name string
}


account := &Account{}
accountID, _ := extractors.BasicField("id").Extract(accountMessage)
// an alias 'B' is provided for readability of basicField extractions
accountName, _ := extractors.B("name").Extract(accountMessage)
account.ID =  accountID.(int64)
account.Name = accountName.(string)
```

Exract a field from a nested `oneof`


```protobuf
message UserEvent {
  oneof event {
    UserCreated user_created = 1;
    UserDeactivatd user_deactivated = 2;
  }
}

message UserCreated {
  google.protobuf.StringValue id = 1;
}

message UserDeactivated {
 google.protobuf.Timestamp occurred_at = 1;
}
```

```go
deactivatedAt, err := extractors.OneOf("event.user_deactivated", extractors.BasicField("occurred_at")).Extract(userEventMessage).(time.Time)
```

## Installation

```bash
go get github.com/zendesk/field-extractors-go
```
table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "username" {
    null    = false
    type    = varchar(25)
  }
  column "password" {
    null    = false
    type    = varchar(255)
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("now()")
  }
  primary_key {
    columns = [column.id]
  }
}

table "receipts" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "user_id" {
    type = int
  }
  column "description" {
    null = true
    type = text
  }
  column "total" {
    null = true
    type = int
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("now()")
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "user_id" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  check "non_empty_description" {
    expr = "description != ''"
  }
}

table "items" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "receipt_id" {
    type = int
  }
  column "price" {
    null = true
    type = int
  }
  column "name" {
    null = true
    type = character_varying(128)
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("now()")
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "receipt_id" {
    columns     = [column.receipt_id]
    ref_columns = [table.receipts.column.id]
    on_update   = NO_ACTION
    on_delete   = CASCADE
  }
  index "FK_items_receipt_id" {
    columns = [column.receipt_id]
  }
}

table "user_orders" {
  schema = schema.public
  column "id" {
    null = false
    type = serial
  }
  column "user_id" {
    null = true
    type = int
  }
  column "item_id" {
    null = true
    type = int
  }
  column "status" {
    null = true
    type = character_varying(128)
  }
  column "created_at" {
    null    = false
    type    = timestamp
    default = sql("now()")
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "user_id" {
    columns     = [column.user_id]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
  foreign_key "item_id" {
    columns     = [column.item_id]
    ref_columns = [table.items.column.id]
    on_update   = NO_ACTION
    on_delete   = NO_ACTION
  }
}

schema "public" {
}

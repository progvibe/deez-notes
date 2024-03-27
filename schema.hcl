schema "main" {

}

table "notes" {
  schema = schema.main
  column "id" {
    type = integer
    auto_increment = true
  }
  column "user_id" {
    type = int
  }
  column "title" {
    type = varchar(255)
    default = ""
  }
  column "content" {
    type = varchar(255)
    default = ""
  }
  primary_key {
    columns = [
      column.id
    ]
  }
  index "idx_user_id" {
    columns = [
      column.user_id
    ]
  }
  foreign_key "user_fk" {
    columns = [
      column.user_id
    ]
    ref_columns = [column.id]
  }
}

table "users" {
  schema = schema.main
  column "id" {
    type = integer
    auto_increment = true
  }
  column "name" {
    type = varchar(255)
  }
  primary_key {
    columns = [
      column.id
    ]
  }
}
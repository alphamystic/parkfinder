package domain



import (
  "fmt"
  "ken/lib/utils"
  ent"ken/lib/entities"

  "errors"
  "database/sql"
  "golang.org/x/crypto/bcrypt"
  _ "github.com/go-sql-driver/mysql"
)

// Add encryption
func (dom *Domain) CreateUser(u ent.UserData) error {
  var passwordHash []byte
  //create their hashes
  passwordHash,err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error creating password hash: %s",err))
    utils.Logerror(e)
    return err
  }
  var ins *sql.Stmt
  query := "INSERT INTO `park_finder`.`users` (userid, role, phone, name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?);"
  ins,err = dom.Dbs.Prepare(query)
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error Preparing to create user: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating user, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(u.UserID, u.Role, u.Phone, u.Name, u.Email, passwordHash, u.CreatedAt, u.UpdatedAt)
  if err != nil {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error executing insert user: %s",err))
    utils.Logerror(e)
    return errors.New("Error creating user while executing.")
  }
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error creating user, more than one row afffected: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating user.")
  }
  return nil
}


func (dom *Domain) ListUsers() ([]ent.UserData,error) {
  stmt := "SELECT * FROM `park_finder`.`users`;"
  rows,err := dom.Dbs.Query(stmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPLAPAs: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing all users.")
  }
  defer rows.Close()
  var users []ent.UserData
  for rows.Next(){
    var u ent.UserData
    err = rows.Scan(&u.ID,&u.UserID,&u.Role,&u.Phone,&u.Name,&u.Email,&u.Password,&u.CreatedAt,&u.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning for user: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all users.")
    }
    users = append(users,u)
  }
  return users,nil
}

func (dom *Domain) ViewUser(uid string) (*ent.UserData,error) {
  var u ent.UserData
  row := dom.Dbs.QueryRow("SELECT * FROM `park_finder`.`user` WHERE userid	 = ?;",uid)
  err := row.Scan(&u.ID,&u.UserID,&u.Role,&u.Phone,&u.Name,&u.Email,&u.Password,&u.CreatedAt,&u.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error viewing user with id %s %s",uid,err))
    utils.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing user with id of %s",uid))
  }
  return &u,nil
}

// Authenticates to the System
func (dom *Domain) Authenticate(email,password string) (*ent.UserData,bool){
  var u ent.UserData
  stmt := "SELECT * FROM `park_finder`.`users` WHERE email = ?;"
  row := dom.Dbs.QueryRow(stmt,email)
  fmt.Println("DB mail: ",email)
  err := row.Scan(&u.ID,&u.UserID,&u.Role,&u.Phone,&u.Name,&u.Email,&u.Password,&u.CreatedAt,&u.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning rows for authentication %s",err))
    utils.Logerror(e)
    return &u,false
  }
  err = bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(password))
  if err != nil{
    e := utils.LogErrorToFile("auth",fmt.Sprintf("Wrong login attempt for email %s with password %s  %s",email,password,err))
    utils.Logerror(e)
    return &u,false
  }
  u.Password = ""
  return &u,true
}

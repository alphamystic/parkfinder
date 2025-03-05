package domain



import (
  "fmt"
  "ken/lib/utils"
  ent"ken/lib/entities"

  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func (dom *Domain) CreateContactUS(cu ent.ContactUs) error {
  var ins *sql.Stmt
  ins,err := dom.Dbs.Prepare("INSERT INTO `park_finder`.`contactus` (cuid,name,email,subject,message,handled,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?);")
  if err !=  nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("EPCN: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating contact us, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(cu.CuID,cu.Name,cu.Email,cu.Email,cu.Subject,cu.Message,cu.Handled,cu.CreatedAt,cu.UpdatedAt)
  if err != nil {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error executing insert contact us: %s",err))
    utils.Logerror(e)
    return errors.New("Error creating contact us while executing.")
  }
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error more than one row affected in creating contact us: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating contact us.")
  }
  return nil
}

func (dom *Domain) ViewCU(cuid string)  (*ent.ContactUs,error){
  var cu ent.ContactUs
  row := dom.Dbs.QueryRow("SELECT * FROM `park_finder`.`contactus` WHERE (`cuid` = ?);",cuid)
  err := row.Scan(&cu.CuID,&cu.Name,&cu.Email,&cu.Email,&cu.Subject,&cu.Message,&cu.Handled,&cu.CreatedAt,&cu.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error viewing contact us with id: %s  %s",cuid,err))
    utils.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing contact us with the id of %s",cuid))
  }
  return &cu,nil
}

func (dom *Domain) MarlCUHandled(cid string,handled bool) error {
  upStmt := "UPDATE `park_finder`.`contactus` SET `handled` = ? WHERE (`cuid` = ?);";
  stmt,err := dom.Dbs.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error updating contact us: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while updating contact us.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(handled,cid)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error updating contact us with id: %s  %s",cid,err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while updating contact us.")
  }
  return nil
}

func (dom *Domain) ListCU(handled bool) ([]ent.ContactUs,error) {
  stmt := "SELECT * FROM `park_finder`.`contactus` WHERE (`handled` = ? ) ORDER BY updated_at DESC;"
  rows,err := dom.Dbs.Query(stmt,handled)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error listing contact us: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing contct us.")
  }
  defer rows.Close()
  var cus []ent.ContactUs
  for rows.Next(){
    var cu ent.ContactUs
    err = rows.Scan(&cu.CuID,&cu.Name,&cu.Email,&cu.Email,&cu.Subject,&cu.Message,&cu.Handled,&cu.CreatedAt,&cu.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning for contatc us: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing contact us.")
    }
    cus = append(cus,cu)
  }
  return cus,nil
}

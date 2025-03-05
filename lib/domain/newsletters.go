package domain



import (
  "fmt"
  "ken/lib/utils"
  ent"ken/lib/entities"

  "errors"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

func (dom *Domain) CreateNewsLetter(nl ent.Newsletter) error {
  var ins *sql.Stmt
  ins,err := dom.Dbs.Prepare("INSERT INTO `park_finder`.`newsletter` (newslt_id,email,subscribed,created_at,updated_at) VALUES(?,?,?,?,?);")
  if err !=  nil {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error Proccessing News Letter: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating news letter, Try again later :).")
  }
  defer ins.Close()
  res,err := ins.Exec(nl.NewsID,nl.Email,nl.Subscribed,nl.CreatedAt,nl.UpdatedAt)
  if err != nil {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error executing insert newsletter: %s",err))
    utils.Logerror(e)
    return errors.New("Error creating news letter while executing.")
  }
  rowsAffec, _  := res.RowsAffected()
  if err != nil || rowsAffec != 1{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error more than one row affected in creating news letter: %s",err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while creating news letter.")
  }
  return nil
}

func (dom *Domain) ViewNewsLetter(nid string) (*ent.Newsletter,error) {
  var nl ent.Newsletter
  row := dom.Dbs.QueryRow("SELECT * FROM `park_finder`.`newsletter` WHERE (`newslt_id` = ?);",nid)
  err := row.Scan(&nl.NewsID,&nl.Email,&nl.Subscribed,&nl.CreatedAt,&nl.UpdatedAt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error viewing newsletter id: %s  %s",nid,err))
    utils.Logerror(e)
    return nil,errors.New(fmt.Sprintf("Server encountered an error while viewing newsletter with the id of %s",nid))
  }
  return &nl,nil
}


func (dom *Domain) MarkNLAsHandled(nid string, subscribed bool) error {
  upStmt := "UPDATE `park_finder`.`newsletter` SET `subscribed` = ? WHERE (`newslt_id` = ?);";
  stmt,err := dom.Dbs.Prepare(upStmt)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error updating newsletter with id %s: %s",nid,err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while updating news letter.")
  }
  defer stmt.Close()
  var res sql.Result
  res,err = stmt.Exec(subscribed,nid)
  rowsAffec,_ := res.RowsAffected()
  if err != nil || rowsAffec != 1 {
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error marking newsletter newsletter id: %s  %s",nid,err))
    utils.Logerror(e)
    return errors.New("Server encountered an error while marking news; as handled.")
  }
  return nil
}

func (dom *Domain) ListNL(subscribed bool) ([]ent.Newsletter,error) {
  stmt := "SELECT * FROM `park_finder`.`newsletter` WHERE (`subscribed` = ? ) ORDER BY updated_at DESC;"
  rows,err := dom.Dbs.Query(stmt,subscribed)
  if err != nil{
    e := utils.LogErrorToFile("sql",fmt.Sprintf("Error listing News letter newsletters: %s",err))
    utils.Logerror(e)
    return nil,errors.New("Server encountered an error while listing newsletters.")
  }
  defer rows.Close()
  var nls []ent.Newsletter
  for rows.Next(){
    var nl ent.Newsletter
    err = rows.Scan(&nl.NewsID,&nl.Email,&nl.Subscribed,&nl.CreatedAt,&nl.UpdatedAt)
    if err != nil{
      e := utils.LogErrorToFile("sql",fmt.Sprintf("Error scanning for news letters: %s",err))
      utils.Logerror(e)
      return nil,errors.New("Server encountered an error while listing all handled news.")
    }
    nls = append(nls,nl)
  }
  return nls,nil
}

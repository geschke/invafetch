package invdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type ProcessdataEntry struct {
	ID          int64
	DateCreated string
	ProcessData string
}

type Repository struct {
	db *sql.DB
}

// NewRepository creates a new database representation
func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// GetShares makes something special
func (r *Repository) GetProcessdata() []ProcessdataEntry {

	//db := dbconn.ConnectDB(dsn)
	var items []ProcessdataEntry
	results, err := r.db.Query("SELECT * from solardata LIMIT 0,1")

	if err != nil {
		log.Println("Database problem in GetShares: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	//defer r.db.Close()

	log.Println("items in database:")
	//fmt.Printf("%-"+fmt.Sprintf("%d", maxStrlen)+"s%-21s%-21s\n", "Host", "Created", "Updated")
	var id int64

	var dt_created, processdata string

	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(&id, &dt_created, &processdata)
		if err != nil {
			log.Println(err.Error()) // proper error handling instead of panic in your app
			os.Exit(1)
		}
		// and then print out the tag's Name attribute
		//fmt.Printf("%-"+fmt.Sprintf("%d", maxStrlen)+"s%-21s%-21s\n", host+"."+domainname, dtCreated, dtUpdated)
		log.Printf("%d %s %s\n", id, dt_created, processdata)
		log.Println("Request: ")

		items = append(items, ProcessdataEntry{ID: id, DateCreated: dt_created, ProcessData: processdata})

	}
	return items

}

// AddAssetTransaction adds dataset to transactions_assets table, which is a representation of asset-wide manual actions
func (r *Repository) AddData(payload string) (int64, error) {

	log.Println("in addData...")

	fmt.Println("mit payload:", payload)

	result, err := r.db.Prepare("INSERT INTO solardata (processdata) VALUES(?)")
	if err != nil {

		fmt.Println(err.Error())
		return -1, err
		//os.Exit(1)
	}

	res, insertErr := result.Exec(payload)
	if insertErr != nil {
		//fmt.Println(insertErr.Error())
		return -1, insertErr
		//fmt.Println(insertErr.Error())
		//os.Exit(1)
	}

	//defer db.Close()
	return res.LastInsertId()
}

/*
// GetAssetEntry returns an AssetEntry type for an asset found in database
func (r *StockRepository) GetAssetEntry(assetID int) AssetEntry {

	//err := db.QueryRowContext(ctx, "SELECT username, created_at FROM users WHERE id=?", id).Scan(&username, &created)

	var id, datasourceID, assetTypeID int64
	var isin, wkn, symbolName, symbolShort, description string
	var resultEntry AssetEntry

	err := r.db.QueryRow("SELECT sm.id, sm.isin, sm.wkn, sm.symbol_name, sm.symbol_short, sm.description, sm.datasource_id, sm.asset_type FROM symbolmeta sm WHERE sm.id=?", assetID).Scan(&id, &isin, &wkn, &symbolName, &symbolShort, &description, &datasourceID, &assetTypeID)

	if err != nil {
		log.Println("Database problem in GetAssetEntry: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	resultEntry.ID = id
	resultEntry.Isin = isin
	resultEntry.SymbolName = symbolName
	resultEntry.SymbolShort = symbolShort
	resultEntry.Description = description
	resultEntry.Wkn = wkn
	resultEntry.DatasourceID = datasourceID
	resultEntry.AssetTypeID = assetTypeID

	return resultEntry

}
*/

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

type HomeConsumption struct {
	ID          int64
	DateCreated string
	HomeOwnP    string
	HomePvP     string
	HomeP       string
	HomeGridP   string
	HomeBatP    string
}

type DevicesLocalBatteryLast struct {
	ID              int64
	DateCreated     string
	BatManufacturer string
	BatModel        string
	BatSerialNo     string
	BatVersionFW    string
	Cycles          string
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

// GetHomeConsumption loads home consumption values from database as average of the last 1 minute
func (r *Repository) GetHomeConsumption() HomeConsumption {

	//db := dbconn.ConnectDB(dsn)
	var values HomeConsumption
	//results, err := r.db.Query("SELECT * from solardata LIMIT 0,1")

	results, err := r.db.Query("SELECT id, dt_created, avg(JSON_VALUE(processdata,'$.devices:local.HomeOwn_P.value')) AS home_own_p, avg(JSON_VALUE(processdata,'$.devices:local.HomePv_P.value')) AS home_pv_p, avg(JSON_VALUE(processdata,'$.devices:local.Home_P.value')) AS home_p, avg(JSON_VALUE(processdata,'$.devices:local.HomeGrid_P.value')) AS home_grid_p, avg(JSON_VALUE(processdata,'$.devices:local.HomeBat_P.value')) AS home_bat_p FROM solardata WHERE dt_created < NOW() AND dt_created > NOW() - INTERVAL 1 Minute")

	if err != nil {
		log.Println("Database problem in GetProcessdata: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	//defer r.db.Close()

	log.Println("items in database:")
	//fmt.Printf("%-"+fmt.Sprintf("%d", maxStrlen)+"s%-21s%-21s\n", "Host", "Created", "Updated")
	var id int64
	var dt_created string
	var home_own_p, home_pv_p, home_p, home_grid_p, home_bat_p string

	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(&id, &dt_created, &home_own_p, &home_pv_p, &home_p, &home_grid_p, &home_bat_p)
		if err != nil {
			log.Println(err.Error()) // proper error handling instead of panic in your app
			os.Exit(1)
		}
		// and then print out the tag's Name attribute
		//fmt.Printf("%-"+fmt.Sprintf("%d", maxStrlen)+"s%-21s%-21s\n", host+"."+domainname, dtCreated, dtUpdated)
		log.Printf("%d %s %s %s %s %s %s\n", id, dt_created, home_own_p, home_pv_p, home_p, home_grid_p, home_bat_p)
		log.Println("Request: ")

		values = HomeConsumption{ID: id, DateCreated: dt_created, HomeOwnP: home_own_p, HomePvP: home_pv_p, HomeP: home_p, HomeGridP: home_grid_p, HomeBatP: home_bat_p}

	}
	return values

}

// GetHomeConsumption loads home consumption values from database as average of the last 1 minute
func (r *Repository) GetDevicesLocalBatteryLast() DevicesLocalBatteryLast {

	var values DevicesLocalBatteryLast

	results, err := r.db.Query("SELECT id, dt_created, JSON_VALUE(processdata,'$.devices:local:battery.BatManufacturer.value') AS bat_manufacturer, JSON_VALUE(processdata,'$.devices:local:battery.BatModel.value') AS bat_model, JSON_VALUE(processdata,'$.devices:local:battery.BatSerialNo.value') AS bat_serial_no, JSON_VALUE(processdata,'$.devices:local:battery.BatVersionFW.value') AS bat_version_fw, JSON_VALUE(processdata,'$.devices:local:battery.Cycles.value') AS cycles FROM solardata order by dt_created desc LIMIT 0,1")

	if err != nil {
		log.Println("Database problem in GetDevicesLocalBatteryLast: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	//defer r.db.Close()

	log.Println("items in database:")
	//fmt.Printf("%-"+fmt.Sprintf("%d", maxStrlen)+"s%-21s%-21s\n", "Host", "Created", "Updated")
	var id int64
	var dt_created string
	var bat_manufacturer, bat_model, bat_serial_no, bat_version_fw, cycles string

	for results.Next() {

		// for each row, scan the result into our tag composite object
		err = results.Scan(&id, &dt_created, &bat_manufacturer, &bat_model, &bat_serial_no, &bat_version_fw, &cycles)
		if err != nil {
			log.Println(err.Error()) // proper error handling instead of panic in your app
			os.Exit(1)
		}
		// and then print out the tag's Name attribute
		//fmt.Printf("%-"+fmt.Sprintf("%d", maxStrlen)+"s%-21s%-21s\n", host+"."+domainname, dtCreated, dtUpdated)
		log.Printf("%d %s %s %s %s %s %s\n", id, dt_created, bat_manufacturer, bat_model, bat_serial_no, bat_version_fw, cycles)
		log.Println("Request: ")

		values = DevicesLocalBatteryLast{ID: id, DateCreated: dt_created, BatManufacturer: bat_manufacturer, BatModel: bat_model, BatSerialNo: bat_serial_no, BatVersionFW: bat_version_fw, Cycles: cycles}

	}
	return values

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

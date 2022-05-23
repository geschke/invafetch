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

type DevicesLocal struct {
	ID          int64
	DateCreated string
	Bat2Grid_P  string
	Dc_P        string
	DigitalIn   string
	EM_State    string
	Grid2Bat_P  string
	Grid_L1_I   string
	Grid_L1_P   string
	Grid_L2_I   string
	Grid_L2_P   string
	Grid_L3_I   string
	Grid_L3_P   string
	Grid_P      string
	Grid_Q      string
	Grid_S      string
	HomeBat_P   string
	HomeGrid_P  string
	HomeOwn_P   string
	HomePv_P    string
	Home_P      string
	Iso_R       string
	LimitEvuRel string
	PV2Bat_P    string
}

type DevicesLocalAc struct {
	ID            int64
	DateCreated   string
	CosPhi        string
	Frequency     string
	InvIn_P       string
	InvOut_P      string
	L1_I          string
	L1_P          string
	L1_U          string
	L2_I          string
	L2_P          string
	L2_U          string
	L3_I          string
	L3_P          string
	L3_U          string
	P             string
	Q             string
	ResidualCDc_I string
	S             string
}

type DevicesLocalLast struct {
	ID            int64
	DateCreated   string
	InverterState string
	SinkMax_P     string
	SourceMax_P   string
	WorkTime      string
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

type DevicesLocalBattery struct {
	ID              int64
	DateCreated     string
	FullChargeCap_E string
	I               string
	P               string
	SoC             string
	U               string
	WorkCapacity    string
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

// GetDevicesLocalBattery
func (r *Repository) GetDevicesLocalBattery() DevicesLocalBattery {

	var values DevicesLocalBattery
	var id int64
	var dt_created string
	var full_charge_cap_e, i, p, soc, u, work_capacity string

	err := r.db.QueryRow("SELECT id, dt_created, avg(JSON_VALUE(processdata,'$.devices:local:battery.FullChargeCap_E.value')) AS full_charge_cap_e, avg(JSON_VALUE(processdata,'$.devices:local:battery.I.value')) AS i, avg(JSON_VALUE(processdata,'$.devices:local:battery.P.value')) AS p, avg(JSON_VALUE(processdata,'$.devices:local:battery.SoC.value')) AS soc, avg(JSON_VALUE(processdata,'$.devices:local:battery.U.value')) AS u, avg(JSON_VALUE(processdata,'$.devices:local:battery.WorkCapacity.value')) AS work_capacity FROM solardata WHERE dt_created < NOW() AND dt_created > NOW() - INTERVAL 1 Minute").Scan(&id, &dt_created, &full_charge_cap_e, &i, &p, &soc, &u, &work_capacity)

	if err != nil {
		log.Println("Database problem in GetDevicesLocalBattery: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	values = DevicesLocalBattery{FullChargeCap_E: full_charge_cap_e, I: i, P: p, SoC: soc, U: u, WorkCapacity: work_capacity}

	return values

}

// GetDevicesLocalBatteryLast
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

// GetDevicesLocal
func (r *Repository) GetDevicesLocal() DevicesLocal {

	var values DevicesLocal
	var id int64
	var dt_created string
	var bat2grid_p, dc_p, digital_in, em_state, grid2bat_p, grid_l1_i, grid_l1_p, grid_l2_i, grid_l2_p, grid_l3_i, grid_l3_p, grid_p, grid_q, grid_s, home_bat_p, home_grid_p, home_own_p, home_pv_p, home_p, iso_r, limit_evu_rel, pv2bat_p string

	err := r.db.QueryRow("SELECT id, dt_created, avg(JSON_VALUE(processdata,'$.devices:local.Bat2Grid_P.value')) AS bat2grid_p, avg(JSON_VALUE(processdata,'$.devices:local.Dc_P.value')) AS dc_p, avg(JSON_VALUE(processdata,'$.devices:local.DigitalIn.value')) AS digital_in, avg(JSON_VALUE(processdata,'$.devices:local.EM_State.value')) AS em_state, avg(JSON_VALUE(processdata,'$.devices:local.Grid2Bat_P.value')) AS grid2bat_p, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L1_I.value')) AS grid_l1_i, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L1_P.value')) AS grid_l1_p, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L1_P.value')) AS grid_l1_p, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L2_I.value')) AS grid_l2_i, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L2_P.value')) AS grid_l2_p, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L3_I.value')) AS grid_l3_i, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L3_P.value')) AS grid_l3_p, avg(JSON_VALUE(processdata,'$.devices:local.Grid_P.value')) AS grid_p, avg(JSON_VALUE(processdata,'$.devices:local.Grid_Q.value')) AS grid_q, avg(JSON_VALUE(processdata,'$.devices:local.Grid_S.value')) AS grid_s, avg(JSON_VALUE(processdata,'$.devices:local.HomeBat_P.value')) AS home_bat_p, avg(JSON_VALUE(processdata,'$.devices:local.HomeGrid_P.value')) AS home_grid_p, avg(JSON_VALUE(processdata,'$.devices:local.HomeOwn_P.value')) AS home_own_p, avg(JSON_VALUE(processdata,'$.devices:local.HomePv_P.value')) AS home_pv_p, avg(JSON_VALUE(processdata,'$.devices:local.Home_P.value')) AS home_p, avg(JSON_VALUE(processdata,'$.devices:local.Iso_R.value')) AS iso_r, avg(JSON_VALUE(processdata,'$.devices:local.LimitEvuRel.value')) AS limit_evu_rel, avg(JSON_VALUE(processdata,'$.devices:local.PV2Bat_P.value')) AS pv2bat_p FROM solardata WHERE dt_created < NOW() AND dt_created > NOW() - INTERVAL 5 Minute").Scan(&id, &dt_created, &bat2grid_p, &dc_p, &digital_in, &em_state, &grid2bat_p, &grid_l1_i, &grid_l1_p, &grid_l2_i, &grid_l2_p, &grid_l3_i, &grid_l3_p, &grid_p, &grid_q, &grid_s, &home_bat_p, &home_grid_p, &home_own_p, &home_own_p, &home_pv_p, &home_p, &iso_r, &limit_evu_rel, &pv2bat_p)

	if err != nil {
		log.Println("Database problem in GetDevicesLocal: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	values = DevicesLocal{Bat2Grid_P: bat2grid_p, Dc_P: dc_p, DigitalIn: digital_in, EM_State: em_state, Grid2Bat_P: grid2bat_p, Grid_L1_I: grid_l1_i, Grid_L1_P: grid_l1_p, Grid_L2_I: grid_l2_i, Grid_L2_P: grid_l2_p, Grid_L3_I: grid_l3_i, Grid_L3_P: grid_l3_p, Grid_P: grid_p, Grid_Q: grid_q, Grid_S: grid_s, HomeBat_P: home_bat_p, HomeGrid_P: home_grid_p, HomeOwn_P: home_own_p, HomePv_P: home_pv_p, Home_P: home_p, Iso_R: iso_r, LimitEvuRel: limit_evu_rel, PV2Bat_P: pv2bat_p}

	return values

}

// GetDevicesLocal
func (r *Repository) GetDevicesLocalAc() DevicesLocalAc {

	var values DevicesLocalAc
	var id int64
	var dt_created string
	var cos_phi, frequency, inv_in_p, inv_out_p, l1_i, l1_p, l1_u, l2_i, l2_p, l2_u, l3_i, l3_p, l3_u, p, q, residual_cdc_i, s string

	err := r.db.QueryRow("SELECT id, dt_created, avg(JSON_VALUE(processdata,'$.devices:local:ac.CosPhi.value')) AS cos_phi, avg(JSON_VALUE(processdata,'$.devices:local:ac.Frequency.value')) AS frequency, avg(JSON_VALUE(processdata,'$.devices:local:ac.InvIn_P.value')) AS inv_in_p, avg(JSON_VALUE(processdata,'$.devices:local:ac.InvOut_P.value')) AS inv_out_p, avg(JSON_VALUE(processdata,'$.devices:local:ac.L1_I.value')) AS l1_i, avg(JSON_VALUE(processdata,'$.devices:local:ac.L1_P.value')) AS l1_p, avg(JSON_VALUE(processdata,'$.devices:local:ac.L1_U.value')) AS l1_u, avg(JSON_VALUE(processdata,'$.devices:local:ac.L2_I.value')) AS l2_i, avg(JSON_VALUE(processdata,'$.devices:local:ac.L2_P.value')) AS l2_p, avg(JSON_VALUE(processdata,'$.devices:local:ac.L2_U.value')) AS l2_u, avg(JSON_VALUE(processdata,'$.devices:local:ac.L3_I.value')) AS l3_i, avg(JSON_VALUE(processdata,'$.devices:local:ac.L3_P.value')) AS l3_p, avg(JSON_VALUE(processdata,'$.devices:local:ac.L3_U.value')) AS l3_u, avg(JSON_VALUE(processdata,'$.devices:local:ac.P.value')) AS p, avg(JSON_VALUE(processdata,'$.devices:local:ac.Q.value')) AS q, avg(JSON_VALUE(processdata,'$.devices:local:ac.ResidualCDc_I.value')) AS residual_cdc_i, avg(JSON_VALUE(processdata,'$.devices:local:ac.S.value')) AS s FROM solardata WHERE dt_created < NOW() AND dt_created > NOW() - INTERVAL 5 Minute").Scan(&id, &dt_created, &cos_phi, &frequency, &inv_in_p, &inv_out_p, &l1_i, &l1_p, &l1_u, &l2_i, &l2_p, &l2_u, &l3_i, &l3_p, &l3_u, &p, &q, &residual_cdc_i, &s)

	if err != nil {
		log.Println("Database problem in GetDevicesLocal: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	values = DevicesLocalAc{CosPhi: cos_phi, Frequency: frequency, InvIn_P: inv_in_p, InvOut_P: inv_out_p, L1_I: l1_i, L1_P: l1_p, L1_U: l1_u, L2_I: l2_i, L2_P: l2_p, L2_U: l2_u, L3_I: l3_i, L3_P: l3_p, L3_U: l3_u, P: p, Q: q, ResidualCDc_I: residual_cdc_i, S: s}

	return values

}

// GetDevicesLocalLast
func (r *Repository) GetDevicesLocalLast() DevicesLocalLast {

	var values DevicesLocalLast
	var id int64
	var dt_created string
	var inverter_state, sink_max_p, source_max_p, work_time string

	err := r.db.QueryRow("SELECT id, dt_created, JSON_VALUE(processdata,'$.devices:local.Inverter:State.value') AS inverter_state, JSON_VALUE(processdata,'$.devices:local.SinkMax_P.value') AS sink_max_p, JSON_VALUE(processdata,'$.devices:local.SourceMax_P.value') AS source_max_p, JSON_VALUE(processdata,'$.devices:local.WorkTime.value') AS work_time FROM solardata order by dt_created desc LIMIT 0,1").Scan(&id, &dt_created, &inverter_state, &sink_max_p, &source_max_p, &work_time)

	if err != nil {
		log.Println("Database problem in GetDevicesLocalLast: " + err.Error())
		os.Exit(1)
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	values = DevicesLocalLast{InverterState: inverter_state, SinkMax_P: sink_max_p, SourceMax_P: source_max_p, WorkTime: work_time}

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

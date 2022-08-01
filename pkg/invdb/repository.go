// Copyright 2022 Ralf Geschke. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package invdb

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	"github.com/geschke/invafetch/pkg/dbconn"
	_ "github.com/go-sql-driver/mysql"
)

type ProcessdataEntry struct {
	ID          int64
	DateCreated string
	ProcessData string
}

type DevicesLocal struct {
	ID          int64
	DateCreated string
	Bat2Grid_P  sql.NullString // string
	Dc_P        sql.NullString // string
	DigitalIn   sql.NullString // string
	EM_State    sql.NullString // string
	Grid2Bat_P  sql.NullString // string
	Grid_L1_I   sql.NullString // string
	Grid_L1_P   sql.NullString // string
	Grid_L2_I   sql.NullString // string
	Grid_L2_P   sql.NullString // string
	Grid_L3_I   sql.NullString // string
	Grid_L3_P   sql.NullString // string
	Grid_P      sql.NullString // string
	Grid_Q      sql.NullString // string
	Grid_S      sql.NullString // string
	HomeBat_P   sql.NullString //string
	HomeGrid_P  sql.NullString //string
	HomeOwn_P   sql.NullString //string
	HomePv_P    sql.NullString //string
	Home_P      sql.NullString //string
	Iso_R       sql.NullString // string
	LimitEvuRel sql.NullString // string
	PV2Bat_P    sql.NullString // string
}

type DevicesLocalAc struct {
	ID            int64
	DateCreated   string
	CosPhi        sql.NullString // string
	Frequency     sql.NullString // string
	InvIn_P       sql.NullString // string
	InvOut_P      sql.NullString // string
	L1_I          sql.NullString // string
	L1_P          sql.NullString // string
	L1_U          sql.NullString // string
	L2_I          sql.NullString // string
	L2_P          sql.NullString // string
	L2_U          sql.NullString // string
	L3_I          sql.NullString // string
	L3_P          sql.NullString // string
	L3_U          sql.NullString // string
	P             sql.NullString // string
	Q             sql.NullString // string
	ResidualCDc_I sql.NullString // string
	S             sql.NullString // string
}

type DevicesLocalLast struct {
	ID            int64
	DateCreated   string
	InverterState sql.NullString // string
	SinkMax_P     sql.NullString // string
	SourceMax_P   sql.NullString // string
	WorkTime      sql.NullString // string
}

type DevicesLocalBatteryLast struct {
	ID              int64
	DateCreated     string
	BatManufacturer sql.NullString // string
	BatModel        sql.NullString // string
	BatSerialNo     sql.NullString // string
	BatVersionFW    sql.NullString // string
	Cycles          sql.NullString // string
}

type DevicesLocalBattery struct {
	ID              int64
	DateCreated     string
	FullChargeCap_E sql.NullString // string
	I               sql.NullString // string
	P               sql.NullString // string
	SoC             sql.NullString // string
	U               sql.NullString // string
	WorkCapacity    sql.NullString // string
}

type DevicesLocalPowermeter struct {
	ID          int64
	DateCreated string
	CosPhi      sql.NullString // string
	Frequency   sql.NullString // string
	P           sql.NullString // string
	Q           sql.NullString // string
	S           sql.NullString // string
}

type DevicesLocalPv struct {
	ID          int64
	DateCreated string
	I           sql.NullString // string
	P           sql.NullString // string
	U           sql.NullString // string
}

type StatisticEnergyFlowLast struct {
	ID                                int64
	DateCreated                       string
	StatisticAutarkyDay               sql.NullString // string
	StatisticAutarkyMonth             sql.NullString // string
	StatisticAutarkyTotal             sql.NullString // string
	StatisticAutarkyYear              sql.NullString // string
	StatisticCO2SavingDay             sql.NullString // string
	StatisticCO2SavingMonth           sql.NullString // string
	StatisticCO2SavingTotal           sql.NullString // string
	StatisticCO2SavingYear            sql.NullString // string
	StatisticEnergyChargeGridDay      sql.NullString // string
	StatisticEnergyChargeGridMonth    sql.NullString // string
	StatisticEnergyChargeGridTotal    sql.NullString // string
	StatisticEnergyChargeGridYear     sql.NullString // string
	StatisticEnergyChargeInvInDay     sql.NullString // string
	StatisticEnergyChargeInvInMonth   sql.NullString // string
	StatisticEnergyChargeInvInTotal   sql.NullString // string
	StatisticEnergyChargeInvInYear    sql.NullString // string
	StatisticEnergyChargePvDay        sql.NullString // string
	StatisticEnergyChargePvMonth      sql.NullString // string
	StatisticEnergyChargePvTotal      sql.NullString // string
	StatisticEnergyChargePvYear       sql.NullString // string
	StatisticEnergyDischargeDay       sql.NullString // string
	StatisticEnergyDischargeMonth     sql.NullString // string
	StatisticEnergyDischargeTotal     sql.NullString // string
	StatisticEnergyDischargeYear      sql.NullString // string
	StatisticEnergyDischargeGridDay   sql.NullString // string
	StatisticEnergyDischargeGridMonth sql.NullString // string
	StatisticEnergyDischargeGridTotal sql.NullString // string
	StatisticEnergyDischargeGridYear  sql.NullString // string
	StatisticEnergyHomeDay            sql.NullString // string
	StatisticEnergyHomeMonth          sql.NullString // string
	StatisticEnergyHomeTotal          sql.NullString // string
	StatisticEnergyHomeYear           sql.NullString // string
	StatisticEnergyHomeBatDay         sql.NullString // string
	StatisticEnergyHomeBatMonth       sql.NullString // string
	StatisticEnergyHomeBatTotal       sql.NullString // string
	StatisticEnergyHomeBatYear        sql.NullString // string
	StatisticEnergyHomeGridDay        sql.NullString // string
	StatisticEnergyHomeGridMonth      sql.NullString // string
	StatisticEnergyHomeGridTotal      sql.NullString // string
	StatisticEnergyHomeGridYear       sql.NullString // string
	StatisticEnergyHomeOwnTotal       sql.NullString // string
	StatisticEnergyHomePvDay          sql.NullString // string
	StatisticEnergyHomePvMonth        sql.NullString // string
	StatisticEnergyHomePvTotal        sql.NullString // string
	StatisticEnergyHomePvYear         sql.NullString // string
	StatisticEnergyPv1Day             sql.NullString // string
	StatisticEnergyPv1Month           sql.NullString // string
	StatisticEnergyPv1Total           sql.NullString // string
	StatisticEnergyPv1Year            sql.NullString // string
	StatisticEnergyPv2Day             sql.NullString // string
	StatisticEnergyPv2Month           sql.NullString // string
	StatisticEnergyPv2Total           sql.NullString // string
	StatisticEnergyPv2Year            sql.NullString // string
	StatisticEnergyPv3Day             sql.NullString // string
	StatisticEnergyPv3Month           sql.NullString // string
	StatisticEnergyPv3Total           sql.NullString // string
	StatisticEnergyPv3Year            sql.NullString // string
	StatisticOwnConsumptionRateDay    sql.NullString // string
	StatisticOwnConsumptionRateMonth  sql.NullString // string
	StatisticOwnConsumptionRateTotal  sql.NullString // string
	StatisticOwnConsumptionRateYear   sql.NullString // string
	StatisticYieldDay                 sql.NullString // string
	StatisticYieldMonth               sql.NullString // string
	StatisticYieldTotal               sql.NullString // string
	StatisticYieldYear                sql.NullString // string
}

type Repository struct {
	db *sql.DB
}

// NewRepository creates a new database representation
func NewRepository(db *sql.DB) *Repository {
	var repository Repository
	repository.db = db
	return &repository
	/*return Repository{
		db: db,
	}*/
}

func (r *Repository) Close() error {
	err := dbconn.CloseDB(r.db)
	if err != nil {
		return err
	}
	return nil
}

// ConvertToString converts a value created by reflect to string. It returns an empty string in case of error. So it's possible to use
// sql.NullString and string type in processdata type definition.
func (r *Repository) ConvertToString(payload *reflect.Value) string {

	var checkSqlNullString sql.NullString
	payloadType := payload.Type()

	if payloadType == reflect.TypeOf(checkSqlNullString) { // payload is type sql.NullString
		if !payload.FieldByName("Valid").Bool() { // invalid, return 0
			return "0"
		}
		return payload.FieldByName("String").String()
	} else if payload.Kind().String() == "string" { // payload is a string
		return payload.String()
	}
	return "0" // payload is something else, return 0
}

// GetDevicesLocalBattery loads the values from devices:local:battery section as average of the last 1 minute
func (r *Repository) GetDevicesLocalBattery() (DevicesLocalBattery, error) {

	var values DevicesLocalBattery

	err := r.db.QueryRow("SELECT id, dt_created, avg(JSON_VALUE(processdata,'$.devices:local:battery.FullChargeCap_E.value')) AS full_charge_cap_e, avg(JSON_VALUE(processdata,'$.devices:local:battery.I.value')) AS i, avg(JSON_VALUE(processdata,'$.devices:local:battery.P.value')) AS p, avg(JSON_VALUE(processdata,'$.devices:local:battery.SoC.value')) AS soc, avg(JSON_VALUE(processdata,'$.devices:local:battery.U.value')) AS u, avg(JSON_VALUE(processdata,'$.devices:local:battery.WorkCapacity.value')) AS work_capacity FROM solardata WHERE dt_created < NOW() AND dt_created > NOW() - INTERVAL 30 SECOND").Scan(&values.ID, &values.DateCreated, &values.FullChargeCap_E, &values.I, &values.P, &values.SoC, &values.U, &values.WorkCapacity)

	if err != nil {
		//log.Println("Database problem in GetDevicesLocalBattery: " + err.Error())
		return values, err
	}

	return values, nil

}

// GetDevicesLocalBatteryLast loads the values from devices:local:battery section which are specified as last values
func (r *Repository) GetDevicesLocalBatteryLast() (DevicesLocalBatteryLast, error) {

	var values DevicesLocalBatteryLast

	err := r.db.QueryRow("SELECT id, dt_created, JSON_VALUE(processdata,'$.devices:local:battery.BatManufacturer.value') AS bat_manufacturer, JSON_VALUE(processdata,'$.devices:local:battery.BatModel.value') AS bat_model, JSON_VALUE(processdata,'$.devices:local:battery.BatSerialNo.value') AS bat_serial_no, JSON_VALUE(processdata,'$.devices:local:battery.BatVersionFW.value') AS bat_version_fw, JSON_VALUE(processdata,'$.devices:local:battery.Cycles.value') AS cycles FROM solardata order by dt_created desc LIMIT 0,1").Scan(&values.ID, &values.DateCreated, &values.BatManufacturer, &values.BatModel, &values.BatSerialNo, &values.BatVersionFW, &values.Cycles)

	if err != nil {
		//log.Println("Database problem in GetDevicesLocalBatteryLast: " + err.Error())
		return values, err
	}

	//defer r.db.Close
	return values, nil

}

// GetDevicesLocal loads the values from devices:local section as average of the last 5 minutes
func (r *Repository) GetDevicesLocal() (DevicesLocal, error) {

	var values DevicesLocal

	err := r.db.QueryRow("SELECT id, dt_created, avg(JSON_VALUE(processdata,'$.devices:local.Bat2Grid_P.value')) AS bat2grid_p, avg(JSON_VALUE(processdata,'$.devices:local.Dc_P.value')) AS dc_p, avg(JSON_VALUE(processdata,'$.devices:local.DigitalIn.value')) AS digital_in, avg(JSON_VALUE(processdata,'$.devices:local.EM_State.value')) AS em_state, avg(JSON_VALUE(processdata,'$.devices:local.Grid2Bat_P.value')) AS grid2bat_p, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L1_I.value')) AS grid_l1_i, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L1_P.value')) AS grid_l1_p, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L2_I.value')) AS grid_l2_i, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L2_P.value')) AS grid_l2_p, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L3_I.value')) AS grid_l3_i, avg(JSON_VALUE(processdata,'$.devices:local.Grid_L3_P.value')) AS grid_l3_p, avg(JSON_VALUE(processdata,'$.devices:local.Grid_P.value')) AS grid_p, avg(JSON_VALUE(processdata,'$.devices:local.Grid_Q.value')) AS grid_q, avg(JSON_VALUE(processdata,'$.devices:local.Grid_S.value')) AS grid_s, avg(JSON_VALUE(processdata,'$.devices:local.HomeBat_P.value')) AS home_bat_p, avg(JSON_VALUE(processdata,'$.devices:local.HomeGrid_P.value')) AS home_grid_p, avg(JSON_VALUE(processdata,'$.devices:local.HomeOwn_P.value')) AS home_own_p, avg(JSON_VALUE(processdata,'$.devices:local.HomePv_P.value')) AS home_pv_p, avg(JSON_VALUE(processdata,'$.devices:local.Home_P.value')) AS home_p, avg(JSON_VALUE(processdata,'$.devices:local.Iso_R.value')) AS iso_r, avg(JSON_VALUE(processdata,'$.devices:local.LimitEvuRel.value')) AS limit_evu_rel, avg(JSON_VALUE(processdata,'$.devices:local.PV2Bat_P.value')) AS pv2bat_p FROM solardata WHERE dt_created < NOW() AND dt_created > NOW() - INTERVAL 30 SECOND").Scan(&values.ID, &values.DateCreated, &values.Bat2Grid_P, &values.Dc_P, &values.DigitalIn, &values.EM_State, &values.Grid2Bat_P, &values.Grid_L1_I, &values.Grid_L1_P, &values.Grid_L2_I, &values.Grid_L2_P, &values.Grid_L3_I, &values.Grid_L3_P, &values.Grid_P, &values.Grid_Q, &values.Grid_S, &values.HomeBat_P, &values.HomeGrid_P, &values.HomeOwn_P, &values.HomePv_P, &values.Home_P, &values.Iso_R, &values.LimitEvuRel, &values.PV2Bat_P)

	if err != nil {
		//log.Println("Database problem in GetDevicesLocal: " + err.Error())
		return values, err
	}
	return values, nil
}

// GetDevicesLocalAc loads the values from devices:local:ac section as average of the last 5 minutes
func (r *Repository) GetDevicesLocalAc() (DevicesLocalAc, error) {

	var values DevicesLocalAc

	err := r.db.QueryRow("SELECT id, dt_created, avg(JSON_VALUE(processdata,'$.devices:local:ac.CosPhi.value')) AS cos_phi, avg(JSON_VALUE(processdata,'$.devices:local:ac.Frequency.value')) AS frequency, avg(JSON_VALUE(processdata,'$.devices:local:ac.InvIn_P.value')) AS inv_in_p, avg(JSON_VALUE(processdata,'$.devices:local:ac.InvOut_P.value')) AS inv_out_p, avg(JSON_VALUE(processdata,'$.devices:local:ac.L1_I.value')) AS l1_i, avg(JSON_VALUE(processdata,'$.devices:local:ac.L1_P.value')) AS l1_p, avg(JSON_VALUE(processdata,'$.devices:local:ac.L1_U.value')) AS l1_u, avg(JSON_VALUE(processdata,'$.devices:local:ac.L2_I.value')) AS l2_i, avg(JSON_VALUE(processdata,'$.devices:local:ac.L2_P.value')) AS l2_p, avg(JSON_VALUE(processdata,'$.devices:local:ac.L2_U.value')) AS l2_u, avg(JSON_VALUE(processdata,'$.devices:local:ac.L3_I.value')) AS l3_i, avg(JSON_VALUE(processdata,'$.devices:local:ac.L3_P.value')) AS l3_p, avg(JSON_VALUE(processdata,'$.devices:local:ac.L3_U.value')) AS l3_u, avg(JSON_VALUE(processdata,'$.devices:local:ac.P.value')) AS p, avg(JSON_VALUE(processdata,'$.devices:local:ac.Q.value')) AS q, avg(JSON_VALUE(processdata,'$.devices:local:ac.ResidualCDc_I.value')) AS residual_cdc_i, avg(JSON_VALUE(processdata,'$.devices:local:ac.S.value')) AS s FROM solardata WHERE dt_created < NOW() AND dt_created > NOW() - INTERVAL 30 SECOND").Scan(&values.ID, &values.DateCreated, &values.CosPhi, &values.Frequency, &values.InvIn_P, &values.InvOut_P, &values.L1_I, &values.L1_P, &values.L1_U, &values.L2_I, &values.L2_P, &values.L2_U, &values.L3_I, &values.L3_P, &values.L3_U, &values.P, &values.Q, &values.ResidualCDc_I, &values.S)

	if err != nil {
		//log.Println("Database problem in GetDevicesLocal: " + err.Error())
		return values, err
	}

	return values, nil

}

// GetDevicesLocalPowermeter loads the values from devices:local:powermeter section as average of the last 5 minutes
func (r *Repository) GetDevicesLocalPowermeter() (DevicesLocalPowermeter, error) {

	var values DevicesLocalPowermeter

	err := r.db.QueryRow("SELECT id, dt_created, avg(JSON_VALUE(processdata,'$.devices:local:powermeter.CosPhi.value')) AS cos_phi, avg(JSON_VALUE(processdata,'$.devices:local:powermeter.Frequency.value')) AS frequency, avg(JSON_VALUE(processdata,'$.devices:local:powermeter.P.value')) AS p, avg(JSON_VALUE(processdata,'$.devices:local:powermeter.Q.value')) AS q, avg(JSON_VALUE(processdata,'$.devices:local:powermeter.S.value')) AS s FROM solardata WHERE dt_created < NOW() AND dt_created > NOW() - INTERVAL 30 SECOND").Scan(&values.ID, &values.DateCreated, &values.CosPhi, &values.Frequency, &values.P, &values.Q, &values.S)

	if err != nil {
		log.Println("Database problem in GetDevicesLocal: " + err.Error())
		return values, err
	}
	return values, nil
}

// GetDevicesLocalPv1 loads the values from devices:local:pv1 section as average of the last 5 minutes
func (r *Repository) GetDevicesLocalPv1() (DevicesLocalPv, error) {

	var values DevicesLocalPv

	err := r.db.QueryRow("SELECT id, dt_created, avg(JSON_VALUE(processdata,'$.devices:local:pv1.I.value')) AS i, avg(JSON_VALUE(processdata,'$.devices:local:pv1.P.value')) AS p, avg(JSON_VALUE(processdata,'$.devices:local:pv1.U.value')) AS u FROM solardata WHERE dt_created < NOW() AND dt_created > NOW() - INTERVAL 30 SECOND").Scan(&values.ID, &values.DateCreated, &values.I, &values.P, &values.U)

	if err != nil {
		//log.Println("Database problem in GetDevicesPv1: " + err.Error())
		return values, err
	}
	return values, nil
}

// GetDevicesLocalPv2 loads the values from devices:local:pv2 section as average of the last 5 minutes
func (r *Repository) GetDevicesLocalPv2() (DevicesLocalPv, error) {

	var values DevicesLocalPv

	err := r.db.QueryRow("SELECT id, dt_created, avg(JSON_VALUE(processdata,'$.devices:local:pv2.I.value')) AS i, avg(JSON_VALUE(processdata,'$.devices:local:pv2.P.value')) AS p, avg(JSON_VALUE(processdata,'$.devices:local:pv2.U.value')) AS u FROM solardata WHERE dt_created < NOW() AND dt_created > NOW() - INTERVAL 30 SECOND").Scan(&values.ID, &values.DateCreated, &values.I, &values.P, &values.U)

	if err != nil {
		//log.Println("Database problem in GetDevicesPv2: " + err.Error())
		return values, err
	}
	return values, nil
}

// GetDevicesLocalLast loads the values from devices:local section which are specified as last values
func (r *Repository) GetDevicesLocalLast() (DevicesLocalLast, error) {

	var values DevicesLocalLast

	err := r.db.QueryRow("SELECT id, dt_created, JSON_VALUE(processdata,'$.devices:local.Inverter:State.value') AS inverter_state, JSON_VALUE(processdata,'$.devices:local.SinkMax_P.value') AS sink_max_p, JSON_VALUE(processdata,'$.devices:local.SourceMax_P.value') AS source_max_p, JSON_VALUE(processdata,'$.devices:local.WorkTime.value') AS work_time FROM solardata order by dt_created desc LIMIT 0,1").Scan(&values.ID, &values.DateCreated, &values.InverterState, &values.SinkMax_P, &values.SourceMax_P, &values.WorkTime)

	if err != nil {
		//log.Println("Database problem in GetDevicesLocalLast: " + err.Error())
		return values, err
	}

	return values, nil

}

// GetStatisticEnergyFlowLast loads the last values from all scb:Statistics:EnergyFlow sections
func (r *Repository) GetStatisticEnergyFlowLast() (StatisticEnergyFlowLast, error) {

	var v StatisticEnergyFlowLast

	err := r.db.QueryRow("SELECT id, dt_created, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:Autarky:Day.value') AS statistic_autarky_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:Autarky:Month.value') AS statistic_autarky_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:Autarky:Total.value') AS statistic_autarky_total, 	JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:Autarky:Year.value') AS statistic_autarky_year, 	JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:CO2Saving:Day.value') AS statistic_co2saving_day, 	JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:CO2Saving:Month.value') AS statistic_co2saving_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:CO2Saving:Total.value') AS statistic_co2saving_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:CO2Saving:Year.value') AS statistic_co2saving_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargeGrid:Day.value') AS statistic_energy_charge_grid_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargeGrid:Month.value') AS statistic_energy_charge_grid_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargeGrid:Total.value') AS statistic_energy_charge_grid_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargeGrid:Year.value') AS statistic_energy_charge_grid_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargeInvIn:Day.value') AS statistic_energy_charge_inv_in_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargeInvIn:Month.value') AS statistic_energy_charge_inv_in_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargeInvIn:Total.value') AS statistic_energy_charge_inv_in_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargeInvIn:Year.value') AS statistic_energy_charge_inv_in_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargePv:Day.value') AS statistic_energy_charge_pv_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargePv:Month.value') AS statistic_energy_charge_pv_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargePv:Total.value') AS statistic_energy_charge_pv_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyChargePv:Year.value') AS statistic_energy_charge_pv_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyDischarge:Day.value') AS statistic_energy_discharge_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyDischarge:Month.value') AS statistic_energy_discharge_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyDischarge:Total.value') AS statistic_energy_discharge_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyDischarge:Year.value') AS statistic_energy_discharge_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyDischargeGrid:Day.value') AS statistic_energy_discharge_grid_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyDischargeGrid:Month.value') AS statistic_energy_discharge_grid_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyDischargeGrid:Total.value') AS statistic_energy_discharge_grid_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyDischargeGrid:Year.value') AS statistic_energy_discharge_grid_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHome:Day.value') AS statistic_energy_home_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHome:Month.value') AS statistic_energy_home_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHome:Total.value') AS statistic_energy_home_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHome:Year.value') AS statistic_energy_home_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomeBat:Day.value') AS statistic_energy_home_bat_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomeBat:Month.value') AS statistic_energy_home_bat_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomeBat:Total.value') AS statistic_energy_home_bat_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomeBat:Year.value') AS statistic_energy_home_bat_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomeGrid:Day.value') AS statistic_energy_home_grid_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomeGrid:Month.value') AS statistic_energy_home_grid_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomeGrid:Total.value') AS statistic_energy_home_grid_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomeGrid:Year.value') AS statistic_energy_home_grid_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomeOwn:Total.value') AS statistic_energy_home_own_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomePv:Day.value') AS statistic_energy_home_pv_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomePv:Month.value') AS statistic_energy_home_pv_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomePv:Total.value') AS statistic_energy_home_pv_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyHomePv:Year.value') AS statistic_energy_home_pv_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv1:Day.value') AS statistic_energy_pv1_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv1:Month.value') AS statistic_energy_pv1_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv1:Total.value') AS statistic_energy_pv1_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv1:Year.value') AS statistic_energy_pv1_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv2:Day.value') AS statistic_energy_pv2_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv2:Month.value') AS statistic_energy_pv2_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv2:Total.value') AS statistic_energy_pv2_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv2:Year.value') AS statistic_energy_pv2_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv3:Day.value') AS statistic_energy_pv3_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv3:Month.value') AS statistic_energy_pv3_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv3:Total.value') AS statistic_energy_pv3_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:EnergyPv3:Year.value') AS statistic_energy_pv3_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:OwnConsumptionRate:Day.value') AS statistic_own_consumption_rate_day, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:OwnConsumptionRate:Month.value') AS statistic_own_consumption_rate_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:OwnConsumptionRate:Total.value') AS statistic_own_consumption_rate_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:OwnConsumptionRate:Year.value') AS statistic_own_consumption_rate_year, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:Yield:Day.value') AS statistic_yield_day, 	JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:Yield:Month.value') AS statistic_yield_month, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:Yield:Total.value') AS statistic_yield_total, JSON_VALUE(processdata,'$.scb:statistic:EnergyFlow.Statistic:Yield:Year.value') AS statistic_yield_year FROM solardata order by dt_created desc LIMIT 0,1").Scan(&v.ID, &v.DateCreated, &v.StatisticAutarkyDay, &v.StatisticAutarkyMonth, &v.StatisticAutarkyTotal, &v.StatisticAutarkyYear, &v.StatisticCO2SavingDay, &v.StatisticCO2SavingMonth, &v.StatisticCO2SavingTotal, &v.StatisticCO2SavingYear, &v.StatisticEnergyChargeGridDay, &v.StatisticEnergyChargeGridMonth, &v.StatisticEnergyChargeGridTotal, &v.StatisticEnergyChargeGridYear, &v.StatisticEnergyChargeInvInDay, &v.StatisticEnergyChargeInvInMonth, &v.StatisticEnergyChargeInvInTotal, &v.StatisticEnergyChargeInvInYear, &v.StatisticEnergyChargePvDay, &v.StatisticEnergyChargePvMonth, &v.StatisticEnergyChargePvTotal, &v.StatisticEnergyChargePvYear, &v.StatisticEnergyDischargeDay, &v.StatisticEnergyDischargeMonth, &v.StatisticEnergyDischargeTotal, &v.StatisticEnergyDischargeYear, &v.StatisticEnergyDischargeGridDay, &v.StatisticEnergyDischargeGridMonth, &v.StatisticEnergyDischargeGridTotal, &v.StatisticEnergyDischargeGridYear, &v.StatisticEnergyHomeDay, &v.StatisticEnergyHomeMonth, &v.StatisticEnergyHomeTotal, &v.StatisticEnergyHomeYear, &v.StatisticEnergyHomeBatDay, &v.StatisticEnergyHomeBatMonth, &v.StatisticEnergyHomeBatTotal, &v.StatisticEnergyHomeBatYear, &v.StatisticEnergyHomeGridDay, &v.StatisticEnergyHomeGridMonth, &v.StatisticEnergyHomeGridTotal, &v.StatisticEnergyHomeGridYear, &v.StatisticEnergyHomeOwnTotal, &v.StatisticEnergyHomePvDay, &v.StatisticEnergyHomePvMonth, &v.StatisticEnergyHomePvTotal, &v.StatisticEnergyHomePvYear, &v.StatisticEnergyPv1Day, &v.StatisticEnergyPv1Month, &v.StatisticEnergyPv1Total, &v.StatisticEnergyPv1Year, &v.StatisticEnergyPv2Day, &v.StatisticEnergyPv2Month, &v.StatisticEnergyPv2Total, &v.StatisticEnergyPv2Year, &v.StatisticEnergyPv3Day, &v.StatisticEnergyPv3Month, &v.StatisticEnergyPv3Total, &v.StatisticEnergyPv3Year, &v.StatisticOwnConsumptionRateDay, &v.StatisticOwnConsumptionRateMonth, &v.StatisticOwnConsumptionRateTotal, &v.StatisticOwnConsumptionRateYear, &v.StatisticYieldDay, &v.StatisticYieldMonth, &v.StatisticYieldTotal, &v.StatisticYieldYear)

	if err != nil {
		//log.Println("Database problem in GetDevicesLocalLast: " + err.Error())
		return v, err
	}

	return v, nil

}

// AddData adds dataset with JSON payload to solardata table
func (r *Repository) AddData(payload string) (int64, error) {

	stmt, err := r.db.Prepare("INSERT INTO solardata (processdata) VALUES(?)")
	if err != nil {
		fmt.Println(err.Error())
		return -1, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(payload)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// RemoveData deletes entries from solardata table which are older than "olderThanDays" days.
func (r *Repository) RemoveData(olderThanDays int) (int64, error) {
	var rowsAffected int64

	stmt, err := r.db.Prepare("DELETE FROM solardata WHERE dt_created < NOW() - INTERVAL ? DAY")
	if err != nil {
		return rowsAffected, err
	}
	defer stmt.Close()

	deleteResult, err := stmt.Exec(olderThanDays)
	if err != nil {
		return rowsAffected, err
	}

	rowsAffected, err = deleteResult.RowsAffected()
	if err != nil {
		return rowsAffected, err
	}

	return rowsAffected, nil

}

package iryaml

type Driver struct {
	CarIdx                  int     `yaml:"CarIdx"`
	UserName                string  `yaml:"UserName"`
	AbbrevName              string  `yaml:"AbbrevName"`
	Initials                string  `yaml:"Initials"`
	UserID                  int     `yaml:"UserID"`
	TeamID                  int     `yaml:"TeamID"`
	TeamName                string  `yaml:"TeamName"`
	CarNumber               string  `yaml:"CarNumber"`
	CarNumberRaw            int     `yaml:"CarNumberRaw"`
	CarPath                 string  `yaml:"CarPath"`
	CarClassID              int     `yaml:"CarClassID"`
	CarID                   int     `yaml:"CarID"`
	CarIsPaceCar            int     `yaml:"CarIsPaceCar"`
	CarIsAI                 int     `yaml:"CarIsAI"`
	CarIsElectric           int     `yaml:"CarIsElectric"`
	CarScreenName           string  `yaml:"CarScreenName"`
	CarScreenNameShort      string  `yaml:"CarScreenNameShort"`
	CarClassShortName       string  `yaml:"CarClassShortName"`
	CarClassRelSpeed        int     `yaml:"CarClassRelSpeed"`
	CarClassLicenseLevel    int     `yaml:"CarClassLicenseLevel"`
	CarClassMaxFuelPct      string  `yaml:"CarClassMaxFuelPct"`
	CarClassWeightPenalty   string  `yaml:"CarClassWeightPenalty"`
	CarClassPowerAdjust     string  `yaml:"CarClassPowerAdjust"`
	CarClassDryTireSetLimit string  `yaml:"CarClassDryTireSetLimit"`
	CarClassColor           string  `yaml:"CarClassColor"`
	CarClassEstLapTime      float64 `yaml:"CarClassEstLapTime"`
	IRating                 int     `yaml:"IRating"`
	LicLevel                int     `yaml:"LicLevel"`
	LicSubLevel             int     `yaml:"LicSubLevel"`
	LicString               string  `yaml:"LicString"`
	LicColor                string  `yaml:"LicColor"`
	IsSpectator             int     `yaml:"IsSpectator"`
	CarDesignStr            string  `yaml:"CarDesignStr"`
	HelmetDesignStr         string  `yaml:"HelmetDesignStr"`
	SuitDesignStr           string  `yaml:"SuitDesignStr"`
	BodyType                int     `yaml:"BodyType"`
	FaceType                int     `yaml:"FaceType"`
	HelmetType              int     `yaml:"HelmetType"`
	CarNumberDesignStr      string  `yaml:"CarNumberDesignStr"`
	CarSponsor1             int     `yaml:"CarSponsor_1"`
	CarSponsor2             int     `yaml:"CarSponsor_2"`
	ClubName                string  `yaml:"ClubName"`
	ClubID                  int     `yaml:"ClubID"`
	DivisionName            string  `yaml:"DivisionName"`
	DivisionID              int     `yaml:"DivisionID"`
	CurDriverIncidentCount  int     `yaml:"CurDriverIncidentCount"`
	TeamIncidentCount       int     `yaml:"TeamIncidentCount"`
}

type DriverInfo struct {
	DriverCarIdx              int      `yaml:"DriverCarIdx"`
	DriverUserID              int      `yaml:"DriverUserID"`
	PaceCarIdx                int      `yaml:"PaceCarIdx"`
	DriverHeadPosX            float64  `yaml:"DriverHeadPosX"`
	DriverHeadPosY            float64  `yaml:"DriverHeadPosY"`
	DriverHeadPosZ            float64  `yaml:"DriverHeadPosZ"`
	DriverCarIsElectric       int      `yaml:"DriverCarIsElectric"`
	DriverCarIdleRPM          float64  `yaml:"DriverCarIdleRPM"`
	DriverCarRedLine          float64  `yaml:"DriverCarRedLine"`
	DriverCarEngCylinderCount int      `yaml:"DriverCarEngCylinderCount"`
	DriverCarFuelKgPerLtr     float64  `yaml:"DriverCarFuelKgPerLtr"`
	DriverCarFuelMaxLtr       float64  `yaml:"DriverCarFuelMaxLtr"`
	DriverCarMaxFuelPct       float64  `yaml:"DriverCarMaxFuelPct"`
	DriverCarGearNumForward   int      `yaml:"DriverCarGearNumForward"`
	DriverCarGearNeutral      int      `yaml:"DriverCarGearNeutral"`
	DriverCarGearReverse      int      `yaml:"DriverCarGearReverse"`
	DriverCarSLFirstRPM       float64  `yaml:"DriverCarSLFirstRPM"`
	DriverCarSLShiftRPM       float64  `yaml:"DriverCarSLShiftRPM"`
	DriverCarSLLastRPM        float64  `yaml:"DriverCarSLLastRPM"`
	DriverCarSLBlinkRPM       float64  `yaml:"DriverCarSLBlinkRPM"`
	DriverCarVersion          string   `yaml:"DriverCarVersion"`
	DriverPitTrkPct           float64  `yaml:"DriverPitTrkPct"`
	DriverCarEstLapTime       float64  `yaml:"DriverCarEstLapTime"`
	DriverSetupName           string   `yaml:"DriverSetupName"`
	DriverSetupIsModified     int      `yaml:"DriverSetupIsModified"`
	DriverSetupLoadTypeName   string   `yaml:"DriverSetupLoadTypeName"`
	DriverSetupPassedTech     int      `yaml:"DriverSetupPassedTech"`
	DriverIncidentCount       int      `yaml:"DriverIncidentCount"`
	Drivers                   []Driver `yaml:"Drivers"`
}

package members

import (
	"github.com/pkg/errors"
)

type Grade int

const (
	CdtAB = iota
	CdtAmn
	CdtA1C
	CdtSrA
	CdtSSgt
	CdtTSgt
	CdtMSgt
	CdtSMSgt
	CdtCMSgt
	CdtSecondLt
	CdtFirstLt
	CdtCapt
	CdtMaj
	CdtLtCol
	CdtCol
	SM
	SSgt
	TSgt
	MSgt
	SMSgt
	CMSgt
	FO
	TFO
	SFO
	SecondLt
	FirstLt
	Capt
	Maj
	LtCol
	Col
	BrigGen
	MajGen
)

func ParseGrade(s string) (Grade, error) {
	switch s {
	case "Maj Gen":
		return MajGen, nil
	case "Brig Gen":
		return BrigGen, nil
	case "Col":
		return Col, nil
	case "Lt Col":
		return LtCol, nil
	case "Maj":
		return Maj, nil
	case "Capt":
		return Capt, nil
	case "1st Lt":
		return FirstLt, nil
	case "2d Lt":
		return SecondLt, nil
	case "SFO":
		return SFO, nil
	case "TFO":
		return TFO, nil
	case "FO":
		return FO, nil
	case "CSMgt":
		return CMSgt, nil
	case "SMSgt":
		return SMSgt, nil
	case "MSgt":
		return MSgt, nil
	case "TSgt":
		return TSgt, nil
	case "SSgt":
		return SSgt, nil
	case "SM":
		return SM, nil
	case "C/Col":
		return CdtCol, nil
	case "C/Lt Col":
		return CdtLtCol, nil
	case "C/Maj":
		return CdtMaj, nil
	case "C/Capt":
		return CdtCapt, nil
	case "C/1st Lt":
		return CdtFirstLt, nil
	case "C/2d Lt":
		return CdtSecondLt, nil
	case "C/CMSgt":
		return CdtMSgt, nil
	case "C/SMSgt":
		return CdtSMSgt, nil
	case "C/MSgt":
		return CdtMSgt, nil
	case "C/TSgt":
		return CdtTSgt, nil
	case "C/SSgt":
		return CdtSSgt, nil
	case "C/SrA":
		return CdtSrA, nil
	case "C/A1C":
		return CdtA1C, nil
	case "C/Amn":
		return CdtAmn, nil
	case "C/AB":
		fallthrough
	case "CADET":
		return CdtAB, nil
	default:
		return -1, errors.Errorf("unrecognized CAP grade: %s", s)
	}
}

func (g Grade) String() string {
	switch g {
	case CdtAB:
		return "C/AB"
	case CdtAmn:
		return "C/Amn"
	case CdtA1C:
		return "C/A1C"
	case CdtSrA:
		return "C/SrA"
	case CdtSSgt:
		return "C/SSgt"
	case CdtTSgt:
		return "C/TSgt"
	case CdtMSgt:
		return "C/MSgt"
	case CdtSMSgt:
		return "C/SMSgt"
	case CdtCMSgt:
		return "C/CMSgt"
	case CdtSecondLt:
		return "C/2d Lt"
	case CdtFirstLt:
		return "C/1st Lt"
	case CdtCapt:
		return "C/Capt"
	case CdtMaj:
		return "C/Maj"
	case CdtLtCol:
		return "C/Lt Col"
	case CdtCol:
		return "C/Col"
	case SM:
		return "SM"
	case SSgt:
		return "SSgt"
	case TSgt:
		return "TSgt"
	case MSgt:
		return "MSgt"
	case SMSgt:
		return "SMSgt"
	case CMSgt:
		return "CMSgt"
	case FO:
		return "FO"
	case TFO:
		return "TFO"
	case SFO:
		return "SFO"
	case SecondLt:
		return "2d Lt"
	case FirstLt:
		return "1st Lt"
	case Capt:
		return "Capt"
	case Maj:
		return "Maj"
	case LtCol:
		return "Lt Col"
	case Col:
		return "Col"
	case BrigGen:
		return "Brig Gen"
	case MajGen:
		return "Maj Gen"
	default:
		panic(errors.Errorf("invalid value for grade: %d", g))
	}
}

func (g Grade) ExtAbbv() string {
	switch g {
	case CdtAB:
		return "Cadet Airman"
	case CdtAmn:
		return "Cadet Airman"
	case CdtA1C:
		return "Cadet Airman 1st Class"
	case CdtSrA:
		return "Cadet Senior Airman"
	case CdtSSgt:
		return "Cadet Staff Sgt."
	case CdtTSgt:
		return "Cadet Tech. Sgt."
	case CdtMSgt:
		return "Cadet Master Sgt."
	case CdtSMSgt:
		return "Cadet Senior Master Sgt."
	case CdtCMSgt:
		return "Cadet Chief Master Sgt."
	case CdtSecondLt:
		return "Cadet 2nd Lt."
	case CdtFirstLt:
		return "Cadet 1st Lt."
	case CdtCapt:
		return "Cadet Capt."
	case CdtMaj:
		return "Cadet Maj."
	case CdtLtCol:
		return "Cadet Lt. Col."
	case CdtCol:
		return "Cadet Col."
	case SM:
		return ""
	case SSgt:
		return "Staff Sgt."
	case TSgt:
		return "Tech. Sgt."
	case MSgt:
		return "Master Sgt."
	case SMSgt:
		return "Senior Master Sgt."
	case CMSgt:
		return "Chief Master Sgt."
	case FO:
		return "Flight Officer"
	case TFO:
		return "Tech. Flight Officer"
	case SFO:
		return "Senior Flight Officer"
	case SecondLt:
		return "2nd Lt."
	case FirstLt:
		return "1st Lt."
	case Capt:
		return "Capt."
	case Maj:
		return "Maj."
	case LtCol:
		return "Lt. Col."
	case Col:
		return "Col."
	case BrigGen:
		return "Brig. Gen."
	case MajGen:
		return "Maj. Gen."
	default:
		panic(errors.Errorf("invalid grade value: %d", g))
	}
}

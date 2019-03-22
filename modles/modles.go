package modles

import (
	_ "github.com/Go-SQL-Driver/mysql"
	"time"
)

// order table
type Orders struct {
	Id					int 					`json:"order_id"`
	HotelId				int						`json:"hotel_id"`
	Date				time.Time				`json:"date"`
	StartTime			time.Time				`json:"start_time"`
	Duration			time.Duration			`json:"duration"`
	GenderMode			string					`json:"gender_mode"`
	Count				int						`json:"count"`
	CountMale			int						`json:"count_male"`
	CountFemale	 		int						`json:"count_female"`
	CountYet			int						`json:"count_yet"`
	CountMaleYet		int						`json:"count_male_yet"`
	CountFemaleYet		int						`json:"count_female_yet"`
	CountChanged		int						`json:"count_changed"`
	CountMaleChanged	int						`json:"count_male_changed"`
	CountFemaleChanged	int						`json:"count_female_changed"`
	AdviserId 			int						`json:"adviser_id"`
	IsTimeChanged		bool					`json:"is_time_changed"`
	StartTimeChanged	time.Time				`json:"start_time_changed"`
	DurationChanged		time.Time				`json:"duration_changed"`
	IsCountChanged		bool					`json:"is_count_changed"`
	GenderModeChanged	string					`json:"gender_mode_changed"`
	HourlySalary		int						`json:"hourly_salary"`
	IsFloat				bool					`json:"is_float"`
	WorkContent			string					`json:"work_content"`
	Attention			string					`json:"attention"`
	IsClosed			bool					`json:"is_closed"`
}

//// hotel basic information
//type Hotel struct {
//	HotelId					int						`json:"hotel_id"`
//	HotelName				string					`json:"hotel_name"`
//	HotelPhone				string					`json:"hotel_phone"`
//	HotelLatitude			string					`json:"hotel_latitude"`
//	HotelLongitude			string					`json:"hotel_longitude"`
//	HotelIntro				string					`json:"hotel_intro"`
//
//}

// adviser information
type Adviser struct {
	Id						*Orders					`json:"order_id"`
	AdviserId				int						`json:"adviser_id"`
	AdviserPhone			string					`json:"adviser_phone"`
	AdviserCompany			string					`json:"adviser_company"`
	AdviserCompanyIntro 	string					`json:"adviser_company_intro"`

}

// pt information
type PT struct {
	Id 						*Orders					`json:"order_id"`
	PTId					int						`json:"pt_id"`
	PTAvatar				TODO					`json:"pt_avatar"`
	PTName					string					`json:"pt_name"`
	PTPhone					string					`json:"pt_phone"`
	PTGender				string					`json:"pt_gender"`
	PTNickName				string					`json:"pt_nick_name"`
	PTStatus				string					`json:"pt_status"`
	PTAge					int						`json:"pt_age"`
	PTHeight				int						`json:"pt_height"`
	PTWeight				int						`json:"pt_weight"`
	PTWorkTimes				int						`json:"pt_work_times"`
	PTWorkHours				int						`json:"pt_work_hours"`
	PTLatestFive			TODO					`json:"pt_latest_five"`
	PTRegistrationChannel	string					`json:"pt_registration_channel"`
}
